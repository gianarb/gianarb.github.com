---
layout: post
img: /img/1280px-NixOS_logo.png
title: "My workflow with NixOS. How do I work with it"
date: 2022-09-12 10:08:27
categories: [post]
tags: [nixos]
summary: "In the last two years I pick up NixOS as I tool I want to use. The
learning curve is steep but I think I have a workflow that I like"
---

## Some context

Coding is fun when you can figure out the right workflow. There is nothing fun
when it comes to writing software in a way that is not sustainable or that does
not sparks joy.

I started to use Nix and NixOS almost two years ago, in a previous job in
a totally different context.

Back then we had to quickly and often provision operating system, build
software and so on. Since I moved back to write Software and to write Rust I
have to admit that building my code, or shipping operating systems is not
something I have to do very often, but I decided to keep learning and fighting
against NixOS because it fits my mindset.

Recently I resumed a few NUCs I keep in a box because everybody
deserves a home lab, and a good home lab deserves some netbooting, so it was
time to play with NixOS for something that is not my workstation or my laptop.

## The workflow

Nix is code, finally. It means that there are libraries, you can import them,
run tests, and execute such code. YAML, Json in my experience, at some point
are a limitation, or they create friction, you ended up with an easy to break
template engine.

I decided to invest some time to figure out how to use flake. And this is where
I am so far:

```nix
{
  description = "A generic and minimal netbooting OS for my homelab";

  inputs =
    {
      nixpkgs.url = "github:NixOS/nixpkgs/nixos-22.05";
    };

  outputs = { self, nixpkgs, ... }:
    let
      system = "x86_64-linux";
    in
    {
      nixosConfigurations = {
        generic = nixpkgs.lib.nixosSystem {
          inherit system;
          modules = [
            ./configuration.nix
          ];
        };
      };
      packages.${system}.netboot = nixpkgs.legacyPackages.${system}.symlinkJoin {
        name = "netboot";
        paths = with self.nixosConfigurations.generic.config.system.build; [
          netbootRamdisk
          kernel
          netbootIpxeScript
        ];
        preferLocalBuild = true;
      };
    };
}
```

I am not the right person to tell you what all of this does because I am not an
expert and it is the outcome of many videos on YouTube, questions on
discourse.nixos.org, articles and beers, a lot of beers.

