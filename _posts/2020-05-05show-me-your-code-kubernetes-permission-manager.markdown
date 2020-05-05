---
img: /img/show-me-your-code-logo.png
layout: "show-me-your-code"
title: "Show Me Your Code with Enrique Paredes: Kubernetes Permission Manager"
date: 2020-05-05 09:00:27
categories: [post]
tags: ["show-me-your-code", "kubernetes"]
summary: "Enrique will share tips and code around kubernetes permission manager
a project that brings sanity to Kubernetes RBAC and Users management, Web UI
FTW"
twitch: channel=gianarb
addeventID: Yh4790232
changefreq: daily
---

When: Saturday 7th May 10-11pm GMT+2

Kubernetes when it comes to authentication and authorization is extremely
complicated.

I think its philosophy is well described in  the documentation:

> Normal users are assumed to be managed by an outside, independent service. An
> admin distributing private keys, a user store like Keystone or Google
> Accounts, even a file with a list of usernames and passwords. In this regard,
> Kubernetes does not have objects which represent normal user accounts. Normal
> users cannot be added to a cluster through an API call.

Kubernetes has user but they have to come from the outside, it is not its
business to care about them. For authorization it uses RBAC and you have a very
long list of possibilities and combinations between actions like: LIST, WATCH,
CREATE, DELETE and resources: pods, deployments, ingress, services...

Sighup is a company based in Italy and well known for their contributions to the
Cloud Native ecosystem. One of their last project is called
[permission-manager](https://github.com/sighupio/permission-manager), it is open
source and it can be describe as follow: "it is a project that brings sanity to
Kubernetes RBAC and Users management, Web UI FTW".

I will host Enrique ([@twitter](https://twitter.com/iknite)) to talk about the
challenges he had when writing such a crucial project, hoping to see some code!

Links:

* [Kubernetes
  Authentication](https://kubernetes.io/docs/reference/access-authn-authz/authentication/)
* [Sighup webiste](https://sighup.io/)
