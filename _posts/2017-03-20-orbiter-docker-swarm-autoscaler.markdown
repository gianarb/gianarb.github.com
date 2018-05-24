---
layout: post
title:  "Orbiter an OSS Docker Swarm Autoscaler"
date:   2017-04-22 08:08:27
categories: [post]
img: /img/swarm.gif
tags: [docker, influxdb, swarm, orchestration, cloud, kapacitor, chronograf,
tick stack, influxdata, monitoring, open source, oss]
summary: "Orbiter is an open source project design to become a cross provider
autoscaler. At the moment it works like Zero Configuration Autoscaler for Docker
Swarm. It also has a basic implementation to autoscale Digitalocean. This
project is designed with InfluxData a company that provides OSS solution like
InfluxDB, Kapacitor and Telegraf. We are going to use all this tools to create
an autoscaling policy  for your Docker Swarm services."
---
<iframe width="560" height="315" src="https://www.youtube.com/embed/Q1xfmfML8ok"
frameborder="0" allowfullscreen></iframe>
My presentation at the Docker HQ in San Francisco.

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

### Actors

There are different actors in this comedy. First of all our cluster needs to be
manageable by outside via API. In this example I am going to use Docker Swarm,
Orbiter supports a basic implementation for DigitalOcean but it still requires
some toning.

You need to have some time series database or analytics platform that can
trigger webhook to trigger orbiter based on some metrics.

