---
img: /img/cncf-logo.png
layout: post
title: "CNCF Webinar: Continuous Profiling Go Application Running in Kubernetes"
date: 2020-03-27 09:08:27
categories: [post]
tags: [golang, pprof, kubernetes]
summary: "Slides, videos and links from a webinar I have with the CNCF about
kubrenetes, profefe, golang and pprof."
changefreq: daily
---

<div class="embed-responsive embed-responsive-16by9">
    <iframe class="embed-responsive-item" src="https://www.youtube.com/embed/SzhQZQ6VGoY" allowfullscreen></iframe>
</div>

Microservices and Kubernetes help our architecture to scale and to be
independent at the price of running many more applications. Golang provides a
powerful profiling tool called pprof, it is useful to collect information from a
running binary for future investigation. The problem is that you are not always
there to take a profile when needed, sometimes you do not even know when you
need to one, that’s how a continuous profiling strategy helps. Profefe is an
open-source project that collect and organizes profiles. Gianluca wrote a
project called kube-profefe to integrate Kubernetes with Profefe. Kube-profefe
contains a kubectl plugin to capture locally or on profefe profiles from running
pods in Kubernetes. It also provides an operator to discover and continuously
profile applications running inside Pods.

A bunch of links for you:

* Video coming soon
* [My article: Continuous profiling in Go with Profefe](/blog/go-continuous-profiling-profefe)
* [My article: Continuous Profiling Go applications running in Kubernetes](/blog/continuous-profiling-go-apps-in-kubernetes)
* [Google-Wide Profiling: A Continuous Profiling Infrastructure for Data Centers](https://research.google/pubs/pub36575/)
* [Profefe on Github](https://github.com/profefe/profefe)
* [Kube Profefe on Github](https://github.com/profefe/kube-profefe)
* [google/pprof](https://github.com/google/pprof) library on GitHub
* [Work in progress documentation! help me out!](https://kubernetes.profefe.dev)

<div class="embed-responsive embed-responsive-16by9">
    <iframe class="embed-responsive-item" src="//speakerdeck.com/player/ff55e041659945bca5d31013bd999c28" allowfullscreen></iframe>
</div>

