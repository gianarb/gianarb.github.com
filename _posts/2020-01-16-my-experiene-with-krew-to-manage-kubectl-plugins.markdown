---
img: /img/kubernetes.png
layout: post
title: "My experience with Krew to manage kubectl plugins"
date: 2020-01-16 09:08:27
categories: [post]
tags: [kubernetes]
summary: "Kubectl plugins are extremely useful to provide a set of friendly
utilities to interact with kubernetes in your environment. Krew is a project
that helps you managing the plugin lifecycle. I have to add profefe to it and
this is what I learned."
changefreq: daily
---
I wrote a good number of kubectl plugins so far, but there is a lot more I can
do with them. Every time I write a new one I discover something new and that is
why I am always excited to see what will happen with the next one.

["Unit Testing Kubernetes Client in
Go"]([https://gianarb.it/blog/unit-testing-kubernetes-client-in-go](https://gianarb.it/blog/unit-testing-kubernetes-client-in-go))
and ["Kubectl flags in your kubectl
plugin"]([https://gianarb.it/blog/kubectl-flags-in-your-plugin](https://gianarb.it/blog/kubectl-flags-in-your-plugin))
are two of the lessons learned along the way.

With
[kubectl-profefe]([https://github.com/profefe/kube-profefe](https://github.com/profefe/kube-profefe))
I decided to have a look at
[krew]([https://github.com/kubernetes-sigs/krew](https://github.com/kubernetes-sigs/krew)).
It is a package manager for kubectl plugins. It is a plugin itself with the end
goal to help you installing and managing the lifecycle for your plugin.


```
$ kubectl krew install profefe
```


It gives you the ability, with a single command to install, update or delete the
kubectl-profefe cli command.

Twitter got pretty excited recently about
[kubectl-tree]([https://github.com/ahmetb/kubectl-tree](https://github.com/ahmetb/kubectl-tree)),
a plugin from
[@ahmetb]([https://twitter.com/ahmetb](https://twitter.com/ahmetb)) an old
friend of mine and active Kubernetes contributor and maintainer for krew as
well. It helps you to visualize kubernetes resources as a tree to simplify the
comprehension of the hierarchy and the connection between resources.

Two other examples that I would like to mention are from @ahmetb too. Kubectl
plugins don’t need to be extremely complicated, but you always have to keep in
mind the mantra "usability first." It doesn’t matter how many lines of code you
write: the end goal should be to develop something usable and well-integrated
with kubernetes! `kubectl ctx` and `kubectl ns` are fabulous examples of
something easy but helpful. We switch between context and namespace more than
once a day between production clusters, local development, and so on. It is not
a very complicated thing to do natively: for example, changing context with
kubectl is just a matter of typing:


```
$ kubectl config use new-context
```


Worst case scenario for the namespace, you have to type the `-n` flags every
time you run a kubectl command that is not in the namespace you have set by
default for the context you are using.

But `kubectl ctx` and `kubectl ns` simplify this process even more. You only
have to type:

```
$ kubectl ctx new-context
```

Or

```
$ kubectl ns new-namespace
```

If you are developing an open-source kubectl plugin and you need a friendly and
easy way to distribute it, you should have a look at krew. The publication
process is straightforward, [this is the
PR]([https://github.com/kubernetes-sigs/krew-index/pull/415](https://github.com/kubernetes-sigs/krew-index/pull/415))
I had to submit for profefe, you have to type some YAML as usual.
