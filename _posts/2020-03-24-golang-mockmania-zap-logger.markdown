---
img: /img/golang-mockmania.png
layout: post
title: "How to do testing with zap a popular logging library for Golang"
date: 2020-03-24 09:08:27
categories: [post]
tags: [golang, mockmania, logging, zap]
summary: "How you can use logging to build assertions when testing. What the
popular Golang logging library provided by Uber gives you around unit tests."
changefreq: daily
---

<div class="alert alert-dark" role="alert">
   <div class="row">
       <div class="col-md-2 align-self-center">
          <a href="https://link.testproject.io/0ak" target="_blank">
              <img class="img-fluid" src="/img/testproject-logo-small.png">
          </a>
       </div>
       <div class="col-md-8 text-center">
           <a href="https://link.testproject.io/0ak"
           class="alert-link" target="_blank">TestProject</a> is a community all
           about testing and you know how much I love communities! Join us.
       </div>
   </div>
</div>


If you follow me on [twitter](https://twitter.com/gianarb) you know that
I am passionate about o11y, monitoring and code instrumentation.

I see logs not as a random print statement that you use only when something is
wrong, but they have value. Logs are the communication channel our applications
use. As developers it is our job to make them to speak in a comprehensive way.

Logs should be structured and in some way consistent across functions, http
handlers, applications even languages to simplify their use. From algorithms and
human operators.

In Go [zap](https://github.com/uber-go/zap) is a popular logging library provided by Uber, I
use it almost by default for all my applications.

```go
package main

import "go.uber.org/zap"

func main() {
	logger, _ := zap.NewProduction()
	do(logger)
}

func do(logger *zap.Logger) {
	logger.Error("Start doing things")
}
```

So logging and testing? In the same article? I should be really drunk!

{:refdef:.text-center}
![](/img/kermit-frog-drunk.jpg){:.img-fluid}
{:refdef}

When I discovered that `zap` comes with a testing utility package called
`zaptest` I felt in love with this library even more:

```go
package main

import (
	"testing"

	"go.uber.org/zap/zaptest"
)

func Test_do(t *testing.T) {
	logger := zaptest.NewLogger(t)
    do(logger)
}
```

The `go test` command supports the flag `-v` to improve verbosity of test
execution. In practice that's how you forward to `stdout` logs and print
statements during a test execution. `zaptest` works with that as well.

Very cool, and useful if you write smoke tests, pipeline tests, or how ever you
call them and see the logs can be spammy, but helpful to figure out the actual
issue.

```go
package main

import (
	"testing"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"go.uber.org/zap/zaptest"
)

func Test_do(t *testing.T) {
	logger := zaptest.NewLogger(t, zaptest.WrapOptions(zap.Hooks(func(e zapcore.Entry) error {
		if e.Level == zap.ErrorLevel {
			t.Fatal("Error should never happen!")
		}
		return nil
	})))
	do(logger)
}
```
You can use `hooks` to check for expected or unexpected logs.

Hook are executed for every log line:

```go
func(e zapcore.Entry) error {
    if e.Level == zap.ErrorLevel {
        t.Fatal("Error should never happen!")
    }
    return nil
})
```

If you do not expect any error level log line for your execution because you are
testing the happy path, you can do something like that.

**DISCALIMER:** This is another way to write assertion. You will may use them to enforce
other checks, or to validate the workflow from a different point of view that
will may be easier to do as first attempt. As I usually say: "an easy and
partial test is better than no test".

Do not test only logs, it won't age well! Keep writing good tests!
