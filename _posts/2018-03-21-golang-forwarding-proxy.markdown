---
layout: post
title:  "How to use a Forwarding Proxy with golang"
date:   2018-03-21 09:38:27
img: /img/gopher.png
categories: [post]
tags: [golang, network, go, proxy, forwarding, privoxy, docker, kubernetes]
summary: "Cloud, Docker, Kubernetes make your environment extremely dynamic, it
has a lot of advantages but it adds another layer of complexity. This article is
about forward proxy and golang. How to configure your http Client to use an
http, https forward proxy for your golang application to increase security,
scalability and to have a set of public ips for outbound traffic."
priority: 1
---

A forwarding proxy is a proxy configuration that handle requests from a set of
internal clients that are trying to create a connection to the outside.

In practice is a man in the middle between your application and the server that you are
trying to connect. It works over the HTTP(S) protocol and it is implemented at the
edge of your infrastructure.

Usually, you can find it in large organizations or universities and it is used as
additional control mechanism for authorization and security.

I find it useful when you work with containers or in a dynamic cloud environment
because you will have a set of servers for all the outbound network
communication.

If you work in a dynamic environment as AWS, Azure and so on you will end up
having a variable number of servers and also a dynamic number of public IPs.
Same if your application runs on a Kubernetes cluster. Your container can be
everywhere.

Now let's suppose that a customer asks you to provide a range of public IPs
because he needs to set up a firewall... How can
you provide this feature?  In some environments can be very simple, in others
very complicated.

