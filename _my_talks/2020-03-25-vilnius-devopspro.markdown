---
title: Continuous Profiling Go Application Running in Kubernetes
date: 2020-03-25
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

Links:

* ["Continuous profiling in Go with Profefe"](https://gianarb.it/blog/go-continuous-profiling-profefe)
* ["Google-Wide Profiling: A Continuous Profiling Infrastructure for Data Centers"](https://research.google/pubs/pub36575/)
* ["Stackdriver Profiler"](https://cloud.google.com/profiler/)
* ["Continuous profiling Go app running in k8s"](https://gianarb.it/blog/continuous-profiling-go-apps-in-kubernetes)
* [github.com/profefe/profefe](https://github.com/profefe/profefe)
* [github.com/profefe/kube-profefe](https://github.com/profefe/kube-profefe)
