---
heroimg: /img/hero/action.jpg
layout: post
title:  "Serverless means extendibility"
date:   2019-01-22 08:08:27
categories: [post]
tags: [serverless, github action, serverless, api gateway lambda]
summary: "Looking at the GitHub Actions design and connecting the docs I think I
got why serverless is useful. It is a great mechanism to extend platform and
SaaS."
changefreq: daily
---
I wrote an article about a [GitHub Action](/blog/kubernetes-GitHub-action) I
recently created to deploy my code to kubernetes. Very nice.  Writing the action
and the post, I realized what serverless is all about.  I wrote it in the
incipit of the article, but I think this topic deserves its dedicated post.
Serverless is not yet for web applications. I know some of you will probably
disagree but this is my blog, and that's why I have one, to write whatever I
like!

![](/img/brave_dad.png)

I used Lambda and API Gateway to distribute two pdf I wrote about ["how to scale
Docker"](https://scaledocker.com), it looks to me way more complicated than a go
daemon. So I wrote them because I got the free tier and because I like to try
new things.  There are excellent applications written in that way for example
[acloud.guru](https://acloud.guru/) but I am probably not ready for that! My bad

Anyway, I know what I am ready for: We should use serverless to offer
extendibility for our as a service platform.

Good for us distributed system and hispters applications are all based on
events, Kafka and so on. Plus now we have
[runC](https://github.com/opencontainers/runc),
[buildkit](https://github.com/moby/buildkit) and a lot of the building blocks
useful to implement a solid serverless implementation.

It is not easy, at scale this is a complicated problem but we are in a better
situation now, and it is a massive improvement from a product perspective:

1. Using containers, we can offer total isolation, and we can take a very
   carefully and self defensive approach.
2. An API already provides extendibility but, you still need to have your server
   and to run your application by yourself to enjoy them. With a serverless
   approach, it will be much easy for the customer to implement their workflow.
3. You can ask your customer to share their implementation creating a vibrant
   and virtuous ecosystem.

You can use a subset of the event that you write in Kafka as a trigger for the
function, VaultDB to store secrets that will be injected inside the service and
so on.

![](/img/heart.jpg)

There is a lot more, but I am excited! Is somebody doing something like that? If
so, let me know [@gianarb](https://twitter.com/gianarb), I would like to chat!
