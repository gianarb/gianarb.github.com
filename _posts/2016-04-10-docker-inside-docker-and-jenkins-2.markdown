---
layout: post
title:  "Docker inside docker and overview about Jenkins 2"
date:   2016-04-1 10:08:27
categories: [post]
img: /img/docker.png
tags: [docker, ci]
summary: "A little overview about Jenkins 2 but the main topic of the article
is about how run docker inside docker to start a continuous integration system
inside a container"
priority: 0.6
changefreq: yearly
---

<blockquote class="twitter-tweet tw-align-center" data-lang="en"><p lang="en" dir="ltr"><a
href="https://twitter.com/hashtag/docker?src=hash">#docker</a> inside docker
and an overview about Jenkins 2 <a
href="https://t.co/qa5ddjfhrs">https://t.co/qa5ddjfhrs</a> <a
href="https://twitter.com/docker">@docker</a> <a
href="https://twitter.com/jenkinsci">@jenkinsci</a> <a
href="https://twitter.com/hashtag/container?src=hash">#container</a></p>&mdash;
Gianluca Arbezzano (@GianArb) <a
href="https://twitter.com/GianArb/status/727876226875068416">May 4,
2016</a></blockquote> <script async src="//platform.twitter.com/widgets.js"
charset="utf-8"></script>

Jenkins  is one of the most famous
continuous integration and deployment tools, it’s written in Java and it helps
you to manage your pipeline and all tasks that help you to put your code in
production or manage your build.

The announcement of Jenkins release of version 2 few days ago, is one of the
best release of this year in my opinion.

The previous version is very stable but it has a lot of years and the ecosystem
is totally different. I am happy to see a strong refurbishment to get the best
of this powerful tool with a series of new feature like:

* Nice installation wizard
* Refactoring of the design, one of the most critical
  feature of the previous version
* Good and modern set of plugins like [Jenkins
  Pipeline](https://jenkins.io/solutions/pipeline/) to manage your build

Jenkins is truly a wonder but the tool of the moment it’s docker, engine
that allow you to work easier with the containers.

This two tools together are perfect to create an isolated environment to test
and deploy your applications.

The first setup could be install Jenkins on your
server and use a plugin to manage the integration and trigger your test inside
an isolated environment, the container.

Great work but in my opinion reproducibility is one of the critical point when
you deal with plugins if you can not run your build on your local environment
easily then you have a problem.  Secondly if the container could be a good
solution to deploy and maintain a solid and isolated application, why your
Jenkins has not the privilege to run inside a container?  In this perspective
how can we run container inside a container?

Ok, now its the time to figure it out how to solve the problems.

We can use the official Jenkins image to put jenkins inside a container, but I
worked on my personal alpine installation, light and easy, [here is the
dockerfile](https://github.com/gianarb/dockerfile/blob/master/jenkins/2.0/Dockerfile)
and we can pull it:

{% highlight bash %}
docker pull gianarb/jenkins:2.0
{% endhighlight %}

If you are interested the main article to understand how run docker inside
docker is written by
[jpetazzo](https://jpetazzo.github.io/2015/09/03/do-not-use-docker-in-docker-for-ci/),
the idea is run our jenkins container with `-privileged` enabled and share our
docker binary and the socket `/var/run/docker.sock` to manage our
communications.

* `/var/run/docker.sock` is the entrypoint of the docker daemon
* `docker` the command is like a client that sends commands to socket
*  `--privileged` give extended privileges to our container

Translated in a docker command:

{% highlight bash %}
docker run -v /var/run/docker.sock:/var/run/docker.sock \
    -v $(which docker):/usr/local/bin/docker \
    -p 5000:5000 -p 8080:8080 \
    -v /data/jenkins:/var/jenkins \
    --privileged \
    --restart always \
    gianarb/jenkins:2.0
{% endhighlight %}

We connect on `http://docker-ip:8080` and start the new awesome wizard!

<img class="img-responsive" alt="First Jenkins 2 page, grab from the log your key and start" src="/img/docker-in-docker/jenkins2-start.png">

<img class="img-responsive" alt="Jenkins's plugins wizard" src="/img/docker-in-docker/jenkins2-plugin.png">

To verify that all work we can create a new job it only runs `docker ps -a` our
expectation is the same list of containers that we have out of jenkins.

<img class="img-responsive" alt="Result of the first build" src="/img/docker-in-docker/jenkins2-result.png">

Now we can use run command from jenkins to manage our build with docker without
any kind of plugins but anyway you are free to use [Docker
Plugin](https://wiki.jenkins-ci.org/display/JENKINS/Docker+Plugin) to start
your build.

I used Jenkins like an example to run docker inside another container but you
can use the same strategy to do the same with your applications if they require
a strong connection with docker.

{% include docker-planet-newsletter.html %}
