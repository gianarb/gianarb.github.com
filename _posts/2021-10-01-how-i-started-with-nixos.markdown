---
layout: post
title: "How I started with NixOS"
date: 2021-10-01 10:08:27
categories: [post]
tags: [nixos]
summary: "I played with NixOS for the last couple of months. This is a story about how I picked it up, or how I should have done it."
---

I frequently change operating-system and distribution moving between macOS and Linux because I didn't marry any of them yet.

Just before having a MacBook again, I was an ArchLinux user, a happy one. I have to admit it was not that different compared with other distributions, at least as a user.  Yes, fewest packages installed, a few services, please don't freak out, as I wrote I enjoyed it.

I see a value when it comes to describing as code your desires, learning from other people sharing their code, importing or copy-pasting it in different places.

Developers do that all day. I am a representative, I hope what I write will match my desires. After so many years I am full of hope.

With this in mind, Arch, Debian, or Ubuntu do not make a difference. It is all about the package manager. NixOS and Nix looked to me as a step forward in this sense.

I decided to end my vacation with macOS earlier. I picked up my personal Asus Zenbook 3 from its box to install NixOS.

Coming from ArchLinux the NixOS installation process is similar, we are on our own:

1. Format disks
2. Write partition table
3. Mount partitions
4. And so on

The main difference comes when you run `nixos-generate-config`:

```
# nixos-generate-config --root /mnt
```

The command tries its best to detect kernel modules from your hardware, mount points, and so on. This phase is a great time to start your first fight of many with NixOS.
The generated file will be in `/mnt/etc/nixos/configuration.nix` and `/mnt/etc/nixos/hardware-configuration.nix`. Open the generated file to can validate if they have sense. Don't worry. It is a Linux distribution. If something is missed, it will tell us.
The `hardware-configuration.nix` file as the name suggests identifies your hardware.

Not everything can be detected yet, I use `luks` to encrypt my disks; the generated `hardware-configuration` needs a bit of help to figure it out.

```nix
  boot.initrd.luks.devices = {
    root = {
      device = "/dev/nvme0n1p2";
      name = "root";
      preLVM = true;
      allowDiscards = true;
    };
  };
```

Nix as a programming language takes a bit of practice, but NixOS is different. Many people share their configuration in GitHub, a boost in productivity.
I keep a list of NixOS configurations or Nix-related repositories that I look at when I don't know how to solve a particular issue. [I really think you should do the same](https://github.com/gianarb/dotfiles/tree/master/nixos#credits) because nobody wants to spend a day fixing its laptop, even worst if it is the one you use at work.

## Start simple

My end goal was to checkout my NixOS configuration as part of my dotfiles in a git repository. Too much when you don't even know how NixOS works.

I have put aside this goal for a few weeks, and my new goal was to get my laptop working in all its part. The complicated part, that I didn't solve in total is audio, it works but the volume control is not as good as it should be. You can check the configuration I use in my dotfiles but the solution does not matter.

## Check it out

When I was happy with my configuration it was time to finally move it in its final destination. I joined the "stable era" for my Nix configuration, everything was good enough and it was not changing costantly. Perfect time for some refactoring.

