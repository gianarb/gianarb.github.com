---
layout: post
title:  "Influx DB and PHP implementation"
date:   2015-01-06
categories: php, influxdb
tags: php, influxdb
summary: "InfluxDB a time series database and PHP integration"
priority: 1
---
Influx DB is [time series database](http://en.wikipedia.org/wiki/Time_series_database) written in Go.

It supports SQL like queries and it has different entry points, REST API (tcp protocol) and UDP.

<div class="row">
<div class="col-md-4 col-md-offset-3"><img class="img-responsive" src="http://influxdb.com/images/influxdb-light400.png"></div>
</div>
We wrote a [sdk](https://github.com/corley/influxdb-php-sdk) to manage integration between Influx and PHP.

It supports Guzzle Adapter but if you use Zend\Client you can write your implementation.

{% highlight php %}
<?php
$guzzle = new \GuzzleHttp\Client();

$options = new Options();
$adapter = new GuzzleAdapter($guzzle, $options);

$client = new Client();
$client->setAdapter($adapter);
{% endhighlight %}

In this case we are using a Guzzle Client, we communicate with Influx in TPC, but we can speak with it in UDP

{% highlight php %}
<?php
$options = new Options();
$adapter = new UdpAdapter($options);

$client = new Client();
$client->setAdapter($adapter);
{% endhighlight %}

Both of them have the same usage

{% highlight php %}
<?php
$client->mark("app.search", $points, "s");
{% endhighlight %}

The first different between udp and tcp is known, TPC after request expects a response, UDP does not expect anything and in this case does not exist any delivery guarantee.
If you can accept this stuff this is the benchmark:

{% highlight bash %}
Corley\Benchmarks\Influx DB\AdapterEvent
    Method Name                Iterations    Average Time      Ops/second
    ------------------------  ------------  --------------    -------------
    sendDataUsingHttpAdapter: [1,000     ] [0.0026700308323] [374.52751]
    sendDataUsingUdpAdapter : [1,000     ] [0.0000436344147] [22,917.69026]
{% endhighlight %}
