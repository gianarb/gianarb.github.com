---
img: /img/gianarb.png
layout: post
title:  "When do you need a Site Reliability Engineer?"
date:   2019-07-05 08:08:27
categories: [post]
tags: [sre, site reliability engineer]
summary: "I read everyday more job description looking for SRE. In the meantime
I hear and live the frustration about who does not understand what SRE means and
hires somebody that won't fit."
changefreq: daily
---
I have started to work as a Site Reliability Engineer more than two years ago as
first hired SRE at InfluxData.  I survived and learned to all the eras that
every company that onboard a new position live:

1. Lack of knowledge about what the job role means
2. Adjustment
3. Growth
4. Re-adjustment
5. Repeat

You are an SRE not because you care about reliability, everybody cares about
reliability but because the system is too complex to be driven by a person that
also does other things.

There are no differences with any other "first hired" in a company. Even the
first project manager gets hired when the person who was doing that job can't
make it anymore because the company needs somebody 100% focused on the product.

The Site Reliability Engineer as a role should improve **service** reliability.
Visibility, observability, logging, scalability, instrumentation are all areas
when it should step to serve better tooling to troubleshoot, identify issues.
Because as we all know, even not that complex distributed system are difficult
to debug, this complexity is caused by what it is called partial failure. The
idea that a distributed system will never fail drastically alltogheter, but it
is continuously in a condition of failure mitigated by re-try policies and or
redundancy.

The ability to acknowledge a problem before it will get reported by a customer
improves reliability.

It is not a responsibility for the Site Reliability Engineer to fix the actual
bug in the service, it can. For all those reasons the SRE knows how to code, and
it should modify the application, and it needs to be close to the team that
builds the service, just as every heterogeneous group has who takes care about
the design, UI, deploy, management.

## are they the unique people on-call?
Obviously no. It's hard to reach a scale where you can manage a sustainable
rotation only with SREs, and every developer is responsible for the code it
ships. If you managed to have a rotation for every service with different
people, all of the teammates should be on-call.

The SREs other than being part of the rotation is the person responsible for the
MTTR (mean time to repair) and the number of false positive.  The Site
Reliability Engineer needs to be able to make the MTTR as short as possible, and
the number of false positive as low as it can. They should improve how the
service is monitored, instrumented, and easy to debug.

## do I need an SRE in every service team?

It is hard to quantify a number, but the SREs needs to have a structure that
gives them time to hang out together and to see each other as a unique team as
well to share knowledge and to avoid the use of too many technologies across the
company. Even more, if the company is not at a gigantic scale in term of the
number of people.  The amount of SREs per team depends on now crucial, and
complex reliability for the service is, how big the service team is. You can
share SREs with organizations and services if they are not too big or too
complicated or if the unit itself has excellent reliability skills embedded in.

## What SRE is not

SRE  does not replace your ops team; it is not a person with DevOps skills that
knows containers and Kubernetes. It knows cloud, containers, and kubernetes
because it is a pretty new "unicorn" role.

It is a side effect of being a coder that loves to see its code running smoothly
under real load.


