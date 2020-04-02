---
layout: post
title:  "Checklist for a new project"
date:   2020-04-02 09:08:27
categories: [post]
img: /img/myselfie.jpg-large
tags: [developer, checklist, ci, cd, logging]
summary: "A person list I developed along those years that I try to implement
across projects I start or contribute to"
changefreq: daily
---
Back in the day I used to start a lot of projects. From zero on GitHub, some of
them are still there, unused probably.

Recently I started to take part of other people projects like
[testcontainers](https://github.com/testcontainers) or
[profefe](https://github.com/profefe). I wrote about why I do it during the
["2019 year in review"](/blog/year-in-review) post.

In both cases, so even when joining a new existing project or when I start a new
one I try to follow a checklist.

I developed this checklist along the years, moving parts and extending the
number of checks. The main goal for that is to validate that the project has
good answers for a couple of questions, not related at what it does, but to how
it does it.

1. is it easy to onboard as a user?
2. as a new contributor is the project easy to understand?
3. as a maintainer do I have everything I can under control in order to waste as
   less time as possible?

I follow the checklist when working on opensource but also in close source
project and what I like about it is that you can propose a change by youself,
you can try to apply those feedback as a solo-developer, hoping to make
contributors, maintainers and colleagues to buy them spreading joy.

But let's make to the list now.

### Have a place where you can write

When I start a new project, but also during the onboard of an existing one in my
tool chain I look for a written format of it.

I look for a readme, an installation process, a getting started guide, a
contribution document. It does not need to be pretty one, a copy/paste of a few
bash scripts that the maintainer does to set itself up is enough.

Having a place during the early days of a project where I can write what I
think, how I would like to get things done is important in order to design
something usable and to spot sooner misleading assumption.

If you can build the place for all those information it will take you one
second to save them forever, it is just a matter of copy/pasting the command you
run in your terminal to spin up dependencies, build the project and so on.

I like to use the README.md, CONTRIBUTOR.md and a `./docs` folder to save
everything I am thinking about or everything I do that I hope will make my life
easy in a month where I will be back on that piece of code without even knowing
it was there. The feeling you get is the same a new person has when it looks at
your project for the first time.

There is no way you can get it right since the beginning, because there is not a
definition of right. First day everything you write is mainly for yourself, in a
month and some editing it will become the first version of the documentation for
your project.

### Logging and instrumentation library

As I said at the beginning of the article all the checks do not depend on the
business logic of your application or library. All of them has to speak with the
outside world sharing their internal state in a way that is reusable,
comprehensive, configurable.

There are a lot of people that speaks about observability, logging, tracing,
monitoring. Everybody has its own opinion, but form a technical point of your
what you write has to be easy to troubleshoot.

You do it using the right telemetry libraries. For logging I do not have any
doubt. In Go I use [zap](https://github.com/uber-go/zap).

During a workshop about observability I built where I had to instrument 4
applications on different languages I selected:

* [pino](https://github.com/pinojs/pino) for NodeJS
* [monolog](https://github.com/Seldaek/monolog) for PHP
* [log4j](https://logging.apache.org/log4j/2.x/) for JAVA

In general I look for libraries that allow me to do structured logging, so for
the one that enables me to attach key value pairs to log line. I also look for
logging libraries that has the concept of exporters and format. Nothing unusual.

For tracing and events I do not have a favourite one but I would like to see
[Opentelemetry](https://opentelemetry.io) to become the way to go.

### Continuous integration

A project without CI can not be called in that way. Nowadays there are a lot of
free services that you can use, so no excuse. When I am on GitHub I go for
Actions now because they are free and embedded in the VCS itself.

If you didn't write any test at least set the process up and running. Just run
tests, usually they do not fails if empty. And there are static checker, linters
and things like that for every language, set them up!

### Continuous delivery

You made the CI part, you are half way done. Release is important and we have
the tools to get it right since day one. It is a pain to do a release, there are
a lot of potential manual state to get it right:

1. Bump version
2. Changelog
3. Compile and push binaries if it is an application
4. ...

There are tools that helps you to do that in automation. For my apps I use
[goreleaser](https://github.com/goreleaser/goreleaser), for the libraries I use
[Releaser drafter](https://github.com/marketplace/actions/release-drafter).

### Testing framework

Write tests, and when you see repeated code extract it in a testing package.
`zap` has `zaptest`, your project should have `yourprojecttest` as well.

It is useful for yourself, because it will make the job to write more test
effortless, and if you well document your testing package contributors will be
able to use it when opening a PR because you will make writing tests easier for
everybody. As a bonus who ever uses your libraries can use the testing package
to write their own tests for their application.

## Conclusion

This is the list I use, and I will keep it up to date now that I wrote it down,
adding or editing it, so be sure to stay around!

I hope this checklist is general enough and useful to be reusable for you in
some of its part.

What I like about it is that I do not need to be a CTO, a maintainer or
something like that to drive the adoption of those point that I think are
crucial, I drove the adoption of some of them even as a solo contributor.
