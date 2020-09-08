---
layout: post
title:  "About your images, security tips"
date:   2016-12-28 08:08:27
categories: [post]
img: /img/container-security.png
tags: [docker, scaledocker, book, docker captain, ebook, security]
summary: "Everything unnecessary in your system could be a very stupid vulnerability. We
already spoke about this idea in the capability chapter and  the same rule
exists when we build an image. Having  tiny images with only what our
application needs to run is not just a goal in terms of distribution but also
in terms of cost of maintenance and security."
---

Everything unnecessary in your system could be a very stupid vulnerability. We
already spoke about this idea in the capability chapter and  the same rule
exists when we build an image. Having  tiny images with only what our
application needs to run is not just a goal in terms of distribution but also
in terms of cost of maintenance and security.  If you have some small
experience with docker already you probably know the
[alpine](https://hub.docker.com/_/alpine/) image. It is build
from the Alpine distribution and it’s only 5MB size, if your application can
run inside it then this is a very good optimization that you can do.  What
about your binaries? Can your application run standalone? If the answer is yes
you can think about a very very minimal image. scratch is usually used as a
base for other images like debian and ubuntu but you can also use it to run
your golang binary and let me show you something with our micro application.
In the [release page](https://github.com/gianarb/micro/releases/tag/1.0.0),
there are a list of binaries already compiled and ready to be used. In this
case we can download the linux_386 binary.

<img class="img-fluid" src="/img/security-image/micro-release.png">

```bash
curl -SsL https://github.com/gianarb/micro/releases/download/1.0.0/micro_1.0.0_linux_386 > micro
```

And we know we can include this binary in the scratch image with this Dockerfile

```bash
FROM scratch

ADD ./micro /micro
EXPOSE 8000

CMD ["/micro"]
```

```bash
docker build -t micro-scratch .
docker run -p 8000:8000 micro-scratch
```

The expectation is an http application on port 8000 but the main difference is
the size of the image, the old one from alpine is 12M the new one is 5M.

The scratch image is impossibile to use with all applications but if you have a
binary you can remove a lot of unused overhead.

Another way to understand the status of your image is to scan it to detect
security vulnerabilities or exposures. Docker Hub and Docker Cloud can do it
for private images.  This is a great feature to have in your pipeline to scan
an image after a build.

CoreOS provides an open source project called [clair](
https://github.com/coreos/clair) to do the same in your environment.

It is an application in Golang that exposes a set of HTTP API to
pull, push and analyse images. It downloads vulnerabilities from different
sources like [Debian Security
Tracker](https://security-tracker.debian.org/tracker) or [RedHat Security
Data]( https://www.redhat.com/security/data/metrics/). Each vulnerability is
stored in Postgres. Clair works like static analyzer, this means that it
doesn’t need to run our container to scan it but it persists different checks
directly into the filesystem of the image.

```bash
docker run -it -p 5000:5000 registry
```

With this command we are running a private registry to use as a source for the
image to scan

```bash
docker pull gianarb/micro:1.0.0
docker tag gianarb/micro:1.0.0 localhost:5000/gianarb/micro:1.0.0
docker push localhost:5000/gianarb/micro:1.0.0
```

Now that we pushed in our private repo the micro image we can setup clair.

```bash
mkdir $HOME/clair-test/clair_config
cd $HOME/clair-test
curl -L https://raw.githubusercontent.com/coreos/clair/v1.2.2/config.example.yaml -o clair_config/config.yaml
curl -L https://raw.githubusercontent.com/coreos/clair/v1.2.2/docker-compose.yml -o docker-compose.yml
```
Modify `$HOME/clair_config/config.yml` and add the proper source
`postgresql://postgres:password@postgres:5432?sslmode=disable`

Now you can run the following command to start postgres and clair:

```bash
docker-compose up
```

To make our test easier, we will use another CLI called hyperclair that is just
a client to work with this application. If you are using Mac OS, you can follow
the above commands, if you are in another OS you can find the correct url in
the release page

```bash
curl -SSl https://github.com/wemanity-belgium/hyperclair/releases/download/0.5.2/hyperclair-darwin-386 > ~/hyperclair
chmod 755 ~/hyperclair
```

Now we have an executable in ~/hyperclair

```bash
~/hyperclair pull localhost:5000/gianarb/micro:1.0.0
~/hyperclair push localhost:5000/gianarb/micro:1.0.0
~/hyperclair analyze localhost:5000/gianarb/micro:1.0.0
~/hyperclair report localhost:5000/gianarb/micro:1.0.0
```

The generated report looks like this:

<img class="img-fluid" src="/img/security-image/report-clair.png">

Hyperclair is just one of the implementations of clair, you can decide to use
it or build your own implementation in your pipeline.
