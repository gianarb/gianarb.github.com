---
layout: post
title:  "The abstract manifesto"
date:   2018-03-17 10:08:27
categories: [post]
changefreq: monthly
tags: [code, development, abstract, manifesto]
summary: "Often looking at the code I spot a lot of places where it looks too
complicated. Disappointment is the feeling that I get reading classes with weird
names or chain of abstractions or interfaces used only one time. Abstraction is
often the reason for all my sadness."
img: /img/shit-pretty.png
---
This is a personal outburst. Stop to abstract by default.

I worked on too many application abstracted by default. And abstracted I mean
complicated.

The abstraction is easy to understand when you need it. If you need to think too
much about why that crappy code should have an interface, or that made should
implement an interface and not an object you are out of track.

Abstraction is not the answer, code architecture is, unit testing helps,
integration tests are the key to the modern microservices environment.

Don't waste your time creating interfaces that nothing will reuse. If you don't
know what to do run.

There are languages and design pattern that probably set your brain to look for
abstraction everywhere. I worked with Java developer that wasn't able to write a
class without an interface, or without its abstract. My question was: "Why are
we doing that?". Compliance.

> Dude, your world is a very boring one, and you are the root cause.

If you are working in a service-oriented environment with services big enough
to be rewritten easily the abstraction is even more useless.

We are developers, we often don't build rockets. That's life, there are a good
amount of companies that make rockets, apply there or you will put your company
in the condition of paying technical debts for you and they will hire smart
contractors to figure out what you did now that you are not working there
because after probably just one year you locked yourself in that boring project
full of complicated concepts.

btw, I don't think that software to control rockets has a lot of abstractions
sorry.

<blockquote class="twitter-tweet tw-align-center" data-lang="en"><p lang="en"
dir="ltr">Y&#39;all are all about passionate programmers, but honestly I&#39;d
rather programmers than care _just enough_. I could do with less pedantic
arguments about code.</p>&mdash; ｡ 𝕷𝖎𝖓𝖉𝖘𝖊𝖞 𝕭𝖎𝖊𝖉𝖆 ｡ (@lindseybieda) <a
href="https://twitter.com/lindseybieda/status/969296749985779712?ref_src=twsrc%5Etfw">March
1, 2018</a></blockquote>
<script async src="https://platform.twitter.com/widgets.js"
charset="utf-8"></script>

I saw this tweet on my timeline yesterday and I think it really describes my
current mood. The code changes over time and I should spend more time making it
flexible enough to support this continues to grow. Abstraction is not the right
way.

So, passionate code engineer always abstract? That's not the giveaway, Java
engineer always abstract? maybe.
