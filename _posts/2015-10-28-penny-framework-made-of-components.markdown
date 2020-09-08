---
layout: post
title:  "Penny PHP framework made of components"
date:   2015-10-27 23:08:27
categories: [post]
img: /img/penny.jpg
tags: [php]
summary: "Penny a PHP framework made of components, write your microframework
made of symfony, zend framework and other components."
priority: 0.6
---
<blockquote class="twitter-tweet tw-align-center" lang="en"><p lang="en" dir="ltr"><a href="https://twitter.com/hashtag/pennyphp?src=hash">#pennyphp</a> <a href="https://t.co/tsA2nE09GM">https://t.co/tsA2nE09GM</a> Why and what?! o.O <a href="https://twitter.com/hashtag/php?src=hash">#php</a> <a href="https://twitter.com/hashtag/framework?src=hash">#framework</a> to build <a href="https://twitter.com/hashtag/microservices?src=hash">#microservices</a> and application &quot;consciously&quot;</p>&mdash; Gianluca Arbezzano (@GianArb) <a href="https://twitter.com/GianArb/status/659762064446083073">October 29, 2015</a></blockquote>
<script async src="//platform.twitter.com/widgets.js" charset="utf-8"></script>


<p class="text-center">
<iframe src="https://ghbtns.com/github-btn.html?user=pennyphp&repo=penny&type=star&count=true&size=large" frameborder="0" scrolling="0" width="160px" height="30px"></iframe>
</p>

The PHP ecosystem is mature, there are a lot of libraries that help you to write
good and custom applications. Too much libraries require a strong knowledge to
avoid the problem of maintainability and they also open a world made on specific
implementations for specific use cases.

A big framework adds a big overhead under your business logic sometimes, and some
of those unused features could cause maintainability problems and chaos.

Spending too much time reading the docs could be a problem, do you think you are
a system integrator and not a developer?! These are different works!

We are writing [penny](https://github.com/pennyphp/penny) to share this idea.
This is a middleware, event driven framework to build the perfect
implementation for your specific project. The starting point we chose is made of:

* [Zend\Diactoros](https://github.com/zendframework/zend-diactoros) PSR-7 HTTP
library
* [Zend\EventManager](https://github.com/zendframework/zend-eventmanager) to
design the application flow
* [PHP-DI](https://php-di) DiC library
* [FastRouter](https://github.com/nikic/FastRoute) because it is fast and easy to
use

but we are working to replace every part of penny with the libraries perfect
for your use case.

Are you curious to try this idea? We are writing a big documentation around penny.
[docs.pennyphp.org/en/latest](https://docs.pennyphp.org/en/latest/)

And we have a set of use cases:

* [pennyphp/penny-classic-app](https://github.com/pennyphp/penny-classic-app)
builds with plates
* [pennyphp/bookshelf](https://github.com/pennyphp/bookshelf) builds with
doctrine, twig
* [gianarb/twitter-uservice](https://github.com/gianarb/twitter-uservice) gets
the last tweet from `#AngularConf15` hashtag

[Share your experience!](https://github.com/pennyphp/penny/issues?utf8=%E2%9C%93&q=is%3Aissue)
