---
heroimg: /img/hero/monkey.jpg
layout: post
title:  "Chaos Engineering"
date:   2018-08-23 08:08:27
categories: [post]
tags: [chaos, netflix, jazoon, engineer, resiliency]
summary: "I took part at a panel at the Jazoon conference in Switzerland called Chaos
Engineering. Where I had a change to learn about techniques and practices around
this topic that even if I new about it I never had the chance to put my head on
it. In this article I am summarizing my ideas and what I get mainly around the
definition of Chaos Enginnering."
changefreq: daily
---
At the [Jazoon](https://jazoon.com/) conference in Switzerland, I had the chance
to speak at the Chaos Engineering panel with [Russ
Miles](https://twitter.com/russmiles) from [ChaosIQ](https://chaosiq.io) and
[Aaron P Blohowiak](https://twitter.com/aaronblohowiak) from Netflix.

The organizers put me in the panel probably because "chaos" was part of the title
for the talk I just gave in the morning. I was too curious to mention that I
never did it before, at least on purpose!

So I was really out of my comfort zone dealing with these two folks that know
their shit so well!

I am sure that as Engineers we are part of the Chaos, we create entropy inside the
system during every deploy and even if we have all the tests in the world the
first time it is tough to make it work. But I indeed never associated the
word engineering to chaos. And that's the real challenge.

So, let's define Chaos and Engineering altogether.

Let's start with `Chaos` because it is the easy one, as I said we as developers create
chaos, distribution creates chaos, and customers create chaos. If somebody tells you
that his production environment is excellent, you should not listen to him,
Production is a nightmare, complicated and painful place. At least if somebody uses it.

And if it is just a bit more complicated than a static site it never works 100%,
the chaos governs it, and that's where the sentence Engineering becomes
essential.

`Engineering` at least for what I can understand means to be driven by data and
not feeling. So associating these two concepts together you have a powerful way
to measure the chaos.

I think you can't avoid chaos, so the best way to handle it is to learn from
what it generates in your system to anticipate unpredictable situations.

As developers, ops or devops we are pessimistic about our system, and we know
that it will fail: servers crashes, CoreOS auto updates itself, third party
services stop to work. Usual the answer is to wait for it to happen usually
Friday night.

Chaos Engineering is an exercise, a practice to leverage "unusual but possible"
situations as teaching vector to our system.

It is another tool to achieve resiliency and to test scalability.

Chaos Engineering doesn't bring down all your production system in an
unrecoverable way. It designs exercises that you and your team will use to
increase your operational experience and confidence.

Observability is a sort of requirement to understand how a chaotic event changes
the "normal" state of your system. But from another point of view a chaotic even
shed some light for a particular part of your system showing up lack of
monitoring and instrumentation.

There open source framework like [Chaos-tookit](https://github.com/chaostoolkit)
or famous tools like [chaos-money](https://github.com/Netflix/chaosmonkey).

I will try to start with some very simple example without writing too much
code. I will get out from my system these metrics:

1. Number of requests (probably from ingress/nginx)
2. The number of requests with status code > 499
3. Http request latency

After that, I will try to simulate an outage removing or scaling down particular
pods (the one that gets all the traffic) and I will look at how the metrics will
change and how long it takes to recover.
