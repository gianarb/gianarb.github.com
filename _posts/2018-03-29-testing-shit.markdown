---
layout: post
title:  "I don't give a shit about testing"
date:   2018-03-29 10:38:27
img: /img/laziness.jpg
categories: [post]
tags: [testing, development, golang, container]
summary: "This is all about how I approach testing in development. TDD, DDD,
unit test, integration test. It should make my development faster and my code
easy to maintain. We have a lot of different techniques because we need to be
good on picking the right one."
priority: 1
---

That's what I learned during my experience as a developer. Doesn't matter
which languages you end up working if you are making HTTP API or things like
that you don't have an excuse. Write tests make your development faster, and it will
drastically improve maintainability of your project.

In this post, I would like to tell you how I approach testing particularly in
Go obviously.

First of all, when you create a new file, you should write its `_test.go` child.
It's hard to tell you who should be the child of who. Sometimes I start
to write everything inside a test function, just because run the actual test is faster
compared with compile, run the binary, trigger the right entry point and so on.
When I am satisfied, I move the
code to a function, and I leave the assertions I wrote as a new test. **Pretty good**.

> I don't give a shit about automated testing. I write tests.

I use [`vim-go`](https://github.com/fatih/vim-go) and `:GoTestFunc` is probably
the most used shortcut during my day to day job.

When I can choose I don't use assertion libraries, the `testing` package is
enough for me and dependency management in go is a pain, so fewer things I vendor
better I feel about myself.

<img src="/img/laziness.jpg" class="img-responsive" />

I use fixtures, but I don't like them. I prefer to write more small tests than complicated fixtures.

A single test for me is more descriptive, and I don't mind to write redundant code, I can always refactor it later or move it some helper function. A complicated fixture will be hard to maintain.
The name of the function is an excellent way to describe what you are covering in your test and the function itself
creates a beautiful block that improves readability.

```go
func TestCarComposition(t *testing.T) {
    fixtures := []car.Composition{
        {"blue", "europe", 1, nil, "2011-12-05", "ford"},
        {"", "", 35, bool, nil, "ford"},
        {"red", "usa", 0, bool, nil, "fiat"},
        {"white", "", 35, bool, nil, "kia"},
        {"orange", "", 1, bool, "2010-05-12", nil},
        {"", "", 0, bool, nil, "ford"},
    }
}
```

Bonus point, as you can see fixtures are sad to read!

Even unit vs. integration vs. function is a very annoying discussion. Don't tell me
about TDD, BDD, CCC, DDD things. I don't care they are all amazing as
soon as they can make my development simple.

So, CDD is probably my best test methodology: **Comfort driven development**.

Usually when I am writing a computing function when it elaborates maps, strings,
files without using too many external resources I start from unit test, because
it makes iteration faster as I told before. And it won't require too many mocks.
I don't like mocks.


## Let's discuss mocks
Mocks are a pain; you end up to be bored when you write mocks, they won't fail when it's useful for you to see an error and they will fail when you don't care.
So comfort looks very far from mocks!

When mocks becomes too complicated, and I can write another kind of tests I go with that solution. Maybe integration or I will try to write the simple mock
possible, sometimes even the entire web server can be a valuable solution:


```go
func TestInfluxDBSdkGetTheRightValues(ti *testing.T) {
    ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
        data := influxdb.Response{
            Results: []influxdb.Result{
                {
                    Series: []influxdbModels.Row{},
                },
            },
        }
        w.Header().Add("Content-Type", "application/json")
        w.WriteHeader(http.StatusOK)
        _ = json.NewEncoder(w).Encode(data)
    }))
    defer ts.Close()

    config := influxdb.HTTPConfig{Addr: ts.URL}
    client, _ := influxdb.NewHTTPClient(config)
    defer client.Close()

    // Whatever you need to check
}
```
You need to play carefully; these tests are slower and more expensive in
resources.
But I like the idea to take the faster solution when I am developing; you can
come back on your tests later when the feature is more stable and better
designed. Write tests should not slow me down too much, I am looking for a way
to write the implementation and the test fasts to iterate on both of them other than waste time making everything perfect. Nothing will be forever; nothing will be complete in programming, so design your environment to be easy to change.

## Integration tests
I am a CLI kind of person, so I often send HTTP requests via cURL.
Docker is very easy from day one to start and stop your application,
cleaning databases and so on.

[`bats`](https://github.com/sstephenson/bats) combines these two sentences. It is an automation test framework for
bash. It is very simple to setup, and it allows me to copy paste some cURL, and
with jq, you can make the checks you need over your JSON response.

An integration test suite made with bats looks like that:

1. An "init" file in bash where you can run setup and teardown function before and after every test. Usually, you can you that functions to spin up and down the
   containers that you need to tests, this is the one that I wrote for this
   example

```bash
#!/bin/bash

function setup() {
  teardownCallback=$(init)
}

function teardown() {
  eval $teardownCallback
}

function getHost() {
  echo "http://localhost"
}

function init {
  executionID=$(cat /dev/urandom | tr -dc 'a-zA-Z0-9' | fold -w 7 | head -n 1)
  containerLabels="exec=${executionID}"
  $(docker run -d -l $containerLabels -p 80:80 nginx)
  echo "docker ps -aq -f 'label=${containerLabels}' | xargs docker rm -f"
}
```
2. You have a set of `.bats` files with the various scenarios, I wrote one to
   check if the status code 200 for the nginx home

```bash
#!/usr/bin/env bats

load utils

@test "Nginx home return 200" {
      statusCode=$(curl -I -X GET "$(getHost)" 2>/dev/null | head -n 1 | cut
      -d$' ' -f2)
        [ $statusCode -eq 200 ]
}
```

What you are running is a `bats` test to check that `nginx:latest` is serving the right page.
Your use case will be ten times more complicated.

Another reason to take this approach is about bash itself. If you are not a bash
expert, you will probably end up to write straightforward tests, cURL, grep,
regex and some pipes. Nothing more.

And you won't use any code that runs your application. It's important to avoid
weird buggy tests.

## developer happiness
Tests are a methodology to decrease the cost of maintenance and to improve your
ability to write code.

It should not be a fashion way to show how good you are as a developer. You will
be a good developer as a side effect.

I look at all the different way to tests my code as a tool set, AI is becoming
very smart. So we need to be less "server" and more human been. 100% coverage
for unit tests looks a lot like something that a server can do. Pick the right
method based on your feeling.

<script>
$(document).ready(function() {
	$('body').css("background", "#F5F3E6");
});
</script>
