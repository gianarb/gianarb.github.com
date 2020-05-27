---
layout: talk
title: 'Continuous Profiling Go applications'
date: 2020-05-19T21:00:00Z
slide: "https://speakerdeck.com/gianarb/gogetcommunity-continuous-profiling-go-application-f0f9b2e0-57ff-4e16-b5c9-077d28d31749"
embedSlide: "7edbe091473b4b648db26193fd35c5bc"
eventName: GoGetCommunity
eventLink: https://www.gogetcommunity.com
city: ""
links: {}

---

I use profiles to better describe post mortems, to enrich observability and
monitoring signals with concrete information from the binary itself. They are
the perfect bridge between ops and developers when somebody reaches out to me
asking why this application eats all that memory I can translate that to a
function that I can check out in my editor. I find myself looking for outages
that happened in the past because cloud providers and Kubernetes increased my
resiliency budget the application gets restarted when it reaches a certain
threshold and the system keeps running, but that leak is still a problem that as
to be fixed. Having profiles well organized and easy to retrieve is a valuable
source of information and you never know when you will need them. That's why
continuous profiling is important today more than ever. I use Profefe to collect
and store profiles from all my applications continuously. It is an open-source
project that exposes a friendly API and an interface to concrete storage of your
preference like badger, S3, Minio, and counting. I will describe to you how to
project works, how I use it with Kubernetes, and how I analyze the collected
profiles.
