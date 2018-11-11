---
layout: post
title:  "Server time vs Response time"
date:   2018-05-18 10:38:27
img: /img/pastrami-sf.jpg
categories: [post]
tags: [monitoring, devops, observability, san francisco, metrics, capacity,
cloud]
summary: "How do you dimension an infrastructure? How can you calculate
container limits or how many nodes your application requires to support a
specific load? Response time and server time are two key measurements to monitor
saturation."
priority: 1
---
If you find yourself in San Francisco walking nearby Market Street, you should
consider stopping at the Jewish Museum. There is a charming Pastrami place just
next to it. It is a sandwich place with good lemonade. It only takes 3-4 minutes
to get your meal, and from there it takes no more than 15 minutes walking to be
in front of the Ocean. Very nice!
Now, let's consider this other scenario.
It is lunchtime, and you are starving. You rush outside your office, and you run
to the Pastrami place close to the Jewish Museum. After 35 minutes of wait, you
get your sandwich and start eating it asking yourself: why it took so long this
time? Shall I probably have walked to the next place to get a faster meal?

Something similar can happen to your Services as well! And that's precisely the
phenomena in computer science we try to capture using the concepts of server
time and response time.
Server time aims to measures how much a server takes to run a specific action.
Let's say consider an example operation the generation of a monthly report: it
usually takes 2ms, but if a lot of customers require the same kind of report at
the same time and your system saturates? This situation might very quickly end
up in having a subset of them getting the report in more than 1 minute or
actually in the timeout of the operation. The time it takes for a customer to
get his report is what is typically called response time.

## How can we measure these metrics?
The answer to this question is not easy: it depends on your architecture and
system. The starting point is instrumenting your application to determine how
much time it gets to produce the report. Stress testing is the other important
aspect: generating some load on your application and sampling the average
response time will let you estimate the application's service time. Notice that
to make this measurement the app should NOT soak during this test!

If you control all the chain (from the HTTP app that sends the request to the
server), you can trace the request and simulate the same behavior of your
customers. If you can't do this, you can consider using the frontend edge,
probably a load balancer.

> I would rather have questions that can't be answered than answers that can't
> be questioned. Richard Feynman

## Why does it matter
How many nodes do I need to deploy to accommodate x number of requests per
second? When should I consider scaling out my application? How does scale-out
affect the customer experience?
This is precisely why server time and response time matters! Having an average
response time close to the defined service time is a signal of proper
utilization and health of an application because it indicates that the response
latency is under control and it is far from saturation. Bringing to the limit
these two signals, in addition, is a key metric to estimate the correct sizing
of the applications instances and infrastructure.

<img alt="Market Street San Francisco, Pastrami Resturant Jewish Meseum" src="/img/pastrami-sf.jpg" class="img-fluid" />

Btw the Pastrami place exists! You should try it! I will be in SF in 2 weeks. So
let me know about other places [@gianarb](https://twitter.com/gianarb).
Picture from GMaps. I will take a better one!
