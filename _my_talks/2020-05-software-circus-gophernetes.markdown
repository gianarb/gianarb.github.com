---
layout: talk
title: 'How to become a Gophernetes'
date: 2020-05-19T11:00:00Z
slide: "https://speakerdeck.com/gianarb/how-to-become-a-gophernetes"
embedSlide: "a858d5dc203c456cb52fac3c6b9a31f3"
eventName: Software Circus
eventLink: https://softwarecircus.io
youtubeID: XPSapVJSy6E
video: https://www.youtube.com/watch?v=XPSapVJSy6E
city: ""
links: {}

---
<div class="alert alert-warning" role="alert">
    We got some technical issue with the audio and, I am so bluish! Probably I
    got too much in the cartoon mode!! Sorry about that
</div>

![Image provided by Software Circus](/img/software-circus-alice.png){:.img-fluid.w-25.float-left}

The Go community well knows what a Cryptogopher is! Today is the way where you
will learn about how a Gophernetes looks like! Kubernetes is all about
extendibility. That’s why every cloud provider is able to plug their network
implementation, storage layer or compute platform to it. But in order to do so,
you have to write code to glue your platform or external project with Kubernetes
itself. Gophers are in a unique position when it comes to writing code for
Kubernetes because even if there is an API that gives you the opportunity to
write integration in any language, it is written in Go, and that’s a huge
benefit. This talk is for Gopher that want to become Kubernetes developers also
called gopherneters. I participated in various efforts around integration at
storage layers, with the container storage interface, or container runtime
interface and recently with cluster-api, the abstract that drives the Kubernetes
provisioning in a declarative way. It means that I wrote a good amount of Custom
Resource Definitions (CRDs), Shared Informers and so on. It is a jungle and I
will share what I learned in terms of best practices, testing to write solid
Kubernetes integrations.

{:refdef: class="card-text small"}
This Alice in Wonderland image comes from
[SoftwareCircus.io](https://softwarecircus.io)
{: refdef}
