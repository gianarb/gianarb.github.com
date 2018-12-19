---
heroimg: /img/hero/high-cardinality.png
layout: post
title:  "You need an high cardinality database"
date:   2018-11-28 08:08:27
categories: [post]
tags: [observability, monitoring, data analysis, tracing, logs, metrics, events]
summary: "Monitoring and observability in a dynamic environment on Cloud or
Kubernetes is a new challange we are facing and I think the tool that plays a
big role is an high cardinality database."
changefreq: daily
---
In order to understand how an application performs you need data. Logs, events,
metrics, traces.

Observability and monitoring are expensive because you need to retrieve all this
data across your system. An architeture these days is not a static rock where
nothing happens and everything stays the same. You don't have your 10 VPC, with
always the same hostname that you can filter for.

Today you are on cloud, your instances are going up and down based on your load
and it is easier for you to replace an EC2 that troubleshooting a failure.

Containers wrap your application and they makes it easy to deploy, as side
effect you release more often, it means more data.

But the data are useless if you can not get anything good out from them, so they
can be your silver buffet or a big pain, the difference is all made by your
ability to use them to answer your questions or in the ability for your team
aggregate them together in order to build automation with them.

To do all off this you need to manage high cardinality, this is word that sales
team in tech are scary off because nobody will never sell an infinite high
cardinality database, everything has a limit, and the unique solution is not a
product itself but it is more like a mindset developers should have.

* You need to store the raw data for just the right time, forever is not an
  option.
* You need to give access to these data across the company in order to build
  better aggregation. Build Engineers will probably need data not just from the
  CI pipeline but also from your VCS. SRE to understand how a code change behave
  in prod they need metrics from servers but also from the CI. Spread the
  knowledge

The technologies that gives you the ability to interact with a big set of
unstructured data should support an high wrtie troughpoot and smart indexes that
will allow your query engine to lookup for what you need fast enough!

So that's what I have in my mind when I think about a database that can support
monitoring data.

I am not selling anything mainly because I think a final solution doesn't exist
yet, I can not really tell you what to buy but you should look around for other
companies at your same scale because everyone has this problem:

* Facebook has scuba
* A lot of people use Cassandra and they looks happy at least for its writing
  capabilities.
* There are new time series databases releases in a daily based
* At InfluxData we obviously use InfluxDB for this purpose

The general idea here is that the goal should be to group
data that now are in different sources: NewRelic, InfluxDB, ElasticSearch,
Papertrail in the same place, because it is rare to get
the answer for your question just looking at logs, or metrics, you need an
aggregation or a sample of different data.

This will bring the debugging and troubleshooting capabilities of your team to
the next level, and listen to me, if you are working with a microservices
architecture or with a highly distributed environment you need help from
everything!
