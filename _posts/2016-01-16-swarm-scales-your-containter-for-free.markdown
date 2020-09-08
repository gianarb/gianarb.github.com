---
layout: post
title:  "Swarm scales docker for free"
date:   2015-12-14 10:08:27
categories: [post]
img: /img/docker.png
tags: [devops, docker]
summary: "Docker is an awesome tool to manage your container.
Swarm helps you to scale your containers on more servers."
priority: 0.6
changefreq: daily
---
<blockquote class="twitter-tweet tw-align-center" data-lang="en"><p lang="en" dir="ltr">An ocean of containers! With docker and swarm.. <a href="https://t.co/1dXoZYS3ZA">https://t.co/1dXoZYS3ZA</a> <a href="https://twitter.com/hashtag/docker?src=hash">#docker</a></p>&mdash; Gianluca Arbezzano (@GianArb) <a href="https://twitter.com/GianArb/status/696620821931036672">February 8, 2016</a></blockquote>
<script async src="//platform.twitter.com/widgets.js" charset="utf-8"></script>

[Gourmet](https://github.com/gianarb/gourmet) is a work in progress application
that allow you to execute little applications on an isolated environment, it
dowloads your manifest and runs it in a container.
I start this application to improve my go knowledge and to work with the Docker API
I am happy to share my idea and my tests with Swam an easy way to scale this type of application.

Gourmet exposes an HTTP API available at the `/project` endpoint that accept a JSON request body like:

{% highlight json %}
{
    "img": "gourmet/php",
    "source": "https://ramdom-your-source.net/gourmet.zip",
    "env": [
        "AWS_KEY=EXAMPLE",
        "AWS_SECRET=",
        "AWS_QUEUE=https://sqs.eu-west-1.amazonaws.com/test"
    ]
}
{% endhighlight %}

* `img` is the started point docker image
* `source` is your script
* `env` is a list of environment variables that you can use on your script

During my test I use this [php script](https://github.com/gianarb/gourmet-php-example) that send a message on SQS.

Your script has a console entrypoint executables in this path `/bin/console` and
gourmet uses it to run your program.

To integrate it with Docker I used `fsouza/go-dockerclient` an open source
library written in go.

{% highlight go %}
container, err := dr.Docker.CreateContainer(docker.CreateContainerOptions{
    "",
    &docker.Config{
        Image:        img,
        Cmd:          []string{"sleep", "1000"},
        WorkingDir:   "/tmp",
        AttachStdout: false,
        AttachStderr: false,
        Env:          envVars,
    },
    nil,
})
{% endhighlight %}

This is a snippet that can be used to create a new container.
With the container started I use the exec feature to
extract your source and to run it.

{% highlight go %}
exec, err := dr.Docker.CreateExec(docker.CreateExecOptions{
    Container:    containerId,
    AttachStdin:  true,
    AttachStdout: true,
    AttachStderr: true,
    Tty:          false,
    Cmd:          command,
})

if err != nil {
    return err;
}

err = dr.Docker.StartExec(exec.ID, docker.StartExecOptions{
    Detach:      false,
    Tty:         false,
    RawTerminal: true,
    OutputStream: dr.Stream,
    ErrorStream:  dr.Stream,
})
{% endhighlight %}

After each build Gourmet cleans all and destroies the environment.

{% highlight go %}
err := dr.Docker.KillContainer(docker.KillContainerOptions{ID: containerId})
err = dr.Docker.RemoveContainer(docker.RemoveContainerOptions{ID: containerId, RemoveVolumes: true})
if(err != nil) {
    return err;
}
return nil
{% endhighlight %}

At the moment it is gourmet, It could be different hypothetical use cases:

* high separated task
* run a testsuite
* dispatch specific functions

A microservice to work with docker container easily.

I thought about an easy way to scale this application and I found
[Swarm](https://docs.docker.com/swarm/), it is a native cluster for docker and
it seems awesome in first because  it is compatibile with the docker api.

## Swarm
A Docker Swarm's cluster is very easy to setup, I worked on this project
[vagrant-swarm](https://github.com/gianarb/vagrant-swarm) to create a local
environment but [the official
documentation](https://docs.docker.com/swarm/install-manual/) is easy to follow.

Swarm's cluster has two actors:
* A master is the entrypoint of your requests, it provide an HTTP
  api compatible with docker.
* A series of nodes that communicate with the master.

During this example we will work with 1 master and 2 nodes.
Build this machine with virtualbox , with another tool, or in cloud is not a
problem and [install docker](https://docs.docker.com/engine/installation/).

Into the master pull swarm and create a cluster identifier.

{% highlight bash %}
docker pull swarm
docker run --rm swarm create
docker run --name swarm_master -d -p <manager_port>:2375 swarm manage token://<cluster_id>
{% endhighlight %}

`swarm create` returns a cluster_id use them to start the manager and the
`manager_ip` is the ip of your master server.

Now go into the node, because we must do few things.

{% highlight bash %}
docker daemon -H tcp://0.0.0.0:2375 -H unix:///var/run/docker.sock
docker run -d swarm join --addr=<node_ip:2375> token://<cluster_id>
{% endhighlight %}

When `cluster_id` is the id created in the previous step and the `node_id` is the ip
of  your current node.
Enter into the master and restart your manager container

{% highlight bash %}
docker restart swarm_master
{% endhighlight %}

Now we are ready to test if all it's up.

{% highlight bash %}
docker -H tcp://0.0.0.0:2375 info
{% endhighlight %}

Replace `0.0.0.0.0` with your master ip if you are in the same server.
You'll wait this type of response

{% highlight bash %}
$. sudo docker -H tcp://192.168.13.1:2375 info
Containers: 1
Images: 1
Role: primary
Strategy: spread
Filters: health, port, dependency, affinity, constraint
Nodes: 2
 vagrant-ubuntu-vivid-64: 192.168.13.101:2375
  └ Status: Healthy
  └ Containers: 1
  └ Reserved CPUs: 0 / 1
  └ Reserved Memory: 0 B / 513.5 MiB
  └ Labels: executiondriver=native-0.2, kernelversion=3.19.0-43-generic, operatingsystem=Ubuntu 15.04, storagedriver=aufs
 vagrant-ubuntu-vivid-64: 192.168.13.102:2375
  └ Status: Healthy
  └ Containers: 0
  └ Reserved CPUs: 0 / 1
  └ Reserved Memory: 0 B / 513.5 MiB
  └ Labels: executiondriver=native-0.2, kernelversion=3.19.0-43-generic, operatingsystem=Ubuntu 15.04, storagedriver=aufs
CPUs: 1
Total Memory: 513.5 MiB
Name: f5e23167339e
{% endhighlight %}

Gourmet is a set of environment variables to create a connection with docker
api, in particular this function
[NewClientFromEnv](https://godoc.org/github.com/fsouza/go-dockerclient#NewClientFromEnv)
and the `DOCKER_HOST` parameter.

Docker Swarm supports the same Docker API in this way gourmet uses more nodes.
{% highlight bash %}
$ DOCKER_HOST="tcp://192.168.13.1:2333" ./gourmet api
{% endhighlight %}
