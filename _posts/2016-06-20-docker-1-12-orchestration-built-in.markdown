---
layout: post
title:  "Docker 1.12 orchestration built-in"
date:   2016-06-20 10:08:27
categories: [post]
img: /img/docker.png
tags: [docker, release, swarm, cloud, orchestration, distributed system, docker
captain]
summary: "Docker 1.12 adds different new features around orchestration, scaling
and deployment, in this article I am happy to share some tests I did with this
version"
priority: 0.6
---

<blockquote class="twitter-tweet tw-align-center" data-lang="en"><p lang="en" dir="ltr">Some
tests with Docker 1.12! <a
href="https://t.co/budUOtMuBB">https://t.co/budUOtMuBB</a> <a
href="https://twitter.com/hashtag/docker?src=hash">#docker</a> <a
href="https://twitter.com/hashtag/DockerCon?src=hash">#DockerCon</a>
orchestration, swarm and services.</p>&mdash; Gianluca Arbezzano (@GianArb) <a
href="https://twitter.com/GianArb/status/744977855277309953">June 20,
2016</a></blockquote> <script async src="//platform.twitter.com/widgets.js"
charset="utf-8"></script>

During the DockerCon 2016 docker announced Docker 1.12 release.
One of the news stories around the new version is the orchestration system built directly
inside the engine, this feature allow us to use swarm
without installing it separately from outside, it’s now a feature provided by Docker directly.

<div class="post row">
  <div class="col-md-12">
      {% include book-adv-lb.html %}
  </div>
</div>

Now we have a new set of commands that allow us to orchestrate containers
across a cluster.

{% highlight bash %}
docker swarm
docker node
docker service
{% endhighlight %}

All these commands are focused on increasing our ability to orchestrate our
containers and also join them in services.

{% highlight bash %}
#!/bin/bash

# create swarm manager
docker-machine create -d virtualbox sw1
echo "sudo /etc/init.d/docker stop && \
    curl https://test.docker.com/builds/Linux/x86_64/docker-1.12.0-rc2.tgz | \
    tar xzf - && sudo mv docker/* /usr/local/bin && \
    rm -rf docker/ && sudo /etc/init.d/docker start" | \
    docker-machine ssh sw1 sh -
docker-machine ssh sw1 docker swarm init

# create another swarm node
docker-machine create -d virtualbox sw2
echo "sudo /etc/init.d/docker stop && \
    curl https://test.docker.com/builds/Linux/x86_64/docker-1.12.0-rc2.tgz | \
    tar xzf - && sudo mv docker/* /usr/local/bin && \
    rm -rf docker/ && sudo /etc/init.d/docker start" | \
    docker-machine ssh sw2 sh -
docker-machine ssh sw2 docker swarm join $(docker-machine ip sw1):2377
{% endhighlight %}

 another Captain wrote this script that I just updated to work
with the public Docker 1.12-rc2. We can use this script to create a cluster with
virtual box ready to be used.  After this script you can see the number of
workers and masters, in this case your one and one.

{% highlight bash %}
$ docker node ls
{% endhighlight %}

Docker 1.12 has a built-in set of primitive functions to orchestrate your containers just
like a summary. The main commands that you must run to create a cluster are

{% highlight bash %}
## On the master to start your cluster
$ docker swarm init --listen-addr <master-IP(this ip)>:2377
## on each node to add it into the cluster
$ docker swarm join <master-ip>:2377
{% endhighlight %}

<img class="img-fluid" alt="Docker Swarm architecture" src="/img/posts/swarm_arch.png">

If you are not confident with docker swarm this is the architecture, this graph
is provided by Docker Inc. and explains really well the design around this project.
The principal actors are managers and workers, managers are the brains of the
system, they dispatch schedules and remember services and containers. Workers
execute these commands.

The cluster is secure because each node has a proper TLS
identity and all communications are encrypted end to end by default with a
automatic key rotation in order to increase the security around the keys use in
the cluster.

