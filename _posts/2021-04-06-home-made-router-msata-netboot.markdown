---
layout: post
title: "How I tricked the cable mafia with PXE. Install OpenWRT on APU4d"
date: 2021-04-06 10:08:27
heroimg: "/img/messy-cabling.jpeg"
categories: [post]
tags: [hardware, homelab, openwrt, pixiecore, apu4]
summary: "No matter how many cables or dongle do you have they are never
enough. The best you can do is to trick the system. I tried Pixiecore to PXE
boot Alpine on my APU4d installing OpenWRT to it."
---

I am too lazy to buy a cable or another adapter. But not to buy an APU4d. A specialized for networking hardware with AMD Embedded G series GX-412TC, widely used for routers.

I got it directly from the manufacturer [PC Engine](https://www.pcengines.ch/apu4d4.htm). With a serial-USB cable and a [Huawei Lte miniPCI chip](https://it.aliexpress.com/item/32443776508.html).

I have also got the [16GB mSata SSD module](https://www.pcengines.ch/msata16g.htm) because you never know, having a 16GB SSD, 4GB of RAM router sounds like an opportunity to run more tools on it!

{:refdef:.text-center}
![Picture of the APU4D board from PC Engine](/img/apu4d.jpeg){:.img-fluid.w-75}
{:refdef}

I assembled all of it nicely at my desk. It was too late when I realized I don't know how to flash an mSATA SSD because I don't have proper cabling...

No matter how big the box with all my cables and dongles is, I will never own the one I need. It is a mantra nobody can escape. The best you can do is to trick the system.

Luckily for me, the APU4d supports PXE booting, and we know how cool it is, perfect opportunity to try [pixiecore](https://github.com/danderson/netboot/blob/master/pixiecore/README.api.md) and have some fun with netbooting.

It worked. If all of this sounds unreasonable, you need to remember that most likely you are right. But you know how much I like simple tools. Pixiecore was on my radar.

## Get what you need

First of all, I installed Pixiecore. It is a Go binary, you can run it as a Docker container or, you can compile it with `go build` but I decided to use Nix shell:

```bash
nix-shell -p pixiecore
```

In practice, it is a program that helps you to serve what a piece of hardware needs to PXE boot over the network, it servers IPXE and a TFTP server for example. It is light and not intrusive. You can keep your DHCP server, and if you like, even implement an API to drive how and what to PXE boot dynamically. Today I have to boot only one server in a very boring network. My solution is already too overly engineered. I decided to run it in static mode:

```bash
sudo pixiecore boot ./vmlinuz-vanilla initramfs-vanilla \
    --cmdline='console=ttyS0,115200n8 \
    alpine_repo=http://dl-cdn.alpinelinux.org/alpine/v3.9/main/ \
    modloop=http://dl-cdn.alpinelinux.org/alpine/v3.9/releases/x86/netboot-3.9.6/modloop-vanilla'
```

The first two arguments of the command line are the Alpine init ramdisk and the kernel. I got them directly from the [Alpine repository](http://dl-cdn.alpinelinux.org/alpine/v3.9/releases/x86/netboot-3.9.6).

The `--cmdline` option can be used to pass configuration to the operating system. The [Alpine netboot wiki page](https://wiki.alpinelinux.org/wiki/PXE_boot) to know the various options supported by the init script.

Now that I have set the PXE distribution tool, I powered on the APU4d board. By default, it tries to boot from a couple of different devices. The last one is PXE mode.

```console
sudo pixiecore boot \
    ./vmlinuz-vanilla initramfs-vanilla \
    --cmdline='console=ttyS0,115200n8 \
        ssh_key=https://github.com/gianarb.keys \
        alpine_repo=http://dl-cdn.alpinelinux.org/alpine/v3.9/main/ \
        modloop=http://dl-cdn.alpinelinux.org/alpine/v3.9/releases/x86/netboot-3.9.6/modloop-vanilla'

Password:
[DHCP] Offering to boot 00:0d:b9:5a:3e:10
[DHCP] Offering to boot 00:0d:b9:5a:3e:10
[TFTP] Sent "00:0d:b9:5a:3e:10/4" to 192.168.1.87:55360
[DHCP] Offering to boot 00:0d:b9:5a:3e:10
[HTTP] Sending ipxe boot script to 192.168.1.87:29233
[HTTP] Sent file "kernel" to 192.168.1.87:29233
[HTTP] Sent file "initrd-0" to 192.168.1.87:29233
```

`192.168.1.87` is the IP the APU4 got from my DHCP. Everything is working and from the serial port I see Alpine booting, the `root` password is `root`! Classy!

## Time to install OpenWRT

I never used OpenWRT before. It is a Linux distribution for routers. You can even flash it to TP-LINK or Netgear devices if supported, at your own risk.

Anyway, since I am now running Alpine in memory on my APU4d I have a functional operating system and access to the device. I can use traditional tools like `dd` to write OpenWRT directly to disk, manipulate partitions and, so on... I followed the blog post ["OpenWRT installation instructions for APU2/APU3/APU4 boards"](https://teklager.se/en/knowledge-base/openwrt-installation-instructions/) written by TekLager.

## Conclusion

My router looks up and running. I was able to reach the administrative Web UI. I didn't use it yet because I have to relocate it to my new house. So I am sure you will read more about it in future articles.

Pixiecore was on my TODO list because those days hardware, datacenters automation are taking a good part of my daily working activity. Its support for external API makes it a great alternative to provide an installation environment like [Hook](https://github.com/tinkerbell/hook) (the one we developed with [Tinkerbell](httsp://github.com/tinkerbell)) without having to onboard the full Tinkerbell stack, in particular, I can avoid [boots](https://docs.tinkerbell.org/services/boots/) when not needed.
