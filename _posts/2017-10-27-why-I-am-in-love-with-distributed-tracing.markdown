---
layout: post
title:  "Why I am in love with distributed tracing"
date:   2017-10-27 10:08:27
categories: [post]
tags: [tracing, monitoring]
img: /img/caos.jpg
summary: ""
changefreq: yearly
---

Cloud computing, serverless, microservices everyone now is a little bit more a
Distributed System Engineer.

Everyone was is love with disaster recovery, application and datacenter
distribution on AWS or on different cloud providers. I remember that time! Every
morning I was scared waiting for a fire or for a natural disaster close to the
single datacenter that was hosting my application.

<img src="/img/graph-dots.png" class="img-responsive"/>

A distributed system has a different complexity. I can't say that an old fashion
Java monolithic with threads, concurrency and so on is easy to manage compared
with a modern microservices environment. They have a different complexity.

Logs for example are not giving a useful high level information about what is
happening in you system. You always need a log to get some details down the
stack but using log to discovery problem in a microservices environment can't
work:

1. Where the log is? In which server? container? functions?
2. Which applications or service is failing?
3. Is the log that I am reading the root cause of the problem?

Look at the style of a log, it can be described as stream of information. An
append only file that contains in time ordered record about what my application
is doing. A modern system doesn't look like it anymore. It looks more like a
graph:

* vertices are our applications, functions, containers, services, microservices,
  third-party API, external services or what ever
* Edges are the connection between two applications. One application can
  communicate with one or more than one applications. Obviously as side effect
  one application can be called from one or event more applications.

Understanding how a request traverse this graph, where it fails and re-build
it's path is the challenge.
When you know these answers you will be able to read the right logs.

<img src="/img/ecosystem_graph.png" class="img-responsive"/>

Your frontend communicate with the OAuth service to get authorization
information and with the car service to print the list of cars. It also
uses the Google Analytics sdk to track what the user is doing.

The car service interrogate an external service called "Ford warehouse" to get
the cars available and the 0Auth service to get the verify if the user is
allowed to get the information that is asking for.

All the services are communicating with the internal monitor to store events,
metrics and so on.

The car service is using MongoDB as database, the user one is using MySql to
keep relations between users and all the information like emails, address, phone
number and so on. Auth0 is using a different MySql.

I stopped to design things just for time and space constrains. The frontend is
also providing a search box and it's communicating with the Search Service. It
uses ElasticSearch to index the cars document. It uses Kafka as event sourcing
to keep in sync the index on ElasticSearch based on the new info that Car
Service gets from the Ford Warehouse.

Yes, I am enjoying this game! What about AI? Big Data?? Where is Spark? I will
leave this questions for you.

Anyway as you can see it's not a linear stream of data anymore. It looks more
like a mess.

Understand how and where a request can fails can be very hard. What about
latency between services? Now everyone uses contains and things like Nomad, k8s,
Swarm. The orchestrator splits contains around our infrastructure as magic balls
but what about the network efficiency between two of them?
What happen if 0Auth and the User Service are too far from each other and we
lost half of the time waiting for a request to move between these two services?
How can we measure and visualize these things?

To avoid vendor lock-in and to make the tracing simple and modular CNCF promotes
a project called [Opentracing](http://opentracing.io/). It provides a set of
interfaces and specification that every languages or tracing service should
implement.

There is an entire [github organisation](https://github.com/opentracing) with a
lot of sdk for different languages like: PHP, Java, Go, C++.

Other than sdk and a common interface you need a tracer to manipulate and store
every trace.

* Trace is a collection of spans.
* Span is the smallest unit of time. You can create a span for every
  microservices for example. It has a span_id and a trace_id.Span is the
  smallest unit of time. You can create a span for every microservices for
  example. It has a span_id and a trace_i.

There are two famous open source tracer:

* [Zipkin](https://github.com/openzipkin/zipkin) is an open source project
  written in Java. It's made by Twitter and it now support opentracing.
* [Jaeger](https://github.com/jaegertracing/jaeger) is a Uber's project written
  in golang. It's under the CNCF and it support opentracing.

They are similar, they provides an HTTP api and both of them
supports different backend as storage like: Cassandra, ElasticSearch, Postgres.

Other vendors support Opentracing:

* [NewRelic](https://blog.newrelic.com/2017/09/13/distributed-tracing-opentracing/)
* [AWS Xray](https://aws.amazon.com/xray/?nc1=h_ls)

As you can see the ecosystem of open source and vendors grow fast because this
looks a solution for a problem that we all have. It is very hard to understand
what is going in our applications.
 
Here some codes that we can use as example:

```go
import "github.com/opentracing/opentracing-go"
import ".../some_tracing_impl"

func main() {
    // example zipkin or newrelic
	tracer := some_tracing_impl.New(...),
	opentracing.InitGlobalTracer(tracer)
}
```

As I wrote before the tracer needs to provide a opentracing compatible client,
at this point we can inject it in the opetracing-sdk. This solution allow us to
change the tracer service without modify our application.

Be able to disable tracing without modify our application is useful. You can
use an environment variable and the go sdk providers a
`opentracing.NoopTracer{}`. It's just a Tracer implementation that doesn't do anything.

Every sdk is well documented with code example, you can see the [go-sdk
here](https://github.com/opentracing/opentracing-go).

Be careful during a span creation, you should start checking if the request has
already some tracing information. If it doesn't have anything it means that
it's the frontend service, it's the first application that handles the request.
At this point you can create a new span.
