---
layout: post
title:  "The Go awesomeness"
date:   2018-04-09 10:38:27
img: /img/fight-club.jpg
categories: [post]
tags: [golang, awesome, list, development, pprof, go, google, delve]
summary: "After 1 year writing go every day at work this is why I like to work
with it."
priority: 1
---
It's one year since I started to use Go every day at work. I was using it before
but for fun or OSS projects. I was looking for my next challenge, and I was
mainly working with PHP, JavaScript previously and I knew that a compiled,
statically typed language was my next step.

At my previous job at CurrencyFair, the environment was pretty standard for a
financial tech company so backend in Java, frontend in PHP. But my experience
with all the interfaces and abstract classes that I created in Java at that time
made me hate that language. So I was looking for something different.

I was as I am now involved in automation, cloud and
operational other than development so all the tools like Docker, InfluxDB,
Kubernetes, Consul, Vault was in golang and for me as OSS addicted it become the
natural choice.
Now after all this time I am ready to write why I think Go is the right choice
for me now.

## 1. abstraction and maintainability
I wrote a lot about [this topic](/blog/the-abstract-manifesto), so I am not going to repeat myself. But I
think maintainability is tied together with abstraction. Previously when I was
working with PHP, we always had services, injection and so on. In that
environment it was good, but all that abstraction like in Java doesn't make your
code more flexible. It makes it hard to understand in the long run and code
needs to be written with history in mind because delete code is very hard.
Go with its interface implementation, how it forces you to struct the project
helps the codebase to grow in a better way.

## 2. Stdlib
Community wasted time across languages to identify the right way to indent code.
Go comes with that decision done. Same for testing. How to write automatic
tests, benchmarks is inside the language. No libraries, it is there.
More in general os, net, net.http, img and so on, a lot of stuff are provided by
the language itself. It is great because you don't need anything to start,
other than Go. Compared with other languages you can do a lot more things.
Having all this feature inside Go guarantees compatibility over time, they won't
break compatibility for the next years, and the code is developed and reviewed
by a large number of people.

## 3. pprof
pprof is a profiler, and it is shipped as part of Go. You can use it even via
cli, or it also has an excellent HTTP package under
[net/HTTP/pprof](https://golang.org/pkg/net/http/pprof/).
Just to show you how much power it can be InfluxDB extends it to export a zip
archive with all the information we need to troubleshoot the database behavior:

```go
func (h *Handler) handleProfiles(w http.ResponseWriter, r *http.Request) {
    switch r.URL.Path {
        case "/debug/pprof/cmdline":
            httppprof.Cmdline(w, r)
        case "/debug/pprof/profile":
            httppprof.Profile(w, r)
        case "/debug/pprof/symbol":
            httppprof.Symbol(w, r)
        case "/debug/pprof/all":
            h.archiveProfilesAndQueries(w, r)
        default:
            httppprof.Index(w, r)
    }
}
```
Here all the code
[influxdata/influxdb](https://github.com/influxdata/influxdb/blob/442581d299b7d642e073bbe42112fa9b58fb071a/services/httpd/pprof.go#L21).
This is super useful because we can ask customers or developer in the OSS
community to export and upload the archive to see what is going on.
Having a standard way to troubleshoot and export a profile allows us to build
visualization or static analysis on it for common calculation.

## 4. delve
A good debugging session is the best way to approach a new application or to go
deeper learning a language or a software.
[delve](https://github.com/derekparker/delve) is easy to setup and to use. Even
if you are not gdb/debugger superhero as I am not you will be able to make your
    first steps with delve. So it is a nice starting point too.

## 5. godoc
Other than behind an excellent way to generate documentation from source code I
use it a lot even when I am not designed libraries just to double check that my
package has the comprehensive public methods. I always think about what I am
exposing to the outside when I write code. APIs are not just JSON or HTTP thing,
every object exposes their API, and you need to be aware of how you are building
iteration between the internal state and the outside. Avoid misuse of your
structs is your responsibility as developer and godoc help me to identify poor
decision.

## 6. vim-go
I would like to stay in my terminal all day, and vim-go allows me to write good
code in my comfort zone. In the past I wrote a lot of vim scripts and plugins,
following how fatih and all the other maintainers are developing
[vim-go](https://github.com/fatih/vim-go) is great.
Bonus point they recently added support for delve, so you can now debug golang
application in vim!

## 7. dep
Dependency management is probably the worst things that Go has. The good thing
is that now we have [dep] and it should become the standard way to manage
dependencies. Right now the situation looks a lot like this:

<img class="img-responsive" src="/img/fight-club.jpg">

Govender, go get, glide currently there are a lot of different ways to manage
dependencies, and it generates a lot of confusion, but I hope at the end we will
converge in just one. Probably dep.

## Conclusion

More in general with Go I am learning that the language is one of expect to
become a good developer. A good developer needs to know the language, but the
best way to go deeper in it is writing tests, benchmarks, profiling application
and using the debugger. All these tools make my life as developer easy. Easy
life for me means that I can go deeper solving problems and indirectly it will
make me a better developer.

Go is fun!