I decided to use my [dotfiles repository](https://github.com/gianarb/dotfiles) with a `nixos` subdirectory. This is the one I had when I first moved the configuration from my local environment to Git:

```
$ tree -L 1 ./nixos
./nixos
└── machines
    └── AsusZenbook
        ├── configuration.nix
        └── hardware-configuration.nix
```

Those `*.nix` files are a copy of the one I have in `/etc/nixos`.

Now I had to teach NixOS where the new configuration are, the are various way, I decided to delete everything inside `/etc/nixos/configuration.nix`, leaving only an `import` to the configuration I moved as part of my dotfiles.

NOTE: I clone my dotfiles at `/home/gianarb/.dotfiles`.

I didn't need `/etc/nixos/hardware-configuration.nix` and this is the content for my `/etc/nixos/configuration.nix`:

```nix
{ config, ... }:

{
  imports = ["/home/gianarb/.dotfiles/nixos/machines/AsusZenbook/configuration.nix"];
}

```

## Pick a second use case

I got a new Thelio System76 workstation (thanks to EraDB) and it was the perfect opportunity to re-use my fresh NixOS configuration, and my new skill.

At this point I am still working from my Asus Zenbook, but it is time to get a new `./machines/thelio` directory without a `hardware-configuratio.nix`, but only with the `configuration.nix` in there. The idea is to start extracting what you want to reuse from your first machine in its own files that can be imported everywhere you want.

I started from my user, because it is a common desire to reuse the same user across different machines. That's why I have in my dotfiles a `users` subdirectory.

```
gianarb@huge ~/.dotfiles  (master=) $ cat nixos/users/gianarb/default.nix
{ config, inputs, lib, pkgs, ... }:
with lib;
{
  # Define a user account. Don't forget to set a password with ‘passwd’.
  users.users.gianarb = {
    isNormalUser = true;
    uid = 1000;
    createHome = true;
    extraGroups = [
      "root"
      "wheel"
      "networkmanager"
      "video"
      "dbus"
      "audio"
      "sound"
      "pulse"
      "input"
      "lp"
      "docker"
    ];
    openssh.authorizedKeys.keys = [
      "ssh-ed25519 AAAAC3NzaC1lZDI1NTE5AAAAIEKy/Uk6P2qaDtZJByQ+7i31lqUAw9xMDZ5LFEamIe6l"
    ];
  };
}
```

I have imported it in both machines as we did previously for the all configuration and I splitted other applications like: `i3`, my audio configuration, vscode and so on. You can find all of them inside the `applications` directory:

```
$ tree -L 1 nixos/applications/
nixos/applications/
├── i3.nix
├── sound-pipewire.nix
├── sound-pulse.nix
├── steam.nix
├── sway.nix
├── tailscale.nix
└── vscode.nix
```

Double checking that the refactoring is just a matter of re-building NixOS:

```
# nixos-rebuild test
# nixos-rebuild switch
```

## Time to install NixOS the second target

I had everything I needed to re-install NixOS with my configuration into another target. It was time to setup a USB stick and boot Thelio from the USB.
The system I want is described as Nix configuration, the installation looks the same as the one we have done, or the one described in the documentation, but at this point, we do not need to generated configuration. We have our own one.
The only part we need, the first time, if we want is the `hardware-configuration.nix`.

1. When you have booted from USB you can do what you have done previously, and what it is explained in the [NixOS installation guide](https://nixos.org/manual/nixos/stable/#sec-installation), format and parition disk.
2. When you have the disk layout done you can mount it to `/mnt` and you can clone/download somewhere the git repository with your nix configuration. I usually create an clone it where I want it to end up: `/home/gianarb/.dotfiles`.
3. Time to run `nixos-generate-config` as you do all the time
4. Replace `/ect/nixos/configuration.nix` with the `import`, copy the generated `hardware-configuration.nix` to your `machines` folder
5. Last step is to open the hardware-configuration and figure out if it has sense for your hardware.
6. When you are happy with it you can run `nixos-install` and it will install it from the configuration you have just declared.

If it sounds like a convoluted process, it can be simplfied. But I didn't yet invested into it yet! I don't want to reinstall them all the time. You can read this article if you want to erase your laptop every day: ["Erase your darlings by Graham Christensen"](https://grahamc.com/blog/erase-your-darlings).

## Conclusion

You just read about my journey with NixOS. With a centralized repository I can assemble, compile, and ship images to run on AWS, or ISO I can PXE boot.
I can build and compile a NixOS derivation that I can use as installation driver, for example cloning my dotfiles.

As next project I want to build an ISO that I can flash into a Raspberry PI who will act as media hub for my speakers playing Bluetooth audio or Spotify playlists via [raspotify](https://github.com/dtcooper/raspotify).
