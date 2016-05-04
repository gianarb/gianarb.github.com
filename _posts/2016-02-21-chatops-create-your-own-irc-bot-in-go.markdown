---
layout: post
title:  "ChatOps create your IRC bot in Go"
date:   2016-02-21 10:08:27
img: /img/go.png
categories: [post]
tags: devops
summary: "ChatOps is a strong topic it is growing day by day
because now with the IaaS are allowed a new way to manage your
infrastacture provide for you an API layer. You can implement it
to create your automation layer. A pretty bot is a good assistence"
priority: 0.6
changefreq: yearly
---
The infrastructure as a service (IaaS) opened new ways to manage your
infrastructure.
Use an API to create, destroy and update your virtual machine
is one of the biggest revolutions of our sector.

A lot of companies and a lot of DevOps started to create own assistence to
increase the automation or to check the status of them infrastructure, in top of
all GitHub provided a series of awesome blogpost and tools to describe this
approach that it has a name: ChatOps.

* [HuBot](https://hubot.github.com/) is a beautiful tools written in node.js to provide smart bot.
* [So, What is ChatOps? And How do I Get Started?](https://www.pagerduty.com/blog/what-is-chatops/) by PagerDuty
* [Say Hello to Hubot](https://github.com/blog/968-say-hello-to-hubot) by GitHub

<div class="row">
    <div class="col-md-12 text-center">
        <iframe width="560" height="315" src="https://www.youtube.com/embed/IhzxnY7FIvg" frameborder="0" allowfullscreen></iframe>
    </div>
</div>

IRC is an application layer protocol that facilitates communication. One of the
most famouse open IRC server is freenode all most important open source projects
use it to chat.

This concept is already applyed it because most projects are your personal bot,
for example Zend use Zend\Bot a good assistence written by DASPRiD.

The ChatOps is an assistence oriented to decrease the distance between your
infrastacture and your communication channels.

I wrote a low level library to communicate on IRC protocol, we can try to use it to
write our dummy bot.

{% highlight go %}
package main

import (
    "log"
    "fmt"
    "regexp"
    "bufio"
    "net/textproto"
    "github.com/gianarb/go-irc"
)

func main(){
    secretary := NewBot(
        "irc.freenode.net",
        "6667",
        "SybilBot",
        "SybilBot",
        "#channel-name",
        "",
    )
    conn, _ := secretary.Connect()
    defer conn.Close()

    reader := bufio.NewReader(bot.conn)
    tp := textproto.NewReader(reader)
    for {
        line, err := tp.ReadLine()
        if err != nil {
            log.Fatal("unable to connect to IRC server ", err)
        }

        isPing, _ := regexp.MatchString("PING", line)
        if isPing  == true {
            bot.Send("PONG");
        }

        fmt.Printf("%s\n", line)
    }
}
{% endhighlight %}

With this code you have a bot, in this case her name is SybilBot and at the
moment it suppot only the PING PONG flow, without this helth system your bot go
down after few time.

You can use the same log to add other actions

{% highlight go %}
yourAction, _ := regexp.MatchString("CheckSomething", line)
if yourAction  == true {
    // Do Something
}
{% endhighlight %}

[go-irc](https://github.com/gianarb/go-irc) allow you to communicate over IRC protocol, our but is very stupid I
like the idea! If you are working on this topic, in go or in other language
please ping me! I am very happy to know your bot!
