---
img: /img/golang-mockmania.png
layout: post
title: "The awesomeness of the httptest package in Go"
date: 2020-02-25 09:08:27
categories: [post]
tags: [golang, mockmania, http, httptest]
summary: "One of the reasons why testing in Go is friendly is driven by the fact
that the core team already provides useful testing package as part of the stdlib
that you can use, as they do to test packages that depend on them. This article
explains how to use the httptest package to mock HTTP servers and to test sdks
that use the http.Client."
changefreq: daily
---

Go has a nice http package. I am able to say that because I am not aware of any
other implementation of it in Go other than the one provided by the standard
library. This is for me a good sign.

```go
resp, err := http.Get("http://example.com/")
if err != nil {
	// handle error
}
defer resp.Body.Close()
body, err := ioutil.ReadAll(resp.Body)
```

This example comes from the [documentation](https://golang.org/pkg/net/http/)
itself.

We are here to read about testing, so who cares about the http package itself!
What matters is the [httptest](https://golang.org/pkg/net/http/httptest/)
package! Way cooler.

This article is not the first one for the MockMania series, I wrote about titled
[“InfluxDB Client
v2”](https://gianarb.it/blog/golang-mockmania-influxdb-v2-client), it uses the
httptest service already! But hey it deserves its own blog post.

## Server Side

The http package provides a client and a server. The server is made of handlers.
The handler takes a request and based on that it returns a response. This is its
interface:

```go
type Handler interface {
    ServeHTTP(ResponseWriter, *Request)
}
```

As you can see if gets a ResponseWriter to compose a response based on the
Request it gets. This process can be as complicated as you like, it can reaches
databases, third party services but in the end, it writes a response.

It means that mocking all the dependencies to get the right scenario we use the
ResponseWriter to figure out if the handler made what we want.

The httptest package provides a replacement for the ResponseWriter called
ResponseRecorder. We can pass it to the handler and check how it looks like
after its execution:

```go
handler := func(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "ping")
}

req := httptest.NewRequest("GET", "http://example.com/foo", nil)
w := httptest.NewRecorder()
handler(w, req)

resp := w.Result()
body, _ := ioutil.ReadAll(resp.Body)

fmt.Println(resp.StatusCode)
fmt.Println(string(body))
```

This handler is very simple, it just manipulates the response body. If your
handler is more complicated and it has dependencies you have to be sure to
replace them as well, injecting the appropriate one.

## Client-Side

Handlers are useful if you can’t use them. The Go http package provides an http
client as well that you can use to interact with an http server. An http client
by itself is useless, but it is the entry point for all the manipulation and
transformation you do on the information you get via HTTP. With the
proliferation of microservices, it is a very common situation.

The workflow is well understood, you have an HTTP backend to interact with, you
fetch data from there are you manipulate them with your business logic. When
testing what you can do is to mock the http backend in order to return what you
want, testing that your business logic does what it is supposed to do based on
the input you get from the HTTP server.

During our first example, the handler was the subject of our testing, this is
not the case anymore, we are testing the consumer this time, so we have to mimic
and handler in order to get what we expect to return

```go
ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "I am a super server")
}))
defer ts.Close()
```

As you can see we are creating a new HTTP server via the httptest. It accepts a
handler. The goal for this handler returns what we would like to gest our code
on. In theory, it should just use the ResponseWriter to compose the response we
expect.

The server has a bunch of methods, the one you are looking for is the URL one.
Because we can pass it to an http.Client, the one we will use as a mock for our
function

```go
res, err := http.Get(ts.URL)
if err != nil {
	log.Fatal(err)
}
bb, err := ioutil.ReadAll(res.Body)
res.Body.Close()
```

That’s it, as you can see `ts.URL` points the http.Client to the mock server we
created.

## Conclusion

I use the httptest package a lot even when writing SDKs for services that do not
have integration with Go because I can follow their documentation mocking their
server and I do not need to reach them until I am confident with the code I
wrote.

My suggestion is to test your client code for edge cases as well because of the
httptest.Server gives you the flexibility to write any response you can think
about. You can mimic an authorized response to seeing how your code with handle
it, or an empty body or a rate limit. The only limit is our laziness.
