---
img: /img/got-your-back.jpg
layout: post
title: "Why code instrumentation?"
date: 2020-03-29 09:08:27
categories: [post]
tags: [golang, codeinstrumentation, o11y, opentelemetry, opentracing, prometheus, sre]
summary: "I decided to finally create a category about code instrumentation.
Because I am a develop. And I think it matters. It is important to write better
code and more reliability application that we can learn from."
changefreq: daily
---

I am writing this blog post as a common introduction for a new category I would
like to write about consistently on my blog. If this is the first time you land
here this is my blog, and I write about everything that catches my attention but
sometime earlier sometime later on I realize I can group my posts in categories,
and that's what I am doing now.

Some of them are: [Assemble Kubernetes](/planet/assemble-kubernetes.html),
[Docker](/planet/docker.html), [MockMania](/planet/mockmania.html). This one
will be called `Code Instrumentation`.

There are a lot of people writing about observability, monitoring and I did it
for the last 3 years as well. I learned a lot along the way but what I think is
crucial is that developers has to write code that is
understandable and easy to debug where it is more valuable, in production. And
if an application or a system is hard to figure out we as a developer play a
mojor role on it.

That's why Site Reliability Engineering (SRE) is not related to ops, servers,
Kubernetes but it is something that plays its match in your code.

What's why I think SRE and DevOps are different, not at all connected.

The technologies that are leading the landscape are:

1. Prometheus but not the time series database, their client libraries and the
   exposition format, now branded from the community and the Cloud Native
   Computing Foundation (CNCF) as OpenMetrics
2. OpenTracing, OpenCensus and OpenTelemetry. They are part of the same bullet
   points because I think about them as the consequence of each other since what
   I hope is "THE LAST ONE", OpenTelemetry. They are instrumentation libraries
   and specification to increase interoperability and to avoid vendor lock-in
   for what concerns distributed tracing and metrics. I hope logs will jump
   onboard at some point

## Prometheus and OpenMetrics

I wrote about this topic previously, so have a look there if you do not know
what I am speaking about.

I think they are worth to mention here because that's how I learned the effect
of good or bad code instrumentation, and the fact that it has to happen in your
code, when you develop it.

It has the same weight has writing a good data structure, writing solid unit
tests, or picking the right design pattern.

## OpenTelemetry (otel)

As I said I will refer to otel when I can, not because I think OpenTracing or
OpenCensus is bad, but because I do not see this as a religion, but for me it is
a technical problem, they is well spread and it has to find a good answer.

Those communities decided to merge to otel in the way they are doing, good or
bad? We can get a beer at some point and I will tell you. It is out of scope.

## What I am gonna talk about

This is a long new category introduction blog post probably but that's it. Over
the last two years I tried to share what I experienced around this topic with a
workshop called: "Application Monitoring". A lot of the articles that I will
write comes from there, and it is an attempt to share what I think worked or
failed.

## Links

* [All about Code Instrumentation](/planet/code-instrumentation.html) from my blog
* [ShopMany](https://github.com/gianarb/shopmany) is the application I developed for the workshop
* [Workshop notes](https://github.com/gianarb/workshop-observability) contains notes, exercises and solutions for the lessons I
  proposed in the workshop itself
* [honeycomb](https://www.honeycomb.io/blog/) because when you speak about o11y you have to quote them!
* [My newsletter](/tinyletter.html) is probably the best way to stay in touch with the content I
  create
* [Twitter](https://twitter.com/gianarb) is the best way to stay in touch with me
