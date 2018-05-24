---
layout: post
title:  "Docker Bench Security"
date:   2016-11-15 10:08:27
categories: [post]
img: /img/docker.png
tags: [docker, scaledocker, security, automation, devops, open source, docker
captain]
summary: "Container security is a hot topic because today containers
are everywhere also in production. It means that we need to trust this
technology and start to think about best practices and tools to make our
container environment safe."
---
Frequently, best practices help you to have a safe environment,
[docker-bench-security](https://github.com/docker/docker-bench-security) is an
open source project that runs in a container and scans your environment to
report a set of common mistakes like:

* Your kernel is too old
* Your docker is not up to date
* Some Docker daemon configurations are not good enough to run a production environment
* Your container runs 2 processes
* and others

<div class="post row">
  <div class="col-md-12">
      {% include book-adv-lb.html %}
  </div>
</div>

It’s a great idea to run it at some stage in each host to have an idea about
the status of your environment. To do that you can just use this command when
running a container

```bash
$ docker run -it --net host --pid host --cap-add audit_control \
    -v /var/lib:/var/lib \
    -v /var/run/docker.sock:/var/run/docker.sock \
    -v /usr/lib/systemd:/usr/lib/systemd \
    -v /etc:/etc --label docker_bench_security \
    docker/docker-bench-security
```

A good way to start is your run  it in your local environment. Run the command
and check what you can do to make your local environment safe.  This tool is
open source on GitHub and it’s also a great example of collaboration and how a
community can share experiences to help other members to improve an
environment.  This is a partial output:

```bash
Initializing Thu Nov 24 21:35:24 GMT 2016

[INFO] 1 - Host Configuration
[WARN] 1.1  - Create a separate partition for containers
[PASS] 1.2  - Use an updated Linux Kernel
[PASS] 1.4  - Remove all non-essential services from the host - Network
[PASS] 1.5  - Keep Docker up to date
[INFO]       * Using 1.13.01 which is current as of 2016-10-26
[INFO]       * Check with your operating system vendor for support and security maintenance for docker
[INFO] 1.6  - Only allow trusted users to control Docker daemon
[INFO]      * docker:x:999:gianarb
[WARN] 1.7  - Failed to inspect: auditctl command not found.
[WARN] 1.8  - Failed to inspect: auditctl command not found.
[WARN] 1.9  - Failed to inspect: auditctl command not found.
[INFO] 1.10 - Audit Docker files and directories - docker.service
[INFO]      * File not found
[INFO] 1.11 - Audit Docker files and directories - docker.socket
[INFO]      * File not found
```
Sometime to have a good result you just need to run a single command.

This article is part of "Drive your boat like a Captain". It's a book about
Docker in production, how manage a cluster of Docker Engine with Swarm and what
it means to manage a production environment today.

Keep in touch to receive news about the book
[scaledocker.com](http://scaledocker.com).  If you are looking for a Docker
Getting Started you can also look on the first chapter that I released [Docker
The
Fundamentals](http://localhost:4000/blog/docker-the-fundamentals)
