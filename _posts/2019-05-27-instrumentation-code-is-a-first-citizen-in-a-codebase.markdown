---
img: /img/gianarb.png
layout: post
title:  "Instrumentation code is a first citizen in a codebase"
date:   2019-05-27 08:08:27
categories: [post]
tags: [code instrumentation, library, opentracing, opencensus, opentelemetry,
influxdb, prometheus, opensource, honeycomb]
summary: "How fast we are capable of instrumenting an application decrease the
out of time requires to understand and fix a bug."
changefreq: daily
---
A few years ago a log was very similar to a printf statement with a message that
in some way we're trying to communicate to the outside the current situation of
a specific procedure.  The format and composition of the message were not
crucial, the main purpose was to make it easy enough to read.  A full-text
search engine is capable of tokenizing and indexing every message for easy
lookup and aggregation and it was enough to fix the gap between a human
understanding log and something that a program can parse and visualize.  Cloud
Computing and containers changed the way we architect, visualize and deploy
software:

1. **The distribution of our applications**. Compared with a more traditional
   approach our application runs in the smallest but much-replicated units
   (container, pods, ec2 and so on).
2. **The size** of our application (microservices) and by consequence the
   interaction between them, over a not perfect communication layer (the
   network).
3. Applications come and go much more frequently because we have automation that
   takes care of the number of replicas running inside a system. They are more
   **dynamic** and we do not really have a stable identifier as before:
   hostnames, IPs change more often.

These points increase the importance for us to get applications metrics out from
our code because that's the language our application speaks. We rely on them in
order to understand what is going on.  We need to realize that logs and metrics
have different purposes:

* To understand what is going right now
* To verify what happened in the past (even from a legal perspective)
* To compare

