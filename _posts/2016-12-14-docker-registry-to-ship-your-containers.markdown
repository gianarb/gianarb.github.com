---
layout: post
title:  "Docker registry to ship and manage your containers."
date:   2016-12-14 08:08:27
categories: [post]
img: /img/docker.png
tags: [docker, scaledocker, docker captain, security, book, ebook, learn]
summary: "Build and Run containers is important but ship them out of your laptop is the
best part! A Registry is used to store and manage your images and all your
layers. You can use a storage to upload and download them across your servers
and to share them with your colleagues."
---
Build and Run containers is important but ship them out of your laptop is the
best part! The Registry is a very important tool that requires a bit more
attention. A Registry is used to store and manage your images and all your
layers. You can use a storage to upload and download them across your servers
and to share them with your colleagues.

The most popular one is hub.docker.com
it contains different kind of images: public, official and private.  You can
create an account and push your images or build them for example from a github
or bitbucket repository. The integration with GitHub and Bitbucket is called
“Automated Builds”. it allows you to create a continuous integration
environment for your images, when you select “Create” and “Automated Builds”
you can specify a repository and a path of your Dockerfile. You can specify
more that one path from the same repository to build more that one image tag.
In this way you can centralize and build your images every time that a new
change is pushed into the repository. It also supports organizations to split
your images in different groups and manage visibility of them in case of
private images.

By default any developer can push their image to registry and
they'll be public and free for other developers to use.  Official images are
those public images selected and maintained from specific organization or
member of the communities, the idea is that they have a better quality or who
provides them are usually involved into the product development. A set of
official images are: Nginx, Redis, MySql, PHP, Go and so on
[https://hub.docker.com/explore](https://hub.docker.com/explore).

Docker Hub offers different plan to store
private images, all people has one for free but if you need more you can pay a
plan and store more.

Registry is not just  a tool but it’s a specification, it
describe how expose capabilities has pull, push, search and so on. This
solution allowed the ecosystem to implement these rules in other projects and
save the compatibility with the Docker Client and with the other runtime engine
that use this capability. It’s for this reason that other providers as
Kubernetes, Cloud Foundry supports download from Docker Hub. This specification
has 2 version, v1 and v2 the most famous registries implement both standard and
they fallback from v2 to v1 for  features that are not supported yet. For
example Search is not supported at the moment into the v2 but only in v1.

If you are looking for an In-House solution you have different tools available
online. The first one is distribution. It is provided by Docker, it’s open
source and offers a very tiny registry that you can start and store in your
server. It also supports different storage like the local filesystem and S3.
This feature is very interesting because usually the size of the images and the
number of layers increase very fast and you also need to keep them safe with
backup and redundancy policies for high availability. This is very important if
your environment is based on containers it means that your register is a core
part of your company. Let’s start a Docker Distribution:

```bash
$ docker run -d -p 5000:5000 --name registry registry:2
```

In docker the default registry is hub.docker.com it means that when we push or
pull an image we are reaching this registry:

```bash
$ docker pull alpine
```

To push our images in another registry we need to tag them:

```bash
$ docker tag alpine 127.0.0.1:5000/alpine
```

With this command you tagged the alpine to a registry 127.0.0.1:5000 because as
we said in previous chapters the name of the image contains a lot of
information:

```
REGISTRY/NAME:VERSION
```

The default registry is hub.docker.com a name could be simple as alpine or with
a username matt/alpine and you can pin a specific build with a version you can
use semver or for example the sha of the commit the default VERSION is latest.

Now that we have a new tag we can push and pull it in from our registry:

```bash
$ docker push 127.0.0.1:5000/alpine
$ docker pull 127.0.0.1:5000/alpine
```

A very important information to remember when you start a customer registry is
that every layers, every build is stored and it’s very easy to have a big
registry, you need to monitor the instance to be sure that your server has
enough disks space and also take care about high availability. In a real
environment the registry it the core of your infrastructure, developers use it
to pull and push build and also to put a version in production. Take care of
your registry.

Other that Docker provided registry there are few alternatives. [Nexus](
https://www.sonatype.com/nexus-repository-sonatype) is a registry manager that
support a lot of languages and packages if you are a Java developer you know
it. Nexus supports Docker Registry API v1 and v2. The Docker registry
specification is young but it has 2 version already.

We can use the image provided by Sonatype and start our Nexus repository:

```bash
$ docker run -d -p 8082:8082 -p 8081:8081 \
    -v /tmp/sonata:/sonatype-work --name nexus sonatype/nexus3
$ docker logs -f nexus
```

When our log tells us that Nexus is ready we can reach the ui from our browser
http://localhost:8081/ or with the IP of your Docker Machine if you are using
Docker for Mac/Windows or Docker in Linux. The default credentials are username
admin and password admin123.

<img class="img-fluid" src="/img/docker-registry/nexus-image-loaded.png">

First of all we need to create a new Hosted Repository for Docker, we need to
press the Settings Icon top left of the page, Repositories and Create
Repository. I called mine mydocker and you need to specify an HTTP port for
that repository, we shared port 8082 during the run and for this reason I chose
8082.

<img class="img-fluid" src="/img/docker-registry/nexus-create-repo.png">

Nexus has different kind of repositories Host means that it’s self hosted but
you can also create a Proxy Repository to proxy for example the official Docker
Hub.
Now we need to login to out docker registry:

```bash
$ docker login 127.0.0.1:8082
```

Now we can tag an alpine and push the tag into the repository

```bash
$ docker tag alpine 127.0.0.1:8082/alpine
$ docker push 127.0.0.1:8082/alpine
```

You can go in Assets click on mydocker repository and see that your image is
correctly stored.

[GitLab](https://about.gitlab.com/) has also a container registry. GitLab uses
it to manage build and it’s available for you from version 8.8 if you are
already using this tool.

<p class="text-muted">Thanks <a
href="https://twitter.com/kishoreyekkanti" target="_blank">Kishore
Yekkanti</a>, <a href="https://twitter.com/liuggio" target="_blank">Giulio De
Donato</a> for your review.</p>

<div class="post row">
  <div class="col-md-12">
      {% include docker-the-fundamentals.html %}
  </div>
</div>
