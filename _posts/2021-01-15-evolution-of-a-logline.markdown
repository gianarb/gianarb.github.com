---
layout: post
title:  "Evolution of a logline"
date:   2021-01-15 10:08:27
heroimg: /img/hero/dna.jpg
categories: [post]
tags: [observability, monitoring]
summary: "This is a story that represent the evolution I had in thinking and
writing logs for an application. It highlights why they are important as a
communication mechanism from your application and the outside. Explaining what I
think are the responsibility we have as a developer when writing logs."
---

This story represents the evolution of a logline for myself when I write and
interpret it.

Late in 2012, I worked as a developer for a software agency in Turin,
specialized in software for tour operators. It was my second "real" job and the
first one not as a solo developer. Exciting time!  An AJAX application with PHP
and MySQL backend running as a service developed mainly by a single person, the
lead developer.  I interpreted the log back then as the equivalent of a save
point in a game. The tail was the primary tool to figure out what was going on;
a logline was helpful to figure out that a lot of customers were reaching a
particular line of code. The interpretation of the situation was up to humans.

Every developer involved in the project participates in adding the logline in
the codebase directly, developing the code, or indirectly when chatting about a
particular feature over lunch or doing a code review.  Building a context from
an unknown log line was an unrequired exercise because the lead was always there
to help you figure out what that logline was supposed to tell.

Even the stream's speed was a crucial metric to figure out the sanity of the
application. Where the logs too fast? The application was under heavy load, and
probably it was slow, not fast enough, or not smooth as usual, well something
was going on, and it was not good!

You can judge this story as unpractical but not as unusual. This approach does
not scale; it has an unmeasurable risk of "bus factor" but, if you don't have a
panel in your Grafana dashboard representing the distribution of loglines, you
should look at it. Just for fun.


# Bus Factor

Bus factor represents the risk of knowledge and responsibility centralization in
a single location. If the lead in my story resigns or gets hit by a bus, nobody
will build context from "not that descriptive" log lines quickly like him. And
the "speed of tail" requires to be very familiar with the stream. Sharing
knowledge and responsibility across the company, writing documentation, and
doing staff rotations are standard techniques that mitigate such risk.

## Automation

When your application state's interpretation requires a human, it is tough to
build automation for it. Standardization in the way your application
communicates to the outside is another way to spread the knowledge in a team,
allowing you to write automation for it.

The format of a logline is the protocol to develop.

The format has to be parsable and useable by automation. You have to see logs as
a point in time, as time series more than as something that I should carefully
watch and try to interpret by myself.

A logline that looks like this:

```
1610107485 New user inserted in the DB with ID=1234
```

Will become:

```
time=1610107485 service="db" action="insert" id=1234 resource="user"
```

You can add a message that can be used to communicate with a person: msg="new
user registered." but not sure if it is mandatory, you can combine it later.

We do this exercise with ElasticSearch, applying full-text algorithms, and
tokenizing on the message. It is expensive, and it hides the developer's
responsibility when it comes to consciously describe the current state of the
system with a logline. No, they are not random printf anymore.  You can even see
it as JSON if you prefer.

Or, more in general, as a description of a particular point in time via
key-value pairs that you can aggregate, visualize, and use to drive powerful
automation, I work a lot in the cloud field. For me, reconciliation or a
system's ability to repair itself is often based on those pieces of information.
If you want to go deeper on this topic, structured logging is what you should
look for.

## Flexibility

Having a logging library that allows you to do structured logging is a
must-have, and there is no answer about the number of key-value pairs you need.
The overall goal is to derive issues to learn the behavior of an application
from those points. It is not something that you use when having problems. The
application exposes the tool we have to figure out what a piece of code we
didn't write is doing in production.  In a highly distributed system, a logline
with the right fields such as hostname, PID, region, Git SHA, or version can
distinguish where the application having problems is running without looking
across many dashboards, Kubernetes UIs, and CLI.

Parsing and manipulating a structured log is more convenient than a random text
that has to be parsed and tokenized, but everything has a limit, so you have to
find the right balance based on experience. It is another never-ending iterative
process that we can call the evolution logline!

![This picture represents the time it takes to learn something new. It is a
picture of an open book with an old clock.](/img/watch-4638673_1280.jpg){:class="img-fluid"}

## Conclusion

* Logline is the way your tech your application on how to communicate with the
  outside.
* Communication is useful in many fields. It is an opportunity to learn
  something new or a way to communicate that we are in trouble. Same as logs,
  use them as an opportunity to learn how a system works overall.
* As a developer, do not see a logline as a random printf. The way it is
  structured and articulated improves the communication quality between your
  application the outside world.
* A logline is not a fire and forget but an entity that evolves in time.
* Logs represent the internal state of your application at some point in time
  and somewhere in your codebase.

Recently I spoke with [Liz](https://twitter.com/lizthegrey) and
[Shelby](https://twitter.com/shelbyspees) from HoneyComb about observability and
monitoring during
[o11ycast](https://www.heavybit.com/library/podcasts/o11ycast/ep-32-managing-hardware-with-gianluca-arbezzano-of-equinix-metal/?utm_campaign=coschedule&utm_source=twitter&utm_medium=heavybit&utm_content=Ep.%20%2332,%20Managing%20Hardware%20with%20Gianluca%20Arbezzano%20of%20Equinix%20Metal)
a podcast about observability if you want to know more about this topic.

{:.small}
Hero image via [Pixabay](https://pixabay.com/illustrations/dna-string-biology-3d-1811955/)
