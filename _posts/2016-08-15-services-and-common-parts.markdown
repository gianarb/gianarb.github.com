---
layout: post
title:  “Microservices and common parts"
date:   2016-08-14 12:08:27
categories: [post]
img: /img/distributed_system_planet.png
tags: [distributed_system, distributed system, monitoring, microservices, open
source, containers, cloud]
summary: "When you think about microservices and distributed system there are a
    lot of parts that usually all your services require. Logging, monitoring,
    testing, distribution. Manage them in the best way it's one of the reason of
    success for your distributed system. In this article I shared few of this parts
    with some feedback to design them in a good way."
priority: 0.6
changefreq: yearly
---
Changing my glossary and replacing the concept of application with
service could be a buzzword but this takes me to built a
new approach to my work.

Nowadays many products require more services than before to
work: perhaps they could be modules, libraries directly integrate in one or
more services  or applications that communicate and provide a feature, it
doesn’t matter cause in any case the product’ll have some dependences.

If you start to follow this path, a lot of redundant concepts will show up in your
product:

* Monitoring
* Logging
* Authentication
* Scaling
* High availability
* Distribution
* Testing (unit, functional, integration..)
* And others

Some of them, like monitoring or logs, require architecture and tools
selection: you can use some B2B tools or host something in house. It’s not
only a problem of tooling, the other face of this “redundant money” is how your
services can communicate logs, metrics with outside in a clean and reusable
way.

In this post I will try to share the common part of a
microservices ecosystem and some possible approach to solve this issues.

## Logging

All applications require a good and strong log system.  There are few
libraries able to help you in managing this section but the minimum requirement, in my opinion, includes:

* Support for multiple stream: usually, I use stdout or file and I move
  them in a database with a separate pipeline, but a lot of good libraries allow you to manage your logs in different collectors.
* Different layers like: INFO, DEBUG, WARNING, FATAL.
* Provide a way to change this layer runtime, for example with a RPC call.

The third point is really important: if your application start to have a big
traffic, the amount of logs you must manage will be relevant; so, changing this
level runtime allows you to manage the amount of logs that you store and, for example, allows DEBUG
information only if you need to do some specific debugging in production. This strategy save storage and money.

There are a lot of services and open source tools able to manage and storage this data. The real issue is decide which street follow.

Are you interested to manage your logs or it’s a big effort for your company? You can move all to log entries and forgot about elastic search and kibana
and similar in this case. Think about your environment and catch the best solution. Remember that it could be just a temporary solution. When you start a business you have different thoughts, start slim and easy.

## Monitoring

Several services require several time and energy to be monitored and to be
maintained alive.

The best way to do that is with a time series
database like prometheus, InfluxDB or other as a service solution like NewRelic
or AppDynamics.

The real problem is how your application can provide
metrics readable and usable from external systems. You can find a very good solution to this problem in Docker: they provide different streams and events to grab this kind of informations.

If you take a lot on how it manage this part
you can implement a good system in your application.  A stream of events is
also a good API to allow other services to enjoy features provided from your
service.

## Heahtcheck

Understand with a single request if your application has all what it needs to
work is really important.

The microservices ecosystem contains a lot of micro applications that change and have dependencies to work. How can you understand if all your system is up and runs without spend a lot of time?

You can create for each service the call `/heath` that return 200 if all it’s fine and 500 if there is something that it’s not working properly.

During a release you can use this endpoint in order to understand if your
service is ready to be attached into the production pool.

In practice, if you have one service called Users that depends from MySQL and
from another service like Emailer, the health entrypoint for Users’s service
will check whether it can connect to the MySql and also you can call `/health` for
Emailer in order to check if the service is up.

Your orchestration and deploy framework can check after each deploy if the health is up and running and manage your release, it can revert or it doesn’t include your new release into the production pool.

## Authentication

Your microservice is not public, sometime you have a set of firewall’s rules or
a strong network settings to manage the security of your environment but for other
services the authentication layer is a requirement and usually there are few
services that need to know which is the identity of the user that is persisting
an action.

Think about a To Do service, it need to know the identity of the user in order
to fetch the correct items.

For this reason this layer could be common between your services and it’s also a
critical section of your architecture because usually from it depends the security of
your application and users.

Oauth2 is a framework to manage authentication, I recommend it
because it has a documentation already done, it’s a standard.
You don't reivent anything, there are a lot of libraries and use cases about it that make it solid, flexible and reusable.

## Automation and Deploy

A good layer of automation is important in every ecosystem to make your work less
bored but also to decrease chance for a human to make a mistake during a
repetitive task.

If you are thinking about a microservices ecosystem all this problems are
multiplied for a big number of applications.

Without a good layer of automation and a good deploy’s flow you will spend all
your day to put line of code in production without have time to stay focused on
new features or other business’s requests.

## Documentation

* Describe the topology of your ecosystem,
* how match microservices you have?
* where they are and how they are distributed across your datacenters
* Make it extensible and easy to read and update.
* How a single service works?
* Which APIs it expose
* how another service can communicate with it.
* Single dependencies for each microservices is also important to know.

All common part like, logs, auth, metrics help you to have a
common documentation easy to maintain, read and implement but for each service
you must provide a specific documentation because all it’s clear today but
between few months when you worked on ten other services the situation could be
really different.

One of the goal about microservices is the possibility to add
and integrate them easily. Documentation is one of the goal to make this
possible and efficient.

## Communication Layer

A lot of companies have one communication layer in the
environment, JSON and REST. It’s a good choice, easy to implement and there are
also a lot of tools to test, document and create client libraries.

But HTTP/REST is not the unique way to expose features out of your service, this is
really important to know.

There are other efficient and less expensive solution, binary protocol is
one of them.


For all this topics we can stay here to speech for years for this reason I have in
plan other posts to analyze some points better.

Please let me know if in your experience there are other common part between
your services.
