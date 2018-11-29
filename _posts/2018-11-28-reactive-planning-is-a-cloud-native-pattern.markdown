---
heroimg: /img/hero/timeclock.jpg
layout: post
title:  "Reactive planning is a cloud native pattern"
date:   2018-11-28 08:08:27
categories: [post]
tags: [software design, cost, cloud, scale, design pattern, reactive planning,
cloud native]
summary: "I discovered how a reactive plan works recently during a major
refactoring for a custom orchestrator that we write at InfluxData to serve our
SaaS offer. In this article I will explain why I think reactive planning is
perfect to build cloud native application like container orchestrators and
provisioning tools."
changefreq: daily
---
Probably this title can sound a bit weird to anyone that already know what
reactive plan is and how far it can look from all the cloud-native and
distributed system hipster movement but recently one of my colleagues [Chris
Goller](https://twitter.com/goller) pushed this pattern to one of the projects
that we have at [InfluxData](https://influxdata.com) and I find it glorious!

“In artificial intelligence, reactive planning denotes a group of techniques for
action selection by autonomous agents. These techniques differ from classical
planning in two aspects. First, they operate in a timely fashion and hence can
cope with highly dynamic and unpredictable environments. Second, they compute
just one next action in every instant, based on the current context.”
([Wikipedia](https://en.wikipedia.org/wiki/Reactive_planning))

The Wikipedia definition of reactive planning as you can see is perfect to
handle a system where the current status can change very frequently based on
external and unpredictable events.

This is a perfect approach for provisioning/orchestrator tool like Mesos, Cloud
Formation, Kubernetes, Swarm, Terraform. Some of them are working like this
already.

The general idea is that before any action you need a plan because for these
tools an action means: cloud interaction, spin up of resources that cost money.
You need to be proactive avoiding useless execution.

A plan is made of a serious of steps and every step can return other steps if it
needs. The plan is complete when there are no steps anymore.  The plan gets
executed at least twice, the second time it should return zero steps because the
first attempts built everything you need, this is the signal that determines its
conclusion. If it keeps returning steps it means that there is something to do
and it tries again.

Let’s start with an example. Think about what Cloud Formation does. You can
declare a set of resources and before to take action it needs to understand what
to do. It is making a plan checking the current state of the system. This first
part makes the flow idempotent and solid because you always start from the
current state of the system. It doesn’t matter if it changes over time because
of somebody that removed one of the resources. If something doesn’t exist it
creates or modify it. Very solid.

Every single step is very small. Let’s take another example like creating a pod
in Kubernetes. When you create a pod there are a lot of actions to do:

* Validation
* Generate the pod id, the pod name
* Register the pod to the DNS, you
* Store it to etcd
Reach to CNI to configure the network
* Reach to docker, container or whatever you use to get the container
* Maybe reach to AWS to create a persistent volume
* Attach the PV

If you try to design all interaction in a single “controller” you will end up
with a lot of * if/else, error handling and so on. Mainly because as you can see
almost every step interact over the network with something: database, DNS, CNI,
docker and so on. So it can fail, it needs circuit breaking, retry policy and
much more complexity.

It is a lot better to design the code where every point is a small step if the
step that reaches docker fails it can return itself as “retry” or it can return
other steps to abort everything and clean up. You will end up with small
reusable (or not that much reusable) steps.

All the steps are combined within a plan,  the “PodCreation” plan. There is a
scheduler that takes and execute every step in the plan recursively.

> This freedom allows you to use an incremental approach

The scheduler as first call a create method for the plan, the create method
checks what to do based on the current state of the system, it is the
responsibility of this function to return no steps when there is nothing to do.

I think this Reactive Planning is one of the best ways to organize the code in a
cloud-native ecosystem for its reactive nature as I said and for the fact that
it forces you to check the state of the system if you don’t do that the plan
will keep executing forever.  Obviously, you can use a high-level check to skip
a lot of steps, this requires a balance if the plan you are executing is
critical and frequently used you should check for every step if it requires an
effort that won’t pay back you can implement deepest and preciser checks. You
can check for the PodStatus. If it is running we are good nothing to do. Or you
can check if Docker has a container running and if it has the right network
configuration. If it is running but with no network, you can return the step
that interacts with CNI to set the right interface. This freedom allows you to
use an incremental approach, you start with an easy creation method with checks
for only critical and high-level signal demanding a more solid and sophisticated
set of checks for later, when you will have the best knowledge about where
the system fails.

{:.small}

Hero image via
[Pixabay](https://pixabay.com/en/time-time-management-stopwatch-3222267/)
