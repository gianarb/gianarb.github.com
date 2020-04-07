---
img: /img/logo/otel-black-stacked.svg
layout: post
title: "How to start tracing with OpenTelemetry in NodeJS?"
date: 2020-04-07 09:08:27
categories: [post]
tags: [nodejs, codeinstrumentation, opentelemetry, tracing]
summary: "I developed an eight hours workshop about application monitoring and
code instrumentation two years ago. This year I updated it to use OpenTelemetry
and that's what I learned to instrument a NodeJS application."
changefreq: daily
---

This post is to celebrate the first beta release for the OpenTelemetry NodeJS
application <i class="fas fa-glass-cheers"></i>

Recently I developed a workshop about code instrumentation and application
monitoring. It is an 8 hours full immersion on logs, metrics, tracing and so on.
I developed it last year and I gave it twice. Let me know if you are looking for
something like that.

Almost all of it is opensource but I didn't figure out a good way to make it
usable without my brain for now. This year I updated it to use OpenTelemetry and
InfluxDB v2.

Anyway the application is called
[ShopMany](https://github.com/gianarb/shopmany). This application does
not return any useful information about its state. It is an e-commerce made of a
bunch of services in various languages. Obviously one of them is in NodeJS and
that's the one I am gonna show you today.

**Discaimer**: I can not define myself as a NodeJS developer. I wrote a bunch of
AngularJS single page application back in the day, I wrote some Cordova mobile
applications ages ago. I am not writing any JS production code since 2015 more
or less.

## First approach

I concluded the application instrumentation it was the day the maintainers
tagged the first beta release. Overnight I had to update libraries and test
code. Very luckily.

The way I learned about how to properly instrument
[discount](https://github.com/gianarb/shopmany/tree/master/discount) required a
lot of digging in the actual
[opentelemery-js](https://github.com/open-telemetry/opentelemetry-js) but
luckily for us it has a lot of examples and the library is designed to load a
bunch of useful modules that are able to instrument the application by itself.
The community is very helpful and you can chat via
[Gitter](http://gitter.im/open-telemetry/opentelemetry-js).

## Getting Started

I am using ExpressJS and OpenTelemetry has a plugin for it that you can load,
and it instruments the app by itself, same for MongoDB that is the packaged I am
using.

Those are the dependencies I installed in my applications, all of them are
provided by the repository I linked above:

```
"@opentelemetry/api": "^0.5.0",
"@opentelemetry/exporter-jaeger": "^0.5.0",
"@opentelemetry/node": "^0.5.0",
"@opentelemetry/plugin-http": "^0.5.0",
"@opentelemetry/plugin-mongodb": "^0.5.0",
"@opentelemetry/tracing": "^0.5.0",
"@opentelemetry/plugin-express": "^0.5.0",
```

I created a `./tracer.js` file that initialize the tracer, I have added inline
documentation to explain the crucial part of it:

```js
'use strict';

const opentelemetry = require('@opentelemetry/api');
const { NodeTracerProvider } = require('@opentelemetry/node');
const { SimpleSpanProcessor } = require('@opentelemetry/tracing');
// I am using Jaeger as exporter
const { JaegerExporter } = require('@opentelemetry/exporter-jaeger');

// This is not mandatory, by default httptrace propagation is used
// but it is not well supported by the PHP ecosystem and I have
// a PHP service to instrument. I discovered B3 is supported
// form all the languages I where intrumenting
const { B3Propagator } = require('@opentelemetry/core');

module.exports = (serviceName, jaegerHost, logger) => {
  // A lot of those plugins are automatically loaded when you install them
  // So if you do not use express for example you do not have to enable all
  // those plugins manually. But Express is not auto enabled so I had to add them
  // all
  const provider = new NodeTracerProvider({
    plugins: {
      mongodb: {
        enabled: true,
        path: '@opentelemetry/plugin-mongodb',
      },
      http: {
        enabled: true,
        path: '@opentelemetry/plugin-http',
          // I didn't do it in my example but it is a good idea to ignore health
          // endpoint or others if you do not need to trace them.
          ignoreIncomingPaths: [
            '/',
            '/health'
          ]
      },
      express: {
        enabled: true,
        path: '@opentelemetry/plugin-express',
      },
    }
  });

  // Here is where I configured the exporter, setting the service name
  // and the jaeger host. The logger is helpful to track errors from the
  // exporter itself
  let exporter = new JaegerExporter({
    logger: logger,
    serviceName: serviceName,
    host: jaegerHost
  });

  provider.addSpanProcessor(new SimpleSpanProcessor(exporter));
  provider.register({
    propagator: new B3Propagator(),
  });
  // Set the global tracer so you can retrieve it from everywhere else in the
  // app
  return opentelemetry.trace.getTracer("discount");
};
```

You will be thinking, that's too easy! You are right, the nature of NodeJS
makes tracing very code agnostic. With this configuration you get a lot "for
free".

You get a bunch of spans for every http request that ExpressJS serves, plus a
span for every MongoDB query. All of them with useful information like the
status code, path, user agents, query statements and so on.

We have to include it in our `./server.js` the entrypoint for our nodejs
application:

```js
'use strict';

const url = process.env.DISCOUNT_MONGODB_URL || 'mongodb://discountdb:27017';
const jaegerHost = process.env.JAEGER_HOST || 'jaeger';

const logger = require('pino')()

// Import and initialize the tracer
const tracer = require('./tracer')('discount', jaegerHost, logger);

var express = require("express");
var app = express();

const MongoClient = require('mongodb').MongoClient;
const dbName = 'shopmany';
const client = new MongoClient(url, { useNewUrlParser: true });

const expressPino = require('express-pino-logger')({
  logger: logger.child({"service": "httpd"})
})
```

As I told you, that's it! With this code you have enough to make your NodeJS
application to show up in your trace.

The instrumented version of the application is available here
[github.com/gianarb/shopmany/tree/discount/opentelemetry](https://github.com/gianarb/shopmany/tree/discount/opentelemetry/discount)

## understand the project

I tend to checkout projects when in the process of learning how they work.
Documentation is useful but always incomplete for such a high moving projects.

I have to say that the scaffolding is clear even for an not fluent NodeJS
developer like me.

```
$ tree -L 1
.
├── benchmark
├── CHANGELOG.md
├── codecov.yml
├── CONTRIBUTING.md
├── doc
├── examples
├── getting-started
├── karma.base.js
├── karma.webpack.js
├── lerna.json
├── LICENSE
├── package.json
├── packages
├── README.md
├── RELEASING.md
├── scripts
├── tslint.base.js
└── webpack.node-polyfills.js
```

I would like to define it as a monorepo, and it uses
[lerne](https://github.com/lerna/lerna) to delivery multiple packages from the
same repository.

`examples` contains workable example of how to use the different `packages`.

```
$ tree -L 1 ./examples/
./examples/
├── basic-tracer-node
├── dns
├── express
├── grpc
├── grpc_dynamic_codegen
├── http
├── https
├── ioredis
├── metrics
├── mysql
├── opentracing-shim
├── postgres
├── prometheus
├── redis
└── tracer-web

$ tree -L 1 ./packages/
./packages/
├── opentelemetry-api
├── opentelemetry-base
├── opentelemetry-context-async-hooks
├── opentelemetry-context-base
├── opentelemetry-context-zone
├── opentelemetry-context-zone-peer-dep
├── opentelemetry-core
├── opentelemetry-exporter-collector
├── opentelemetry-exporter-jaeger
├── opentelemetry-exporter-prometheus
├── opentelemetry-exporter-zipkin
├── opentelemetry-metrics
├── opentelemetry-node
├── opentelemetry-plugin-dns
├── opentelemetry-plugin-document-load
├── opentelemetry-plugin-express
├── opentelemetry-plugin-grpc
├── opentelemetry-plugin-http
├── opentelemetry-plugin-https
├── opentelemetry-plugin-ioredis
├── opentelemetry-plugin-mongodb
├── opentelemetry-plugin-mysql
├── opentelemetry-plugin-postgres
├── opentelemetry-plugin-redis
├── opentelemetry-plugin-user-interaction
├── opentelemetry-plugin-xml-http-request
├── opentelemetry-propagator-jaeger
├── opentelemetry-resources
├── opentelemetry-shim-opentracing
├── opentelemetry-test-utils
├── opentelemetry-tracing
├── opentelemetry-web
└── tsconfig.base.json
```

The suffix of the package helps you to figure out what they are:

* `opentelemetry-plugin-*` usually contains the code that instrument a specific
  library, you can see here `express`, `http`, `https`, `dns`. Some plugins are
  loaded by the `NodeTracerProvider` by default. Other has to be specified. You
  can to relay on the code or read the documentation to figure it out. For
  example `http` is loaded by default but if you need `express` you have to load
  them up by yourself, figuring out the right dependencies. At least or now.
* `opentelemetry-exporter-*` contains various exporters for now Jaeger,
  Prometheus, Zipkin the and otel-collector.

Anyway, what I am trying to say is that it is very intuitive and looking here it
is clear what you can get from this project.

## Plugin

NodeJS sounds very easy to instrument and on the right path to get automatic
instrumentation right,  because you can listen from the outside to function
call, you do not need to specifically change your code where you do a request or
where you get one, you can add tracing in a centralized location. That's how the
provided plugins work.

[Shimmer](https://github.com/othiym23/shimmer) is the library that simplify the
trick. I recently had a chat with [Walter](https://twitter.com/walterdalmut)
because I know he works in NodeJS and during the experiments otel were easy
enough to fit his use case. He is currently trying it and as he discovered that
[mongoose](https://github.com/Automattic/mongoose), the ORM library he uses does
not use the officially provided [mongodb
driver](https://mongodb.github.io/node-mongodb-native/) so the
otel-plugin-mongodb where not magically tracing his requests to mongodb, sadly.
But he is currently writing a [plugin for
that](https://github.com/wdalmut/opentelemetry-plugin-mongoose), so it won't be
a problem pretty soon.
