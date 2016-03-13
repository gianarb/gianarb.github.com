---
layout: post
title:  "Some days of work vs Jenkins CI"
date:   2016-02-21 10:08:27
categories: devops
img: /img/jenkins.png
tags: jenkins, continuous integration, go
summary: "I love Jenkins CI is a beautiful and stable project to run job and
manage continuous integration and deploy pipeline, few days ago I worked to
improve the delivery pipeline in CurrencyFair and I started to do some thought
about this topic, here my internal battle vs Jenkins CI"
priority: 0.6
changefreq: yearly
---
Guys please move down your hands, I love JenkinsCI! I am not here to write a
bad post about it!  
I am here to share few days of reasonings about continuous
inteagrations, Jenkins CI and all this strong topic.  

## Reproducible
There are a lot of tools that you can use to run tasks,
ant, make, grunt. Use it to run section or all your build on your local
environment, this appraoch increase the value your tasks becuase you use and
test them more times.
To have a reproducible build help you to maintain splitted your flow by your
runner maybe Jenkins is perfect but the are other tools and services:
Travis-CI, CircleCI, Drone, don't create a big dependency with your environment.

## Speedy
A slow test suite is a bad idea, in 1 minutes I can maintain the
focus on the execution but 5 minutes are a lot, you can take a coffé or
start to think about another task and return on your old task require an
effort. This several focus switch it's not good, and at the same time you
lost 5 minutes for each build and for every engineer after 1 week are a lot
of money.

## Versionable
I lost more time about this point and I am not sure if it is a required point
or not but TravisCI for example use a yml specification file, this file doesn't
describe only your build but it part of the story of your application, if you
include it into the VCS. Could be it a value for your pipeline?

## Maintainable
There are a lot of tools that you can choice to create the perfect pipeline ant
it's very easy lost your focus and start to use too much tools, you must try
all but it's your task to create the perfect sub-set of tools the point 1
(Reusability) increase the value, use tools that you can reuse during the daily
work of your team to increase the develop and go the flow better.  
Each tools that you add seems perfect until they don't becase a problem.

# Scalable
An easy way to decrease the time for your job is split it in different little
jobs and run them in parallel, you can check the codestyle and run your test
suite in the same time for example.
Another good reason to create a scalable environment for your jobs is because
your company would grow and the continuous integration system burns to helps it
to grow and not to stop it.

# Unique
Jenkins, vagrant, ant, make, drone, docker are only a list of amazing tools to
create the perfect pipeline to deploy and test your code but they are only a
means the goal is indeed the best pipeline for your code and for your team.
Observe how your team works, which the requirments and criticalities and design
the best pipeline for your use case.

## Communication layer
One goal for your team is understand the status of the build without logged in
any application, because enter into the Jenkins site (at first because it is
not beautiful :P ) and it is another step to do other: create feature branch,
submit pull request, write code lalala..  
Use directly the pull request to create a connection with your job, you
continuous integration system can submit a new comment or if you are working
with GitHub you can use the status check, in this way you can help your
colleagues during them work and remove a jump.

With JenkinsCi you can do all but if you lost more time to create your best
pipeline? Maybe you don't know it or maybe it is not the best tool for your use
case. Jenkins is flexible, but the flexibility is only the number of plugins
that you can install?

I don't know I use it but I am happy to experiment and there are a lot of new
technologies and tools that maybe can help us to do a good work, with or
without JenkinsCI.  

## Slimmer, proof of concept
I tried to create a runner for my test suite, [slimmer](https://github.com/gianarb/slimmer),
to implement this thought with docker and go.
Go offers a lot of libraries and tools to create something in a bit of time and
docker it's perfect because it creates isolated environment and it's very easy
to scale with Swarm.  
In practice at the moment this console app exec a `build.slimmer` a bash script
executable flexible and versionable.  
[TravisCI](https://travis-ci.org) is powerful but the YML file is it a good way
to describe a build? It's flexible? Maybe yes but I am curious to try a "low
level" approach, because finally all becames a series of commands.  
I created also a series of agent to trigger notification quicly:
[ircer](https://github.com/gianarb/ircer),
[slacker](https://github.com/gianarb/slacker).  You can use them to notify the
result of your build.

{% highlight bash %}
composer install
vendor/bin/phpunit
RESULT=$?
curl -LSs https://github.com/gianarb/ircer/releases/download/0.1.0/ircer_0.1.0_linux_386 > ircer
chmod 755 ircer
if [ $RESULT = 0 ]; then
    ./ircer -j tech-team -m "You are a great develop. Your build works"
else
    ./ircer -j tech-team -m "No bad but your build doesn't work"
fi
{% endhighlight %}

This is an example of `build.slimmer` with an IRC notification, it is a PoC and
I prepared a little [presentation](http://gianarb.it/slimmer-poc-slide/#/) to
receive some feedback and I presented it during a Dublin Go Meetup

<div class="row">
    <div class="col-md-12 text-center">
        <iframe width="560" height="315" src="https://www.youtube.com/embed/CWCHT3GClMM" frameborder="0" allowfullscreen></iframe>
    </div>
</div>

I wait some feedback if you are interested about continuous integration and
continuous delivery.
