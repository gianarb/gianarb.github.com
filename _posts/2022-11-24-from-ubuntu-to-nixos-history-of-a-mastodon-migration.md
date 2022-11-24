---
layout: post
img: /img/1280px-NixOS_logo.png
title: "From Ubuntu to NixOS the story of a mastodon migration"
date: 2022-11-24 10:08:27
categories: [post]
tags: [nixos, mastodon]
summary: "Twitter is not at its best. Developers are looking for an alternative as many others. Mastodon with its decentralized and feel of ownership is raising in popularity. I started with a hands crafted self hosted Ubuntu server because I felt the pressure about joining as early as possible but the end goal was to use NixOS for that. This is the story of how I moved my Mastodon instance to NixOS"
---
Do you know that Elon Musk bought Twitter for a lot of money? As a consequence many people are trying to figure out what to do. Developers quickly turned to Mastodon.

I decided to self host my server, you can interact with me on Mastodon as [@gianarb@m.gianarb.it](https://m.gianarb.it/@gianarb).

I do not have strong opinions about a decentralized system, I think it is another way to build a distributed system, experimenting with it is an opportunity, nothing more right now. I never liked the idea to sell my identity for free on social media but having a presence online proved to be crucial for my career and I don’t want to miss that.

Mastodon pushes many people, myself included, to think about: “should I host my own server?”. In my opinion it is an important question because it forces us to make our hands dirty again. We all know how comfortable GitHub pages are. You can set up your own static website in a minute, for free but it lowered my enthusiasm for technology because it makes things too easy. If you answered “yes” and now you are hands down trying to run your own Mastodon I hope you are having fun and that you are learning something that raises your excitement for how computers work. “Hosting more of my own things” was on my bucket list and Mastodon pushed me down the stairs.


## At the beginning it was all about Ubuntu

NixOS has been my way to go for everything since the last two years, but I am not good at it. I tried to run my own Mastodon for a few days but I was not getting anywhere, got stuck trying to figure out how to properly manage secrets, and the machine lifecycle, how to deploy, how to interact with the tootctl, everything was a big unknown. Mastodon itself was a big unknown too. So I decided to step back and run my own instance following a random blog post:[ “How to Install Mastodon on Ubuntu 22.04/20.04 Serves”](https://www.linuxbabe.com/ubuntu/how-to-install-mastodon-on-ubuntu). Not sure if it is the best one out there but it gave me a Mastodon to play with in 10 minutes. Don’t need to tell me about infrastructure such as code, immutability and so on, this environment teached me all of this crap works, Mastodon is a bit more familiar and my end goal is still to figure out how to run it with NixOS.


## Build a migration plan

[“Setting up your own Mastodon instance with Hetzner and NixOS”](https://page.romeov.me/posts/setting-up-mastodon-with-nixos/) by romeov explained how to get Mastodon running on NixOS. A few lines of configuration and the NixOS Mastodon Module configures Postgres, Redis, Nginx with TLS, and Mastodon itself for me. It is not the only way to go, the module supports running dedicated pools of those services as well but for my single user and single server configuration it is more than enough. So I started planning how to migrate my own server following the official [Mastodon documentation](https://docs.joinmastodon.org/admin/migrating/) and it ended up looking like this:

1. Provision a very basic NixOS instance (called beetroot from now on)
2. Stop mastodon services (web, sidekiq, stream) in the Ubuntu box
3. Take a backup of Postgres with the suggested command:` pg_dump -Fc mastodon_production -f backup.dump`
4. Create a tar.gz archive for the system directory in Mastodon
5. Move the archive and the sql backup to beetroot via tailscale file:` tailscale file cp` public-system.tar.gz beetroot: 
6. Get the two files from beetroot via tailscale: `tailscale file get .`
7. Untar the system directory
8. Stop the mastodon systemd services, drop the mastodon database from beetroot and replace it with the backup from the Ubuntu server
9. Restart the mastodon services via systemd and have fun


## How it went

The plan was solid! [CULTPONY](https://pony.social/@cult) looked at it briefly as well, so we are good!

But you know, in reality there are many unknowns. There is only one way to figure them out, time to stop making plans, it is time to break them!


```nix
services.mastodon = {
  enable = true;
  localDomain = "PUT-YOUR-DOMAIN-HERE e.g. computing.social";
  configureNginx = true;
  smtp.fromAddress = "";
};
```


First, when I initialized the NixOS Mastodon module it starts an Nginx server because Mastodon requires TLS, it uses Let’s Encrypt for that and this requires the DNS record to point to the NixOS instance otherwise Let’s Encrypt won’t be able to close the loop, but I can’t point the DNS to a not yet ready instance because who knows if I am gonna be able to make it today, tomorrow or never! I decided to tell the Mastodon module to skip Nginx configuration for now setting` services.mastodon.configureNginx=false;`

Technically there is [another way](https://discourse.nixos.org/t/nixos-deploy-in-a-vm-how-to-test-https-website-acme-lets-encrypt/8876) to do it but it did not work for me and I still don’t know why. Let me know if you figure it out because it will be way more comfortable to get a self signed certificate so we can test without having to change DNS.

In the process of making the tar archive for the system directory I saw it contained a directory called cache, huge like multiple GBs. Cache to me means ephemeral, easy to rebuild, safe to wipe. So I did it! To be fair, I knew I was doing something stupid. And I knew the dirty way to go requires at least to move the directory, to keep it around until realization of the silly fact that cache means something important that should not be lost! Too late! I lost it, my Mastodon instance was then empty of all the avatars and profiles images, not a great start. After some googling and some struggling, [I was able to build it back](https://github.com/mastodon/mastodon/discussions/21305#discussioncomment-4218030) (if you have a more official answer for this issue let us know there). My 2 cents: move the cache folder around, way easier than figure out how to get it back. Oh get another 2 cents, remember to check file permissions when you do this (guess why I know).

Everything now was set and ready to receive traffic, so I pointed the DNS to the new server, I set my host file to get routed to it quickly, I changed services.mastodon.configureNginx to true and I waited.

Ok! This is how it went after a day of struggling obviously! Last time I used postgres was probably 6 years ago. pg_dump, pg_restore are easy but I had to figure out how to authenticate properly, Ubuntu was set up to run over 127.0.0.1, the NixOS Mastodon Module by default provisions Postgres with [auth trust](https://www.postgresql.org/docs/current/auth-trust.html) trust and with a socket entrypoint. It means that authentication does not require a password and it is based on a UNIX user. For example the Linux user Postgres owns and has access to the database owned and managed by Postgres. The NixOS Mastodon module creates a mastodon user in Linux with access to the mastodon file (the system directory for example) and with access to its own mastodon Postgres database. Nothing that looks like rocket science but still, it took me some time to figure it all out.

How to manage password in NixOS is a question I don’t feel comfortable answering yet and it blocked me at the beginning when I was trying to setup my own instance because I wanted to manage tailscale auth key for example automatically, or when thinking about how to manage the connection between mastodon web and postgres. Currently my answer is to avoid passwords. It works for now, but I know it won’t be the right answer for the following articles in this series that will probably title: “Mastodon monitoring a success story” where I will share how to configure the monitoring and observability pipeline for my instance with Grafana Cloud, but this is a story for another time.

Point 7 of the migration plan was about untar the system directory, but I realized I didn’t know where to place it. [Looking at the NixOS module](https://github.com/NixOS/nixpkgs/blob/master/nixos/modules/services/web-apps/mastodon.nix#L32) there is a path for that: 


```terminal
PAPERCLIP_ROOT_PATH = "/var/lib/mastodon/public-system";
```


But what does it look like? And what is PAPERCLIP_ROOT_PATH? Is it really what I think it is? It was not clear to me and only `var/lib/mastodon` was there in the system because the public-system folder gets created when Mastodon is actually in use. So I had to take a step back and I created a vanilla e2e working Mastodon instance to figure it out. At the end it **obviously** look like it should be, but who knew that!


```terminal
[nix-shell:/var/lib/mastodon]# tree -L 2
.
├── public-system
│   ├── accounts
│   ├── cache
│   ├── custom_emojis
│   └── media_attachments
└── secrets
```



## Show me the code

Currently I published the NixOS configuration for beetroot as part of my [dotfiles](https://github.com/gianarb/dotfiles/tree/main/nixos/machines/beetroot) along with the other NixOS configurations for my Thelio workstation and for the Asus Zenbook I use at home. It uses [flake](https://nixos.wiki/wiki/Flakes) and [deploy-rs](https://github.com/serokell/deploy-rs). It targets a Linode shared CPU virtual machine and that’s why, as you can see in the hardware-configuration NixOS detected Qemu as hardware.


```nix
deploy.nodes.beetroot = {
      hostname = "139.162.167.171";
      sshUser = "root";

      profiles.system = {
        user = "root";
        path = deploy-rs.lib.x86_64-linux.activate.nixos
          self.nixosConfigurations.production;
      };
    };
```


Do not ask me about my deploy preference when it comes to Nix, deploy-rs is just the one I figured out, I may switch to Nixops because it is a bit more standard, they work similarly from a configuration standpoint but in theory deploy-rs is designed with profiles in mind, to deploy single users, something that I don’t think I need. But it works well enough for now.

If you look inside the flake.nix file you see two different nixosConfigurations, production and vm, both importing the same `configuration.nix`. Production is deployed via deploy-rs and vm is used for testing purposes with: `nixos-build build-vm -flake .#vm`

I didn’t find a good use of it just yet, I am currently blocked by the acme certificate and because I am lazy. I am not sure if it is needed for Linode Shared CPU since it is a VM as well and it detects Qemu as a hypervisor. Time will help me figure it out.

At the beginning I developed this configuration outside of my dotfiles. Mainly because I didn’t know what to expect from it. Now that Mastodon is up and running and this configuration is in use I feel more confident. Even if I have a lot I want to do I decided to move it in my dotfiles to have access to other NixOS components there. I need to add a secret to authenticate to Grafana Cloud, probably with [agenix](https://github.com/ryantm/agenix) in its own private repo imported via flake so I won’t have my password shared with you all (forgive me, it is not you, it is me or something else I don’t know), I want to move the cache directory and postgres data to a ZFS pool as well, but not now, right now I want to enjoy my running instance.

## Now what?

This is everything I have learned for now migrating from Ubuntu to NixOS. I want to be clear, even if the core of this article looks like a bunch of mistakes I am not frustrated, I think the NixOS Mastodon Module is comfortable to use and well written. The challenges I described come from a rusty and inexperienced ops person. The module lacks documentation around operational experience, how to use it and what it provides but it is reasonable and I hope those notes will help to improve it and will push me to contribute back to the official documentation.

When I mentioned Prometheus and Grafana I shared that I am thinking of writing a series of posts about this topic, those are the one I have currently ongoing:

* Monitoring success story (probably with a deep dive in Password management on its own)
* NixOS configuration, GitOps and machine lifecycle (this is about how I manage my NixOS configuration, how I deploy NixOS and so on)
* Data management with ZFS
* Mastodon update from 3.x to 4.0

Your support and interest will push me forward writing all of them so let me know what you think about this one, the following topics and if you would like to read something else, like my journey with Linode, since I decided to try it out running this Mastodon instance there.

I would like to thanks all the writers behind the documentations, articles, GitHub discussions I have linked, and all the GitHub issues, StackOverflow questions, and GitHub repositories I have looked at to resolve my unknowns, sharing is caring! Thanks [@hazelweakly](https://hachyderm.io/@hazelweakly) for your early review!
