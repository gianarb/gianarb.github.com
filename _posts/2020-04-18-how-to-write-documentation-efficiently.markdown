---
img: /img/me.jpg
layout: post
title: "How to write documentation efficiently"
date: 2020-04-18 09:00:27
categories: [post]
tags: ["documentation", "development"]
summary: "A bunch of experiences and considerations about how to write
documentation efficiently. Without wasting too much time or even more important,
without getting too bored or stressed out."
changefreq: daily
---

You have to remember two things to effectively read this article:

1. I have a blog post where I create content with a good frequency, and I do it
   for fun, so I like to write.
2. I work in remote, the HQ for InfluxData is in San Francisco, it means I am +9
   from a lot of my colleagues. Writing is a solid communication channel I use
   every day at work because I think it is great, and because I do not have
   other alternatives.

{:refdef:.text-center}
![Child writing on paper](/img/child-writes.jpg){:.img-fluid}
{:refdef}

## Develop a workflow

If you do not like to clean your apartment a strategy you have is to try to keep
it as clean as possible, and in order day by day, in this way you won't have to
spend a full weekend cleaning every corner of it. Spread a boring task in
a way that won't make you too tired.
An effective way is to write along the way, side by side with the code you are
developing.

I can highlight a few steps in the process of writing code: analysis, design,
validation, PoC, rollout. Those phases are not unique, they go continuously over
many iteration. I write during all of those steps, many times. Iterations do not
help only your code, they make documentation solid, you can check for typos an
so on.

If you make writing an ongoing process you will find yourself at the end where
the only thing left is to organize and move what you wrote a way that readers
will find familiar.

## Find the right place

There are many time of documentation, because there are a lot of stakeholders
and many phases to document (some of them where listed previously).

If I have to think about my stakeholders they are:

1. project managers
2. documentation team if you are lucky, otherwise let's say customers or end users.
3. VP or tech leads.
4. your teammate or reviewers

All those people will enjoy reading a specific point of view, or phase of work.

I think teammate or reviewers are kind of happy to read the process you followed
to design and implement what you wrote, and they will really appreciate to read
inline documentation for your code, doc blocks and so on.

Project Managers will enjoy reading considerations on issues and things like
that, they are super valuable and I end up copy pasting a lot from those
discussions.

End user obviously need a function documentation they they can follow and also a
bit about internal design, monitoring mainly to get them onboard with the work
you developed. It really depends on your audience. We are lucky and we have a
team that is capable of reading code and figure out what we did, but it is a
nice exercise to help them explaining in a good way your work.

VP and tech leaders are usually focused on the design, why you did something in
a way other than another, the trade off you accepted, the one you avoided, why
and how. I like the idea to write this kind of documentation in the code itself.

I am fascinated when I open C codebases where the first thousen lines of code
are documentation. In Go packages can have a file called `./doc.go` that `godoc`
will render as a package introduction. If you work with the kind of tech lead or
VP that are not used to read code anymore, you can always copy paste it to
google doc.

## Write a lot

This point self explains itself. More you write during all the phases if your
work, less you will have to do all together at the end of the code iteration.

Where I usually end up tired about the code I wrote, even more when it takes
weeks, and it is not easy to work on.

## Pair on documentation

I am not a fan of pair programming but recently I changed my mind a little bit,
probably caused by all this social isolation. Before jumping straight on writing
code with my teammate two solid hours over two iterations writing the `./doc.go`
file together. The outcome made me happy, I hope it will work the same for you.

{:refdef:.text-center}
![Child writing on paper](/img/toomany-files.jpg){:.img-fluid}
{:refdef}

## Conclusion

This is my experience when writing documentation, but as I said, I love to do
it! Do you have anything to share about it? I am particularly curious about how
and if you READ somebody else documentation, internally written by your teammate
how do you evaluate it and if you have any suggestion to make it more friendly.
Because it is good to write but people has to be able to read it and get what
they need out of it without wasting too much time.
