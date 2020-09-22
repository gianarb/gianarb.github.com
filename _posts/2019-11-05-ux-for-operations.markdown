---
img: /img/gianarb.png
layout: post
title: "UX for Operations"
date: 2019-11-05 08:08:27
categories: [post]
tags: [devops, ux, operations, kubernetes]
summary: ""
changefreq: daily
---
> UX for Ops? are you joking?! Ops is ops! It can not look nice or it can be
> made buy other people! They will screw it out my lovely servers! by a pure ops
> guy

I am not a strong sysadmin or a person with ninja skills in operations. Where I
have started to work AWS was already there, I worked with a company that has a
few racks in its basement but I never had the joy to interact with then.

--PICTURE MY BASEMENT--

That's probably why I can't wait to build my own datacenter!

Anyway, I made some infrastructure as code, I got passionate about how many
powerful things I was able to do with a bunch of requests to a cloud providers.
Security groups as firewall, VPC, subnets and route tables. That's my
background.

I think operation has a lot to explore and to improve. I try to make people
excited about "UX for operations".

Everybody tries to empower developers to take care of their own mess. I think it
is the right approach. I think ops people have the knowledge to help us to get
there.. They should find a nice way to outsource they work to developers in a
safe and friendly way.

Just to be clear, runbooks are not friendly, random scripts are not safe. You
can do way better.

I recently wrote a report for O'Reilly called [Extending
Kubernetes](https://get.oreilly.com/ind_extending-kubernetes.html) where the
introduction is about "UX for operations". Let's think about the most popular
"opsy" tools:

* terraform
* kubectl

Both of them has the ability to create plugins. If your company use one of those
almost all the developers has their CLI installed. Even just for validate the
terraform module they are writing or to do a `terraform plan` or a `kubectl
diff`.

Nowadays there is not much value for operators to do by their own, they should
provide tools and APIs that other entities, they can be human or robot can
leverage.

## I have a dream

`kubectl` or `terraform` are installed in the developer laptops if your company
uses those tools.  Developers time to time uses them, so they know the basic,
let's make the core of the operations. And you as a member of the ops/devops/sre
team should drive your colleagues to extend and use them. It will become the
entrypoint for operations.

Having a single entrypoint is cool because you can take good care of it, you can
make it safe and friendly to use, and you will know how your users interact with
the infrastructure.

I am not an expert about terraform, so I will leave it out from this blog post
and I will keep doing my speculation with Kubernetes.

I am gonna assume every that developer, if you use Kubernetes, has the `kubectl`
installed. And if this is true it means that they have a user because you need
one to interact with Kube API. If this users comes from a global identity
provider that has API you have a bridge from the developer user used to
authenticate to Kubernetes to other services it can connect to, AWS or any cloud
provider sounds like a useful one.

What I am picturing is a kubectl plugin that is an entrypoint for the
infrastructure or for operations controlled by the ops team:

```
kubectl infra status
```

They can be an easy wrapper around Kubernetes resources, or you can
programmatically provide shortcuts or common actions useful for developers.

When you write a kubectl plugin you can use the k8s clients to authenticate
against the current context used by the developer or by yourself. You can get
the node or you can agree on a label like: `owned-by: <team-name>` and return
the list of resources owned by the team where the user works with:

```
kubectl property
```

How do you know in witch team the person is?! As I wrote every request to
kubernetes is authenticated with RBAC you can have groups and from there you can
build a convention that will take you right to the team it is in. Or you can go
to your Identity provider and via API ask for information about the who is using
the command.

You can also develop plugins that reaches multiple providers: Kubernetes, with
AWS, with Mailchimp, with your monitoring platform or even cooler with a
service running inside the infrastucture because you can reach Kubernetes, so
you can get the service and the endpoint of everything running inside the
environment selected by the context in use.

I think this was a nice trip! But it is time to get to an end!

## Conclusion

I recently wrote an
[article](https://gianarb.it/blog/kubectl-flags-in-your-plugin) if you use Go
about how to re-use `kubectl` like flags such as `-L` for labels, or
`--context`, `-n` and `--namespace`.

Who owns the infrastructure needs to understand that there is not value about
keeping ops for itself. It is not the right way to look smarter anymore.

Just to be clear, sysadmin or ops people with a deep knowledge about a specific
domain, I am thinking about DBA are crucial if your business depends on a
database. But I see a lot of companies that will benefit from a proper UX when
doing ops tasks. Ops won't be replaced, but I think they can do a lot more
sharing the boring part of their work with who will benefit of doing it by
itself.
