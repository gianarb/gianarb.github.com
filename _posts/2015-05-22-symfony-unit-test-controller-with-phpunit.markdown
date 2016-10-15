---
layout: post
title:  "Test your Symfony Controller and your service with PhpUnit"
date:   2015-05-21 23:08:27
categories: [post]
img: /img/symfony.png
tags: [php]
summary: Test your Symfony Controller with PhpUnit. You expect that if one parameter is true your action get a service by Dependence Injcation and use it!
priority: 0.6
changefreq: yearly
---
<blockquote align="center" class="twitter-tweet" lang="en"><p lang="en" dir="ltr">Unit <a href="https://twitter.com/hashtag/test?src=hash">#test</a> for your <a href="https://twitter.com/hashtag/Controller?src=hash">#Controller</a> with <a href="https://twitter.com/hashtag/PhpUnit?src=hash">#PhpUnit</a> and <a href="https://twitter.com/hashtag/Symfony?src=hash">#Symfony</a>.. With a little use case of <a href="https://twitter.com/hashtag/DepedenceInjaction?src=hash">#DepedenceInjaction</a> test <a href="http://t.co/JNb39EyRly">http://t.co/JNb39EyRly</a> <a href="https://twitter.com/hashtag/php?src=hash">#php</a></p>&mdash; Gianluca Arbezzano (@GianArb) <a href="https://twitter.com/GianArb/status/601526550438215680">May 21, 2015</a></blockquote>
<script async src="//platform.twitter.com/widgets.js" charset="utf-8"></script>

In this article I would like share with you a little experience with:

* Symfony MVC
* PhpUnit
* Symfony Dependence Injaction

This is an example of very easy controller.

{% highlight php %}
<?php
namespace AppBundle\Controller;

use Sensio\Bundle\FrameworkExtraBundle\Configuration\Route;
use Symfony\Bundle\FrameworkBundle\Controller\Controller;
use Symfony\Component\HttpFoundation\Request;

class SomeStuffController extends FOSRestController
{
    /**
     * @Rest\Post("/go")
     * @return array
     */
    public function goAction(Request $request)
    {
        if($this->container->getParameter("do_stuff")) {
            $body = $this->container->get("stuff.service")->splash($request->getContent());
        }
        return [];
    }
}
{% endhighlight %}

`$this->container->getParameter("do_stuff")` is a boolean parameter that enable or disable a feature, How can I test this snippet?
I can try to write a functional test but in my opinion is easier write a series of unit tests with PhpUnit to validate my expectations.

<div class="post row">
  <div class="col-md-12">
      <a href="http://scaledocker.com?from=gianarb" target="_blank"><img src="http://scaledocker.com/adv/leaderboard.gif"></a>
  </div>
</div>

## Expectations
* If `do_stuff` parameter is false function get by my container will be call zero times
* If `do_stuff` parameter is true function get by my container will be call one times

{% highlight php %}
<?php

namespace AppBundle\Tests\Controller;

use Liip\FunctionalTestBundle\Test\WebTestCase;
use AppBundle\Controller\SomeStuffController;

class SomeStuffControllerTest extends WebTestCase
{
    public function testDoStuffIsTrue()
    {
        $request = $this->getMock("Symfony\Component\HttpFoundation\Request");
        $container = $this->getMock("Symfony\Component\DependencyInjection\ContainerInterface");
        $service = $this->getMockBuilder("Some\Stuff")->disableOriginalConstructor()->getMock();
        $container->expects($this->once())
            ->method("getParameter")
            ->with($this->equalTo('do_stuff'))
            ->will($this->returnValue(true));

        $container->expects($this->once())
            ->method("get")
            ->with($this->equalTo('stuff.service'))
            ->will($this->returnValue($service));

        $controller = new SameStuffController();
        $controller->setContainer($container);

        $controller->goAction($request);

    }
}
{% endhighlight %}
This is my first expetection "If `do_stuff` param is true I call `stuff.service`".
In this controller I use a few objects, Http\Request, Container and `stuff.service` in this example is a `Some\Stuff` class.
In the first step I have created one mock for each object.

{% highlight php %}
<?php
$request = $this->getMock("Symfony\Component\HttpFoundation\Request");
$container = $this->getMock("Symfony\Component\DependencyInjection\ContainerInterface");
$service = $this->getMockBuilder("Some\Stuff")->disableOriginalConstructor()->getMock();
{% endhighlight %}

In the second step I have written my first expetctation, "Call only one time function `getParameter` from `$container` with argument do_stuff and it returns true".

{% highlight php %}
<?php
$container->expects($this->once())
    ->method("getParameter")
    ->with($this->equalTo('do_stuff'))
    ->will($this->returnValue(true));
{% endhighlight %}
Thanks at this definitions I know that there will be another effect, my action will call only one time `$container->get("stuff.service")` and it will be return an Some\Stuff object.

The second test that we can write is "if `do_stuff` is false `$contaner->get("stuff.service")` it will not be called.

{% highlight php %}
<?php
public function testDoStuffIsFalse()
{
    $request = $this->getMock("Symfony\Component\HttpFoundation\Request");
    $container = $this->getMock("Symfony\Component\DependencyInjection\ContainerInterface");
    $service = $this->getMockBuilder("Some\Stuff")->disableOriginalConstructor()->getMock();
    $container->expects($this->once())
        ->method("getParameter")
        ->with($this->equalTo('do_stuff'))
        ->will($this->returnValue(false));

    $container->expects($this->never())
        ->method("get")
        ->with($this->equalTo('stuff.service'))
        ->will($this->returnValue($service));

    $controller = new SameStuffController();
    $controller->setContainer($container);
    $controller->goAction($request);
}
{% endhighlight %}

