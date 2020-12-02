---
layout: post
title:  "Lessons learned working as site reliability engineer"
date:   2020-11-26 10:08:27
categories: [post]
tags: [golang]
summary: "I want to share a few lessons I learned working three years as site
reliability engineer. I kept the focus on the one I think are reusable, and that
made me a better developer because reliability just as everything is
an everybody business"
heroimg: /img/amr-taha-coffee.jpg
---

I worked for three years at InfluxData as Site Reliability Engineer. When
onboarding a new role in this jungle called "career" in information technology,
you should be ready to learn what that role means.

Along the way, I developed new skills and mastered a few that I already had, but
they were not widely used elsewhere.

My title is not Site Reliability Engineer anymore, but the IT career is like a
roulette, and everything you learn will get back at some point. So I want to
share skills I think are essential when working as a Site Reliability Engineer
and that I feel they are useful to keep in your toolchain even when your title
changes.

### Ability to develop a friendly environment for yourself

One of my goals as a Site Reliability Engineer is to quickly support developers
having trouble with their code at scale.

Another one is to figure out criticalities when it comes to on-call.

I am far from my local environment in both cases, usually interacting with an
environment much more complicated. Having something you can call familiar help.
It can be whatever:

-   A few bash scripts that wrap other commands with a UX hard to remember
-   A CLI tool you wrote with your team
-   A directory where you can quickly go and write notes about what is going on
    for future use
-   A set of bullet points or a runbook you know is rock solid and can drive you
    where you want to go.

Those are just a few tricks I use. If you have yours and you want to share them
as a comment, please do it!

This is an essential skill that everyone has to master, but as an SRE, when you
have to act quickly, I really learned it, and now I do my best to develop my
workflows and a working environment I like. Those days I am giving a try with
Nix and, in particular, nix-shell because it helps me customize my environment
without the overhead of a Docker container.

This may sound time-consuming. Many projects have a README describing tools and
requirements to contribute or build the project. Why I need my way? Well, I am
not saying you should start from zero, but when I glue it with the flavor I
like, I code better and am happier. So for me, it is a big YES!

### Troubleshoot like a ninja

Starting from the same purpose as before, a Site Reliability Engineer looks at
the code when it runs in production, and production is a scary and dangerous
place. As a developer, if you are lucky and smart, you try to focus on one
application at a time, yes it probably has many dependencies, but still, code
moves one line at a time.

In production with concurrency and thousands of requests happening almost
simultaneously, things get pretty messy. Having the ability and the right tools
to slice and dice from different points of view and prospectives, from an entire
region to a specific application requires operational experience and training.

A good exercise is to have the desire to troubleshoot everything. Does a
teammate have a question about a system in production? Go and help him. Visit
logs, traces, and dashboards even when everything looks quiet if you are so
lucky to have a definition of it.

Another thing I do, but more in general, is to follow the best. There is a lot
about the topic in the forms of books, talks, and similar. Read them but even
more importantly, follow who master these topics every day:
[@rakyll](https://twitter.com/rakyll),
[@brendangregg](https://twitter.com/brendangregg),
[@relix42](https://twitter.com/relix42),
[@lizthegrey](https://twitter.com/lizthegrey),
[@lauralifts](https://twitter.com/lauralifts). Please do not follow them on
Twitter only, but look at their GitHub as well; sometimes, a small project that
works reliably and well for us is gold.

### Think about code debuggability in production

Like everything you read so far, it is always essential because, as I said, Site
Reliability Engineer is just a role that has a subset of responsibilities and
objectives, but we are not silos; everything matters a feature has to be usable;
good looking, functional. Code review for me become a lot more about: "is this
code understandable in production?", "what do I want it to tell me when running
at scale?", "how this trace looks like?" "is this log useful, and how does it
impact the overall context?".

Those questions strongly show my mind when working as an SRE, but they made me a
better developer. I still try to answer them when coding or when doing code
reviews.

### Conclusion

Titles are just titles; reading them will help to know a set of skills you
leveraged most, but that's it, and it is not always true. You are not married to
your title, and if you are curious about various aspects of our work, you will
change many of them. The right balance of all those skills will make you unique.