[Raft](https://raft.github.io/) is the consensual protocol used to distribute
message around the cluster and check the number of nodes, it’s complex
algorithm but really interested I have in plan another article about it but the
offical site contains a lot of details about it.

We already saw the concept of services in docker-compose they are a single or a
group of containers to describe your ecosystem, you can scale a specific
service or orchestrate it across your cluster. It’s the same here, you don't have
a specification file like compose at the moment but anyway you can run a bunch
of commands to create your service.

{% highlight bash %}
$ docker service create --name helloworld --replicas 1 alpine ping docker.com
{% endhighlight %}

With this example we push up a new service helloworld. It has one container from
the alpine image and it pings docker.com site.

{% highlight bash %}
docker service ls
{% endhighlight %}

To watch all our services, we can also inspect a service

{% highlight bash %}
docker service inspect <container_id>
{% endhighlight %}

There is a new concept, when you run a service you are also creating a task,
this task represents the container/s under your service, in this case we have
just one task

{% highlight bash %}
docker service tasks helloworld
{% endhighlight %}

When you scale your service you are creating new tasks

{% highlight bash %}
docker service scale helloworld=10
{% endhighlight %}

Now you can see 10 tasks that are running and you can inspect one of them,
inside you can find the containerId and you can, for example, follow logs

{% highlight bash %}
22:17 $ docker inspect  6fhfse4it8lwzlsk1t5sd5jbk
[
    {
        "ID": "6fhfse4it8lwzlsk1t5sd5jbk",
        "Version": {
            "Index": 67
        },
        "CreatedAt": "2016-06-18T21:06:36.707664178Z",
        "UpdatedAt": "2016-06-18T21:06:39.241942781Z",
        "Spec": {
            "ContainerSpec": {
                "Image": "alpine",
                "Args": [
                    "ping",
                    "docker.com"
                ]
            },
            "Resources": {
                "Limits": {},
                "Reservations": {}
            },
            "RestartPolicy": {
                "Condition": "any",
                "MaxAttempts": 0
            },
            "Placement": {}
        },
        "ServiceID": "24e0pojscuj2irvlxvx2baiid",
        "Slot": 2,
        "NodeID": "55v4jjzf56mcwnhbwvn4cq1rs",
        "Status": {
            "Timestamp": "2016-06-18T21:06:36.7110425Z",
            "State": "running",
            "Message": "started",
            "ContainerStatus": {
                "ContainerID": "4ec69142e3e886098915140663737f4176c6de5afe9f2fad1f5b2439d8fc336d",
                "PID": 3627
            }
        },
        "DesiredState": "running"
    }
]
22:17 $ docker logs -f 6fhfse4it8lwzlsk1t5sd5jbk
{% endhighlight %}

At this point it is a normal container and it’s running on your cluster.
Well I tried to explain the main concept around this big feature provided by
Docker 1.12, the last example is just to cover the DNS topic.

I created an application that serve an http server and print the current IP.
Each server has an internal load balancer that dispatches traffic in round robin
between the different tasks.
In this way it’s totally transparent, you can just
resolve your service with a normal URL, docker will do the rest for you.

{% highlight bash %}
$. docker service create —-name micro —-replicas 10 —-publish 8000/tcp gianarb/micro
{% endhighlight %}

[Micro](https://github.com/gianarb/micro) is an application that exposes an
http server on port 8000 and print the current ip, now we have 10 tasks with
this service.
To grab the current entry point for our service we can inspect it
and search for this information:

{% highlight bash %}
$. docker service inspect <id-service>

...
      "Endpoint": {
            "Spec": {},
            "Ports": [
                {
                    "Protocol": "tcp",
                    "TargetPort": 8000,
                    "PublishedPort": 30000
                }
            ],
            "VirtualIPs": [
                {
                    "NetworkID": "890fivvc6od3pa4rxd281lobb",
                    "Addr": "10.255.0.5/16"
                }
            ]
       }
...
{% endhighlight %}

In this case our published port is 3000, we can call <ip>:3000 to resolve our
service, if you try to do multi request you can see your IP chances because the
internal DNS is calling different containers.

This is just an overview about the features but there are other powerful news
like DAB, stacks and how do an easy update of your containers, this could be
the topic around my next article. Please stay in touch follow me on
[Twitter](https://github.com/gianarb) to chat and receive news about the next articles.

<blockquote>
  <p>Thanks <a href="https://twitter.com/gpelly">@gpelly</a> for your review!</p>
</blockquote>


{% include docker-planet-newsletter.html %}
