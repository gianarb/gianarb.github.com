---
layout: post
title:  "InfluxDB PHP 1.3.0 is ready to go"
date:   2016-02-18 10:08:27
categories: [post]
img: /img/influx.jpg
tags: [devops, influxdb, php, oss, open source, metrics, monitoring,
observability]
summary: "InfluxDB is a time series database, it helps us to manage matrics,
point and offert a stack of tools to collect and see this type of data. I am a
maintainer of InfluxDB PHP integration. In this past I describe the news
provided by new relesae 1.3.0"
priority: 0.6
changefreq: weekly
---

<blockquote class="twitter-tweet tw-align-center" data-lang="en"><p lang="en" dir="ltr">Shout out to <a href="https://twitter.com/GianArb">@GianArb</a> for shipping a new release of the InfluxDB-PHP library! Here&#39;s what&#39;s new: <a href="https://t.co/tJQIu9OCbL">https://t.co/tJQIu9OCbL</a></p>&mdash; InfluxData (@InfluxDB) <a href="https://twitter.com/InfluxDB/status/704403294592970752">February 29, 2016</a></blockquote>
<script async src="//platform.twitter.com/widgets.js" charset="utf-8"></script>

We are happy to annuonce a new minor release, [Influxdb-php library](https://github.com/influxdata/influxdb-php) 1.3.0.

This is a list of PRs merged in 1.3.0 since 1.2.2:

* [#36](https://github.com/influxdata/influxdb-php/pull/36) Added quoting of dbname in queries
* [#35](https://github.com/influxdata/influxdb-php/pull/35) Added orderBy to query builder
* [#37](https://github.com/influxdata/influxdb-php/pull/37) Fixed wrong orderby tests
* [#38](https://github.com/influxdata/influxdb-php/pull/38) Travis container-infra and php 7

The `QueryBuilder` now support the orderBy function to order our data, InfluxDB supports it from version 0.9.4.

{% highlight sql %}
select * from cpu order by value desc
{% endhighlight %}

Now you can do it in PHP

{% highlight php %}
$this->database->getQueryBuilder()
  ->from('cpu')
  ->orderBy('value', 'DESC')->getQuery();
{% endhighlight %}

We are increase our Continuous Integration system in order to check our code with PHP7, it's perfect!

We escape our query to support reserved keyword like `database`, `servers` personally I prefer avoid this type of word but you are free to use them.

Please we are very happy to understand as the PHP community use this library and InfluxDB, please share your experience and your problem into the repository, on IRC (join influxdb on freenode) and we wait you on [Twitter](https://twitter.com/influxdata).

Remeber to update your `composer.json`!

```json
{% highlight json %}
{
    "require": {
        "influxdb/influxdb-php": "~1.3"
    }
}
{% endhighlight %}

A big thanks at all our contributors!
