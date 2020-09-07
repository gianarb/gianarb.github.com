---
img: /img/me.jpg
layout: post
title: "Maintainer life, be an enabler"
date:   2020-08-20 09:08:27
categories: [post]
tags: [opensource, oss, maintainer]
summary: "Be an enabler is important in my daily job. It is a skill I learned as
open source maintainer."
changefreq: daily
---

This is not something important only in open source, or as maintainer. But it is
a skill I have personally learned as one.

When I have to build a sustainable open source project, but it applies with
teams as well developing the right code or new feature is often not that useful
in the long term. I think you get quickly to something better when people
collaborate together in an effective way.

The maintainer's role is to enable other people to contribute successfully.

You have to switch from "let's write documentation" to "how do I create a
workflow that enables contributors to write documentation." I started with doc
because I think it is crucial. It is easier to write documentation when you are
writing code or a new feature. And we know a developer prefers to write code; as
a maintainer, you need to create a workflow that allows the contributor to write
documentation quickly when writing code. In practice, mark a PR with a label
`needs-doc` and make it a requirement for the PR to be merged. The
maintainer has to design a rock-solid structure for the documentation. In this
way, the contributor won't spend two days trying to figure out where to add the
documentation for its feature.

You can't ask a contributor to create an entire test suite or to write
documentation if you don't have one. But from a solid foundation is reasonable
to ask a contributor to keep at least the same quality level.

You don't write all the tests; you create and maintain the continuous delivery
pipeline required to help contributors to stay compliant. Is your project suffering
from low test coverage? Do not waste time writing all the tests yourself;
codebase is significant, and pull requests are flowing continuously. You have to
stay focused on developing a system that brings and keeps you where you want:
good coverage in this case.

In practice, you can create another label `needs-tests` to notify the
contributor that its work won't be merged until tests will be added (the plural
is crucial, tests!). You can use something like [codecov](https://codecov.io/)
in your CI to evaluate with numbers the situation. Invest time being sure that
tests are easy to write; write a doc in the contributor file highlighting how to
write a good test. If a package is too hard to test, you can write a few tests
that other people can use as a starting point or a reusable set of utility
functions.

Being a facilitator or an enabler is a lot of work. If you feel less
effective because you wrote 90% of the codebase at this point and you can write
documentation and tests by yourself in a couple of days you are wrong. Or at
least from my experience you can do it but the outcome will be worst in quality
compared with the one you can build from a solid foundation in a collaborative
environment.
