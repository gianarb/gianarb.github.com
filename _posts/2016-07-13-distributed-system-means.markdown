---
layout: post
title:  "What Distributed System means"
date:   2016-07-12 16:08:27
categories: [post]
img: /img/distributed_system_planet.png
tags: [distributed_system, distributed system, discovery, orchestration,
container, open source]
summary: "I will speak about service discovery, micro services, container, virtual machines, schedulers, cloud, scalability and latency I hope to have, at the end of this experience a good number of posts in order to share what I know and how I work and approach this kind of challenges."
priority: 0.6
---
I choose  to put my experience about distributed system in a serie of blog
posts on which I’ll  cover different topics.

I will speak about service discovery, micro services, container, virtual
machines, schedulers, cloud, scalability and latency I hope to have, at the end
of this experience a good number of posts in order to share what I know and how
I work and approach this kind of challenges.

In first I will not speak about nothing new, in fact distributed system means:

<blockquote>A distributed system consists of a collection of autonomous computers,
connected through a network and distribution middleware, which enables
computers to coordinate their activities and to share the resources of the
system, so that users perceive the system as a single, integrated computing
facility.
<p><a href="http://www0.cs.ucl.ac.uk/staff/ucacwxe/lectures/ds98-99/dsee3.pdf" target="_blank">Wolfgang Emmerich, 1997</a></p>
</blockquote>


Internet is a distributed system, you infrastructure is usually a distributed
system if you follow the minimum requirements to make high availability for
your services.

In first of all I love what service means, your application is a service,
microservices is just a way to remind to people that a little application is
easy to maintain, deploy and control but the idea in my opinion is just make
something autonomous and useful for your customers. Sometimes your customer is
a human in other case could be another service provided by yourself or from a
third-party, it is not really important. It’s important that your service must
be ready to communicate with the extern.

Distributed your system is important to make it available, if you close your
service inside a single datacenter in a single part of the world you take the
risk to make it unavailable in case of problem in that particular area, if you
distribute your service in different location you are increasing the chances to
stay up.

You are also mitigating the latency around your system because you are bringing
your application near your customers and if you have a world wide traffic
that’s param is really important.

<img alt="Internet Global Submarine map" src="/img/global-submarine-cable.jpg" class="img-fluid">

This is the map of the submarine cable (2014) and all know that internet is not
in the air and serve different point in the world require different amount of
time to have a response and also is not just a problem of distance but traffic
and quality of the network have them weight. Akamai is a expert about this
topic, he provide a service of content delivery (CDN) and also it’s a
monitoring system for the status of the network, they provide different data,
one of them describe the [high level status of Internet](https://www.akamai.com/us/en/solutions/intelligent-platform/visualizing-akamai/real-time-web-monitor.jsp).

Virtualisation, container,  cloud computing and in general the low price to
design an infrastructure and the growth of internet’s users allow little
company with a little budget to create something of stable, secure and
available in different part of worlds. I think that for this reason micro
services and distributed system start to have a big impact in the industry.

A good exercise to understand the current situation could be design a little
infrastructure cross provider in multi datacenter to support a normal blog,
with a database and an application. With a couple of servers on different cloud
providers you can create an high available and distributed system across
multiple datacenter and avoid a lot of point of failure like: Geography
disaster Provider errror...

Docker, openstack, AWS, Consul, Prometheus, Elasticsearch, MongoDB are just a
set of products that help us to create something really stable and useful.
Continue Delivery, High Availability, disaster recovery, monitoring, Continuous
Integration, reliable are a subset of topic that you must resolve when you
think about distributed system because you can not care about where the
instances of your applications are around the world and the network is not a
paradise of stability.  Microservices helps you to create better and stable
application, allow your company to create more rooms for more developers and to
replace single pieces and features but they create other kind of problems like
architecture complexity, good knowledge of in different layers (DevOps point of
view), network and chain of failures. All the topics that we already know must
be adapter for this new architecture, monitoring, logging, deploy.
