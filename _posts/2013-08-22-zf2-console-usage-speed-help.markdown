---
layout: post
title:  "Zend Framework 2 - Console usage a speed help"
date:   2013-08-22 08:08:27
img: /img/zf.jpg
categories: php
categories: [post]
tags: [zf2, php, cli, Zend Framework, development, debug]
summary: "CLI tools are an easy way to interact with an application because you
can drive users or even other developers in a well know direction. It is a very
good way to decrease possible mistakes. Zend Framework 2 a PHP open source
framework has a Console package that helps you to address common issue like
argument management, command parsing and to format a colored and nice output."
---
With Zend Framework is very easy to write a command line tool to manage
different things. But what if there are more commands? How do you remeber them
all?

{% highlight php %}
<?php
namespace ModuleTest;
use Zend\Console\Adapter\AdapterInterface;
class Module {
	public function getConsoleUsage(AdapterInterface $console)
	{
		return array(
			array('test <params1> <params2> [--params=]', 'Description of test command'),
			array('run <action>', 'Start anction')
		);
	}
}
{% endhighlight %}

You can write this function in a Module.php file, and create a basic helper to
help you see when you write a bad command.

English by Rali :smile: Thanks!!!! :smile:
