---
layout: post
title:  "LinuxKit operating system built for container"
date:   2017-04-18 10:08:27
categories: [post]
img: /img/builder.gif
tags: [docker, linuxkit, container]
summary: "LinuxKit is a new tool presented during the DockerCon 2017 built by
Docker to manage cross architecture and cross kernel testing. LinuxKit is a
secure, portable and lean operating system built for containers. It supports
different hypervisor as MacOS hyper or QEMU to run testsuite on different
architectures. In this article I am showing you some basic concept above this
tool. How it works and why it can be useful."
changefreq: yearly
---
Linuxkit is a new project presented by Docker during the DockerCon 2017. If we
look at the description of the project on
[GitHub](https://github.com/linuxkit/linuxkit):

> A secure, portable and lean operating system built for containers

I am feeling already exited. I was an observer of the project when [Justin
Cormack](https://twitter.com/justincormack) and the other
[contributors](https://github.com/linuxkit/linuxkit/graphs/contributors) was
working on a private repository. I was invited as part of ci-wg group into the
CNCF and I loved this project from the first day.

You can think about linuxkit as a builder for Linux operating system everything
based on containers.

It's a project that can stay behind your continuous integration system to allow
us to test on different kernel version and distribution. You can a light kernels
with all the services that you need and you can create different outputs
runnable on cloud providers as Google Cloud Platform, with Docker or with QEMU.

## Continuous delivery, new model

I am not really confident about Google Cloud Platform but just to move over I am
going to do some math with AWS as provider.
Let's suppose that I have the most common continuous integration system, one big
box always up an running configured to support all your projects or if you are
already good you are running containers to have separated and isolated
environment.

Let's suppose that you Jenkins is running all times on m3.xlarge:

`m3.xlarge` used 100% every months costs 194.72$.

Let's have a dream. You have a very small server with just a frontend
application for your CI and all jobs are running in a separate instance, tiny as
a t2.small.

`t2.small` used only 1 hour costs 0.72$ .

I calculated 1 hour because it's the minimum that you can pay and I hope that
your CI job can run for less than 1 hour.
Easy math to calculate the number of builds that you need to run to pay as you
was paying before.

194.72 / 0.72 ~ 270 builds every month.

If you are running less than 270 builds a months you can save some money
too. But you have other benefits:

1. More jobs, more instances. Very easy to scale. Easier that Jenkins
   master/slave and so on.
2. How many times during holidays your Jenkins is still up and running without
   to have nothing to do? During these days you are just paying for the frontend
   app.

And these are just the benefit to have a different setup for your continuous
delivery.

## LinuxKit CI implementation

There is a directory called
[./test](https://github.com/linuxkit/linuxkit/tree/master/test) that contains
some linuxkit use case but I am going to explain in practice how linuxkit is
tested. Because it uses itself, awesome!

In first you need to download and compile linuxkit:
```shell
git clone github.com:linuxkit/linuxkit $GOPATH/src/github.com/linuxkit/linuxkit
make
./bin/moby
```
You can move it in your `$PATH` with `make install`.

```
$ moby
Please specify a command.

USAGE: moby [options] COMMAND

Commands:
  build       Build a Moby image from a YAML file
  run         Run a Moby image on a local hypervisor or remote cloud
  version     Print version information
  help        Print this message

Run 'moby COMMAND --help' for more information on the command

Options:
  -q    Quiet execution
  -v    Verbose execution
```

At the moment the CLI is very simple, the most important commands are build and
run.  linuxkit is based on YAML file that you can use to describe your kernel,
with all applications and all the services that you need.  Let's start with the
[linuxkit/test/test.yml](https://github.com/linuxkit/linuxkit/blob/master/test/test.yml).

```yaml
kernel:
  image: "mobylinux/kernel:4.9.x"
  cmdline: "console=ttyS0"
init:
  - mobylinux/init:8375addb923b8b88b2209740309c92aa5f2a4f9d
  - mobylinux/runc:b0fb122e10dbb7e4e45115177a61a3f8d68c19a9
  - mobylinux/containerd:18eaf72f3f4f9a9f29ca1951f66df701f873060b
  - mobylinux/ca-certificates:eabc5a6e59f05aa91529d80e9a595b85b046f935
onboot:
  - name: dhcpcd
	image: "mobylinux/dhcpcd:0d4012269cb142972fed8542fbdc3ff5a7b695cd"
	binds:
	 - /var:/var
	 - /tmp:/etc
	capabilities:
	 - CAP_NET_ADMIN
	 - CAP_NET_BIND_SERVICE
	 - CAP_NET_RAW
	net: host
	command: ["/sbin/dhcpcd", "--nobackground", "-f", "/dhcpcd.conf", "-1"]
  - name: check
	image: "mobylinux/check:c9e41ab96b3ea6a3ced97634751e20d12a5bf52f"
	pid: host
	capabilities:
	 - CAP_SYS_BOOT
	readonly: true
outputs:
  - format: kernel+initrd
  - format: iso-bios
  - format: iso-efi
  - format: gcp-img
```

Linuxkit builds everythings inside a container, it means that you don't need a
lot of dependencies it's very easy to use. It generates different `output` in
this case `kernel+initrd`, `iso-bios`, `iso-efi`, `gpc-img`  depends of the
platform that you are interested to use to run your kernel.

I am trying to explain a bit how this YAML works. You can see that there are
different primary section: `kernel`, `init`, `onboot`, `service` and so on.

Pretty much all of them contains the keyword `image` because as I said before
everything is applied on containers, in this example they are store in
[hub.docker.com/u/mobylinux/](https://hub.docker.com/u/mobylinux/).

The based kernel is `mobylinux/kernel:4.9.x`, I am just reporting what the
[README.md](https://github.com/linuxkit/linuxkit#yaml-specification) said:


- `kernel` specifies a kernel Docker image, containing a kernel and a
filesystem tarball, eg containing modules. The example kernels are built from
`kernel/`
- `init` is the base `init` process Docker image, which is unpacked as the base
system, containing `init`, `containerd`, `runc` and a few tools. Built from
`pkg/init/`
- `onboot` are the system containers, executed sequentially in order. They
should terminate quickly when done.
- `services` is the system services, which normally run for the whole time the
system is up
- `files` are additional files to add to the image
- `outputs` are descriptions of what to build, such as ISOs.

At this point we can try it. If you are on MacOS as I was you don't need to
install anything one of the runner supported by `linuxkit` is `hyperkit` it
means that everything is available in your system.

`./test` contains different test suite but now we will stay focused on
`./test/check` directory. It contains a set of checks to validate how the
kernel went build by LinuxKit. They are the smoke tests that are running on each
new pull request created on the repository for example.

As I said everything runs inside a container, if you look into the check
directory there is a makefile that build a mobylinux/check image, that image
went run in LinuxKit, into the `test.yml` file:

```yaml
onboot:
  - name: check
	image: "mobylinux/check:c9e41ab96b3ea6a3ced97634751e20d12a5bf52f"
	pid: host
	capabilities:
	 - CAP_SYS_BOOT
	readonly: true
```

You can use the
[Makefile](https://github.com/linuxkit/linuxkit/blob/master/test/check/Makefile)
inside the check directory to build a new version of check, you can just use
the command `make`.

When you have the right version of your test we can build the image used by moby:

```
cd $GOPATH/src/github.com/linuxkit/linuxkit
moby build test/test.yml
```

Part of the output is:

```shell
Create outputs:
  test-bzImage test-initrd.img test-cmdline
  test.iso
  test-efi.iso
  test.img.tar.gz
```

And if you look into the directory you can see that there are all these files
into the root. These files can be run from qemu, google cloud platform,
hyperkit and so on.

```shell
moby run test
```
On MacOS with this command LinuxKit is using hyperkit to start a VM, I can not copy
paste all the output but you can see the hypervisor logs:

```
virtio-net-vpnkit: initialising, opts="path=/Users/gianlucaarbezzano/Library/Containers/com.docker.docker/Data/s50"
virtio-net-vpnkit: magic=VMN3T version=1 commit=0123456789012345678901234567890123456789
Connection established with MAC=02:50:00:00:00:04 and MTU 1500
early console in extract_kernel
input_data: 0x0000000001f2c3b4
input_len: 0x000000000067b1e5
output: 0x0000000001000000
output_len: 0x0000000001595280
kernel_total_size: 0x000000000118a000
booted via startup_32()
Physical KASLR using RDRAND RDTSC...
Virtual KASLR using RDRAND RDTSC...

Decompressing Linux... Parsing ELF... Performing relocations... done.
Booting the kernel.
[    0.000000] Linux version 4.9.21-moby (root@84baa8e89c00) (gcc version 6.2.1 20160822 (Alpine 6.2.1) ) #1 SMP Sun Apr 9 22:21:32 UTC 2017
[    0.000000] Command line: earlyprintk=serial console=ttyS0
[    0.000000] x86/fpu: Supporting XSAVE feature 0x001: 'x87 floating point registers'
[    0.000000] x86/fpu: Supporting XSAVE feature 0x002: 'SSE registers'
[    0.000000] x86/fpu: Supporting XSAVE feature 0x004: 'AVX registers'
[    0.000000] x86/fpu: xstate_offset[2]:  576, xstate_sizes[2]:  256
[    0.000000] x86/fpu: Enabled xstate features 0x7, context size is 832 bytes, using 'standard' format.
[    0.000000] x86/fpu: Using 'eager' FPU context switches.
[    0.000000] e820: BIOS-provided physical RAM map:
[    0.000000] BIOS-e820: [mem 0x0000000000000000-0x000000000009fbff] usable
[    0.000000] BIOS-e820: [mem 0x0000000000100000-0x000000003fffffff] usable
```
When the VM is ready LinuxKit is starting all the `init`, `onboot`, the logs is
easy to understand as the `test.yml` is starting `containerd`, `runc`:

```
init:
  - mobylinux/init:8375addb923b8b88b2209740309c92aa5f2a4f9d
  - mobylinux/runc:b0fb122e10dbb7e4e45115177a61a3f8d68c19a9
  - mobylinux/containerd:18eaf72f3f4f9a9f29ca1951f66df701f873060b
  - mobylinux/ca-certificates:eabc5a6e59f05aa91529d80e9a595b85b046f935
onboot:
  - name: dhcpcd
	image: "mobylinux/dhcpcd:0d4012269cb142972fed8542fbdc3ff5a7b695cd"
	binds:
	 - /var:/var
	 - /tmp:/etc
	capabilities:
	 - CAP_NET_ADMIN
	 - CAP_NET_BIND_SERVICE
	 - CAP_NET_RAW
	net: host
	command: ["/sbin/dhcpcd", "--nobackground", "-f", "/dhcpcd.conf", "-1"]
  - name: check
	image: "mobylinux/check:c9e41ab96b3ea6a3ced97634751e20d12a5bf52f"
	pid: host
	capabilities:
	 - CAP_SYS_BOOT
	readonly: true
```

```
Welcome to LinuxKit

			##         .
		  ## ## ##        ==
		   ## ## ## ## ##    ===
	   /"""""""""""""""""\___/ ===
	  ~~~ {~~ ~~~~ ~~~ ~~~~ ~~~ ~ /  ===- ~~~
	   \______ o           __/
		 \    \         __/
		  \____\_______/


/ # INFO[0000] starting containerd boot...                   module=containerd
INFO[0000] starting debug API...                         debug="/run/containerd/debug.sock" module=containerd
INFO[0000] loading monitor plugin "cgroups"...           module=containerd
INFO[0000] loading runtime plugin "linux"...             module=containerd
INFO[0000] loading snapshot plugin "snapshot-overlay"...  module=containerd
INFO[0000] loading grpc service plugin "healthcheck-grpc"...  module=containerd
INFO[0000] loading grpc service plugin "images-grpc"...  module=containerd
INFO[0000] loading grpc service plugin "metrics-grpc"...  module=containerd
```
The last step is the `check` that runs the real test suite:

```
kernel config test succeeded!
info: reading kernel config from /proc/config.gz ...

Generally Necessary:
- cgroup hierarchy: properly mounted [/sys/fs/cgroup]
- CONFIG_NAMESPACES: enabled
- CONFIG_NET_NS: enabled
- CONFIG_PID_NS: enabled
- CONFIG_IPC_NS: enabled
- CONFIG_UTS_NS: enabled
- CONFIG_CGROUPS: enabled
- CONFIG_CGROUP_CPUACCT: enabled
- CONFIG_CGROUP_DEVICE: enabled
- CONFIG_CGROUP_FREEZER: enabled
- CONFIG_CGROUP_SCHED: enabled

........
.......

Moby test suite PASSED

			##         .
		  ## ## ##        ==
		   ## ## ## ## ##    ===
	   /"""""""""""""""""\___/ ===
	  ~~~ {~~ ~~~~ ~~~ ~~~~ ~~~ ~ /  ===- ~~~
	   \______ o           __/
		 \    \         __/
		  \____\_______/

[    3.578681] ACPI: Preparing to enter system sleep state S5
[    3.579063] reboot: Power down
```

The last log is the output of
[check-kernel-config.sh](https://github.com/linuxkit/linuxkit/blob/master/test/check/check-kernel-config.sh)
files.

If you are on linux you can do the same command but by the default you are going
to use [qemu](http://www.qemu-project.org/) an open source machine emulator.

```bash
sudo apt-get install qemu
```

I did some test in my Asus Zenbook with Ubuntu, when you run `moby run` this is
the command executed with qemu:

```
/usr/bin/qemu-system-x86_64 -device virtio-rng-pci -smp 1 -m 1024 -enable-kvm
	-machine q35,accel=kvm:tcg -kernel test-bzImage -initrd test-initrd.img -append
	console=ttyS0 -nographic
```

By default is testing on `x86_64` but qemu supports a lot of other archs and
devices. You can simulate an arm and a rasperry pi for example. At the
moment LinuxKit is not ready to emulate other architecture. But this is the main
scope for this project. It's just a problem of time. It will be able soon!

Detect if the build succeed or failed is not easy as you probably expect. The
status inside the VM is not the one that you get in your laptop. At the moment
to understand if the code in your PR is good or bad we are parsing the output:

```
define check_test_log
	@cat $1 |grep -q 'Moby test suite PASSED'
endef
```
[./linuxkit/Makefile](https://github.com/linuxkit/linuxkit/blob/master/Makefile)

Explain how linuxkit tests itself at the moment is the best way to get how it
works. It is just one piece of the puzzle, if you have a look here [every
pr](https://github.com/linuxkit/linuxkit/pulls) has a GitHub Status that point to
a website that contains logs related that particular build. That part is not
managed by linuxkit because it's only the builder used to create the
environment. All the rest is managed by
[datakit](https://github.com/docker/datakit). I will speak about it probably in
another blogpost.

## Conclusion

runc, docker, containerd, rkt but also Prometheus, InfluxDB, Telegraf a lot of
projects supports different architecture and they need to run on different
kernels with different configuration and capabilities. They need to run on your
laptop, in your IBM server and in a Raspberry Pi.

This project is in an early state but I understand why Docker needs something
similar and also, other projects as I said are probably going to get some
benefits from a solution like this one. Have it open source it's very good and
I am honored to be part of the amazing group that put this together. I just did
some final tests and I tried to understand how it's designed and how it works.
This is the result of my test. I hope that can be helpful to start in the right
mindset.

My plan is to create a configuration to test InfluxDB and play a bit with `qemu`
to test it on different architectures and devices. Stay around a blogpost will
come!

<p class="text-muted">
    Reviewers: <a href="https://twitter.com/justincormack">Justin Cormack</a>
</p>

<div class="post row">
  <div class="col-md-12">
      {% include docker-the-fundamentals.html %}
  </div>
</div>
