---
layout: post
title:  "FastEventManager, only an event manager"
pate:   2015-11-01 10:08:27
categories: [post]
img: /img/github.png
tags: [php eventmanager, framework, open source, oss]
summary: "FastEventManager is a PHP library designed to be a smart and light
event manager. You can use it in your applications or as a base component for
your framework. It adds capabilities around events as attach and triggering of
events."
priority: 0.6
changefreq: weekly
---

> The Event-Driven Messaging is a design pattern, applied within the
> service-orientation design paradigm in order to enable the service consumers,
> which are interested in events that occur within the periphery of a service
> provider, to get notifications about these events as and when they occur
> without resorting to the traditional inefficientpolling based mechanism.
> by. [wiki](https://en.wikipedia.org/wiki/Event-Driven_Messaging)

In PHP there are different implementation of this pattern, but [I tried to write
my idea](https://github.com/gianarb/fast-event-manager).
An easy to understand and to extends event manager based on regex.

Why? Because it is a good way to match strings, it is flexible and powerful.
As it is smart and little and it can be used as basis for custom implementation.
it resolves a regex and triggers events It supports a priority to order
triggered listeners.

## Install
{% highlight bash %}
composer require gianarb/fast-event-manager
{% endhighlight %}

## Getting Started
{% highlight php %}
<?php
require __DIR__."/vendor/autoload.php";
use GianArb\FastEventManager;
$eventManager = new FastEventManager();
$eventManager->attach("user_saved", function($event) {
});
$user = new Entity\User();
$eventManager->trigger("/user_saved/", $event);
{% endhighlight %}

Each listener has a priority (default = 0), it describe the order of execution

{% highlight php %}
<?php
$eventManager->attach("wellcome", function() {
    echo " dev!";
}, 100);
$eventManager->attach("wellcome", function() {
    echo "Hello";
}, 345);
$eventManager->trigger("/wellcome/");
//output "Hello dev!"
{% endhighlight %}

I wrote this library because there are a lot of solutions that implement this
pattern but they are verbose, this is only an event manager if you search other
features you can extends it or you can use differents implementations.
On top of this library you can write your library to build an event manager ready
to use with your team in your applications.

This is a good solution because it is easy, ~31 line of code to trigger events
without fear to inherit many line of codes and unused features to maintain.
