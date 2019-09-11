---
img: /img/gianarb.png
layout: post
title: "Control Theory is dope"
date: 2019-09-04 08:08:27
categories: [post]
tags: [software design, design pattern]
summary: "This is an introductive article about control theory applied to
microservices and cloud computing. It is a very high level overview about
control theory driven by what I loved most about it."
changefreq: daily
---
For the last two years at InfluxData I worked on our custom orchestrator that
empower InfluxCloud v1 to run. I have some talk about it at InfluxDays, but they
are not recorded so, I can't really post them here, sadly.

If you are thinking: "Why you should write your own orchestrator?", I have few
answers for you.

1. Back in the day Kubernetes was not so popular, 4 years ago when InfluxCloud
   started it was not at least.
2. We had since the beginning to manage data and state, people still says that
   Kubernetes is not for them today, image how it was 4 years ago.

Btw now InfluxCloud v2 leverages Kubernetes.

Writing a good orchestrator is super fun! When I started but still today a big
part of it are frustrating and not so good but the one we wrote following
reactive planning and control theory are lovely! This article is an introduction
about Control Theory. [Chris Goller](https://twitter.com/goller) Solution
Architect at InfluxData was the first person that told me about how Control
Theory works in theory, and he pushed me to try reactive planning for our
orchestrator.

As Kubernetes contributor I recognized some of those patterns as looking at
shared informers, controller and so on. So I understood since the beginning that
those patterns was everywhere around me!

[Colm MacCárthaigh](https://twitter.com/colmmacc) from Amazon Web Service with
his  talks (like the one posted here) helped me to find resources to read, more
patterns and use cases for it.


<div class="embed-responsive embed-responsive-16by9 col-xs-12 text-center">
    <iframe width="560" height="315" src="https://www.youtube.com/embed/O8xLxNje30M"
    frameborder="0" allow="accelerometer; autoplay; encrypted-media; gyroscope;
    picture-in-picture" allowfullscreen></iframe>
</div>

## Why it works

When I started to work as a Web Developer, designing APIs or websites I had
different challenges to face. To write a solid CRUD you put all your effort
when a request comes to your API, you validate it, apply transformation to
sanitize the input and if it is valid you save it
in your database. You need to build good UX, complex validations systems and so
on. But what lands in the database is right and rock solid.

There are other systems where you do not have a database that tells you what is
right or not. You need to **measure** the current state, **calculate** what
needs to get back to your desired state and you need to **apply** what you
calculated.

Those systems are everywhere:

* The boiler you have at home to keep the water warm needs to constantly check if the
  desired temperature you set is the current one. What it is stored in its
  memory is what you desire, not the truth.
* The example Colm MacCárthaigh used is the Autoscaler. It keeps checking the
  state of your system based on the scalability rules you set. For example if
  CPU is over 70% spin up 3 nodes. The autoscaler measures the current state of
  your CPUs and when it is over it calculates what needs to be done and it
  executes the scale up or down.
* When you read Kubernetes documentation is will see reference to Controller,
  reconciliation loop, desired state and so on. All of those concepts come from
  Control Theory.

Orchestrator but more in general big microservices environment do not have the
concept of data locality as we used to have in the past. The data you need can
change continously, and they need to collected from different sources and
combined in order to calculate what needs to be done.

I think this is the main reason about why patterns coming from Control Theory
works well.

If you need to write a program that provisions 3 virtual machines and attach them
to a random DNS record you can approach this problems in 2 ways. You can write a
procedure that:

1. Creates 3 instances.
2. Takes the public IPs.
3. Creates the DNS record with the IPs as A record.

Another way you have to fix this issue is to start from checking what you have,
making a plan to matches what it is not as you desire. So it will look like
this:

1. Check how many instances there are and mark what you need to do, if there are
   2 of them you need one, if there are 5 you need to delete 2, you there are 0
   of them you need to create all of them.
2. Check if the DNS record is already there and how many IPs are assigned to it.
3. If it does exist you do not need to create it but you need to check if the
   IPs assigned to it are the same of the instances, If they are not you need to
   reconcile the DNS record fixing the IPs.
4. The record does not exist? You can create it.

If you are wondering how all those checks makes the system more reliable is because
you never know what you already created or what it is already where. Let's
assume you are on AWS. API requests can fails at the middle of your process and
you need to know where you are. AWS itself can stop or terminate instances, or
some other procedures can do it or for manual mistake.

Approaching the problem in this way allows you to repeat the flow over and over
because it idempotent and at every retry the process will be able to reconcile
any divergence between what you asked for (3 VMs and one DNS record) and what it
is actually running. This process is called reconciliation loop.

## 101 architecture

Colm MacCárthaigh highlights three major areas around how a successful Control
Theory implementation looks like:

1. Measurement process
2. Controller
3. Actuator

## Measurement process

The way you retrieve the current state of the system is crucial in order to have
a low latency. They are crucial in order to calculate what needs to be done
because from the current state your program get different decisions.

## Controller

This section is where I have more experience with. The desired state is stored
and clear usually. You know here to go. You get the measurements and with this
information you need to write a procedure capable of making a plan stating from
your current state to get to the desired one.

I wrote a few weeks ago an introduction about [reactive
planning](https://gianarb.it/blog/reactive-planning-is-a-cloud-native-pattern)
it is the way I used to calculate a plan.

I am also preparing a PoC in Golang with actual code you can run and test to
share in practice what means reactive planning.

## Actuator

It is the part that take a calculated plan, and it executes it. I worked a lot
with schedulers that are able to take a set of steps and execute them one by one
or in parallel based on needs.

## Conclusion

Think about one of them problem you have a try to think in a more reactive way,
starting from checking where you are and not from doing things. Reliability and
stability for your code will improve drastically.
