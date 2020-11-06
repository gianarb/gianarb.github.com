---
layout: post
title:  "What is Tinkerbell?"
date:   2020-11-06 10:08:27
categories: [post]
tags: [tinkerbell, bare metal]
summary: "I want to share with you what is Tinkerbell. An open source project I
help maintaining, developed by Equinix Metal. Tinkerbell helps you to manage
your hardware and datacenter programmatically via an API."
heroimg: /img/hero-tinkerbell.jpeg
---
First things first, Tinkerbell is an open-source project mainly written in Go
that comes from PacketHost, now Equinix Metal. Equinix Metal is a cloud provider
that serves bare metal servers. No virtual machines, no high-level services, I
said bare metal! Imagine a colocation that you can rent per hour.

Tinkerbell is the software Equinix Metal dreamed about as an internal
provisioner for datacenter automation. They took their internal provisioner and
removed any PacketHost references of business specific code, and pushed it to
GitHub for the community to enjoy the same technologies.

The project is a number of micro-services that provide various functionality to
configure hardware and provision bother Operating System and additional software
through it’s workflow engine.

## What the project provides

*   The first micro-service is called
    [boots](https://github.com/tinkerbell/boots). This Tinkerbell service
    provides a DHCP and a TFTP server to tell a piece of hardware (server) what
    to do when netbooting, it provides this information through the iPXE
    project.
*   Tinkerbell serves a CLI that you can use to interact with a control plane
    that serves HTTP and gRPC API. The service which does all those things is in
    the [tink](https://github.com/tinkerbell/tink) repository, and it provides
    three binaries: tink-server (the control plane), tink-CLI, the command line
    interface, and tink-worker.
*   Tinkerbell provides an operating system that runs in memory, it is based on
    Alpine, and it is called [Osie](https://github.com/tinkerbell/osie). This
    in-memory operating system runs directly on the hardware you want to
    provision, and it runs the tink-worker.
*   Once the in-memory Osie has started it begins the tink-worker which in turn
    communicates with the control plane (tink-server), asking for any work that
    has to be done on that server. This unit of work is called a workflow.
*   [Hegel](https://github.com/tinkebell/hegel) is a metadata server, comparable
    to the AWS EC2 metadata or the Equinix Metal one; the majority of cloud
    vendors provide this type of service , so you should have it as well! It is
    crucial when running scripts in a particular server because you can get
    concrete variables from it, such as the operating system it runs, its IPs,
    location, and so on.

## The end goal

The Tinkerbell end work is to bring to life a piece of hardware.

## Workflow and template

A template is a specification file that describes what we want to execute. A
workflow starts from a template, and it has a particular target. Templates are
reusable; workflows are a single execution and can't be reused. The single unit
of work in a template is called action. You can get as many actions you want in
a template, and each action runs in its own Docker container.

## Action

As mentioned above, actions are Docker containers and that means that you can
build each action in isolation in the language you want. It can use python,
bash, Golang, Rust, or whatever you can run in a container.

You may think that Docker would sound like an overhead, however we took a
natural decision based on how we could use the container concept in operations.
The concept of build, pull, and push has become commonplace within development
environments, and we think it could also work well in operational environments
too. Building containers to contain operational tasks in isolation and enhancing
that with testing and simplified execution of a container is a clear benefit. It
is an effective way to move code around in a reusable way without having to
reinvent the distribution model. Some of the actions you will see very often in
a Tinkerbell workflow may be:

*   Disk related actions: mounting a disk, wiping it, or setting up a partition
    table to boot an operating system
*   Downloading an Operating System like Ubuntu, Debian, NixOS, CentOS
*   Copy an operating system in a partition

But you will be able to write actions related to your business:

*   notify a particular API when provisioning fails
*   Attempt a recovery
*   Observe and mark the status of your provisioning
*   Who knows! There are no limitations here.


## How a template and a workflow looks like

Unfortunately, there are not many examples, but as maintainers, the next three
months will be all about public workflows and reusable actions.

Kinvolk wrote a blog post about [how to provision
Flatcar](https://kinvolk.io/blog/2020/10/provisioning-flatcar-container-linux-with-tinkerbell/)
on bare metal with Tinkerbell.

The Tinkerbell documentation [has an
example](https://tinkerbell.org/examples/hello-world/) of a "hello world."
template.

[Frans van Berckel](https://www.fransvanberckel.nl/) wrote a workflow for
[CentOS](https://github.com/fransvanberckel/debian-workflow) and
[Debian](https://github.com/fransvanberckel/debian-workflow).

One of my next projects will be to write a workflow that won't install an
operating system. It will start something like k3s or k8s directly on Osie for
my ephemeral homelab! I am not sure it has a sense or will ever work, but I
think it is an excellent example: "it is not all about having a persisted and
traditional operating system those days."


## How to get started

We put a fair amount of effort into a
[sandbox](https://github.com/tinkerbell/sandbox) project and setup guide. You
can run it [locally with Vagrant
](https://tinkerbell.org/docs/setup/local-with-vagrant/)or on [Equinix
Metal](https://tinkerbell.org/docs/setup/terraform/).

Aaron Ramblings wrote a blog post, ["Tinkerbell or iPXE boot on
OVH"](https://geekgonecrazy.com/2020/09/07/tinkerbell-or-ipxe-boot-on-ovh/)
using the sandbox to run Tinkerbell on OVH! I am still surprised when I read it
because he experimented with the sandbox in a very early stage of the project,
and in the same way, he was able to run sandbox on OVH; it can run almost
wherever else (at least for the control plane part).

## Next steps

With the help of our community we recently improved our continuous integration
pipeline to build all the projects for various architecture: `linux/386`,
`linux/amd64`,` linux/arm/v6`, `linux/arm/v7`, `linux/arm64` levering Docker
buildx, Qemu, and GitHub Actions. My goal was to be able to run the provisioner
in a Raspberry Pi. Because as I wrote before, my homelab tends to go away, get
moved, disconnected, and I think I can keep running reliably only a Raspberry PI
as it is today. So I want to run the control plane on a RaspberryPI. I presume
there are smarter things to do with multi-arch, but let's be honest; we all have
a RaspberryPI leftover somewhere.

We use the sandbox project as a way to release Tinkerbell's version as an all
project. We are pinning all the various dependencies such as Boots, Hegel,
Tink-Server, CLI, Osie, and when they all pass the integration tests, we tag a
new release. The generated artifacts are containers for now. We want to get
binaries in this way. You can run Tinkerbell as you like, even without
containers. At some point, we will tag and manage each component independently,
but for now, it is a lot of effort.

Releasing new workflows is something we are working on already. So stay tuned!

Another project is available in the Tinkerbell GitHub organization that I didn't
mention because it is not hooked yet as part of the stack. After all, we are
working at its version two. [PBNJ](https://github.com/tinkerbell/pbnj) provides
a standard API to interact with various BMCs and IPMIs (Intelligent Platform
Management Interface). Having this kind of ability in a datacenter is essential
because we want to pilot things like reboot, restart, switch off for each server
programmatically, and even as part of a workflow.

## Conclusion

There already exists huge demand for bare-metal usage, which with the growth
caused by things like 5G, dedicated GPUs/FPGAs, HPC, constant and expected
performance and security boundaries is only going to grow. A recent report by
the [Mordor Intelligence
company](https://www.mordorintelligence.com/industry-reports/bare-metal-cloud-market)
reports  “The bare metal cloud market was valued at USD 1.75 billion in 2019 and
expected to reach USD 10.56 billion by 2025” which clearly shows that there is
growing demand for a modern platform to provision their bare-metal
infrastructure.

Datacenter management is hard, and that's why the public cloud got so much
traction. For companies and products, managing hardware is unnecessary and a
distraction, but when it becomes a requirement or when you think it is strategic
to manage your own hardware Tinkerbell and its community comes to rescue you.

{:.alert.alert-info}
A big thank you goes to [Dan](https://twitter.com/thebsdbox) for his review and
support writing this article!

## More, I want more!

[Dan and Jeremy had a conversation](https://www.youtube.com/watch?v=Y04eCSKaQCc)
about netbooting and bare metal provisioning.  It is available on YouTube, you
should really have a look at it!

[Alex Ellis and Mark Coleman recorded a
video](https://www.youtube.com/watch?v=QxpKnMGywTU) setting up and using Tinkerbell.
The video is a bit out of date and they did not use the new sandbox project
because it was not available at that time. But still a good and valuable!
