---
layout: post
title:  "Continuous Integration and silent check. You are looking in the wrong place"
date:   2016-11-18 10:08:27
categories: [post]
img: /img/jenkins.png
tags: [ci, devops]
summary: ""
changefreq: yearly
---
Continuous Integration is a process of merging all developer working copies to
shared mainline several times a day. In practice is when you have in place a
system that allow you to trust all changes that all developers are doing in a
short period of time in order to have that code complaint and ready to be
pushed in production.

There are a lot of different way to do CI but I will stay focused on a very
important expect, you need a policy that contains a series of checks that you
can easy automate. All this steps persisted on every change allow you to mark
that new code as `ready`.

Automation is a key part to keep your integration continuous, usually what
people do is a human review of the code, if one or more people mark your code
as complaint and the continuous integration system is agree with them your code
can be merged. This is the unique manual step.

But let's talk about what I call "Silent Checks" they are really one of the
best invention that I never saw in my like, they are like laws in Italy, we
have a lot of them but no body cares.

Usually your CI system use exit code to understand if a check is good or bad,
your command come back with `0` in case of success or with another number if
something fails. Sometime you can find in your continuous integration checks
that put the status code in a silent mode. The check fails but it's not important enough.
It's like a party!

<img class="img-responsive" src="/img/the-wolf-ci.jpeg" alt="continuous integration party">

You have a check that runs but you are not asking people to care about the
result. There are few disadvantages about this approach:

* That check is making your job slow.
* If the job doesn't fail no one care about that optional check.
* When a job fails you just need to scroll and jump over all the logs generated
  by the optional check because usually they fails, because nobody cares. There
  is more, usually your coworkers forget about this check and they ping you
  about that errors.

Analyse your code is very important but there is another simple strategy that
you can use to avoid this inconvenient. Usually the silent checks are in place
in a period of migration, maybe they are important. They are just in the bad
position.
You can move them in a separated job, collected them and analyse what you need
to analyse and monitoring trends about how your team works.

I saw a TEDx Talk by Adam Tornhill. He talked about Analyzing Software with
forensic psychology. This topic is great! You can gram a lot of information
about your application from who is writing that code.

<div style="text-align:center">
<iframe width="640" height="360"
src="https://www.youtube.com/embed/qJ_hplxTYJw" frameborder="0"
allowfullscreen></iframe>
</div>

Trends and monitoring not just to understand how your application works but
they are fundamentals to understand how your team is working, how they
feel and also to catch how your codebase is moving. They are really important
and if you are strong enough to have a good monitoring system for that metric
you are really in a good position!  You just need to understand that insert
them into the continuous integration flow is not a good idea.
