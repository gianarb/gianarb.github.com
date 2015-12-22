---
layout: post
title:  "Docker and wordpress for a better world"
date:   2015-12-14 10:08:27
categories: devops
img: /img/docker.png
tags: devops, docker
summary: "Docker and wordpress to guarantee scalability, flexibilty and isolation.
A lot of webagencies install all wordpress in the same server but how can they manage a disaster?
AWS with Elastic Container Service could be a more professional solution."
priority: 0.6
changefreq: yearly
---

<blockquote class="twitter-tweet tw-align-center" lang="en"><p lang="en" dir="ltr"><a href="https://twitter.com/hashtag/docker?src=hash">#docker</a> and <a href="https://twitter.com/hashtag/wordpress?src=hash">#wordpress</a> for a better world.. <a href="https://t.co/o9c6YXvsl3">https://t.co/o9c6YXvsl3</a> Blogpost after my talk <a href="https://twitter.com/CodemotionIT">@CodemotionIT</a> How and Why? <a href="https://twitter.com/awscloud">@awscloud</a></p>&mdash; Gianluca Arbezzano (@GianArb) <a href="https://twitter.com/GianArb/status/679241680797700096">December 22, 2015</a></blockquote>
<script async src="//platform.twitter.com/widgets.js" charset="utf-8"></script>

I am trying to represent a typical wordpress infrastructure

![Wordpress typical infrastructure](/img/posts/2015-12-16/wp-infra.png)

**Isolation**: every single wordpress share all with the others, filesystem,
memory, database.

This lack of isolation causes different problems:

* The monitoring of each installation is harder.
* We share security problems
* We don't have the freedom to work without the fear or blocking 100 customers

We are overwhelmed by the problems

![Problem](/img/posts/2015-12-16/problem.png)

## [LXC Container](https://github.com/opencontainers/runc)

> it is an operating-system-level virtualization environment for running multiple
> isolated Linux systems (containers) on a single Linux control host.
>
> by wikipedia

Wikipedia helps me to resolve one problem (theory), container is **isolated
Linux System**

## Docker

Docker wrap LXC container to serve your application ready to go in an isolate
environment, with own filesystem and dependencies.

Worpdress in this implemetation has two containers, one to provide apache and
php and one for mysql database.  This is an example of Dockerfile, it describes
how a docker container works it is very simple to understand, from this example
there are different keywords

* `FROM` describes the image that we use as start point.
* `RUN` run a command.
* `EXPOSE` describes ports to open during a link, in this case MySql runs on
  the default port 3306.
* `CMD` is the default command used during the run console command.

{% highlight bash %}
FROM ubuntu
RUN dpkg-divert --local --rename --add /sbin/initctl
RUN ln -s /bin/true /sbin/initctl
RUN echo "deb http://archive.ubuntu.com/ubuntu precise main universe" > /etc/apt/sources.list
RUN apt-get update
RUN apt-get -y install mysql-server
EXPOSE 3306
CMD ["/usr/bin/mysqld_safe"]
{% endhighlight %}

Very easy to read, it is a list of commands!
We are only write a container definition, now we can build it!

{% highlight bash %}
docker build -t gianarb/mysql .
{% endhighlight %}

