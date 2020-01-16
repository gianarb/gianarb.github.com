---
title: Continuous Profiling Go Application Running in Kubernetes
date: 2020-03-24
slide:
embedSlide:
video:
embedVideo:
eventName: DevOps Pro Europe
eventLink: https://devopspro.lt/
city: Vilnius, Lithuania
---
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
