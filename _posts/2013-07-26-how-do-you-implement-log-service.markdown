---
layout: post
title:  "Zend Framework 2 - How do you implement log service?"
date:   2013-07-26 23:08:27
categories: php
tags: zf2, logger, log, zend framework, service manager
summary: Implementation of logger service in Zend Framework 2
---

A log system is an essential element for any application. It is a way to check the status and use of the application. For a basic implementation you can refer to the  fig-stanrdars organization [PSR-3](https://github.com/php-fig/fig-standards/blob/master/accepted/PSR-3-logger-interface.md) article, that describes th elogger interface.

Zend Framework 2 implement a [Logger Component](https://github.com/zendframework/zf2/tree/master/library/Zend/Log), the following is an example of how to use it with service manager.

{% highlight php %}
<?php
return array(
  'service_manager' => array(
    'abstract_factories' => array(
      'Zend\Log\LoggerAbstractServiceFactory',
    ),
  ),
  'log' => array(
      'Log\App' => array(
        'writers' => array(
      	  array(
    				'name' => 'stream',
   					'priority' => 1000,
   					'options' => array(
    					'stream' => 'data/app.log',
    				),
    			),
    		),
   		),
    ),
);
{% endhighlight %}
[LoggerAbstractServiceFactory](https://github.com/zendframework/zf2/blob/master/library/Zend/Log/LoggerServiceFactory.php) is a Service Factory, as an example,  into service Manager class Logger and will be used in the whole application. Log/App is the name of a single logger, and writer is an adapter that is used to choose the method of writing, in this case everything is written to file, but you can use a DB adapter and write your log into database.

{% highlight php %}
<?php
namespace GianArb\Controller;
class GeneralController
  extends AbastractActionController
{
  public function testAction(){
    $logger = $this->getServiceLocator()->get('Log\App');
    $logger->log(\Zend\Log\Logger::INFO, "This is a little log!");
  }
}
{% endhighlight %}

With this configuration Log\App writes a string into data/app.log file, with INFO property. By default you can use an array of properties.

{% highlight php %}
<?php
protected $priorities = array(
  self::EMERG  => 'EMERG',
  self::ALERT  => 'ALERT',
  self::CRIT   => 'CRIT',
  self::ERR    => 'ERR',
  self::WARN   => 'WARN',
  self::NOTICE => 'NOTICE',
  self::INFO   => 'INFO',
  self::DEBUG  => 'DEBUG',
);
{% endhighlight %}

Usage of different keys is a good practice because it is very easy to write filter or log categories.

Another good practice, valid for all services in general, is to create your class extending single service.

{% highlight php %}
<?php
use Zend\Log\Logger
class MyLogger extends Logger
{% endhighlight %}

This choice helps managing future customizations  of services and is another important layer for managing unexpected updates.

Rali, thanks for your help with my robotic english! :P
