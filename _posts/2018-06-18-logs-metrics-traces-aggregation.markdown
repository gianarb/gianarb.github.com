---
layout: post
title:  "Logs, metrics and traces are equally useless"
date:   2018-06-18 10:38:27
img: /img/dna.jpg
categories: [post]
tags: [devops, observability, opentracing, tracing, distributed system,
monitoring, logs, cloud, tracing]
summary: "The key monitoring a distributed system is not logs, metrics or traces
but how you are able to aggregate them. You can not observe and monitor a
complex system looking at single signals."
priority: 1
---
Every signal from applications or infrastructure is useless, in the distributed
system era aggregation matters.

The ability to combine logs, metrics and, traces together is the key takeaway
here.

Kubernetes spin ups too many containers to allow us to stream or tail a log
fail.

Even cloud providers offer too many virtual machines to enable us to tail
logs.

A centralized place where to store all of them is a great start, but you
need to experience and learn how to combine the metrics you are ingesting to
increase the visibility over your system.

If you instrument your code with
opentracing, for example, you can get the `trace_id` and attach it to your log
to associate it with the trace itself. It can also work as the lookup key for
troubleshooting.

If you get some weird logs, you will know from where it comes.
With opentracing, this is still a bit of a mess the specification recently
[added explicit support to extract TraceId and SpanId from the
SpanContext](https://github.com/opentracing/specification/blob/master/rfc/trace_identifiers.md).
It is currently not implemented in a lot of implementation. I recently started a
conversation in the
[opentracing-go](https://github.com/opentracing/opentracing-go/issues/188)
project to figure out how to apply it because currently it depends from what
tracer you are using and it is an essential regression for the specification
itself that should hide it by design.

Using Jaeger this is the way to do it:
```
if sc, ok := span.Context().(jaeger.SpanContext); ok {
  sc.TraceID()
}
```
Using Zipkin:
```
zipkinSpan, ok := sp.Context().(zipkin.SpanContext)
if ok == true && zipkinSpan.TraceID.Empty() == false {
  w.Header().Add("X-Trace-ID", zipkinSpan.TraceID.ToHex())
}
```

To get back in track, I wrote this article because I saw this problem and this
inclination speaking with friends, colleagues and other devs, we are now good
(or just better) storing high cardinality values but save them inside a database
doesn't give us any value it is all about how we use them.

Correlation brings your alert to a different level. You probably have an alarm
to measure how much disks you still have.

An alert on the only CPU usage can be very frustrated
even more if it happens too often and a lot of time you restart a container or a
node to make it work because at 2 am you can't fix the cause. You can
investigate what matters to fill an issue on GitHub.

Every automation tools can make your work leaving you free to sleep. It can
probably fill out the issue.

Combining the CPU with the time for the system to recover from a node restart
can make your alert smart enough to wake you up when it is not able to fix
itself leaving you ready for more acute and trivial problems.

## Conclusion
It is a pretty straightforward concept, but yes, everything is useless if you
store data without getting values out of them doesn't matter if they are logs,
metrics or traces.  The real value is not in a single one of them, it is in how
do you aggregate them together because a complex simple doesn't explain itself
    over one signal.