We ran a demo with the TICKStack (InfluxDB, Telegraf, and Kapacitor) days ago.
It's available [at this
link](https://www.influxdata.com/resources/influxdata-helps-docker-auto-scale-monitoring/?ao_campid=70137000000Jgw7).

In the end you need to deploy [orbiter](https://github.com/gianarb/orbiter).

### Orbiter, design and arch

Orbiter is an open source tool designed to be a cross platform autoscaler. It is
in go and it provides a REST API to handle scale requests.

It provides one entrypoint:

```sh
curl -v -d '{"direction": true}' \
    http://localhost:8000/handle/infra_scale/docker
```

* `direction` represent how to scale your service, true means up, false means
  down.
* `/handle/infra_scale/docker` identify the autoscaling group.
  `infra_scale` is the autoscaler name, `docker` is the policy name.

`infra_scale` for example contains information about the cluster manager, where
it is, what is it? Docker or Digitalocean or what ever?

The policy describes how an application scale. If you know a bit Docker Swarm
`docker` is the name of the service.

Orbiter supports two different boot methods. One is via configuration:

```yaml
autoscalers:
  infra_scale:
    provider: swarm
    parameters:
    policies:
      docker:
        up: 4
        down: 3
```

The second one is actually only supported by Docker Swarm and it's called
autodetection. In practice when you start orbiter, it's looking for a Docker
Swarm up and running. If it finds Swarm it's going to list all the services
deployed and it's going to manage all the services labeled with `orbiter=true`.

By default up and down are set to 1 but you can override them with the label
orbiter.up=3 and orbiter.down=2.

Let's suppose to have a Docker Swarm cluster with 3 nodes.

```bash
$ docker node ls
ID                           HOSTNAME  STATUS  AVAILABILITY  MANAGER STATUS
11btq767ecqhelidu8ah1osfp *  node1     Ready   Active        Leader
ptre8d4bjccqi6ml6z445u0mz    node2     Ready   Active
q5rwi3cej9gc1vqyscwfau640    node3     Ready   Active
```

I deployed a service called [gianarb/micro](https://github.com/gianarb/micro).
It is an open source demo application. There are different versions, I deployed
the version 1.0.0. It only shows the current IP of the container/server.

```bash
docker service create --label orbiter=true \
    --name micro --replicas 3 \
    -p 8080:8000 gianarb/micro:1.0.0
```

You can check the number of tasks running with the command:

```bash
$ docker service ps micro
ID                  NAME                IMAGE                 NODE
DESIRED STATE       CURRENT STATE            ERROR
         PORTS
         onsqgriv3nel        micro.1             gianarb/micro:1.0.0   node3
         Running             Running 51 seconds ago

         yxtxyder7bs3        micro.2             gianarb/micro:1.0.0   node1
         Running             Running 51 seconds ago

         lyzxxdc00052        micro.3             gianarb/micro:1.0.0   node2
         Running             Running 52 seconds ago

```

At this point you can visit port `8080` of your cluster to have a look of the
service but for this demo doesn't really matter. We are going to start orbiter
and we are going to trigger a scaling policy to simulate a request made by our
monitoring tool.

```bash
docker service create --name orbiter \
    --mount type=bind,source=/var/run/docker.sock,destination=/var/run/docker.sock \
    -p 8000:8000 --constraint node.role==manager \
    -e DOCKER_HOST=unix:///var/run/docker.sock \
    gianarb/orbiter daemon --debug
```

I am using Docker to deploy orbiter as service. I am using the Unix Socket to
communicate with Docker Swarm and I am deploying this service into the `manager`
because it needs to have write permission to start and stop tasks. This can be
done only into the manager. You can configure orbiter with the variable
`DOCKER_HOST` to use REST API. In this way you don't have this constraint. This
configuration in very easy to show in a demo like this one.

```bash
$ docker service logs orbiter
orbiter.1.zop1qkwa1qxy@node1    | time="2017-04-18T09:24:56Z" level=info
msg="orbiter started"
orbiter.1.zop1qkwa1qxy@node1    | time="2017-04-18T09:24:56Z" level=debug
msg="Daemon started in debug mode"
orbiter.1.zop1qkwa1qxy@node1    | time="2017-04-18T09:24:56Z" level=info
msg="Starting in auto-detection mode."
orbiter.1.zop1qkwa1qxy@node1    | time="2017-04-18T09:24:56Z" level=info
msg="Successfully connected to a Docker daemon"
orbiter.1.zop1qkwa1qxy@node1    | time="2017-04-18T09:24:56Z" level=debug
msg="autodetect_swarm/micro added to orbiter. UP 1, DOWN 1"
orbiter.1.zop1qkwa1qxy@node1    | time="2017-04-18T09:24:56Z" level=info
msg="API Server run on port :8000"
```
As you can see into the logs the API are running on port 8000 and orbiter
already detected a service called `micro`, the one that we deployed before and
it auto-created a autoscaling group called `autodetection_swarm/micro`.
This is the unique name that we can use when we trigger our scale request.

```bash
$ curl -d '{"direction": true}' -v
http://10.0.57.3:8000/handle/autodetect_swarm/micro
*   Trying 10.0.57.3...
* TCP_NODELAY set
* Connected to 10.0.57.3 (10.0.57.3) port 8000 (#0)
> POST /handle/autodetect_swarm/micro HTTP/1.1
> Host: 10.0.57.3:8000
> User-Agent: curl/7.52.1
> Accept: */*
> Content-Length: 19
> Content-Type: application/x-www-form-urlencoded
>
* upload completely sent off: 19 out of 19 bytes
< HTTP/1.1 200 OK
< Content-Type: application/json
< Date: Tue, 18 Apr 2017 09:30:35 GMT
< Content-Length: 0
<
* Curl_http_done: called premature == 0
* Connection #0 to host 10.0.57.3 left intact
```

With that cURL I simulated a scale request and as you can see in the log above
orbiter detected the request and it scaled up 1 task for our service called
`macro`

```bash
$ docker service logs orbiter
orbiter.1.zop1qkwa1qxy@node1    | POST /handle/autodetect_swarm/micro HTTP/1.1
orbiter.1.zop1qkwa1qxy@node1    | Host: 10.0.57.3:8000
orbiter.1.zop1qkwa1qxy@node1    | Accept: */*
orbiter.1.zop1qkwa1qxy@node1    | Content-Length: 19
orbiter.1.zop1qkwa1qxy@node1    | Content-Type:
application/x-www-form-urlencoded
orbiter.1.zop1qkwa1qxy@node1    | User-Agent: curl/7.52.1
orbiter.1.zop1qkwa1qxy@node1    |
orbiter.1.zop1qkwa1qxy@node1    | {"direction": true}
orbiter.1.zop1qkwa1qxy@node1    | time="2017-04-18T09:30:35Z" level=info
msg="Received a new request to scale up micro with 1 task." direc
tion=true service=micro
orbiter.1.zop1qkwa1qxy@node1    | time="2017-04-18T09:30:35Z" level=debug
msg="Service micro scaled from 3 to 4" provider=swarm
orbiter.1.zop1qkwa1qxy@node1    | time="2017-04-18T09:30:35Z" level=info
msg="Service micro scaled up." direction=true service=micro
```

We can verify the current number of tasks that are running for `micro` and we
can see that it's not 3 as before but 4.

```bash
$ docker service ls
ID                  NAME                MODE                REPLICAS
IMAGE
azi8zyeor5eb        micro               replicated          4/4
gianarb/micro:1.0.0
ezklgb6uak8b        orbiter             replicated          1/1
gianarb/orbiter:latest
```

This project is open source on
[github.com/gianarb/orbiter](https://github.com/gianarb/orbiter) you can have a look on
it, try and leave some feedback or request if you need something different.

PR are also open if you are working with a different cluster manager or with a
different provider, add a new one is very easy. It's just a new interface to
implement.

