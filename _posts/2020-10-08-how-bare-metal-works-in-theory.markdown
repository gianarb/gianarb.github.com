---
layout: post
title:  "How bare metal provisioning works in theory"
date:   2020-10-08 10:08:27
categories: [post]
tags: [bare metal, tinkerbell, opensource]
summary: "What I learned about how bare metal provisioning works
developing tinkerbell."
---

I am sure you heard about bare metal. Clouds are made of bare metal, for example.

The art of bringing to life an unanimated piece of metal like a server to something useful is something I am learning since I joined [Equinix Metal](https://metal.equinix.com) in May.

Let me make a comparison with something you are probably familiar with. Do you know why Kubernetes is hard? Because there is not one Kubernetes. It is a glue of an unknown number of pieces working together to help you deploy your application.

Bare Metal is almost the same, hundreds of different providers, server size, architectures, chips that in some way you have to bring to life.


We have to work with some common concepts we have. When a server boots, it runs a BIOS that looks in different places for something to run:

1. It looks for a hard drive
2. It looks for external storage like a USB stick or a CD-Rom
3. It looks for help from your network (netbooting)

Options one and two are not realistic if the end goal is to get to a handsfree, reliable solution. I am sure cloud providers do not have people running around with a USB stick containing operating systems and firmware.

## Netbooting

I spoke about [my first experience netbooting Ubuntu](https://gianarb.it/blog/first-journeys-with-netboot-ipxe) on my blog. That article is efficient with reproducible code. Here the theory.

When it comes to netbooting, you have to know what PXE means. Preboot Execution Environment is a standardized client/server environment that boots when no operating system is found, and it helps an administrator boot an operating system remotely. Don't think about this OS as the one you have in your laptop, I mean, technically it is, but the one your run there or in a server is persisted, that's why you can have files that survive a reboot.

The one you start with PXE runs in memory, and from there, you have to figure out how to get the persisted OS you will run in your machine.

When the in-memory operation system is up and running, you can do everything you are capable of with Ubuntu, Alpine, CentOS, or Debian. In practice, what people tend to do is to run applications and scripts to format a disk with the right partition and to install the end operation system.

Pretty cool. PXE is kind of old, and for that reason, it is burned into a lot of different NICs. You will hear a lot more about iPXE, a "new" PXE implementation. What is cool about those is the `chain` function. From one PXE/iPXE environment, you can chain another PXE/iPXE environment. That's how you get from PXE (that usually runs by default in a lot of hardware (if you have a NUC you run it)) to iPXE.

```
chain --autofree https://boot.netboot.xyz/ipxe/netboot.xyz.lkrn
```

iPXE supports a lot more protocols usable to download OS from such as TFTP, FTP, HTTP/S, NFC...

This is an example of iPXE script:

```
#!ipxe
dhcp net0

set base-url http://archive.ubuntu.com/ubuntu/dists/focal/main/installer-amd64/current/legacy-images/netboot/ubuntu-installer/amd64/
kernel ${base-url}/linux console=ttyS1,115200n8
initrd ${base-url}/initrd.gz
boot
```

The first command, `dhcp net0`, gets an IP for your hardware from the DHCP. `kernel` and `initrd` set the kernel and the initial ramdisk to run in memory.

`boot` starts the `kernel` and the `initrd` you just set.

There is more; this is what I find myself using more often.

### Infrastructure

To netboot successfully, you need to distribute a couple of things:

1. An iPXE script
2. The operating system you want to run (kernel and initrd)

### Workflow

1. Server starts
2. There is nothing to boot in the HD
3. Starts netbooting
4. It makes a DHCP request to get network configuration, and  the DHCP returns the TFTP address with the location of the iPXE binary
5. iPXE starts and makes another DHCP request; the response contains the URL of the iPXE scripts with the commands you saw above
6. At this point, iPXE runs the script, downloads the kernel, and the initrd with the protocol you specified, and it runs the in-memory operating system.

Pretty cool!

## The in-memory operating system

The in-memory operating system can be as smart as you like; you can build your one, for example, starting from Ubuntu or Alpine. Size counts here because it has to fit in memory.

When the operating system starts, it runs as PID 1, what is called `init.` It is an executable located in the ramdisk called `/init.` That script can be as complicated as you like. It can be a problematic binary that downloads from a centralized location commands to execute, or it can be bash scripts that format the local disk and installs the final operating system.

What I am trying to say is that you have to make the in-memory operating system useful for your purpose. If you use native Alpine or Ubuntu, the init script will start a bash shell, not that useful.

## DHCP

As you saw, the DHCP plays an important role. It is the first point of contact between unanimated hardware and the world. If you can control what the DHCP can do, you can, for example, register and monitor the healthiness of a server.

Imagine you are at your laptop, and you are expecting a hundred new servers in one of your datacenters, monitoring the DHCP requests. You will know when they are plugged into the network.

## Containers what?

Containers are a comfortable way to distribute and run applications without having to know how to run them. Think about this scenario. Your in-memory operating system at boot runs Docker. The `init` script at this point can pull and run a Docker container with your logic for partitioning the disk and installing an operating system, or it runs some workload and exit leaving space for the next boot (a bit like serverless, but with servers, way cooler).

Or the Docker Container can run a more complex application that reaches a centralized server that dispatches a list of actions to execute via a REST or gRPC API. Those actions can be declared and stored from you.

## Conclusion

The chain of tools and interactions to get from a piece of metal to something that runs some workload is not that long. Controlling all the steps and the tools along the way gives provides the ability to provisioning cold servers from zero to something that developers know better how to use.

Ok, I lied to you. This is not just theory. This is how [Tinkerbell](https://tinkerbell.org) works.


This post was originally posted on [dev.to](https://dev.to/gianarb/how-bare-metal-provisioning-works-in-theory-1e4e).
{:.small}
