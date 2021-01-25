---
layout: post
title: "Nix for developers"
date: 2021-01-25 10:08:27
heroimg: /img/hero/nix.png
categories: [post]
tags: [nix]
summary: "Nix is slowly sneaking in my toolchain as a developer bringing back
the joy of provisioning. Not as an exercise of translation between technologies
and environment but as the art of building your own environment. "
---

## Nix is slowly sneaking in my toolchain.

Currently, a lot of my colleagues use Nix. It is a package manager that runs on Linux and macOS. It is versatile. I will show you more about it moving forward, but for now, think about it as a replacement for APT, YUM, and HomeBrew can work on both Mac and Linux. It is also a build system, but I didn't use it much for it just yet.

### Not tight to an operating system

For me, this is already a huge benefit. From time to time, for no reason, I end up switching from Mac to Linux and vice versa. It usually happens because I change the place where I work, and the policy forces me to make some weird decisions.

When I got off college, my parents bought me my first laptop; it was a Macbook Pro, but for the first two, three jobs, I used Linux because Mac was too expensive for my employers, and the main reason for owning a Mac was because back then I was not a fully flashed developers. I used to do video editing, playing with Photoshop, and so on. I quickly learned that I couldn't match two colors nicely together, and Linux was just enough for me as a developer logging CLI, VIM, and tools like that. When I was in Dublin, Macbook Pro was the only available option; at InfluxData, I had a Thinkpad (the best laptop I ever had). Currently, I work on a Macbook Pro again because the non-apple available option was pretty low in terms of performance.

Now that you know my struggle for laptop and operating system consistency, a tool that works on both sounds appealing.

Nix has its own Linux distribution called NixOS. I slowly have a lot at it, but it is not a topic for this article.

### Declarative environment

