---
layout: post
title:  "Watch demo about Docker 1.12 made during Docker Meetup"
date:   2016-08-24 12:08:27
categories: [post]
img: /img/docker.png
tags: [docker]
summary: "Docker 1.12 contains a lot of news about orchestration and
    production. During August Docker Meetup in Dublin I presented with a demo a set
    of new features around this new release."
priority: 0.6
changefreq: yearly
---
In August during the Docker Meetup I presented with a demo some new
features provided by Docker 1.12.

It's an important release because it improves your experience with Docker
in production with an orchestration framework included into the product.

Docker provides a new set of commands to create a cluster of Docker
deamon and manage a production enviroment.

<div class="post row">
  <div class="col-md-12">
      {% include book-adv-lb.html %}
  </div>
</div>

It's something like Kubernetes, Mesos, Swarm but it is included and
built-in Docker.

I wrote an article about it few months ago ["Docker 1.12 orchestration
built-in"](http://gianarb.it/blog/docker-1-12-orchestration-built-in).


In this demo I do an introduction of some new features like:

<div style="    text-align: center;">
<iframe width="420" height="315"
src="https://www.youtube.com/embed/h7a7vhzjElo" frameborder="0"
allowfullscreen></iframe>
</div>

* How create a SwarmMode docker cluster
* What is a service? What tasks means?
* How Docker SwarmKit manage a node down?
* I tried to show the HealthCheck feature :)
* How docker swarmkit manage containers update
* service discovery

{% include docker-planet-newsletter.html %}