In order to increase the value of this article and to use stable images I will
use the official [mysql](https://hub.docker.com/_/mysql/) and
[wordpress](https://hub.docker.com/_/wordpress/) images.

Download this images
{% highlight bash %}
docker pull wordpress
docker pull mysql
{% endhighlight %}

We are ready to run all! Dockerfile is only a way to describe each single
container, and the pull command downloads online container ready to work, it is
a good way to reuse your or other containers.

We downloaded mysql and wordpress, with the run command we start them and we
define our connections

{% highlight bash %}
docker run \
    --name mysql \
    -p 3306:3306 \
    -e MYSQL_ROOT_PASSWORD=passwd  mysql

docker run -e WORDPRESS_DB_HOST=wp1.database.prod \
    -e WORDPRESS_DB_USER=root \
    -e WORDPRESS_DB_PASSWORD=help_me \
    -p 8080:80 \
    -d --name wp1 \
    --link wp.database.prod:mysql wordpress
{% endhighlight %}

I can try to explain this commands, it run two containers:

* The name of the first container is mysql and it uses the `mysql` image, we
  use -p flag to expose mysql port now you can use phpmyadmin or other client
  to fetch the data but remember that is not a good practice.
* The second container called wp1 uses the image `gianarb/wordpress` forward
  the container port 80 (apache) on host 8080, that in this case it is the way
  to see the site.  --link flag is the correct way to consume mysql outside the
  main container, in this particular case we could use wp.database.prod how url
  to connect at mysql from our worpdress container, awesome!
* Docker image supports environment variable `ENV` for example we can use them
  to configure our services, in this case to set root password in mysql and to
  configure worpdress's database connection

We are ready! Now you have a worpdress ready to go on port 8080.

## Docker Compose
To save time and to increase reusability we can use
[docker-compose](https://docs.docker.com/compose/) tool
that helps us to manage multi-container infrastructures, in this case one for
mysql and one for wordpress.
In practice we can describe all work did above in a `docker-compose.yml` file:

{% highlight yaml %}
wp:
  image: wordpress
  ports:
    - 8081:80
  environment:
      WORDPRESS_DB_HOST: wp1.database.prod
      WORDPRESS_DB_USER: root
      WORDPRESS_DB_PASSWORD: help_me
  links:
    - wp1.database.prod:mysql
mysql:
  image: mysql:5.7
  environment:
    MYSQL_ROOT_PASSWORD: help_me
{% endhighlight %}

Now we can run

{% highlight bash %}
docker-compose build .
docker-compose up
{% endhighlight %}

To prepare and start our infrastructure. Now we have one wordpress with own
mysql that run on port 8081. We can change wordpress port to start new isolate
wordpress installation.

<p class='text-center'>
<iframe src="//giphy.com/embed/l41lYCDgxP6OFBruE" width="480" height="268"
frameBorder="0" class="giphy-embed" allowFullScreen></iframe><p><a
href="http://giphy.com/gifs/foxtv-win-ricky-gervais-emmys-2015-l41lYCDgxP6OFBruE">via
GIPHY</a></p>
</p>

## In Cloud with AWS ECS
We won a battle but the war is too long, we can not use our PC as server.  In
this article I propose [AWS Elastic Container
Service](http://docs.aws.amazon.com/AmazonECS/latest/developerguide/Welcome.html)
a new AWS service that helps us to manage containers, why this service? Because
it is Docker and Docker Composer like, it's managed by AWS, maybe there are
more flexible solutions, Swarm, Kubernetes but it is a good start point.

![AWS Elastic Container Service](/img/posts/2015-12-16/ecs.png)

A services of keywords to understand how it works:

* **Container instance**: An Amazon EC2 that is running the Amazon ECS Agent. It has been registered into the ECS.
* **Cluster**: It is a pool of Container instances
* **Task definition**: A description of an application that contains one or more container definitions
* Each Task definition running is a **Task**

### In practice

1. Create a cluster

{% highlight bash %}
ecs-cli configure \
    --region eu-west-1 \
    --cluster wps \
    --access-key apikey \
    --secret-key secreyKey
{% endhighlight %}

2. Up nodes (one in this case)
{% highlight bash %}
ecs-cli up --keypair key-ecs \
    --capability-iam \
    --size 1 \
    --instance-type t2.medium
{% endhighlight %}

3. Push your first task!
{% highlight bash %}
ecs-cli compose --file docker-compose.yml  \
    --project-name wp1 up
{% endhighlight %}

4. Follow the status of your tasks
{% highlight bash %}
ecs-cli ps
{% endhighlight %}

You can use another docker-compose.yml with a different wordpress port to build
another task with another worpdress!

## Now is only a problem of URL
We are different isolated worpdress online, but they are an ip and different
ports, maybe our customers would use a domain name for example.
I don't know if this solution is ready to run in production and it is good to
run more and more wordpress but a good service to turn and proxy requests is
HaProxy. This is an example of configuration for our use case:

wp1.gianarb.it and wp1.gianarb.it are two our customers and 54.229.190.73:8080,
54.229.190.73:8081 are our wordpress.

{% highlight bash %}
...
frontend wp_mananger
        bind :80
        acl host_wp1 hdr(host) -i wp1.gianarb.it
        acl host_wp2 hdr(host) -i wp2.gianarb.it
        use_backend backend_wp1 if host_wp1
        use_backend backend_wp2 if host_wp2
backend backend_wp1
        server server1 54.229.190.73:8080 check
backend backend_wp2
        server server2 54.229.190.73:8081 check
{% endhighlight %}

Note: This configuration increase the scalability of our system, because we can
add other service in order to support more traffic.

{% highlight bash %}
backend backend_wp1
        server server1 54.229.190.73:8080 check
        server server1 54.229.190.12:8085 check
        server server1 54.229.190.15:80 check
{% endhighlight %}

### There are other solutions
* Nginx
* Consul to increase the stability and the scalability of our endpoint

<div class="alert alert-info" role="alert">
This article is based on my presentation at <a href='http://gianarb.it/codemotion-2015/' target='_blank'>Codemotion 2015</a>
</div>

<div class="alert alert-success" role="alert">
Thanks for review <a href='https://twitter.com/fntlnz'
target='_blank'>Lorenzo</a>! I'm in Ireland from 3 weeks but I am not ready to
write an article without your english review!
</div>
