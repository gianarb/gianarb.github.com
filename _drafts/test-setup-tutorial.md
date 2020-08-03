[Tinkerbell](https://tinkerbell.org) is a tool open sourced recently by
[Packet](https://packet.com), the company I work for.

It is a provisioner for bare metal. It has an API that you can use to switch on
and off servers via IPMI or to execute workflows and install operating systems
on server that does not have one!!

Tinkerbell is in its early days as open source project but the concept is battle
tested from 6 years of production use internally at Packet.

I am excited to lean a lot of the cool technologies that are making datacenters
working, but I am not here to write about it[^1].

One of my recent tasks[^2] was about end to end testing the Vagrant Setup
tutorial[^3] we wrote.

I like the idea! The Setup tutorial is important for our community because it is
the entry point for a lot of people and having a consistent way to test its
accuracy is crucial.

It is also a quick way to get a valuable end to end test running that covers the
all project.

Tinkerbell is under development and it is easy to make mistakes and break
things at this point, we have to know when it happens. Tinkerbell requires
virtualisation capabilities, and we do not have an end to end testing framework
for that yet.

## Tell me more about the test itself

It is a long tasks but let's summarize it (have a look at the tutorial, it helps
to read this article moving forward):

1. The script has to start a vagrant machine called provisioner
2. When the provisioner is up it has to exec via ssh a docker-compose command
   that starts a bunch of containers, one of those is Tink grpc server
3. When Tinkerbell is up and running we have to do a bunch of things like:
    a. Register a new hardware
    b. Create a template
    c. Create the workflow that will get executed in the worker from a template
4. Start the worker
5. Wait and check if the workflows executes as expected.

NOTE: the test should  clean up after itself, Vagrant is not ideal
to get parallelization of VMs, and we do not support it. A dirty environment
will break future tests as it is today.

## How to write this test

There are a million way to write end to end test the one I evaluated are bash
and Go. 

The project is in Go, Tinkerbell serves a gRPC server and a client, I thought it
was a good idea to write everything in Go to try the client itself and because
it is easier to coordinate long running actions with channels and context
compared with bash for example. Or at least that's what I think.

I can also keep the code inside the `testing` framework that Go provides keeping
the test closer to the code and the developers that contribute to the project,
compared with a random `scripts.sh`.

I am not sure if this will be useful in the future but one of my goal was to
serve a clean API and a small framework that can be used to write other tests
that starts from the Vagrant setup. This is the API I designed:

```go
type Vagrant struct {}

func Up(ctx context.Context, opts ...VagrantOpt) (*Vagrant, error) {}

func (v *Vagrant) Destroy(ctx context.Context) error {}

func (v *Vagrant) Exec(ctx context.Context, args ...string) ([]byte, error) {}
```

Consistency is important, developers who knows vagrant or that will have to fix
the tests coming from the tutorial will know `Up`, `Destroy` and `Exec` because
those verbs are used by Vagrant and in the documentation itself.

Even for Go developers `Exec` is not a new function, `os/exec`[^4] exists and it
does a similar job, the one I wrote is over ssh.

## Go challenges and tips and tricks

I would like to share some of the challenges I faced when writing the Vagrant
framework and some tips useful for this task.

## Opt are great!

I have to say options are great! It is a well known pattern in Go and it
translates to:

```go
ctx := context.Background()

machine, err := vagrant.Up(ctx,
    vagrant.WithLogger(t.Logf),
    vagrant.WithMachineName("provisioner"),
    vagrant.WithWorkdir("../../deploy/vagrant"),
)
if err != nil {
    t.Fatal(err)
}
```

It allowed me to add new options and to tune the Vagrant struct with strong
default. If you never used it, do it! It is pretty easy, you need an interface
like this:

```go
type VagrantOpt func(*Vagrant)
```

In this way you can write as many `With`function you need:

```go
func WithStderr(s io.ReadWriter) VagrantOpt {
	return func(v *Vagrant) {
		v.Stderr = s
	}
}

func RunAsync() VagrantOpt {
	return func(v *Vagrant) {
		v.async = true
	}
}
```

I execute the opts as part of the `Up` function:

```go
func Up(ctx context.Context, opts ...VagrantOpt) (*Vagrant, error) {
	const (
		defaultVagrantBin = "vagrant"
		defaultName       = "vagrant"
		defaultWorkdir    = "."
	)
	v := &Vagrant{
		VagrantBinPath: defaultVagrantBin,
		Name:           defaultName,
		Workdir:        defaultWorkdir,
		log: func(format string, args ...interface{}) {
			fmt.Println(fmt.Sprintf(format, args))
		},
	}
	for _, opt := range opts {
		opt(v)
	}

    // ...
}
```

### test segmentation with packages

I don't want to run the vagrant end to end tests as part of the default test
suite because they take too much time and they require Vagrant installed. They
do not even run in CI in the same way unit test works, but I will get to it
later.

I learned that packages that starts with `_` does not get executed when using
something like `./...`.

I wrote the framework and tests as part of the package:

```console
./test/_vagrant/
    ./vagrant.go
    ./vagrant_test.go
```

In this way to run the tests you have to explicitly call the package out:

```console
$ go test ./test/_vagrant
```

### Observability or "what is going on?"

Go has its own way to print logs during the execution of the tests:

```console
$ go test -v ./...
```

It works because `testing` has a function called `t.Log` and `t.Logf`. Those
functions watches the `-v` flags. To be complaint with that and to keep the
`Vagrant` struct agnostic I wrote a `WithLogger`:

```go
func WithLogger(log func(string, ...interface{})) VagrantOpt {
	return func(v *Vagrant) {
		v.log = log
	}
}
```

The function it accepts as a argument is `t.Logf`.

Continuous Integrations runs with verbosity enabled for this task because it is
long and complicated, the logging prints all the outputs from the `vagrant up`
and `destroy` command, and the stdout for the `exec` over ssh, it gives a very
good overview about what is going on.


### Stdout and Stdin, buffer and loggers

I don't have a lot to say about this other than: "it was very hard to do!!".
The code that fixed my problems can be summarized in this way:

```go
stderrPipe, err := cmd.StderrPipe()
if err != nil {
    return nil, fmt.Errorf("exec error: %v", err)
}
stdoutPipe, err := cmd.StdoutPipe()
if err != nil {
    return nil, fmt.Errorf("exec error: %v", err)
}

go v.pipeOutput(ctx, fmt.Sprintf("%s stderr", cmd.String()), bufio.NewScanner(stderrPipe))
go v.pipeOutput(ctx, fmt.Sprintf("%s stdout", cmd.String()), bufio.NewScanner(stdoutPipe))

err = cmd.Start()
```

```go
func (v *Vagrant) pipeOutput(ctx context.Context, name string, scanner *bufio.Scanner) {
	for scanner.Scan() {
		select {
		case <-ctx.Done():
			return
		default:
			v.log("[pipeOutput %s] %s", name, scanner.Text())
		}
	}
}
```

### Kill process and subprocess

There are a lot of process going on when creating or destroying a VM with
Vagrant. There is VirtualBox for example, and we have an edge case for the
worker machine because the `up` commands technically never ends, it is in
pending until you `destroy` the machine. But you can't run multiple commands
against the same machine because `up` hols a lock and if blocks `destroy` to
execute. `os/exec` helps here but you have to tune it a little bit:

```go
cmd := exec.CommandContext(ctx, v.VagrantBinPath, args...)
cmd.Dir = v.Workdir
cmd.Stdout = v.Stdout
cmd.Stderr = v.Stderr
cmd.SysProcAttr = &syscall.SysProcAttr{Setpgid: true}
```

Now when killing `cmd` the subprocess terminates as well.

## Continuous Integration

We decided to go with GitHub Actions with a self running runner, in this way we
can use Packet bare metal that supports virtualisation.

As I told you I don't want this test to run for all the commit, or for all the
pull request because it is time and resource consuming. It is also risky, so I
want maintainers to decide when to trigger it.

That's why it gets triggered with a GitHub label:

```yaml
name: Setup with Vagrant on Packet
on:
  push:
  pull_request:
    types: [labeled]

jobs:
  vagrant-setup:
    if: contains(github.event.pull_request.labels.*.name, 'ci-check/vagrant-setup')
    runs-on: self-hosted
    steps:
    - name: Checkout
      uses: actions/checkout@v2
    - name: Vagrant Test
      run: |
        export VAGRANT_DEFAULT_PROVIDER="virtualbox"
        go test -v ./test/_vagrant
```

This is what it takes to make the process working!! And I am still surprised it
is so easy! When a contributor label a PR with `ci-check/vagrant-setup` the
process starts. My idea was to remove the label straight away, but I am
[blocked](https://github.community/t/actions-ecosystem-action-remove-labels-fails-resource-not-accessible-by-integration/124188).

An alternative that we are evaluating is to run it as a cronjob[^5] as well.

## Conclusion

[^1]: If you are curious ask me any question on Twitter @gianarb
[^2]: https://github.com/tinkerbell/tink/pull/224
[^3]: https://tinkerbell.org/setup/local-with-vagrant/
[^4]: https://golang.org/pkg/os/exec/#pkg-examples
[^5]: https://docs.github.com/en/actions/reference/workflow-syntax-for-github-actions#onschedule
