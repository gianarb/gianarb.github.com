---
layout: post
title:  "Your release workflow is code, it is just about time"
date:   2020-10-21 10:08:27
categories: [post]
tags: [kubernetes, release, code]
summary: "I think code is some way has to win against specification languages or
DSL or even from languages that are not easy to move around like bash. The
Kubernetes sig-release is migrating a bunch of scripts from bash to Go and I
think it is the right way to go. You will do the same, it is just about time."
heroimg: /img/elephant-5647723_1920.jpg
---

20th October 2020 is the day I released Kubernetes for the first time. To be
precise, I piloted with the help of sig-release Kubernetes `v1.20.0alpha3`. One of
the reasons I am happy to work with the sig-release is to learn and see how such
a significant process gets released reliably and continuously by a group of
people coming from different backgrounds, jobs, and locations.

The first lesson learned you would notice as soon as you join the SIG meeting
more frequently or as soon as you start contributing is the general effort in
converting what used to be bash scripts to Go.

Now, don't fight against the languages by themself, but I think the story is
reasonable. You start small, and when it comes to releasing code, a lot happens
in somebody's terminal. That's why many of the release workflows I saw in my
life are a mix of Makefile and bash script.

I don't think it scales because it is hard to get error handling, retry logic,
and testing made right in bash. Maybe I am just not good enough with BASH, and I
know there are testing libraries for it like
[bats](https://github.com/sstephenson/bats), for example.

Anyway, I have to admit, I feel good enough with BASH, but I code way better in
Go, PHP, and probably even JavaScript. Also, I am sure this is a feeling a share
with many people, and more in general, the Kubernetes development community is
very fluent with Golang.

Anyway, let's treat the code that empowers the release lifecycle as application
code, just as the Sig Release is doing with Kubernetes. Documentation, testing,
user experience, and so on. Develop useful libraries that can be encapsulated in
command-line tools, or API, or bots.

There is a BASH script that takes snapshots from
[Testgrid](http://testgrid.k8s.io/) called
[testgridshot](https://github.com/kubernetes/release/blob/master/testgridshot),
uploads them to Google Cloud, and outputs a markdown that can be copy-pasted as
a comment in [the issue we use to track every
release](https://github.com/kubernetes/sig-release/issues/1296). We run it to
take a snapshot of the various testing pipeline status at the time of a release.

testgridshot is the unique one in BASH I had to interact with for now, and it
didn't work because of some environmental issues with my laptop. Coincidence? It
can be solved by running it as a container and having a binary with statically
compiled with all the needed dependencies.

[Carlos](https://twitter.com/comedordexis) is currently working on rewriting
testgridshot in Golang; it will use as a command-line interface, and I think it
will be even better to encapsulate it as prow capability.

[Prow](https://github.com/kubernetes/test-infra/tree/master/prow) is the
Kubernetes CI/CD system. It can trigger jobs for particular actions, and almost
everything you see happening in GitHub when using `/` commands like `/open` `/assign`
and so on is a Prow responsibility.

Testgridshot is useful during a release cut. The cut starts from a GitHub issue;
as we saw, it sounds very comfortable to have a command available like
/testgridshot and leaving Prow the responsibility to comment.

Now the takeaway hidden by the word **encapsulates**, it is great to have both a
CLI and a Prew command. Go becomes your baseline where the operational
experience live. All the rest is a UX, and you can have as many of them.

I am not writing this because you should stop and move all your BASH to
something else, but I experienced by myself. I see this little story with the
Kubernetes SIG release to confirm that it's easy to block ourselves as release
engineering because there is a BASH script that we don't want to rewrite. After
all, it is like that since forever. The project is not the same since day one,
the team grew or changed, and it is reasonable for a workflow to follow this
evolution.
