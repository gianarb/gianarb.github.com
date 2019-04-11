---
img: /img/gianarb.png
layout: post
title:  "Workshop Design"
date:   2019-04-19 08:08:27
categories: [post]
tags: [go, monitoring, observability, workshop, development, conference,
tracing, jaeger]
summary: "I recently developed a workshop about application instrumentation. I
ran it at the CloudConf in Turin. I developed it in open source and I thought it
was a nice idea to share more about why I did it and how."
changefreq: daily
---

Hello sweet internet! At this point, you should know that I am far to be a
lovely happy code. I like to share what I learn and to have a chat about what
you are doing. That's as it is! Feel free to follow my wheeze on
[twitter/gianarb](https://twitter.com/gianarb).

If you don't know what I am gonna speak about, I can tell you this is another
way to enjoy coding!

Recently a friend of mine that organizes the [CloudConf](https://cloudconf.it)
in Turin, Italy asked me if I was able to deliver a workshop. Let me say THE
workshop, 8 hours of chatting with exercises and questions.
I did something like that years ago about AngularJS but hey, this sounds like a
challenge, and I love challenges! So I took it.
![](/img/got-your-back.jpg)
If you read my recent posts you know I have a passion nowadays:

* [Observability is for troubleshooting](/blog/go-observability-is-for-troubleshooting)
* [You need an high cardinality database](/blog/high-cardinality-database)
* [Logs, metrics and traces are equally useless](/blog/logs-metrics-traces-aggregation)

The topic was clear, I have called it "Application instrumentation". Lovely!

I am driven by passion and purpose. My passion for troubleshooting and the
purpose of figuring out what the f happen in production. I was ready to work on
it!

![](/img/passion-fruit.jpg)

## Workshop?

This article is about how I prepared the workshop and I hope it can help
somebody to avoid the same mistake and also to use some of the material I
developed.

I made everything in open source. There are two new repositories on my GitHub one with fake
e-commerce I made using 4 different programming languages:

* Golang as frontend proxy with a UI in HTTP/JQuery.
* Java to do the most secure part of the e-commerce obviously the payment
  service.
* NodeJS to get discounts for the items.
* PHP to get the list of items currently available.

You can find the code on
[github.com/gianarb/shopmany](https://github.com/gianarb/shopmany).

I decided to develop a minimum version of the application in order to have it
reusable for another purpose. It can be used to build a use case for Kubernetes
deployment for example or a CI lesson.

The branch `master` contains the minimum set of features that I need to have an
application that has some sense. But for example, the services are without logs,
metrics and tracing because they will be added as exercises from the attendees.

If you check out the workshop you will be able to see in the history a commit for
every exercise and applications.

The lessons are available on
[github.com/gianarb/workshop-observability](https://github.com/gianarb/workshop-observability),
every directory is a lesson. The readme contains a couple of information about
what where we are, why we should care and one or more exercise to do in practice
in order to familiarize with the concepts.

The lessons I developed for the purpose of the CloudConf workshop are:

1. lesson1 designing a health check endpoint. Adding a single endpoint is a good
   way to familiarize with a new application and there is so much to learn about
   how to design a good health check endpoint!
2. lesson2 is about logging and [structure
   logging](https://charity.wtf/2019/02/05/logs-vs-structured-events/). I tried
   to pick the most popular logging libraries for the languages. Logging using
   JSON format to open the door for future serialization as an event.
3. lesson3 is about InfluxDB v1 and the TICK stack. The goal was to serve a
   monitoring stack that can work with a different structure such as events and
   traces.
4. lesson4 is about tracing. Using Jaeger we instrumented and build a trace for
   the application.

I have also reported an idea of a possible timeline (the one I used at the
CloudConf):

09.00 Registration and presentation
09.30 - 13.00 Theory

* Observability vs monitoring
* Logs, events, and traces
* How a monitoring infrastructure looks like: InfluxDB, Prometheus, Jaeger,
  Zipkin, Kapacitor, Telegraf...
* Deep dive on InfluxDB and the TICK Stack
* Deep dive on Distributed Tracing

13.00 - 14.00 Launch
14.00 - 17.00 Let's make our hands dirty
17.30 - 18.00 Recap, questions and so on

## Learning during the development

I like to prepare slides, posts, and workshop because I learn a lot along the
way about concepts that I usually develop during a long set and frustrating
attempts. Or reading a lot of blog posts, books, code. Writing about it helps me
to put together what I learned developing easy to understand materials.

This workshop was not a special case. It is not clear for me that even if there
is a lot going on with OpenCensus, OpenTracing, and other instrumentation
libraries there is still room for improvement.

Instrumenting an application is not anymore just a matter of adding `printf`
around the execution of the code. But it is the way we have to write an
application capable of behind debugged and that speaks with the outside in an
understandable way.

The course has two different sections: theory and practice.

The theory went well. I do not have a lot to say about it and for me, it is where
I am most comfortable with because it looks like a long talk.

The practical part was for sure a bit too long and I didn't have time to walk
all the people over it but the fact that there are all the solutions, the
purpose written down helped them to feel less lonely and everyone can follow the
resolution is it can do the exercise in practice.

This usually happens because of different skills set or for trouble configuring
the environment.

`git` helped me a lot, every commit has a diff that I used to explain the
solution of the lessons. People that were not confidently writing the solution in a
particular language had to just `cherry-pick` the commit in the language they
didn't know.

## Collaboration

The practical part was designed to be a collaboration between people. IMHO it
helps to feel less "at school" but more as a team that it is something we should
feel more comfortable with at work.

I think it worked but not that well. People were supporting and helping each
other. But I probably need to cut the lessons in a different way. I think I will
remove the `influxdb` lessons injecting the learning process of only what
matters for the course along the other lessons.  Next time and I will develop a
new lesson about how to parse the logs and push them to InfluxDB for example.
(let me know if you would like to help me!)

## Feedback

I asked them to do a survey before the end of the course in order to help me
get their feeling. There is a lot to do and some of their feedback is part
of this article. But in general, I am happy because I have all the material in
order and this for me was just a first iteration. I hope to make it better, to
grant more feedback from the open source and to run it again! So let me know if
you would like to have me on board!

## Next

As I said instrumentation is hard and I still hoping to get an easier solution
across languages. I tried OpenCensus but I didn't manage to have it running at I
was in the rush so I used Jaeger.

I will develop something about structured logging as I said for sure.

I hope to get a lesson from some as a service provider like HoneyComb for
example.

## Fun Fact
The youngest person in the room was a student in high school! Wow!
