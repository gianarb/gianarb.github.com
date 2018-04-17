---
layout: post
title:  "Go testing tricks"
date:   2018-04-17 10:38:27
img: /img/gopher.png
categories: [post]
tags: [golang]
summary: "This post contains some feedback about how to write tests in Go."
priority: 1
---
I recently wrote a blog post with my [point of view about
testing](/blog/testing-shit). I used Go as the language to concretize it. I had
good feedback about that article, and this is all about how I write tests in
Go.

## Fixtures
I wrote that I don't like them, but I think they are useful. You can use them
to verify the same function checking the same assertion with different input.
So let's say you are testing a function that returns the multiplication of two
numbers if the first is even, it returns the division if not.

I will write two tests, one to test events number and one to test the other
case, and I will set up two fixtures one for every test. I won't write just one
test with elaborate fixtures because they are hard to read and the name of the
test function will help a lot to understand the assertion. Small example good
for blogging purpose. But I hope you got the idea.

```golang
package test

import "testing"

func MagicFunction(f int, s int) int {
    if f%2 == 0 {
        return f * s
    }
    return f / s
}

func TestEventInputsShouldReturnMoltiplication(t *testing.T) {
    table := []struct {
        first  int
        second int
        result int
    }{
        {2, 1, 2},
        {4, 10, 40},
    }
    for _, s := range table {
        if r := MagicFunction(s.first, s.second); r != s.result {
            t.Errorf("Got %d, expected %d. They should be the same.", r, s.result)
        }
    }
}

func TestOddInputsShouldReturnDivision(t *testing.T) {
    table := []struct {
        first  int
        second int
        result int
    }{
        {15, 3, 5},
        {21, 7, 3},
    }
    for _, s := range table {
        if r := MagicFunction(s.first, s.second); r != s.result {
            t.Errorf("Got %d, expected %d. They should be the same.", r, s.result)
        }
    }
}
```

## sub-test

To make the fixtures a bit better I use the `t.Run` function a lot. It is a
feature introduced in Go 1.9 as part of the `testing` package.

```go
package test

import (
    "fmt"
    "testing"
)

func MagicFunction(f int, s int) int {
    if f%2 == 0 {
        return f * s
    }
    return f / s
}

func TestEventInputsShouldReturnMoltiplication(t *testing.T) {
    table := []struct {
        first  int
        second int
        result int
    }{
        {2, 1, 2},
        {4, 10, 40},
    }
    for _, s := range table {
        t.Run(fmt.Sprintf("%d * %d", s.first, s.second), func(t *testing.T) {
            if r := MagicFunction(s.first, s.second); r != s.result {
                t.Errorf("Got %d, expected %d. They should be the same.", r, s.result)
            }
        })
    }
}
```

`vim-go` has an option `let g:go_test_show_name=1` to allow the name of the
test as part of the output for :GoTest or :GoTestFunc. This helps a lot to
enjoy this feature.

## Golden files

Golden files are something used in different packages in the Go standard
library, and Michael Hashimoto spoke about it during his brilliant talk about
testing at the [GopherCon 2017](https://www.youtube.com/watch?v=8hQG7QlcLBk).
In case of complex output, you can verify the result of the tests with the
content of a file. It improves order and readability.  When you declare a
global flag in your test, it becomes available inside `go test` so if you use
the update flags all the tests will pass, but you will update all the golden
files. So this is very useful if you need to compare a lot of bytes.

```go
update := flag.Bool("update-golden-files", false, "Update golden files.")
```

```sh
go test -update-golden-files
```
I was using this trick a lot when I was writing PHP code, and I was testing HTTP responses.

## Test helper and return function
When you have repeatable code across tests, you can create a helper function,
and you can use it in your tests. There are two general rules about this
approach:

1. The helper function should have access to *t testing.T variable.
2. Your helper never returns an error; it marks the test as failed. That's why
it needs access to `*t. testing.T`.

Another good trick is to return a function from the helper to clean up what you
did in the helper. So let's say that your helper starts an HTTP server. You can
return the HTTP Close function as a callback.

```go
func () testHelperStartHTTPServer(t *testing.T) func() {
    ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
        // long and complex mock maybe with a golden file and so on
    }))
    return func() { ts.Close() }
}

func TestYourTest(t *testing.T) {
    hclose := testhelperStartHTTPServer()
    // All your logic and checks
    defer hclose.Close()
}
```
I used the same practice when I was writing integration tests using bash and
[bats](https://github.com/sstephenson/bats). It is a very clean and easy to
read approach.

## parallel
You can use the function `t.Parallel()` to notify at the test runner that your
case can run in parallel with other tests marked as parallel.  When you write
unit tests, you can almost always run them in parallel because they should be
completely isolated.

## Short and verbose
`-short` and `-v` are two flags available when you run `go test`. You can use
them in your tests:

```
import "testing"

func TestVeryLongAndExpensiveCapability(t *testing.T) {
  if testing.Short() {
    t.Skip("skipping this testsVeryLongAndExpensive is too expensive")
  }
  // ... other code
}
```
`-short` describes itself pretty well, you can skip tests that are too long and expensive.

`-v` allows you to print more:
```
import "testing"
func TestVeryLongAndExpensiveCapability(t *testing.T) {
  if testing.Verbose() {
  }
  // ... other code
}
```

## testing/quick
[testing/quick](https://golang.org/pkg/testing/quick/) is a nice package that
offers a set of utilities to write test quick. Go has not an assertion library
inside the stdlib but this can help if you are like me and you are happy to not
vendor assertion libraries because `if { }` with some sugar is what I need.

So that's it, have fun and write tests!
