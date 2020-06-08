---
img: /img/me.jpg
layout: post
title: "Netboot and iPXE let's make our hands dirty installing Ubuntu"
categories: [post]
tags: ["ipxe", "netboot", "packet", "bare metal"]
summary: ""
changefreq: daily
---

I recently joined Packet, a company acquired by Equinix, finally after 7 years
working with cloud computing I can see how they look like from the other side!

{:refdef:.blockquote}
> Spoiler alert: clouds are made of servers
{:refdef}

As a first task I had to revamp the kubernetes cluster-api implementation,
moving it from v1alpha1 to v1alpha3. Kind of cool and in a domain I know very
well. We tagged the first release and I can't wait to see how it will be used.
I got some meetups and webinars planned about it, so stay in touch on
[twitter](https://twitter.com/gianarb) to know more about it.

Anyway, one of the topics I am curious about is hardware automation. The idea to
get up and running, in a repeatable and autonomous way a piece of inanimate
metal well known as rack, switch, server is a topic I never touched and I would
like to know more! Obviously, this is only one of the article I will write about
the topic. Mainly because there is much to learn.

As you can image when we buy a server it does not do much, it's great to keep
your door open, and as a table. In order to make something good with it
has to be configurable, in our case customers can do it via API, so it means
that there is some code and automation involved! I want to know more!!

For sure there are a couple of things that you have to do manually, assemble the sever,
power it on, plug the ethernet cable in, pick the right location and things like
that.

But as you can imagine it comes without operating system, even more complicated to
install because the customers can select the one they like most, or even push
their own one. This is for sure something that has to be done magically, I doubt
we have people running with USB stick in a datacenter installing operating
systems.

{:refdef:.text-center}
![Forst Gump picture](https://i0.wp.com/www.anonimacinefili.it/wp-content/uploads/2019/07/forrest-gump-25-anni.jpg?fit=1200%2C600){:.img-fluid}
{:refdef}

We also know that one of the things that happen when booting a laptop or server
is the bootloader. The one that requires a master skill to get in because there
is a timeout, and I never know what to press! I have to be honest I thought they
were pretty static and not that fun.

BUT, there are smart bootloader! We know smart means does days: internet! In
practice there is a bootloader capable of booting not from USB, not from disk
but from the internet.

Usually a private one but it is not mandatory, it goes and take the kernel, the
`initrd` and it `boot` an operating system. It is like having a USB stick that
starts the live installation of Ubuntu, there is a lot more after that because
we have to persist the installation on disk, format them and so on, but that's
what the point where I am now. I have to manually do the installation wizard to get to a
running OS but still, not too bad.

The bootloader is called PXE, and the new generation I used is called iPXE.

The cool things about PXE/iPXE is that you can chain scripts from one to
another. [Graham Christensen](https://twitter.com/grhmc) told me that the main goal when you work with PXE is to escape
from it and get to iPXE, that is way cooler. Time for a recap:

1. The machine starts and enters PXE
2. PXE download and chains iPXE in this way we are on iPXE bootload
3. From iPXE you can download kernel, initrd and boot the OS in RAM.

IPXE supports different ways to download what it needs from the internet, the
one I used so far are TFTP and HTTP.

## What is this PXE/iPXE?

Such a nice question, I had the same one few days ago. A couple of links:

1. [Wikipedia: Preboot Execution Environment](https://en.wikipedia.org/wiki/Preboot_Execution_Environment)
2. [iPXE: open soure boot firmware](https://ipxe.org/)

Roughly you can think about iPXE as a shell that has a bunch of commands like:

1. dhcp: to require an IP from a DHCP server and configures the network
   interface
2. route: to figure out if the network interface is configured (if it has an IP
   already)
3. chain: gets an argument (a URL) and it executes the content, it is a good way
   to pass scripts
4. You can see variables `set name value`
5. kernet: downloads the kernel from a source and load it
6. initrd: download the init ramdisk
7. boot: triggers the boot

And [many more](https://ipxe.org/cmd) that I did not use yet but docs lists
them.

It also has support for building a menu like this one:

{:refdef:.text-center}
![netboot menu](https://netboot.xyz/images/netboot.xyz.gif){:.img-fluid}
{:refdef}

The image comes from [netboot.xyz](https://netboot.xyz/) and as you can see from
their website is a project that simplify the process of installing a lot of
different operating systems via PXE. I started with it at first for my
experiments. Obviously menus and automations does not play nice together but in
the process of learning I took this extra step.

## Hello world

To give you some context this is the script I used to start the installation
wizard for Ubuntu:

```
#!ipxe
dhcp net0

set base-url http://archive.ubuntu.com/ubuntu/dists/focal/main/installer-amd64/current/legacy-images/netboot/ubuntu-installer/amd64/
kernel ${base-url}/linux console=ttyS1,115200n8
initrd ${base-url}/initrd.gz
boot
```

I hope it looks familiar, some code at least. In order to
reach the internet you need an IP, and to get one if you are lazy like me is to
use a DHCP (the alternative is to set one statically). The first command does that, asks the DHCP an IP for the network
interface `net0`.

When the IP is set, iPXE reachs `ubuntu.com` to get the kernel and the initrd.
Everything I need to boot an OS in RAM.

## Try yourself

I am obviously using Packet for my test. You can register with []() to get some
credit for free.

When you request a device (a server) you can select the operating system, we
don't need to do it, so you can select `Custom iPXE` because we are going to
install it ourselves.

{:refdef:.text-center}
![A screenshot from packet.com about how to create an on demand device with Custom iPXE](/img/packet-create-device.png){:.img-fluid.w-75}
{:refdef}

There are two ways we can inject out script to iPXE in order to teach the server
what to do when it boot, first one is giving a URL (I use a gist (raw link)), or
via user data. The script I used is the one pasted above. You can create a gist
and paste the link in "iPXE Script URL" or you can use the user data, as I am
doing right now.

{:refdef:.text-center}
![A screenshot form packet.com about how to pass a user data to a server](/img/packet-user-data.png){:.img-fluid.w-75}
{:refdef}

As soon as the machine starts you can click on its name to enter the get its
details and you get a ssh into the "Out-of-Band Console" console:

{:refdef:.text-center}
![A screenshot from packet.com that shows where to locate the out-of-band
console](/img/packet-out-of-band.png){:.img-fluid.w-75}
{:refdef}

{:refdef:.text-center}
![A screenshot from packet.com that shows how to get the ssh command to use
the out-of-bound console](/img/packet-out-of-band-ssh.png){:.img-fluid.w-75}
{:refdef}

ALERT: if you are doing this activity, remember to enable OpenSSH when you
follow the installation wizard otherwise you won't be able to SSH in the server
at the end!

When deploying the server the code you passed gets chained from the Packet iPXE.
And you should see the Ubuntu wizard ready for you:

{:refdef:.text-center}
![A screenshot from my terminal that shows the first wizard for ubuntu](/img/packet-ubuntu-install-wizard.png){:.img-fluid}
{:refdef}

At the end of the wizard you will get a persisted operating system in the server
itself, it will survive the reboot and it will be just as any other server you
used in the past, but better because you know how you installed the OS! Get its
IP and ssh in!

## Conclusion

That's it, as I told you I have no idea about how to get to a fully automated
workflow but I will get there I am sure, it is not easy but Packet has an open
source project called [Tinkerbell](http://tinkerbell.org/), that does that. But
I want to know what it does under the hood! In practice is an open source
version of the provisioner used internally to set up servers. We are moving to it
as well!

Let me know what you think and if know how to get into a fully persisted and
working operating system programmatically from here let me know, I am open
to any pointer!
