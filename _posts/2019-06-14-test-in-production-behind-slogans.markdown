---
img: /img/gianarb.png
layout: post
title:  "Test in production behind slogans"
date:   2019-05-27 08:08:27
categories: [post]
tags: [code instrumentation, library, opentracing, opencensus, opentelemetry,
influxdb, prometheus, opensource, honeycomb]
summary: "How fast we are capable of instrumenting an application decrease the
out of time requires to understand and fix a bug."
changefreq: daily
---
<blockquote class="tw-align-center twitter-tweet"><p lang="en" dir="ltr">What do we test before
prod? We do our known unknowns -- does it work? (unit tests). does it fail in
ways I can predict?<br><br>We need to test our unknown unknowns in production
with ✨observability✨. and experiment upon them with chaos engineering! <a
href="https://twitter.com/hashtag/VelocityConf?src=hash&amp;ref_src=twsrc%5Etfw">#VelocityConf</a></p>&mdash;
Liz Fong-Jones (方禮真) (@lizthegrey) <a
href="https://twitter.com/lizthegrey/status/1139273082412027904?ref_src=twsrc%5Etfw">June
13, 2019</a></blockquote> <script async
src="https://platform.twitter.com/widgets.js" charset="utf-8"></script>

I got inspired by Liz's tweet recently, and I am writing this post as a reminder
for everybody. "Test in prod" is a slogan, a trademark. It doesn't explain all
the concepts behind a sentence as "things bo getter with Coke" hides why or
how. Slogans are great as a quick reminder for more articulated ideas. They
are useful because in one sentence you can recall to more profound contents
inside your brain.

<img class="img-fluid" src="/img/coke-slogan.jpg">

"You" do not test unknown unknowns in production, mainly because you do not know
your unknowns. In production, you as a developer **validate** three kinds of
things:

* Complicated parts of your system that are not well covered by tests are
  working.
* Something you are working on and you would like to be sure that it is working
  fine, even if it has a unit test, integration tests and so on.
* Crucial part of the system that needs to work or your boss will kick your ass,
  and you are afraid that you test them even if you just changed a line of CSS.

What "test in prod" means is the fact that somebody, a random customer human or
not randomly will trigger an unknown action that will cause an issue. It doesn't
even be to be triggered, it can be an environmental issue. For example, what
Twitch call ["the refresh
storm"](https://blog.twitch.tv/go-memory-ballast-how-i-learnt-to-stop-worrying-and-love-the-heap-26c2462549a2)
is an excellent example of an environmental issue. When a broadcaster has a
connectivity issue, all the watcher starts to refresh the page multiple times
thinking to solve the problem. As a side effect, the Twitch infrastructure can
suffer about a high number of requests. This is a no-Twitch problem that becomes
a Twitch problem.

We need to learn and onboard tools and mindset that will help us to improve how
fast we can track, record, fix, and learn from an issue. All the question that
matters happens in production, and by consequence, we need to stay focused on
it.  I think a lot of people test in prod in some way.

When your laptop starts, but it restarts by itself after some point you have a
problem. You look around, and you notice that your fan doesn't run anymore. It
is a pretty simple issue to solve and detect. You hear that the fan doesn't make
any noise, so you replace it.

I am sorry! Everybody got distracted by distributed system, containers, cloud.
90% of our failures if you know how to design a fault tolerance application are
a partial failure! They are a disaster to figure out, understand, and fix! Only
a subset of our system may break, for a subset of customers, but the same part
works correctly for another subgroup, and you need to figure out why! You should
also be able to message that subgroup of customers to say "I am sorry! Shit
happens, we are working on it", proactively!

## Conclusion

"test in prod" means all the things I wrote and probably way more! It is
reasonable to say that nobody can do anything to avoid "test in prod" to happen,
so have fun!

