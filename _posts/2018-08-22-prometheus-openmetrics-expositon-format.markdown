---
layout: post
title:  "OpenMetrics and the future of the prometheus exposition format"
date:   2018-08-23 08:08:27
categories: [post]
tags: [codeinstrumentation, go, open source, metrics, prometheus, influxdb, openmetrics, cncf, oss,
exposition format, snmp]
summary: "This post explain my point of view around prometheus exposition format
and it summarise the next step with OpenMetrics behing supported by CNCF and
other big companies."
changefreq: daily
---
Who am I to tell you the future about the prometheus exposition format? Nobody!

I was at the PromCon in Munich in August 2018 and I found the conference great!
A lot of use cases about metrics, monitoring and prometheus itself. I work
at InfluxData and we was there as sponsor but I followed a lot of talks and I
had the chance to attend the developer summit the next day with a lot of
promehteus maintainers. Really good conversarsations!

<blockquote class="twitter-tweet tw-align-center" data-lang="en"><p lang="en" dir="ltr">I just
realized how lucky I was these days having chance to be so welcomed by the <a
href="https://twitter.com/hashtag/prometheus?src=hash&amp;ref_src=twsrc%5Etfw">#prometheus</a>
community. I love my work. Thanks <a
href="https://twitter.com/juliusvolz?ref_src=twsrc%5Etfw">@juliusvolz</a> <a
href="https://twitter.com/TwitchiH?ref_src=twsrc%5Etfw">@TwitchiH</a> <a
href="https://twitter.com/tom_wilkie?ref_src=twsrc%5Etfw">@tom_wilkie</a> and
everyone.. I feel regenerated</p>&mdash; :w !sudo tee % (@GianArb) <a
href="https://twitter.com/GianArb/status/1028414240535793664?ref_src=twsrc%5Etfw">August
11, 2018</a></blockquote>
<script async src="https://platform.twitter.com/widgets.js" charset="utf-8"></script>


To be honest my scope few years ago was very different, I was working in PHP
writing webapplication that yes I was deploying but I wasn't digging to much
around them and I was not smart enough to uderstand that all the pull vs push situation was just all garbage.
Smoke in the eyes that luckily I left behind me pretty soon because I had the
chance to meet smart people that drove me out.

Provide a comfortable way for me to expose and store metrics is a vital
request and the library needs to expose the RIGHT data it doesn't matter if they
are pushing or pulling.

RIGHT means the best I can get to have more observability from an ops point of
view, but also from a business intelligence prospetive probably just
manipulating again the same data.

It is safe to say that a pull based exposition format is easy to pack together
because it works even if the server that should grab the exposed endpoint is
unavailable or even if nothing will grab them. A push based service will always
create some network noice even if nobody has interest on getting the metrics.

Back in the day we had SNMP but other than being an internet standard the
adoption is not comparable with the prometheus one, if we had how old it is and
how fast prometheus growed the situation gets even worst.

```
.1.0.0.0.1.1.0 octet_str "foo"
.1.0.0.0.1.1.1 octet_str "bar"
.1.0.0.0.1.102 octet_str "bad"
.1.0.0.0.1.2.0 integer 1
.1.0.0.0.1.2.1 integer 2
.1.0.0.0.1.3.0 octet_str "0.123"
.1.0.0.0.1.3.1 octet_str "0.456"
.1.0.0.0.1.3.2 octet_str "9.999"
.1.0.0.1.1 octet_str "baz"
.1.0.0.1.2 uinteger 54321
.1.0.0.1.3 uinteger 234
```

It also started as network exposing format, so it doesn't express really well
other kind of metrics.

