---
layout: post
title:  "Build your Zend Framework Console Application"
date:   2015-05-21 23:08:27
img: /img/zf.jpg
categories: [post]
tags: php, console, automation, zend framework
summary: ZF\Console is a component written by zf-campus and Apigility organization that help you to build console application using different Zend Framework components
priority: 0.6
changefreq: yearly
---
<blockquote class="twitter-tweet tw-align-center" lang="en"><p lang="en" dir="ltr">Blogpost about console-skeleton-app for your console application <a href="http://t.co/WuVq0GZlxE">http://t.co/WuVq0GZlxE</a> <a href="https://twitter.com/hashtag/PHP?src=hash">#PHP</a> <a href="https://twitter.com/hashtag/ZF?src=hash">#ZF</a> <a href="https://twitter.com/hashtag/console?src=hash">#console</a> <a href="https://twitter.com/hashtag/develop?src=hash">#develop</a></p>&mdash; Gianluca Arbezzano (@GianArb) <a href="https://twitter.com/GianArb/status/613292048708468736">June 23, 2015</a></blockquote>
<script async src="//platform.twitter.com/widgets.js" charset="utf-8"></script>

<div class="alert alert-success" role="alert"><strong>Github: </strong>Article written about <a target="_blank" href="https://github.com/gianarb/console-skeleton-app">console-skeleton-app</a> 1.0.0</div>

I'm writing a skeleton app to build console/bash application in PHP.
This project is very easy and it depends on ZF\Console a zfcampus project and Zend\Console builds by ZF community.
I have a todo list for the future but for the time being it's just a blog post about these two modules.

* Integration with container system to manage dependency injection
* Docs to test your command
* Use cases and different implementations

## ZF\Console and other components

* [ZF\Console](https://github.com/zfcampus/zf-console) is maintained by zfcampus and it is used by Apigility
* [zendframework\zend-console](https://github.com/zendframework/zend-console) is maintained by zendframework, all the info are in the [documantation](http://framework.zend.com/manual/current/en/modules/zend.console.introduction.html)

## Tree

This is my folders structure proposal, there are three entrypoint in the `bin` directory, one for bash, one for php and a bat for Window.
I use composer to manage my dependencies and I included .lock file because this project is an APPLICATION not a library..
`/config` directory contains only routing definitions but in the future we can add services and other configurations.
`src/Command/` contains my commands.

{% highlight bash %}
├── bin
│   └── console.php
├── composer.json
├── composer.lock
├── config
│   └── routes.php
├── src
│   └── Command
│       ├── Conf.php
│       ├── Database.php
│       └── Download.php
└── vendor
    └── ...
{% endhighlight %}

## Bootstrap

The Application's entrypoints are just example and they require few changes.
First we have to change the version in the parameters.php configuration file and also change the application name `'app'` to what fits.
To load configurations from different sources I will use the well known `Zend\Config` component.

{% highlight php %}
<?php
require __DIR__.'/../vendor/autoload.php';

use Zend\Console\Console;
use ZF\Console\Application;
use ZF\Console\Dispatcher;

$version = '0.0.1';

$application = new Application(
    'app',
    $version,
    include __DIR__ . '/../config/routes.php',
    Console::getInstance(),
    new Dispatcher()
);

$exit = $application->run();
exit($exit);
{% endhighlight %}

## Routes
`config/routes.php` contains router configurations. This is just an example but you can see all options [here](https://github.com/zfcampus/zf-console#defining-console-routes).

{% highlight php %}
<?php
return [
    [
        'name'  => 'hello',
        'route' => "--name=",
        'short_description' => "Good morning!! This is a beautiful day",
        "handler" => ['App\Command\Hello', 'run'],
    ],
];
{% endhighlight %}

## Command

Basic command to wish you a good day!
I decided that a command doesn't extends any class because in my opinion is a good way to impart readability and simplicity.

{% highlight php %}
<?php
namespace App\Command;

use ZF\Console\Route;
use Zend\Console\Adapter\AdapterInterface;

class Hello
{
    public static function run(Route $route, AdapterInterface $console)
    {
        $name = $route->getMatchedParam("name", "@gianarb");
        $console->writeLine("Hi {$name}, you have call me. Now this is an awesome day!");
    }
}
{% endhighlight %}

## Troubleshooting and tricks
* OSx return an error because zf-console use a function blocked into the mac os php installation. Have a look at  PR[#22](https://github.com/zfcampus/zf-console/pull/22)
* See [this](http://www.sitepoint.com/packaging-your-apps-with-phar/) article to package your application in a phar archive..

<br/>
<br/>
<br/>

<div class="well"><a target="_blank" href="https://twitter.com/__debo">@__debo</a> thanks for trying to fix my bad English</div>
