---
img: /img/gianarb.png
layout: post
title: "OpenTelemetry the instrumentation library, I hope"
date: 2019-11-20 08:08:27
categories: [post]
tags: [o11y, tracing, observability, opentracing, opencensus, opentelemetry,
cncf]
summary: "OpenTelemetry, OpenCensus, OpenTracing, Open your heart"
changefreq: daily
---
Hello! If you follow my rumbling here or on
[twitter](https://twitter.com/gianarb) you know that I like to speak about
observability and tracing.

If you don't know what I am speaking about this is an
[FAQ](/blog/faq-distributed-tracing) about distributed tracing and something
about [OpenTracing and OpenCensus](/blog/what-is-distributed-tracing-opentracing-opencensus).

Observability is the ability to figure out what is going on in your application
from the outside. In order to do that you need to instrument your applications
in order to expose the right information.

The instrumentation is not easy, there are too many developers and too many
opinions, too many languages but in order observe a system that cross all the
applications and services everything needs to come together in the same way.
Otherwise the aggregation will become very a complicated job.

When you instrument an application there is a lot of code to write and inject,
you can not do or change it based on the vendor or services you are using to
store your telemetry: Zipkin, InfluxDB, NewRelic, HoneyComb and so on.

That's why over the last couple of years big foundations and companies such as
LightSteps, Google, CNCF, Uber tried to get their hands on the democratization
of code instrumentation. First with OpenTracing, after that with OpenCensus and
now with OpenTelemetry that is the merge between OpenCensus and OpenTracing.

At the beginning when this project went out a was very tired and stressed about
the topic. I made a workshop last year about code instrumentation at [the
CloudConf](https://cloudconf.it) and I wish it was easier to prepare and
develop. At the end attendees were satisfied btw, so I am happy enough.

Since the beginning I had a very bad feeling about OpenTracing and OpenCensus,
it is necessary as a project but the fact that we had two ways because they
didn't want to agree on only one for me was unbelievable.

Anyway, now that I pushed that feeling back I will give it another try. I will
get my [observability
workshop](https://github.com/gianarb/workshop-observability) and I will refresh
it to use OpenTelemetry because as I said, we need a way to instrument
applications, cross vendor and cross languages.

Here some link about it:

* [opentelemetry.io](https://opentelemetry.io/)
* [github.com/open-telemetry](https://github.com/open-telemetry)
* [Mailing List](https://lists.cncf.io/g/cncf-opentelemetry-community)

At KubeCon 2019 [lizthegrey](https://twitter.com/lizthegrey) gave a demo about
OpenTelemetry and I am confident that my experience will be a bit better.

It is not easy to democratize something, even less when you need to change the
habit for developers across programming languages. But that's the goal for
OpenTelemetry and I think we need to get there and to make it a commodity. It is
not a joke!

If you would like to help me, let me know!
