---
layout: post
title:  "Zf2 Event, base use"
date:   2013-11-21 12:38:27
img: /img/zf.jpg
categories: zf2, php, event
tags: zf2, zend framework, php, events
summary: "Zend Framework 2 base use of Events and Event manager"
priority: 1
---

Hi! Some mouths ago I have writted a gist for help me to remember a base use of Events and Event Manager into Zend Fremework,
in this article I report this small tutorial.

{% highlight php %}
<?php
require_once __DIR__."/vendor/autoload.php";

class Foo
{
    /* @var \Zend\EventManager\EventManagerInterface */
    protected $eventManager;

    public function getEventManager()
    {
        if(!$this->eventManager instanceof \Zend\EventManager\EventManagerInterface){
            $this->eventManager = new \Zend\EventManager\EventManager();
        }
        return $this->eventManager;
    }

    public function echoHello()
    {
        $this->getEventManager()->trigger(__FUNCTION__."_pre", $this);
        echo "Hello";
        $this->getEventManager()->trigger(__FUNCTION__."_post", $this);
    }
}

$foo = new Foo();
$foo->getEventManager()->attach('echoHello_pre', function($e){
    echo "Wow! ";
});
$foo->getEventManager()->attach('echoHello_post', function($e){
    echo ". This example is very good! \n";
});
$foo->getEventManager()->attach('echoHello_post', function($e){
    echo "\nby gianarb92@gmail.com \n";
}, -10);
$foo->echoHello();
{% endhighlight %}

The result

{% highlight bash %}
gianarb@GianArb-2 eventTest :) $ php try.php
Wow! Hello. This example is very good!

by gianarb92@gmail.com
{% endhighlight %}


[@see Zend Event Manager Ref](http://framework.zend.com/manual/2.0/en/modules/zend.event-manager.event-manager.html)
