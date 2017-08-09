---
layout: post
title:  "Orbiter the Docker Swarm autoscaler on the road to BETA-1"
date:   2017-08-09 10:08:27
categories: [post]
tags: [docker]
img: /img/docker.png
summary: "Orbiter is a project written in go. It is an autoscaler for Docker
containers. In particular it works with Docker Swarm. It provides autoscaling
capabilities for your services."
changefreq: yearly
---
Orbiter is an open source project written in go hosted on
[GitHub](https://github.com/gianarb/orbiter). It provides autoscaling
capabilities in your Docker Swarm Cluster.

As you probably know at the moment autoscaling is not a feature supported
natively by Docker Swarm but this is not a problem at all.

Docker Swarm provides a useful API that helps you improving its capabilities.

I created Orbiter months ago as use case with InfluxDB and to allow services to
scale automatically based on signal `up` or `down`. You can follow the webinar
that I made with InfluDB
[here](https://www.influxdata.com/resources/influxdata-helps-docker-auto-scale-monitoring/?ao_campid=70137000000Jgw7).

This article is not about "How it works". You can [read more here about how it
works](https://gianarb.it/blog/orbiter-docker-swarm-autoscaler) and you can
watch the embedded video that I made in the Docker HQ in San Francisco.

Yesterday we made some very good improvements and we are moving forward to tag
the first beta release. I need to say a big thanks to [Manuel
Bovo](https://github.com/mbovo). He coded pretty much all the features listed
here.

0. [PR #26](https://github.com/gianarb/orbiter/pull/26) e2e working example. [Please try
it](https://github.com/gianarb/orbiter/tree/master/contrib/swarm).

1. [PR #27](https://github.com/gianarb/orbiter/pull/27) Now Orbiter has
   background job that listen on the Docker Swarm event API and register and
   de-register new services [deployed with right
   labels](https://github.com/gianarb/orbiter#autodetect). You don't need to
   restart orbiter anymore. It detect new services automatically.

2. [PR #29](https://github.com/gianarb/orbiter/pull/29) Fixed the up/down range.
   Now we can not scale under 1 tasks but we can scale up services with 0 tasks.

3. [PR #31](https://github.com/gianarb/orbiter/pull/31) We have a cooldown
   period configurable via label `orbiter.cooldown`. This fix avoid multiple
   scaling in a short amount of time.

4. [PR #32](https://github.com/gianarb/orbiter/pull/32) We are migrating our API
   base root. Now all the API are `/v1/orbiter/.....`. At the moment we are
   supporting old and new routes. **In October I will remove the old one. Please
   migrate to `/v1/orbiter/....` now!**.

## Now?

That's a good question, but I have part of the answer. In October the plan is to
release a BETA and finally the first stable version but what we need to do to go
there?

* Offer a proper auth method. Manuel started this
  [PR](https://github.com/gianarb/orbiter/pull/33). I have some concerns but
  we are on the right path.
* Make orbiter "Only-Swarm". The project started with the vision to become a
  general purpose autoscaler. But this is not in line with the idea of single
  responsibility and we designed a very clean API for Docker Swarm, make it
  usable in other context is not going to work. We tried it with DigitalOcean
  but the api and the project looks too complex and I love simplicity.
* Get other feedback from the community to merge valuable features before the
  stable release.

That's it! Share it and give it a try! For any question I am available on
twitter (@gianarb) or open an issue.
