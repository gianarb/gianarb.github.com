---
layout: post
title:  "Orbiter an OSS Docker Swarm Autoscaler"
date:   2017-03-12 08:08:27
categories: [post]
img: /img/container-security.png
tags: [docker]
summary: "Orbiter is an open source project design to become a cross provider
autoscaler. At the moment it works like Zero Configuration Autoscaler for Docker
Swarm. It also has a basic implementation to autoscale Digitalocean. This
project is designed with InfluxData a company that provides OSS solution like
InfluxDB, Kapacitor and Telegraf. We are going to use all this tools to create
an autoscaling policy  for your Docker Swarm services."
changefreq: yearly
---
## Autoscaling
One of the Cloud's dreams is a nice world when everything magically happen. You
have unlimited resources and you are just going to use what you need and to pay
to you use.
To do what AWS provides a service called autoscaling-group for example. You can
specify some limitation and some expectation about a group of servers and AWS is
matching your expectation for you.
If you are able to make an automatic provision of a node you can use Cloudwatch
to set some alerts. When AWS trigger these alerts the austocaling-group is
creating or removing one or more instance.

### let's try with an example
You have a web service and you know that for 2 hours every day you don't need 4
EC2 because you have a lot of traffic, you need 10 of them.
You can create an autoscaling group, set some alerts:

1. When the memory usage is more than 65% for 3 minutes start 3 new servers.
2. When the memory usage is less than 30% for 5 minutes stop 2 servers.

Just to have an idea. In this way AWS knows what do you and you don't need to
stay in front of your laptop to wait something happen. You can just do something
funny.

It's something useful, if you think about a daily magazine, they usually has a
lot of traffic in the beginning of the day when all the people are usually
reading news. At that's an easy scenario.

But it can also happen than a new shared on reddit or HackerNews is getting a
lot of traffic and the last thing that you are looking for is to go down just
during that spike!




```
                       +-----------------------------------+
                       |                                   |
+----------------------v-----+                             |
|    Alert/monitoring system |                             |
|                            +------------------+          |
|                            |                  |          |
| like InfluxDB and Kapacitor|                  |          |
+-----+----^------+-----+-------------------+---v----+     |
      ^    +---+  |     |                   |  orbiter     |
      |        +--+     |                   |        |     |
      |        +--------+ Cluster   Swarm   +--------+     |
      |        |                                     |     |
      |        |                                     |     |
      |        +------+-------+          +------+----+     |
      |        |      |       |          |      |    +-----+
      +--------+      +-------+          |      +----+
               |              |          |           |
               +--------------+----------+-----------+

```
