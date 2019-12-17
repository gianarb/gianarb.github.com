---
img: /img/gianarb.png
layout: post
title: "Kubernetes is not for operations"
date: 2019-09-18 08:08:27
categories: [post]
tags: [kubernetes, developer, kubeassemble]
summary: "Kubernetes it not for operations. It democratize resources and
workloads. It can be the solution to bring developers closer to ops. But YAML is
not the answer."
changefreq: daily
---
I work in tech from 8 years. It is not a lot but it is something.
I started as a PHP developer doing CMS with MySQL and  things like that.

Where I saw what I was capable to do with a set of API requests to AWS I enjoyed
it and I moved to what people called DevOps probably.

I like communities and people so Docker was everywhere and I became a Docker
Captain for my passion about delivery, and development workflow, containers but
with always developers in mind. That's what I like do to. Write code.

> The complexity not hidden behind Kubernetes, or not solved by who runs
> Kubernetes in your company creates that friction.

Everyone that was/is in the containers space more or less touched Kubernetes.
I did it, I enjoyed to look at the patterns used by it like control theory,
reconciliation loops and so on.

In the last couple of years I saw a lot of company moving to Kubernetes
and I worked on that path in InfluxData as well. Yes we use Kubernetes obviously!

I have always sawed friction from developers forced to onboard Kubernetes (no
developer will do it otherwise). First because everybody uses YAML and I
think yaml is just the wrong answer for your problem, nothing personal with it.

What developers are happy to do is to **write code** that runs in production and
that gives them good challenge to debug and fix. **write code** is in bold
because that's what we like most. At least the majority of us.

The complexity not hidden behind Kubernetes, or not solved by who runs
Kubernetes in your company creates that friction.

Running Kubernetes is not hard, we have tutorials, companies, contractors, cloud
providers that can help us out. It is a set of binaries and a database. We run
them since ages! There are a good amounts of them, and they need to be configured,
connected and there are also a lot of different combinations, but that's fine.
We are used to playing with mobile apps, wordpress plugins and so on.

When I think about myself as a developer I understand why there is this
friction, if I was not passionate about containers at the right time to try out
Kubernetes I probably even had that friction myself.

It does not help me to write better code or to do something different compared
with updating systemd service one by one via `ssh`. I bet developers working with
Kubernetes in a system under real load will likely get back to `ssh` to the
servers one by one deploying their new version of the application to have all
the control and visibility they can. That's what a lot of developers
tries to achieve when I look at them using Kubernetes.

What Kubernetes does very well is democratize ops, it provides a common set of
concepts that we can use to run applications and very good API that abstract the
concrete implementation of containers, VMs, workload, ingress, dns and so on.

We should not west our time trying to run it, we should spend time to make it
usable in our company because that's we can get from k8s.

## my recipe

I do not have a recipe, a product or something ready to go. But I think there
are two directions I would like to see and to try with a team.

### leave yaml

YAML is the wrong answer, it is good to make an impact and to write a
document that everyone can read, but your company is not "everybody", you are
pretty unique. You should use the K8S API. I didn't have time to make a public
prototype yet but I will do I promise. You should use the language you know
better! I have a lot of experience with go, so my suggestion is to replace yaml
with real code, real function and so on. From Kubernetes 1.16 `kubectl diff`
runs server side. Sweet!

### split spec file by team

It is very easy to end up with a single Kubernetes YAML file that is crazy long.
That file contains everything you run. Across teams, responsabilities and
people. Do not do it. Split it in different files or repositories by team or
application owners.

DevOps, SRE, Sysadmin, reliability, penguins or what ever you call the team that
owns the underline architecture will have the Yaml related to the foundation of
the infrastructure. The content of it is not important for other teams, they
will only write and see what matters to them.

This approaches will decrease complexity for developers making them probably
less worried to screw up part of the infrastructure that is not related to their
work.

## Conclusion

If you are a developer please develop good code! If you own Kubernetes in your
company make it to work for your users.
