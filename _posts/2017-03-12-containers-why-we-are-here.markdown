---
layout: post
title:  "Containers why we are here"
date:   2017-03-12 08:08:27
categories: [post]
img: /img/container-security.png
tags: [docker, scaledocker]
summary: "Everything unnecessary in your system could be a very stupid vulnerability. We
already spoke about this idea in the capability chapter and  the same rule
exists when we build an image. Having  tiny images with only what our
application needs to run is not just a goal in terms of distribution but also
in terms of cost of maintenance and security."
changefreq: yearly
---

> “It is change, continuing change, inevitable change, that is the dominant
> factor in society today. No sensible decision can be made any longer without
> taking into account not only the world as it is, but the world as it will be...
> This, in turn, means that our statesmen, our businessmen, our everyman must take
> on a science fictional way of thinking”  Asimov, 1981

# Isolation and Virtualization

I can see clearly two kind of invention: the ones that allow people to do
something they couldn’t do before and the ones that let them do something
better. Fire, for example,  gave people the chance to cook food, push away wild
beasts and warm themselves up during cold nights. Many years later, electricity
let people warm their houses just by pushing a button. After wheels discovery
people began to travel and to trades goods, but was only with car’s invention
that they might do it faster and efficiently.  Similarly, the web creates a huge
network, able to connect people all over the world, web application gave people
tools to use and customise such a complex system. Under this perspective,
container is one of the main revolution of the last years, a unique tool that
helps with app management and development. Let’s discover  something more about
the real story of containers.

We have not a lot of documentation about why Bill Joy 18th March 1982 added
chroot into the BSD probably to emulate him solutions and program is an isolated
root. That’s was amazing but not enough few years later in 1991 Bill Cheswick
extended chroot with security features provided by FreeBSD and implemented the
“jails” and in 2000 he introduced what we know as the proper jails command now
our chroots can not be anything, anywhere out of themself. When you start a
process in a chroot the PID is one and there is only that process but from
outside you can see all processes that are running in a chroot.  Our
applications can not stay in a jail! They need to communicate with outside,
exchange information and so on. To solve this problem in 2002 in the kernel
version 2.4.19 a group of developers like Eric W. Biederman, Pavel Emelyanov
introduced the namespace feature to manage system resources like network,
process and file system.

This is just a bit of history about how the ecosystem spin up, in the end of
this chapter we will try to understand how why Docker arrives on the scene, but
the main goal of this book is on another layer and on another complexity, we are
here to understand how manage all this things in cloud and how to design a
distributed system but you know the past is important to build a solid future.

All this great features are now popular under the name of container, nothing
really news and this is one of the reason about why all this things are amazing!
They are under the hood from a while! Solid and tested feature put together and
made usable.

Nothing to say about the importance for a system to being isolated: isolation
helps us to usefully manage resources, security and monitoring, in the best way,
false problems creation in specific applications, often not even related to our
app.

The most common solution  is virtualization: you can use an hypervisor to create
virtual server in a single machine.  There are different kind of virtualization:

* Full virtualization
* Para virtualization like Virtual Machine, Xen, VMware
* Operating System virtualization like Containers
* Application virtualization like JVM.


<img class="img-responsive" src="/img/virtualization.png">
<a href="https://fntlnz.wtf/post/why-containers/" target="_blank"><small>img from fntlnz's blog. Thanks</small></a>


The main differences between them is how they abstract the layers, application,
processing, network, storage and also about how the superior level interact with
underlying level.  For example into the Full virtualization the hardware is
virtualized, into the para virtualization not.

Container is an operation-system-level virtualization. The main difference
between Container and Virtual Machine is the layer: the first works on the
operating system, the second on the hardware layer.

When we speak about container we are focused on the application virtualization
and on a specific feature provided by the kernel called Linux Containers (LXC):
what we do when we build containers is create new isolated Linux systems into
the same host, it means that we can not change the operation system for example
because our virtualization layer doesn’t allow us to run Linux containers out of
Linux.

# The reasons

Revolutions are not related to a single and specific event but come from
multiple movements and changes: Container is just a piece of the story.

Cloud Computing allowed us to think about our infrastructure as an instable
number of servers that can scale up and down, in a reasonable short amount of
time, with less money and without the investment requested to manage a big
infrastructure made of more than one datacenter across the world.

As a consequence, applications that had been in a cellar, now are on Amazon Web
Service, with a load balancer and maybe different availability zone. This
allowed little teams and medium companies, without datacenter and
infrastructures, to think about concept like distribution, high availability,
redundancy.  Evolution never stop .

Once our applications are running in few virtual machines, our business will
grow up so we start to scale up and down this servers to serve all our users.
We experimented few benefits but also a lot of issues related, for example, to
the time requested for managing this dynamism; moreover big applications are
usually more expensive to scale.

Our application can only grow but the deploy can be really expensive. We
discovered that the behavior of an application is not the same across all of our
services and entrypoint, because few of them receive more traffic that others.
So, we started to split our big applications in order to make them easy to scale
and monitor. The problem was that, in order to maintain our standard, we need to
find a way to keep them isolated, safe and able to communicate each others.

The Microservices Architecture arrived and companies like Netflix, Amazon,
Google and others counts hundreds and hundreds of little and specific of
services that together work to serve big and profitable products.  Netlix is one
of first companies that started sharing the way they build Netlix.com: with more
that 400 microservices, they managed feature like registration, streaming,
rankins and all what the application provides.  At the moment, Containers are
the best solution for managing a dense and dynamic environment with a good
control, security and for moving your application between servers.

<p class="text-muted">
    Reviewers: Arianna Scarcella, <a href="https://twitter.com/TheBurce">Jenny Burcio</a>
</p>