They are not random printf. All these purposes require methodologies and tools.
This article will stay focused on the first point: "What is going on?" because
it is a question I ask even to myself when I look at the system I wrote or
manage and the answer is a real pain to retrieve.  To troubleshoot a system we
need a very dense amount of information "almost in real time" because that's
when a system is broken "now" and a picture or a sample of older data in order
to compare the current situation with something that we can define as "working".
We can not really use old data because our codebase changes frequently (because
somebody told us that we can break and develop fast). So there is not a lot of
value at looking at high-density data coming from two weeks ago where the
codebase was different. That's why time series databases as InfluxDB have data
retention features built in to keep themselves clean.
[InfluxDB](https://github.com/kapacitor/influxdb) removes the data after a
certain amount of time, but with
[Kapacitor](https://github.com/influxdata/kapacitor) you can aggregate or sample
the data to an older retention policy in order to keep what you need in the
database.  Back in the day, I wrote this article about [Opentracing and
Opencensus](https://gianarb.it/blog/what-is-distributed-tracing-opentracing-opencensus).
This is a follow up after another year of working around code instrumentation,
observability, and monitoring.

First of all both of them are vendor-neutral projects that help you instrument
your applications without lock you with a specific provider. It doesn't really
need to be a bad, evil vendor. If you use the Prometheus client directly in your code,
everywhere, you will be locked to it forever or until you will find the right
time to move over all your codebase. But it sounds like "change your logger":
something you would like to do magically, one shot without wasting your time.

OpenTracing is 100% for tracing, the problem it solves is about how to
instrument your application to send traces. OpenCensus does the same, plus it
also takes care of metrics.

These two projects have a major issue, they are TWO different projects. They
 were not smart enough to agree on the same format and it split the dev community without
any reason, sham of you!. Good for us they will be
[merged](https://medium.com/opentracing/merging-opentracing-and-opencensus-f0fe9c7ca6f0)
together at some point to something called OpenTelemetry. Finally!

Another misunderstanding is around how tracers such as Zipkin, Jager, XRay advertise them self as
"opentracing compatible". When I think about "compatible" I think like a REST
API that follow some rules, and for that reason, the SystemA is compatible with
SystemB and you can change them transparently.

This is not what happens with tracing infrastructure, because you need to
remember that OpenTracing and OpenCensus play in your codebase size, it is not
REST or nothing like that.

Compatibility, in this case, means that the
tracers (Zipkin, Jaeger, AWS X-Ray, NewRelic) ship an OpenTracing compatible
library across many languages that you can change in your codebase in order to
point your application to a different tracer without changing the
instrumentation code you wrote.

NB: OpenCensus has the same goal for metrics as well

```javascript
function initTracer(serviceName) {
  var config = {
    serviceName: serviceName,
    sampler: {
      type: "const",
      param: 1,
    },
    reporter: {
      agentHost: "jaeger-workshop",
      logSpans: true,
    },
  };
  var options = {
    logger: {
      info: function logInfo(msg) {
        logger.info(msg, {
          "service": "tracer"
        })
      },
      error: function logError(msg) {
        logger.error(msg, {
          "service": "tracer"
        })
      },
    },
  };
  return initJaegerTracer(config, options);
}

const tracer = initTracer("discount");
opentracing.initGlobalTracer(tracer);
```
This example comes from
[shopmany](https://github.com/gianarb/shopmany/blob/end/discount/server.js) a
test e-commerce I wrote. In this case, the `tracer` is Jaeger, but if you need
to change to Zipkin you can probably use
[zipkin-javascript-opentracing](https://github.com/DanielMSchmidt/zipkin-javascript-opentracing)

It is important to evaluate an instrumentation library like OpenCensus,
OpenTracing, OpenTelemetry because there is a community that writes and supports
libraries across many languages and tracers. it means that you do not really
need to write your own library, that sounds a bit like too much!  I was very
frustrated about the fact that these two libraries was TWO! I can't wait to see
how the result will look like.  How easy it is to instrument an application is a
key value for a company like Honeycomb.io and this sounds like a good reason for
them to have their own instrumentation library
([go](https://github.com/honeycombio/beeline-go),
[js](https://github.com/honeycombio/beeline-nodejs),
[Ruby](https://github.com/honeycombio/beeline-ruby)), and when they started the
ecosystem was different (it is still a mess today as you read) but I hope that
OpenTelemetry will push everybody to just work together because understanding
what is going in production right now is a hard, messy and amazing challenge.

<blockquote class="twitter-tweet"><p lang="en" dir="ltr">it is so nice to see
how two great open source community such as <a
href="https://twitter.com/InfluxDB?ref_src=twsrc%5Etfw">@InfluxDB</a> and <a
href="https://twitter.com/ntop_org?ref_src=twsrc%5Etfw">@ntop_org</a> can do
togheter. That&#39;s how we can solve observability/monitoring challanges all
togheter <a
href="https://twitter.com/Chris_Churilo?ref_src=twsrc%5Etfw">@Chris_Churilo</a></p>&mdash;
gianarb (@GianArb) <a
href="https://twitter.com/GianArb/status/1126107355895214082?ref_src=twsrc%5Etfw">May
8, 2019</a></blockquote> <script async
src="https://platform.twitter.com/widgets.js" charset="utf-8"></script>


## Keep instrumentation

![The infinite symble](/img/infinite-loop.png)

The ability to instrument an application fast, precisely increase your
troubleshooting capabilities. Fast your iterate on your instrumentation code
faster your will understand what is going on. It is not a one short exercise but
it is something you improve everyday based on what you will learn. But your
ability to learn depends on how well you can read the language that your
applications exposes (let me tell you a secret, it depends on how well you
instrument your code).

More to read:
- [Jaeger and
  OpenTelemetry](https://medium.com/jaegertracing/jaeger-and-opentelemetry-1846f701d9f2)
- [Structured
  logs](https://www.honeycomb.io/blog/how-are-structured-logs-different-from-events/)
- [Logs Metrics Traces are uqually
  useless](https://gianarb.it/blog/logs-metrics-traces-aggregation)