The [prometheus exposition
format](https://github.com/prometheus/docs/blob/master/content/docs/instrumenting/exposition_formats.md)
is extremly valuable and I recently instrumented a legacy application using the
prometheus sdk and my code looks a lot more clean and readable.

At the beginning I was using logs as transport layer for my metrics and time
series but I ended up having a lot of spam in log themself because I was also
streaming a lot of "not logs but metrics" garbage.

The link to the prometheus doc above is the best place to start, here I am just
copy pasting something form there:

```
# HELP http_requests_total The total number of HTTP requests.
# TYPE http_requests_total counter
http_requests_total{method="post",code="200"} 1027 1395066363000
http_requests_total{method="post",code="400"}    3 1395066363000

# Escaping in label values:
msdos_file_access_time_seconds{path="C:\\DIR\\FILE.TXT",error="Cannot find file:\n\"FILE.TXT\""} 1.458255915e9

# Minimalistic line:
metric_without_timestamp_and_labels 12.47

# A weird metric from before the epoch:
something_weird{problem="division by zero"} +Inf -3982045

# A histogram, which has a pretty complex representation in the text format:
# HELP http_request_duration_seconds A histogram of the request duration.
# TYPE http_request_duration_seconds histogram
http_request_duration_seconds_bucket{le="0.05"} 24054
http_request_duration_seconds_bucket{le="0.1"} 33444
http_request_duration_seconds_bucket{le="0.2"} 100392
http_request_duration_seconds_bucket{le="0.5"} 129389
http_request_duration_seconds_bucket{le="1"} 133988
http_request_duration_seconds_bucket{le="+Inf"} 144320
http_request_duration_seconds_sum 53423
http_request_duration_seconds_count 144320

# Finally a summary, which has a complex representation, too:
# HELP rpc_duration_seconds A summary of the RPC duration in seconds.
# TYPE rpc_duration_seconds summary
rpc_duration_seconds{quantile="0.01"} 3102
rpc_duration_seconds{quantile="0.05"} 3272
rpc_duration_seconds{quantile="0.5"} 4773
rpc_duration_seconds{quantile="0.9"} 9001
rpc_duration_seconds{quantile="0.99"} 76656
rpc_duration_seconds_sum 1.7560473e+07
rpc_duration_seconds_count 2693
```

Think about that not as the prometheus way to grab metrics, but as the language
that your application uses to teach the outside world how does it feels.

It is just a plain text entrypoint over HTTP that everyone can parse and re-use.

For example
[kapacitor](https://www.influxdata.com/time-series-platform/kapacitor/) or
[telegraf](https://www.influxdata.com/time-series-platform/telegraf/) have
specific ways to parse and extract metrics from that URL.

If you don't have time to write a parser for that you can use
[prom2json](https://github.com/prometheus/prom2json) to get a JSON version of
that.

In Go you can dig a bit more inside that code and reuse some of functions for
example:

```go
// FetchMetricFamilies retrieves metrics from the provided URL, decodes them
// into MetricFamily proto messages, and sends them to the provided channel. It
// returns after all MetricFamilies have been sent.
func FetchMetricFamilies(
	url string, ch chan<- *dto.MetricFamily,
	certificate string, key string,
	skipServerCertCheck bool,
) error {
	defer close(ch)
	var transport *http.Transport
	if certificate != "" && key != "" {
		cert, err := tls.LoadX509KeyPair(certificate, key)
		if err != nil {
			return err
		}
		tlsConfig := &tls.Config{
			Certificates:       []tls.Certificate{cert},
			InsecureSkipVerify: skipServerCertCheck,
		}
		tlsConfig.BuildNameToCertificate()
		transport = &http.Transport{TLSClientConfig: tlsConfig}
	} else {
		transport = &http.Transport{
			TLSClientConfig: &tls.Config{InsecureSkipVerify: skipServerCertCheck},
		}
	}
	client := &http.Client{Transport: transport}
	return decodeContent(client, url, ch)
}
```
[FetchMetricsFamilies](https://github.com/prometheus/prom2json/blob/master/prom2json.go#L123) can be used to get a channel with all the fetched
metrics. When you have the channel you can make what you desire:

```go
mfChan := make(chan *dto.MetricFamily, 1024)

go func() {
    err := prom2json.FetchMetricFamilies(flag.Args()[0], mfChan, *cert, *key, *skipServerCertCheck)
    if err != nil {
        log.Fatal(err)
    }
}()

result := []*prom2json.Family{}
for mf := range mfChan {
    result = append(result, prom2json.NewFamily(mf))
}
```

As you can see
[`prom2json`](https://github.com/prometheus/prom2json/blob/master/cmd/prom2json/main.go#L42)
converts the result to JSON.

It is pretty fleximple! And it is a common API to read applicatin status. A
common API we all know means automation! Dope automation!

## Future
The prometheus exposition format growed in adoption across the board and a
couple of people leaded by [Richard](https://twitter.com/TwitchiH) are now pushing
to have this format as new Internet Standard!

The project is called [OpenMetrics](https://openmetrics.io/) and it is a Sandbox
project under CNCF.

if you are looking to follow the project here the official repository on
[GitHub](https://github.com/OpenObservability/OpenMetric).

Probably it looks just a political step with no value at all from a tech point of
view but I bet when it will be a standard and not just "the prometheus
exposition" we will start to have routers exposing stats over
`http://192.168.1.1/metrics` and it will be a lot of fun!

It will be obvious that it is not a `only-prometheus` feature and this new group
has people from difference companies and backgrounds. So the exposition format
will be probably not just for operational metrics but more generic.