1st December 2015 a users asked this question on the [CircleCI
forum](https://discuss.circleci.com/t/circleci-source-ip/1202) this request is
still open. This is just an example, CircleCi is great. I am not complaining
about them.

One of the possible ways to fix this problem is via the forwarding proxy. You can
spin up a set of nodes with a static ip and you can offer the list to the
customer.

Almost all cloud providers have a way to do that, floating ip
on DigitalOcean or Elastic IP on AWS.

You can configure your applications to forward the requests to that pool and
the end services will get the ip from the forward proxy nodes and not from the
internal one.

This can be another security layer for your infrastructure because you will be
able to control and scan packages that are going outside from your network in a
really simple way and in a centralized place.

This is not a single point of failure because you can spin up more than one
forward proxies and they scale really well.

Under the hood, a forward proxy is the [HTTP method
`CONNECT`](https://developer.mozilla.org/en-US/docs/Web/HTTP/Methods/CONNECT).

> The CONNECT method converts the request connection to a transparent TCP/IP
tunnel, usually to facilitate SSL-encrypted communication (HTTPS) through an
unencrypted HTTP proxy.

A lot of HTTP Client across languages already support this in a very
transparent way. I built a very small example using golang and
[privoxy](https://www.privoxy.org/) to show you how simple it is.


First of all, let's build an application called `whoyare`. It is an HTTP server
that returns your remote address:

```go
package main

import (
	"encoding/json"
	"net/http"
)

func main() {
	http.HandleFunc("/whoyare", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		body, _ := json.Marshal(map[string]string{
			"addr": r.RemoteAddr,
		})
		w.Write(body)
	})
	http.ListenAndServe(":8080", nil)
}
```

You can call the `GET` the route `/whoyare` and you will receive a JSON like
`{"addr": "34.35.23.54"}` where `34.35.23.54` is your public address. Running
`whoyare` from your laptop if you make a request on your terminal you should get
`localhost` as remote address. You can use curl to try it:

```bash
18:36 $ curl -v http://localhost:8080/whoyare
* TCP_NODELAY set
> GET /whoyare HTTP/1.1
> User-Agent: curl/7.58.0
> Accept: */*
>
< HTTP/1.1 200 OK
< Content-Type: application/json
< Date: Sun, 18 Mar 2018 17:36:40 GMT
< Content-Length: 31
<
* Connection #0 to host localhost left intact
{"addr":"localhost:38606"}
```

I wrote another application, it uses `http.Client` to print the response on
stdout. If you have the server running you can run it:

```go
package main

import (
	"io/ioutil"
	"log"
	"net/http"
	"os"
)

type whoiam struct {
	Addr string
}

func main() {
	url := "http://localhost:8080"
	if "" != os.Getenv("URL") {
		url = os.Getenv("URL")
	}
	log.Printf("Target %s.", url)
	resp, err := http.Get(url + "/whoyare")
	if err != nil {
		log.Fatal(err.Error())
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err.Error())
	}
	println("You are " + string(body))
}
```

So, this is a very simple example, but you can apply this example to more
complex environments.

To make this example a bit more clear I created two virtual machines on
DigitalOcean, one will run privoxy the other one will run `whoyare`.

* **whoyare**: public ip 188.166.17.88
* **privoxy**: public ip 167.99.41.79

Privoxy is a very simple to setup forward proxy, nginx, haproxy doesn't fit very
well for this use case because they do not support the CONNECT method.

I built a docker image
[`gianarb/privoxy`](https://hub.docker.com/r/gianarb/privoxy/), it's on Docker
Hub. You can run and by default, it runs on port 8118.

```bash
core@coreos-s-1vcpu-1gb-ams3-01 ~ $ docker run -it --rm -p 8118:8118
gianarb/privoxy:latest
2018-03-18 17:28:05.589 7fbbf41dab88 Info: Privoxy version 3.0.26
2018-03-18 17:28:05.589 7fbbf41dab88 Info: Program name: privoxy
2018-03-18 17:28:05.591 7fbbf41dab88 Info: Loading filter file:
/etc/privoxy/default.filter
2018-03-18 17:28:05.599 7fbbf41dab88 Info: Loading filter file:
/etc/privoxy/user.filter
2018-03-18 17:28:05.599 7fbbf41dab88 Info: Loading actions file:
/etc/privoxy/match-all.action
2018-03-18 17:28:05.600 7fbbf41dab88 Info: Loading actions file:
/etc/privoxy/default.action
2018-03-18 17:28:05.607 7fbbf41dab88 Info: Loading actions file:
/etc/privoxy/user.action
2018-03-18 17:28:05.611 7fbbf41dab88 Info: Listening on port 8118 on IP address
0.0.0.0
```

The second step is to build and scp `whoyare` in your server. You can
build it using the command:

```
$ CGO_ENABLED=0 GOOS=linux go build -o bin/server_linux -a ./whoyare
```
Now that we have the application up and running we can try via cURL to query it
directly and via privoxy.

Let's try directly as we did previously:

```
$ curl -v http://your-ip:8080/whoyare
```

`cURL` uses an environment variable `http_proxy` to forward the requests through
the proxy:

```
$ http_proxy=http://167.99.41.79:8118 curl -v http://188.166.17.88:8080/whoyare
*   Trying 167.99.41.79...
* TCP_NODELAY set
* Connected to 167.99.41.79 (167.99.41.79) port 8118 (#0)
> GET http://188.166.17.88:8080/whoyare HTTP/1.1
> Host: 188.166.17.88:8080
> User-Agent: curl/7.58.0
> Accept: */*
> Proxy-Connection: Keep-Alive
>
< HTTP/1.1 200 OK
< Content-Type: application/json
< Date: Sun, 18 Mar 2018 17:37:02 GMT
< Content-Length: 29
< Proxy-Connection: keep-alive
<
* Connection #0 to host 167.99.41.79 left intact
{"addr":"167.99.41.79:58920"}
```
As you can see I have set `http_proxy=http://167.99.41.79:8118` and the response
doesn't contain my public ip but the proxy one.

![](/img/frankenstain-jr.jpg)

These are the logs that you should expect from privoxy for the requests crossing it:

```
2018-03-18 17:28:22.886 7fbbf41d5ae8 Request: 188.166.17.88:8080/whoyare
2018-03-18 17:32:29.495 7fbbf41d5ae8 Request: 188.166.17.88:8080/whoyare
```

The client that you run it previously by default it connects to `localhost:8080`
but you can override the target URL via env var `URL=http://188.166.17.88:8080`.
Running the following command I reached directly `whoyare`.

```
$ URL=http://188.166.17.88:8080 ./bin/client_linux
2018/03/18 18:37:59 Target http://188.166.17.88:8080.
You are {"addr":"95.248.202.252:38620"}
```

The golang `HTTP.Client` supports a set of environment
variables to configure the proxy, it makes everything very flexible because
passing
these variables to any service already running it will just work.

```
export HTTP_PROXY=http://http_proxy:port/
export HTTPS_PROXY=http://https_proxy:port/
export NO_PROXY=127.0.0.1, localhost
```
The first two are very simple, one is the proxy for the HTTP requests, the
second for HTTPS. `NO_PROXY` excludes a set of hostname, the hostname listed
there won't cross the proxy.  In my case localhost and 127.0.0.1.

```
HTT_PROXY=http://forwardproxy:8118
     +--------------+           +----------------+         +----------------+
     |              |           |                |         |                |
     |   client     +----------^+ forward proxy  +--------^+    whoyare     |
     |              |           |                |         |                |
     +--------------+           +----------------+         +----^-----------+
                                                                |
                                                                |
    +---------------+                                           |
    |               |                                           |
    |   client      +-------------------------------------------+
    |               |
    +---------------+
   HTTP_PROXY not configured
```
The client with the environment variables configured will cross the forward
proxy. Other client will reach it directly.

This granularity is very important. It's very flexible because other than a
"per-process" you can also select what request to forward and what to exclude.

```
$ HTTP_PROXY=http://167.99.41.79:8118 URL=http://188.166.17.88:8080
./bin/client_linux
2018/03/18 18:39:18 Target http://188.166.17.88:8080.
You are {"addr":"167.99.41.79:58922"}
```
As you can see we just reached `whoyare` via proxy and the `addr` in response is
now ours but the proxy one.

The last command is a bit weird but it is just to show how the `NO_PROXY` works.
We are calling the proxy excluding the `whoyare` URL, and as expected it doesn't
cross the proxy:

```
$ HTTP_PROXY=http://167.99.41.79:8118 URL=http://188.166.17.88:8080 NO_PROXY=188.166.17.88 ./bin/client_linux
2018/03/18 18:42:03 Target http://188.166.17.88:8080.
You are {"addr":"95.248.202.252:38712"}
```
Let's read this article as a practical introduction to golang, forward proxy. You can
subscribe to my [rss feed](/atom.xml) or you can follow me on
[@twitter](https://twitter.com/gianarb). Probably I will write about how to
replace `privoxy` with golang and about how to setup and deploy this solution on
Kubernetes. So let me know what to write first!