The `output` part describes what I want to build and as you can see there are
two outcomes. One is a `nixosConfigurations`, potentially it can contain more
than one NixOS description but right now I have a single one called `generic`
and as you can see it imports a module called `configuration.nix`. You can see
it as a ready to go NixOS provisioned as I want. This is 99% a copy paste of a
traditional `configuration.nix` file as you may know them. The one I use comes
from  ["Netbooting Wiki" in NixOS.org](https://nixos.wiki/wiki/Netboot).

```nix
{ config, pkgs, lib, modulesPath, ... }: with lib; {
  imports = [
    (modulesPath + "/installer/netboot/netboot-base.nix")
  ];
  users.users.root.openssh.authorizedKeys.keys = [
    "ssh-sfdbsrbs"
  ];

  ## Some useful options for setting up a new system
  services.getty.autologinUser = mkForce "root";

  environment.systemPackages = [ pkgs.tailscale ];

  networking.dhcpcd.enable = true;

  services.openssh.enable = true;
  services.tailscale.enable = true;

  hardware.cpu.intel.updateMicrocode =
    lib.mkDefault config.hardware.enableRedistributableFirmware;

  systemd.services.tailscale-autoconnect = {
    description = "Automatic connection to Tailscale";

    # make sure tailscale is running before trying to connect to tailscale
    after = [ "network-pre.target" "tailscale.service" ];
    wants = [ "network-pre.target" "tailscale.service" ];
    wantedBy = [ "multi-user.target" ];

    # set this service as a oneshot job
    serviceConfig.Type = "oneshot";

    # have the job run this shell script
    script = with pkgs; ''
      # wait for tailscaled to settle
      sleep 2

      # check if we are already authenticated to tailscale
      status="$(${tailscale}/bin/tailscale status -json | ${jq}/bin/jq -r .BackendState)"
      if [ $status = "Running" ]; then # if so, then do nothing
        exit 0
      fi

      # otherwise authenticate with tailscale
      ${tailscale}/bin/tailscale up -authkey tskey-really
    '';
  };

  networking.firewall = {
    checkReversePath = "loose";
    enable = true;
    trustedInterfaces = [ "tailscale0" ];
    allowedUDPPorts = [ config.services.tailscale.port ];
  };

  system.stateVersion = "22.05";
}
```

The only difference compared with a traditional non-flake configuration is the import:

```
  imports = [
    (modulesPath + "/installer/netboot/netboot-base.nix")
  ];
```

Flake provides the utility variable `modulesPath` as a shortcut for accessing
the nixpkgs modules described as flake input.

This OS does a few simple things:

* Setup a public ssh key for the root user that I can use to ssh into the server.
* It register itself to Tailscale

The output `nixosConfigurations` is used via `nixos-build`.
It took me some time to figure out that `nixos-build` used in the right wat does not replace my current operating system.
Do not run `nixos-build switch` if you won't want to screw up your local NixOS OS! Instead you can build this
operating system in the `./result` directory via:

```
$ nixos-rebuild build --flake .#generic
```

A single configuration can describe different NixOS, that's why you have to
identify what you want to build with ` .#generic`.

The second output builds the same OS but it shapes the content of the
`./result` directory as I want it (I am not sure if I need it but this is what
the NixOS netbooting wiki does, so far so good).

To build it you can use `nix build`:

```
$ nix build .#netboot
```

Pretty cool! I can tar.gz that and ship it where I want. Straightforward.

## How to run this VM

Do you know how boring and time consuming it is to test a new operating system?

If you want to do it on real hardware you have to set it up, and if you want to
use QEMU you have a few days in front of you to remember all the flags you need, how
to bridge the guest with the host and who knows what. I tried for a few days
and I failed, until I discovered:

```
$ nixos-rebuild build-vm --flake .#generic
building the system configuration...

Done.  The virtual machine can be started by running /nix/store/dk4i22xmacnxxdmgvjhlyain5spb11yn-nixos-vm/bin/run-nixos-vm
```

Pure gold! If you run the `run-nixos-vm` script a QEMU virtual machine will
appear ready for you to test your operating system. Kind of cool! I can even
see it showing up in the Tailscale admin console!

A zero friction experience that boost my ability to try what I am working on.

## Integration tests

Nix provides a testing framework, but I started to use it recently. It spins up
one or more virtual machines and assert that they work as expected. I wrote
a test that looks for the tailscale network inteface:

```nix
let
  nixpkgs = fetchTarball "https://github.com/NixOS/nixpkgs/archive/0f8f64b54ed07966b83db2f20c888d5e035012ef.tar.gz";
  pkgs = import nixpkgs { };
in
pkgs.nixosTest
  ({
    system = "x86_64-linux";

    nodes.machine = import ./configuration.nix;

    testScript = ''
      start_all()
      machine.succeed("sleep 5")
      machine.succeed(
          "ifconfig | grep tailscale0",
      )
    '';
  })
```

This test uses the same `configuration.nix` I used to generate my netbooting
NixOS. It starts a node called `machine` and via python script it runs the bash
command `ifconfig | grep tailscale0`.  I am sure I can do better than `sleep 5`
but as I said, I am far away from being good at this.

You can use this approach to run assertions on multiple nodes,
here an example from Nix.dev ["Integration testing using virtual machines
(VMs)"](https://nix.dev/tutorials/integration-testing-using-virtual-machines).

## Steep learning curve

Everyone agrees that Nix and NixOS are not easy technology to pick up. And I
can confirm, there are articles, blogs, dotfiles available everywhere but they
look all different and it is hard to figure out if they are new, old or how to
apply them to your use case.

Flake is an attempt from the community to standardize all of that, and much
more. We will see!

It is also true that motivation and context can flat the curve. My plan is to
write more about this topic since I am trying to spin up and automated a home
lab.

I have to figure out how to do secret management but as soon as I have
it sorted out I will share my homelab configuration as I share my laptops
configuration in my [dotfiles](https://github.com/gianarb/dotfiles/tree/main/nixos).

Stay tuned.