The open-source project I mainly consistently from years is my [dotfiles](https://github.com/gianarb/dotfiles) repository. I am probably the only person who knows how to run it, but it contains configuration for the various tools I use.

I have to admit that I would like to install it consistently and quickly on any of my on-demand servers I spin up, but I too lazy for it. Anyway, I like that approach because I describe what I want, and I can consistently get it everywhere. Nix gave me the same possibility, and it does not use a specification language like YAML, JSON, or whatever it uses a dialect.

It is a lazy, pure, and functional language. It is pretty awkward; I have to say, at least for my background. I didn't figure it out yet, but the more I use it, and better it sticks to my mind.

I am also not that good when it comes to picking up new languages, it takes me some time, and I have to practice with them.

The good thing is that there are plantly of tutorials, each of them with different stakeholders. Do you like to be driven by example? There is ["Nix by example."](https://nixos.wiki/wiki/Nix_Expression_Language) You have time, and you want a more traditional [reference manual](https://nixos.org/manual/nix/stable/#ch-expression-language); they got you covered.

The fact that I don't have to fight with the template engine makes me happy.

### It is all based on Git.

I use Git since my first day at my first job. I was a solo developer, and my remote repository was not GitHub but a USB stick.

The Nix package manager is a GitHub repository. You can have your one, or you can use [nixpkgs](https://github.com/NixOS/nixpkgs). You can even merge multiple ones. Or import your derivation (the way Nix calls package).

A text editor and Git to clone a repository are what you need to look at all the packages, and their definition gives me a friendly feeling.

Based on how you want to define your environment, you can pin all the packages you are installing to a specific commit SHA from the package manager repository:

```nix
let _pkgs = import <nixpkgs> { };
in
{ pkgs ?
  import
    (_pkgs.fetchFromGitHub {
      owner = "NixOS";
      repo = "nixpkgs";
      #branch@date: nixpkgs-unstable@2021-01-25
      rev = "ce7b327a52d1b82f82ae061754545b1c54b06c66";
      sha256 = "1rc4if8nmy9lrig0ddihdwpzg2s8y36vf20hfywb8hph5hpsg4vj";
    }) { }
}:

with pkgs;
```

Very powerful.

### Environment composition with Nix

I am not sure if environment composition has any sense, but it sounds descriptive to me. Nix is user and project aware.

With nix-env, you can install packages as a user. With nix-shell, you can manipulate your system at the project level. If you add NixOS to this chain, you get free customization at the operating system layer.

Currently, nix-shell is the tool I know more about, and I am in love with.

I didn't experience those composition levels; I am currently writing my home-manager configuration file to solve my dotfiles repository's required dependencies. Right now, I don't have a way to install them automatically. I am not sure if that's the right layer for such a problem yet, but I will figure it out soon.

### Project level sandboxing with nix-shell

With a combination of symlinks and who knows what, nix-shell gives you a sandboxed environment with the only dependencies you need for your project. When you run nix-shell, it looks for a file called shell.nix that describes needed dependencies, environment variables, and so on. By default, you get all the commands and utilities you have in your system plus the one you declared for that project. If you have Go 1.15 in your system but want 1.13 for a single project nix-shell, you want to make it happen, for example. Tinkerbell has a [shell.nix](https://github.com/tinkerbell/tink/blob/master/shell.nix) for almost all the repositories.

For some particular scenarios, I use the Docker container in development. But with Nix, I can remove that extra layer. I use containers and images to ship and run my applications on Kubernetes. Removing that layer decreases the need for volume mounting, port forwarding, the debugger works much more comfortably, and performance is the one your hardware provides to you, without virtualization if you are on Mac.

Containers in development are my way to go when for dependencies that I don't care about or that I will never modify and have a state such as databases. But it is a joy to develop "locally."

### Everything can be "nixyfied"

Passing the flag --pure to nix-shell won't rely on the system installed packages but only on the one specified in nix-shell. It is a great way to validate that the declaration you wrote for your project can work everywhere you can run Nix. It makes continuous delivery what it should be a way to run workflows. It is not like that; for me, it is a constant translation exercise between Jenkinsfile, bash, YAML for GitHub Actions, drone, or Travis. With Nix, you declare the environment, and you can run it everywhere. For example, you can set shebang in your scripts, leaving to nix-shell the responsibility for satisfying the dependencies it needs:

```sh
#!/usr/bin/env nix-shell
#!nix-shell -i bash ../shell.nix

make deploy
```

If you don't want to translate from Nix to GitHub actions, there is an action that installs Nix, combined with the right shebang; you can reuse the shell.nix description for your project. I do that in [gianarb/tinkie](https://github.com/gianarb/tinkie/blob/master/.github/workflows/ci.yaml):

```yaml
name: For each commit and PR
on:
  push:
  pull_request:

jobs:
  validation:
    runs-on: Ubuntu-20.04
    env:
      CGO_ENABLED: 0
    steps:
    - name: Checkout code
      uses: actions/checkout@v2
    - uses: cachix/install-nix-action@v12
      with:
        nix_path: nixpkgs=channel:nixos-unstable
    - run: ./hack/build-and-deploy.sh
```

As you can see, I am not using actions to install the dependencies I need, I use `cachix/install-nix-action@v12` to get Nix, and everything is managed as I do locally. Something I don't have to maintain, I suppose.

Mitchell Hashimoto uses Nix to provision its virtual machine, quickly enjoying the Linux environment when it comes to development compared with MacOS.

{:refdef:.text-center}
![Mitchel Hashimoto tweet: I switched my primary dev environment to a graphical NixOS VM on a macOS host. It has been wonderful. I can keep the great GUI ecosystem of macOS, but all dev tools are in a full screen VM. One person said “it basically replaced your terminal app” which is exactly how it feels.](/img/mitchellh-tweet-nixos.png){:.img-fluid}
{:refdef}

### Conclusion

I tend to avoid complications, and I am picky when it comes to the number of tools I have in my toolchain, but after a few months of observation, I think Nix deserves a place in my daily workflow. I just scratched the Nix surface; I didn't even write my first derivation yet.

It brings back the joy I had a few years ago provisioning infrastructure, which I have lost in the last few years.

{:.small}
Hero image via [Medium.com](https://medium.com/@robinbb/what-is-nix-38375ed59484)
