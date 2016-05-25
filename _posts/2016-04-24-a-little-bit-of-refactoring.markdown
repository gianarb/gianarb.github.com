---
layout: post
title:  "A little bit of refactoring"
date:   2016-04-24 10:08:27
categories: [post]
img: /img/refactoring.jpg
tags: [php]
summary: "Strategy to refactoring your code. Tricks about performance from
PHPKonf Istanbul Conference."
priority: 0.6
changefreq: yearly
---
I wrote this note during the [PHPKonf](http://phpkonf.org/), I spoke about
[Jenkins and Continuous Delivery](http://gianarb.it/jenkins-real-world/#/) but
during the days I followed few interested talks and this is my list of notes.

Is the code reusable? For a few people the response is yes, for other ones is
no. I agree with [Ocramius](https://twitter.com/ocramius) that the response is
no. An abstraction is reusable, an interface is reusable, but it’s very hard to
reuse a final implementation.  First of all because when you finish to write it
your code is already old, your function is already legacy and you start to
maintain it, you search bugs and edge cases.

One of the way to reduce the time dedicated to do refactroring is prevent and
defende your code from bad integration, in OOP usually you have a sort of
visibility (private, public, protected) and other different way to defend your
code, opinable but the ocramius's talk about [Extremly defensive
php](https://ocramius.github.io/extremely-defensive-php/#/) is good to see.

Refactoring is a methodology to make better your code. There are different
improvement topics like readable, performance, solidity.

- make your code readable for the new generation is one of the best stuff that
  you can do to show your love for your team and your company.
- If your site require more time to be loaded usually you lost your client.
  “less performance is a bug” cit. [Fabien Potencier](https://twitter.com/fabpot)
- When your run your code all it’s fine, you are a good developer and your
  feature works. After the deploy in production your code is the same but
  usually there is a bd category of people, your client, that will use it in a
  very strange way, usually it’s synonym of bug or edge case. Each bug fix make
  your code more solid.

Test your code before start to change it, you know automation is good but if
you love seems a machine to it manually.  Setup a continuous integration
system, it can be do just one step like run tests but remember to increase that
with all steps that you usually do to test the compliance of your code like
style, standard, static analysis just to enforce that you are not a machine.
Create a good environment and an automatic lifecycle for your application allow
you to stay focused on the code and not lost your time around stupid task,
remember that when a routine is good the machine fail less respect a human,
usually.

Refactoring is one of the best stuff that you can do for other people and to
make your feature ready for the real world, usually it’s hard for no-tech
company understands it because few times they don’t see any kind of change
create a good environment to save time and use it to do refactoring is a godo
strategy.  Automation is the unique method that I know to do that. There are
different layer of automation just to start my 2coins is just put a make file
on your codebase and when you do something for the second time stop to write it
on your console and write a new make task to share with you team.  After that
install Jenkins and allow it do run this task for you before put on your on the
master branch (for git users, trunk for svn users).

Make your development environment comfortable and increase the conformability's
perception about the lifecycle it’s the best way to do refactoring without the
fear to die.  If you are fear to die usually you don’t do nothing.

<blockquote class="twitter-tweet tw-align-center" data-lang="en"><p lang="en" dir="ltr">Great
talk as always Gianluca :) <a href="https://twitter.com/GianArb">@GianArb</a>
<a href="https://twitter.com/hashtag/phpkonf?src=hash">#phpkonf</a> <a
href="https://t.co/ZW2G1UsXm7">pic.twitter.com/ZW2G1UsXm7</a></p>&mdash;
Fontana Lorenzo (@fntlnz) <a
href="https://twitter.com/fntlnz/status/733986655334486016">May 21,
2016</a></blockquote> <script async src="//platform.twitter.com/widgets.js"
charset="utf-8"></script>

Add the PHPKonf in your list! See you next year!
