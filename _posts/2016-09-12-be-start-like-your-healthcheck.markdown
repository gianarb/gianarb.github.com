---
layout: post
title:  "Be smart like your healthcheck"
date:   2016-08-25 12:08:27
categories: [post]
img: /img/docker.png
tags: [docker, distributed_system, distributed system, automation, devops,
health check]
summary: "In a distributed system environment has a simple way to know the
status of your server help you do understand if it's ready to go in production.
HealthCheck is simple and common but design a good one can help you do avoid
strange behaviors. Docker 1.12 supports healthcheck and we in this blog I share
an example of implementation."
changefreq: yearly
---
I am not a doctor, I am a Software Engineer and this is a tech post! You can
continue to read!

To monitor monolithic what we usually do is install a tool
like [Nagios](https://www.nagios.org/) to centralize all our metrics and to
stay in touch with our infrastructure and our application.  In a distributed
system with more that one services with own metrics the situation is totally
different.  This about how it’s more dynamic respect a monolithic.  Containers
or VM that scale up and down and that move around the network, Nagios is a good
solution to check if our new service after a deploy is safe and ready to be
attached into the production pool?  I love a talk made by [Kelsey
Hightower](https://github.com/kelseyhightower) during the Monitorama event, he
speak about healthcheck watch him to follow a [great demo](
https://vimeo.com/173610242)!

Healthcheck is an API that your service exposes to share it’s status, if you
make it really start it’s a good tool to understand the situation of your
service with just a call.  A service could be ready or not and it’s in the best
situation to communicate its status.  It’s a like a patient, you need to ask
him all what you need to make the best diagnosis and take a decision about it.

<div class="post row">
  <div class="col-md-12">
      {% include book-adv-lb.html %}
  </div>
</div>

We can stay focused on a REST service, it exposes an API under the route
/health. The response could has two different Status Code:

* 200 if all it’s good and you service is ready
* 500 it there is something wrong and your service is not ready

To make an smart HealthCheck what do we need to check?

This is a real implementation:

```php
<?php
echo 1;
```

It’s better that nothing but we are looking for something smart!  We need to
check all dependendencies that our service has and it’s for this reason that
the service itself is the best actor because it knows what it need to be ready.
I wrote a demo service, the name is [micro](
https://github.com/gianarb/micro/blob/master/handle/health.go), it’s in go and
the version 2 use
mysql.

```go
func Health(username string, password string, addr string) func(http.ResponseWriter, *http.Request) {
    return func(w http.ResponseWriter, r *http.Request) {
        res := healtResponse{Status: true}
        httpStatus := 200
        dsn := fmt.Sprintf("%s:%s@tcp(%s:3306)/micro", username, password, addr)
        ddb, err := sql.Open("mysql", dsn)
        if err != nil {
            log.Fatal(err)
        }
        if err := ddb.Ping(); err != nil {
            res.Status = false
            res.Info = map[string]string{"database": err.Error()}
        }
        c, _ := json.Marshal(res)
        if res.Status == false {
            httpStatus = 500
        }
        log.Println("%s called /health", r.Host)
        w.WriteHeader(httpStatus)
        w.Header().Set("Content-Type", "application/json")
        w.Write(c)
    }
}
```
Doesn’t matter how many dependencies you service has, you need to check all of
them, databases, other services that it uses.  In my case I decided to add a
key-value field, I called it `info`, it contains some information about whether
mysql is or is not not working, in order to make the debug flow easy.  If the
service that you are checking has an healthcheck you are lucky! You can use
that entrypoint to know if your dependency is fine.  If you are not so lucky if
you can create a wrapper or just check if you can reach the service, in my case
I just tried to connect to mysql in order to know if my network supports me! I
also using the correct database name in order to avoid edge case like “mysql is
on but the database doesn’t exist”.

The ecosystem supports healthchecks! Nginx looks it to know if a server is
reachable, if the health check doesn’t work for a while it just make the server
out for few times. Same for Kubernetes, Swarm and Docker.  Docker provides a
library in go an [healthcheck
framework](https://github.com/docker/go-healthcheck) that you can use in your
applications, it is also used in Docker 1.12.

You can describe in your  Dockerfile an HealthCheck

```
HEALTHCHECK CMD ./cli health
```

If the exit code is 0 Docker marks you container like healthy if it’s different like unhealthy.
Very easy and flexible, you can check your REST healthcheck in this way

```
HEALTHCHECK --interval=30s --timeout=30s --retries=3 \
  CMD curl -si localhost:8000/health | grep 'HTTP/1.1 200 OK' > /dev/null
```

`--interval` is the timing between two healthcheck, `--timeout` is used to mark
like unhealthy a service that doesn’t come back after 30s in this case,
`--retries` is the attempts to do before make a container unhealthy.

HealtCheck doesn’t replace traditional monitoring system but with a lot of
instances and services has a single point to check and understand the situation
after a deploy make your like easy and your products stable.
