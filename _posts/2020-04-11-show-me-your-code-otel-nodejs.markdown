---
img: /img/show-me-your-code-logo.png
layout: "show-me-your-code"
title: "Show Me Your Code with Dan and Walter: How to contribute to
OpenTelemetry JS"
date: 2020-04-11 09:00:27
categories: [post]
tags: ["show-me-your-code", "codeinstrumentation"]
summary: "Show me your code has two special guest Walter CTO for
CorelyCloud SRL the company behind the CloudConf in Turin and Dan Engineer at
Dynatrace Maintainer of OpenTelemetry JS. This during show we will talk about
OpenTelemetry and NodeJS. Walter wrote a plugin for instrumenting mongoose with
opentelemetry. We are gonna see how he did it, considerations from Dan and so on"
youtubeVideoID: FV4mSApgY60
changefreq: daily
---

When: Thursday 16th 6-7pm GMT+2 (9am PDT)

## OpenTelemetry for JS and how to contribute

OpenTelemetry is a specification and set of instrumentation libraries developed
in open source from multiple companies such as Google, HoneyComb.io, Dynatrace,
LightStep and many more!

OpenTracing and OpenCensus joined the force, and they started a common project
called OpenTelemetry that I hope will become the way to go in terms of code
instrumentation because I really think it is something we need.

Walter and his team develop in Javascript, frontend and backend and back in the
day we experimented OpenTracing but we had some issue and it was not easy to
pick up at that time. When I tried OpenTelemetry I realized that it was for him.

He tried it out and he wrote its first opentelemetry instrumentation plugin
mongoose, a popular library he uses and that it was not instrumented yet.

Dan will help us to figure how they designed the opentelemetry-js implementation
as it is today, the good the bad and the ugly about this experience. I hope to
get some feedback about roadmap and future development as well now that the
library reached its first release beta.

## About Dan

When I was working on my observability workshop Dan gave me a huge help,
drastically increasing my very low experience with NodeJS. Thank you for that.

Dan works as Engineer at Dynatrace, and he maintains the OpenTelemetry JS
library. You can find him on twitter as [@dyladan](https://twitter.com/dyladan)
and in [Gitter](https://gitter.im/open-telemetry/opentelemetry-node) discussing
opentelemetry.

## About Walter

Walter Dal Mut works as a Solutions Architect [@Corley SRL](https://corley.it/).
He is an electronic engineer who moved to Software Engineering and Cloud
Computing Infrastructures. Passionate about technology in general and open
source movement lover.

You can follow him on [Twitter](https://twitter.com/walterdalmut)
and [GitHub](https://github.com/wdalmut).

## Links

* [opentelemetry.io](https://opentelemetry.io/)
* [dynatrace.com](https://www.dynatrace.com)
* [How to start tracing with OpenTelemetry in NodeJS?](https://gianarb.it/blog/how-to-start-with-opentelemetry-in-nodejs)
* [github.com/open-telemetry/opentelemetry-js](https://github.com/open-telemetry/opentelemetry-js)
* [wdalmut/opentelemetry-plugin-mongoose](https://github.com/wdalmut/opentelemetry-plugin-mongoose)
