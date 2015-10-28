---
layout: post
title:  "Penny PHP framework made of components"
date:   2015-10-27 23:08:27
categories: php
tags: framework, php, interoperability 
summary: "Penny a PHP framework made of components, write your microframework
made of symfony, zend framework and other components."
priority: 0.6
changefreq: yearly
---
The PHP ecosystem is mature, there are a lot of libraries that help you to write
good and custom applications.  Too much libraries require a strong knowledge to
avoid the problem of maintanability but they open a world made on specific
implementations for specific use case.

A big framework adds a big overhead under your business logic sometimes did of
unused features cause of maintanability problems and chaos.

Spend too much time reading the docs could be a problem, do you seem a system
integrator and not a developer?! This are different works!

We are writing [penny](https://github.com/pennyphp/penny) to share this idea.
This is a middleware, event driven framework to build the perfect
implementation for your specific project. How start point we choice:

* [Zend\Diactoros](https://github.com/zendframework/zend-diactoros) psr7 http
library
* [Zend\EventManager](https://github.com/zendframework/zend-eventmanager) to
design the application flow
* [PHP-DI](https://php-di) DiC library
* [FastRouter](https://github.com/nikic/FastRoute) because it is fast and easy to
use

but we are working to substitute every part of penny with the libraries perfect
for your use case.

Are you curious to try this idea? We are writing a big documentation around penny.
[docs.pennyphp.org/en/latest](http://docs.pennyphp.org/en/latest/)

And we are a set of use cases:

* [pennyphp/penny-classic-app](https://github.com/pennyphp/penny-classic-app)
builds with plates
* [pennyphp/bookshelf](https://github.com/pennyphp/bookshelf) builds with
doctrine, twig
* [gianarb/twitter-uservice](https://github.com/gianarb/twitter-uservice) gets
the last tweet from `#AngularConf15` hashtag

[Share your experience!](https://github.com/pennyphp/penny/issues?utf8=%E2%9C%93&q=is%3Aissue)
