---
img: /img/golang-mockmania.png
layout: post
title: "How to test CLI commands made with Go and Cobra"
date: 2020-03-09 09:08:27
categories: [post]
tags: [golang, mockmania, cobra, cli]
summary: "CLI commands are common in Go. Testing them is an effective way to run
a big amount of code that is actually very close to the end user. I use Cobra,
pflags and Viper and that's what I do when I write unit test for Cobra commands"
changefreq: daily
---
Almost everything is a CLI application when writing Go. At least for me. Even
when I write an HTTP daemon I still have to design a UX for configuration
injection, environment variables, flags and things like that.

The set of libraries I use is very standard, I use
[Cobra](https://github.com/spf13/cobra),
[pflags](https://github.com/spf13/pflag) and occasionally
[Viper](https://github.com/spf13/viper). I can say, without a doubt that [Steve
Francia](https://twitter.com/spf13) is awesome!

This is how a command looks like directly from the Cobra documentation:

```
var rootCmd = &cobra.Command{
  Use:   "hugo",
  Short: "Hugo is a very fast static site generator",
  Long: `A Fast and Flexible Static Site Generator built with
                love by spf13 and friends in Go.
                Complete documentation is available at http://hugo.spf13.com`,
  Run: func(cmd *cobra.Command, args []string) {
    // Do Stuff Here
  },
}
```

I like to write a constructor function that returns a command, in this case it
will be something like:

```
func NewRootCmd() *cobra.Command {
    return &cobra.Command{
      Use:   "hugo",
      Short: "Hugo is a very fast static site generator",
      Long: `A Fast and Flexible Static Site Generator built with
                love by spf13 and friends in Go.
                Complete documentation is available at http://hugo.spf13.com`,
      Run: func(cmd *cobra.Command, args []string) {
        // Do Stuff Here
      },
  }
}
```

The reason why I like the have this function is because it helps me to clearly
see the dependency my command requires. In this case nothing. I also like to use
not the Run function but the RunE one, it works in the same way but it expects
an error in return.

```
func NewRootCmd(in string) *cobra.Command {
    return &cobra.Command{
      Use:   "hugo",
      Short: "Hugo is a very fast static site generator",
      Long: `A Fast and Flexible Static Site Generator built with
                love by spf13 and friends in Go.
                Complete documentation is available at http://hugo.spf13.com`,
      RunE: func(cmd *cobra.Command, args []string) (error) {
          fmt.Fprintf(cmd.OutOrStdout(), in)
          return nil
      },
  }
}
```

In order to execute the command, I use cmd.Execute().

Let’s write a test function:

The output with `go test -v` contains “hi” because by default cobra prints to
stdout, but we can replace it to assert that automatically

```
func Test_ExecuteCommand(t *testing.T) {
	cmd := NewRootCmd("hi")
	cmd.Execute()
}
```

```
=== RUN   Test_ExecuteCommand
hi--- PASS: Test_ExecuteCommand (0.00s)
PASS
ok      ciao    0.006s
```

The trick here is to replace the stdout with something that we can read
programmatically like a bytes.Buffer for example:

```go
func Test_ExecuteCommand(t *testing.T) {
	cmd := NewRootCmd("hi")
	b := bytes.NewBufferString("")
	cmd.SetOut(b)
	cmd.Execute()
	out, err := ioutil.ReadAll(b)
	if err != nil {
		t.Fatal(err)
	}
	if string(out) != "hi" {
		t.Fatalf("expected \"%s\" got \"%s\"", "hi", string(out))
	}
}
```

Personally I do not think there is much more to know in order to effectively
test CLI commands, they can be very complex, but if you can mock its
dependencies and check what the execution prints out you are very flexible! 

Another thing you have to control when running a command is its arguments and
its flags because based on them you will get different behavior that you have to
test in order to figure out that your commands work with all of them.

The logic works the same for both but arguments are very easy, you just have to
set the argument in the command with the function
`cmd.SetArgs([]string{"hello-by-args"}).`

```go
package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"testing"

	"github.com/spf13/cobra"
)

func NewRootCmd() *cobra.Command {
	return &cobra.Command{
		Use:   "hugo",
		Short: "Hugo is a very fast static site generator",
		RunE: func(cmd *cobra.Command, args []string) error {
			fmt.Fprintf(cmd.OutOrStdout(), args[0])
			return nil
		},
	}
}

func Test_ExecuteCommand(t *testing.T) {
	cmd := NewRootCmd()
	b := bytes.NewBufferString("")
	cmd.SetOut(b)
	cmd.SetArgs([]string{"hi-via-args"})
	cmd.Execute()
	out, err := ioutil.ReadAll(b)
	if err != nil {
		t.Fatal(err)
	}
	if string(out) != "hi-via-args" {
		t.Fatalf("expected \"%s\" got \"%s\"", "hi-via-args", string(out))
	}
}
```


Flags works in the same:


```go
package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"testing"

	"github.com/spf13/cobra"
)

var in string

func NewRootCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "hugo",
		Short: "Hugo is a very fast static site generator",
		RunE: func(cmd *cobra.Command, args []string) error {
			fmt.Fprintf(cmd.OutOrStdout(), in)
			return nil
		},
	}
	cmd.Flags().StringVar(&in, "in", "", "This is a very important input.")
	return cmd
}

func Test_ExecuteCommand(t *testing.T) {
	cmd := NewRootCmd()
	b := bytes.NewBufferString("")
	cmd.SetOut(b)
	cmd.SetArgs([]string{"--in", "testisawesome"})
	cmd.Execute()
	out, err := ioutil.ReadAll(b)
	if err != nil {
		t.Fatal(err)
	}
	if string(out) != "testisawesome" {
		t.Fatalf("expected \"%s\" got \"%s\"", "testisawesome", string(out))
	}
}
```

This is it! I like a lot to write unit tests for cli command because in real
life they are way more complex than the one I used here. It means that they run
a lot more functions but the command is well scoped in terms of dependencies (if
you write a constructor function) and in terms of input and output. So it is
easy to write an assertion and write table tests with different inputs.
