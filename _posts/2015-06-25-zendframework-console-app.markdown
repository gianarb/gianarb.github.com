---
layout: post
title:  "Build your Zend Framework Console Application"
date:   2015-05-21 23:08:27
categories: php
tags: php, console, automation, zend framework
summary: ZF\Console is a component written by zf-campus and Apigility organization that help you to build console application using different Zend Framework components
priority: 0.6
changefreq: yearly
---
I'm writing a skeleton app to build console/bash application in PHP.  
This project is very easy and it depends from ZF\Console a zfcampus project and Zend\Console builds by ZF community.  
In this moment it's only a blog of this two modules but in the future I have a to do list.  

* Integration with container system to manage dependence injaction
* Docs to test your command
* Use case and different implementation

## ZF\Console and other components

* [ZF\Console](https://github.com/zfcampus/zf-console) is maintained by zfcampus and it is used by Apigility
* [zendframework\zend-console](https://github.com/zendframework/zend-console) is maintained by zendframework, all info into the [docs](http://framework.zend.com/manual/current/en/modules/zend.console.introduction.html)

## Tree

This is my foldering's proposal, there are three entrypoint into the `bin` directory, bash, php and bat for Window.  
I use composer to manage dependencies and I have included .lock file because this project is an APPLICATION not a library..  
`conf` directory in this moment contains only routing definitions but in the future we can add services and other configurations.  
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

Applications' entrypoint, in this moment is very simple and it requires features..  
In first we can move on the version into the parameters.php configuration file, same stuff for 'app' the application's name..  
Ro load configurations from different sources I will use `Zend\Config`, very good component IMO.

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
`config/routes.php` contains router configurations this is an example but you can see all options [here](https://github.com/zfcampus/zf-console#defining-console-routes).

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
I have choised that a command doesn't extends any class because in my opinion is a good way to impart readability and simplicity.

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
* OSx return an error because zf-console use a function blocked into the mac os php installation.. Follow [#22](https://github.com/zfcampus/zf-console/pull/22)
* [See this article](http://www.sitepoint.com/packaging-your-apps-with-phar/) to package your application in a phar archive.. 
