---
layout: post
title:  "From Docker to Moby"
date:   2017-10-20 10:08:27
categories: [post]
tags: [docker, moby, docker captain, automation, cncf, ambassador]
img: /img/docker.png
summary: "Docker announced during DockerCon a new project called Moby. Moby will
be the new home for Docker and all the other open source projects like
containerd, linuxkit, vpnkit and so on. Moby is the glue for all that open
source code. It will look as an entire platform to ship, build and run
containers at scale."
---
At DockerCon 2017 Austin
[Moby](https://blog.docker.com/2017/04/introducing-the-moby-project/) was the
big announcement.

It created confusion and some communities are still trying to understand what is
going on. I think it's time to step back and see what we have after seven months
after the announcement.

1. `containerd` is living a new life, the first stable release will happen soon.
   It has been donated to CNCF.
2. `notary` is the project behind `docker trust`. I wrote a full e-book about
   [Docker Security](https://scaledocker.com) if you need to know more. This
   also has been donated to the CNCF.
3. github.com/docker/docker doesn't exist anymore there is a new repository
   called github.com/moby/moby .
4. [CLI](https://github.com/docker/cli) has a separate home.
5. docker-ce is the first example of moby assembling. It is made my Docker Inc.

Containers are not a first class citizen in Linux.

<img class="img-fluid" src="/img/container-is-not-real.jpeg"/>

They are a combination of cgroups, namespaces and other kernel features. They are
also there from a lot of year. LXD is one of the first project that mentioned
container but the API wasn't really friendly and only few people are using it.

Docker created a clean and usable api that human beings are happy to use. It
created an ecosystem with an amazing and complete UX. Distribution, Dockerfile,
`docker run`, `docker image` and so on.

That's what Docker is, in my opinion. Other than a great community and a fast
growing company.

What Docker is doing with Moby is to give the ability to competitors, startups, new
projects to join the ecosystem that we built in all these 4 years.

Moby in other hands is giving the ability at Docker to take ownership of the
clean and usable experience. The `Docker CLI` that we know and use every day
will stay open source, but not the moby project's part. It will be owned by
Docker. As I wrote above, the code is already moved out.

Moby allows other companies and organisations to build their
user interface based on what they need. Or to build their product on top of a
open source project designed to be modular.

Cloud and container moves fast Amazon with ECS, RedHat with OpenShift,
Pivotal with Cloud Foundry, Mesos with Mesosphere, Microsoft with Azure
Container Service, Docker with Docker, they are all pushing hard to build
projects around containers to sell them at big and small corporations to make
legacy projects less bored.

> Legacy is the new buzzword

Docker will continue to assemble and ship docker as we know it. The project is
called `docker-ce`:

```
apt-get install docker-ce
docker run -p 80:80 nginx:latest
```

Everything happens down the street, in the open source ecosystem. Moby won't
contain the CLI that we know.

Moby won't have the swarmkit integration as we know it. It was something that
Docker as company was looking to have. Mainly to inject an orchestrator in
million of laptops. Other companies and projects that are not using swarm don't
need it and they will be able to remove it in some way.

Companies like Pivotal, AWS are working on
`containerd` because other the runtime behind Docker it's what matters for a lot
of projects that are just looking to run containers without all the layers on
top of it to make it friendly. ECS, Cloud Foundry are the actual layers on top
of "what runs a container".

Container orchestrator doesn't really care about how or who spins up a container,
they just need to know that there is something able to do that.

It is what Kubernetes does with CRI. They don't care about Docker, CRI-O,
containerd. It's out of scope they just need a common interface. In this case is
a gRPC interface that every runtime should implement. Here a list of them:

* [cri-o](https://github.com/kubernetes-incubator/cri-o)
* [cri-containerd](https://github.com/kubernetes-incubator/cri-containerd)
* [rktlet](https://github.com/kubernetes-incubator/rktlet)

That's a subset of reasons about why everything is happening:

* Docker Inc. will be free to iterate on their business services and projects
  without breaking every application in the world. And they will have more
  flexibility on what they can do as company.
* The transaction between Docker to Moby is the perfect chance to split the
  project to different repositories we already spoke about docker-cli, containerd
  and so on.
* Separation of concern is popular design pattern. Split
  projects on smallest libraries allow us to be focused on one specific scope of the
  project at the time.
  [buildkit](https://github.com/moby/buildkit) is the perfect example. It's the
  evolution of the `docker build` command. We had a demo at the MobySummit and
  it looks amazing!

That's almost it. Let's summarise:

**Are you a company in the container movement?**
You are competing with Docker building container things and you was complaining
about them breaking compatibility or things like that now you should blame the
Moby community.

**Are you using docker run?**
You are fine! You will be able to do what you was doing before.

**Are you a OpenSource guru?**
Maybe you will be a bit disappointed if you worked hard on docker-cli and now
Docker will bring your code out but you signed a CLA, the CLI will stay open
source. Blame yourself.

That's it! Or at least that's what I understood.
