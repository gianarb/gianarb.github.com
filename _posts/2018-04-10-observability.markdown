---
layout: post
title:  "Observability according to me"
date:   2018-04-04 10:38:27
img: /img/mountain-garbage.jpg
categories: [post]
tags: [observability, cloud, influxdb, monitoring, distributed system, scale]
summary: "Prometheus, InfluxDB, TimescaleDB, Cassandra and all the time series databases
that show up every week is a clear sign that now we need more than just a way to
store metrics. I think is now clear that collecting more metrics is the point.
More data is not directly related to a deeper understanding of our system."
priority: 1
---
I started to read about observability almost one year ago, [Charity
Major](https://twitter.com/mipsytipsy)
comes to my mind when I thinking about this topic and she is the person that is
pushing very hard on this.

This is probably the natural evolution of how we approach monitoring.

Distributed systems require a different way to approach the three monitoring
piles: collect, storage and analytics.

Understand a microservices environment brings a new layer of complexity and the
most obvious consequence is the amount of data that we are storing compared, I
think it is way more than before.

Prometheus, InfluxDB, TimescaleDB, Cassandra and all the time series databases
that show up every week is a clear sign that now we need more than just a way to
store metrics. I think is now clear that collecting more metrics is the point.
More data is not directly related to a deeper understanding of our system.

Observability for a lot of companies look like a new way to sell analytics
platform but according to me it's a scream to bring us back to the problem: "How
can we understand what is happening?" or even better "How should we use the data
we have to understand what's going on?".

All the data should be organized, reliable and usable. Logs, metrics, traces
are part of the resolution, the brain to analyze and get value out from them is
what Observability means to me.

Visualization is one expect, proactive monitoring, correlation, and hierarchy
are other steps. Looking at our old graphs all of them are driven by the
hostname for example. But now we have containers, we have virtual machines and
immutable infrastructure makes rebuilding less costly and more secure than an
incremental update. The name of the server should not be the keyword for our
queries, the focus should be moved to the role of services.

Think about your Kubernetes cluster, you label servers based on what they will
run, if something unusual happens the first things to do is to move the node out
from the production pool, the autoscaler will replace it and you will be
troubleshoot it later.

Before we were looking at processes, we were keeping them alive as the Olympic
flame but containers are making them volatile. We spin them up and down for
every request in some serverless environment. What we care and what we should
monitor are the events that float across our services, that's the new gold. W we
can lose 1000 containers but we can't miss the purchase order made my a
customer. All our effort should be moved on that side.

I love this point of view because it brings us to what really matters, our
applications.

<img src="/img/mountain-garbage.jpg" class="img-responsive">

According to me the mountain of waste showed in the picture explains really well
our current situation, we collected what ended up to be a lot of garbage and now
we need to climb it looking for a better point of view. I think the data in our
time series databases are not garbage but gold, it's just not simple as it
should be.

That's why is great that companies are building tools to fill the gap:
[IFQL](https://github.com/influxdata/ifql) is an example. The idea behind the
project is to build a language to query and manipulate data in an easy way. Same
for company line Honeycomb or open source projects like Grafana and Chronograf
that are trying to make these data easy to use.

We spoke about tools but there is another big expect and it's all a cultural,
distributed teams need different tools to collaborate and troubleshoot problems.
Different UI and way to interact with graph and metrics.
