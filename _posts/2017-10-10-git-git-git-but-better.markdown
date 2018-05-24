---
layout: post
title:  "Git git git, but better"
date:   2017-10-10 10:08:27
categories: [post]
tags: [docker, docker captain, git, coreos, automation, infra as code,
infrastructure]
img: /img/git.png
summary: "Doesn't matter for how much time you are using git or any version
control system, you always have something to learn abut them. Not about the
actual interface but about the right mindset."
changefreq: yearly
---
I can't say that Git is a new topic. Find somebody unable to explain how a
version control system was working was very hard. Now it's almost impossible.

I used SVN and Git for many years and I also put together some unusual use case
for example: ["Splitting Zend Framework Using the Cloud"](https://devzone.zend.com/6134/splitting-zend-framework-using-the-cloud/)
is a project that I made with Corley SRL my previous company and the Zend
Framework team.

It helped me to put my hands down on the Git file system and I discovered a lot
of features and capabilities that are not the usual: commit, checkout, reset,
branch, cherry-pick, rebase and so on.

But during my experience building cloud at InfluxData I need to say that I can
see a change of my mindset, I am sharing this because I am kind of proud of
this. It's probably not super good looking at the time required to achieve this
goals but how cares!

> Sometimes it’s the journey that teaches you a lot about your destination.
> (Drake)

I don't know this Drake, I am not even sure if it's the right author of the
quote but that's not the point.

At InfluxData, just to give you more context, I am working on a sort of
scheduler that provisions and orchestrate servers and containers on our [cloud
product](https://cloud.influxdata.com/). A lot of CoreOS instances, go, Docker and
AWS api calls.

It's a modest codebase in terms of size but it is keeping up a huge amount of
servers, I am actively working on the code base almost by myself and I am kind
of enjoying this. Nate, Goller and all the teams are supporting my approach and
are using it but I am not using Git because hundreds of developers need to
collaborate on the same line of code. I had some experience in that environment
working as contributor in many open source project. But this time is different.

I am mainly alone on a codebase that I didn't start and I don't know very well,
this project is running in production on a good amount of EC2.

I really love the idea of having a clean and readable Git history. I am not
saying that because it's cool. I am saying that because every time I commit my
code I am thinking about which file to add/commit `-a` is not really an option
that I use that much anymore. I think about the title and the message.

I try to avoid the `WIP` message and I use it only if I am sure about a future
squash, rebase and if I need to push my code to ask for ideas and options (as I
said I am writing code almost alone, but I am always looking for support from my
great co-workers).

This has a very big value I think also as remote worker. This is my first
experience in this environment and for a no-native English a good and
self explanatory title can be the hardest part of the work but it will help
other people to understand what I am doing.

When you are working on a new codebase and you have tasks that require
refactoring to be achieved in a fancy and professional way you will find
yourself moving code around without really be able to figure out when and how it
will become useful to close your task and open the PR that your team lead is
waiting for. At the end if you start to write code and you commit your changes
at the end of the day as I was doing at the beginning after a couple of days you
will figure out that your PR is too big and you are scared to merge them.
And probably it's just the PR that is preparing the codebase to get the initial
requests. I hated the situation but if you think about what I wrote you will
find that it's totally wrong.

VCS is not there as saving point, you are not plaining Crash Bandicoot anymore,
you don't need to use Git as your personal "ooga booga". The right commit
contains an atomic information about a feature, bug fix or whatever.

![](/img/crash_bandcioot.jpg)

These are the questions that I am asking myself now before to make a commit:

* am I confident cherry-picking this commit to `master`? This is a good way to
  make your commit small and easy to merge. If one of your PR is becoming too
  big and you have "cherry-picked" commits you can select some of them merge
  them as single PR.
* are deploy and rollback easy actions? This is similar to previous one but I am
  the one that deploy and monitor the service in production. I need to ask this
  question to myself before every merged PR.
* Looking at the name of the branch that in my case the task in my
  viewfinder the commit that I am creating is about it or can I create a new PR
  just for this piece of code? This helps me a log to split my PR and to them
  small. A small PR is easier to review, it has a better scope and it makes me
  less scared to deploy it.

Git is more than a couple of commands that you can execute. You need to
be in the right mindset to enjoy all the power.
