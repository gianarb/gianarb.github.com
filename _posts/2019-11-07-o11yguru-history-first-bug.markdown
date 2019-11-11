---
img: /img/gianarb.png
layout: post
title: "o11y.guru the history of the first bug"
date: 2019-11-07 08:08:27
categories: [post]
tags: [o11y.guru, bug, honeycomb, twitter]
summary: ""
changefreq: daily
---

{% include o11y-series-intro.markdown %}

## The history of the first bug

After the first deploy I used my twitter account
[@gianarb](https://twitter.com/dev_campy) and
[@devcampy](https://twitter.com/dev_campy) to try the application. I have also
asked a friend to try it out.

So far so good, the way I coded the following workflow is very basic, and
probably it will quickly reach its scalability limit. It is a loop with a
`time.Sleep(5 * time.Second)` break between each account to avoid the Twitter
rate limit.

```go

for _, guru := range gurus {
    time.Sleep(5 * time.Second)
    err = newFriendship(ctx, twitterClient, guru)
    if err != nil {
        logger.Warn(err.Error(), zap.Error(err))
    }
}
```

No retry or things like that for now. Very simple. I hope to iterate on it in
the future when it will start to not working well enough anymore.

It does not report any error if the Twitter API request to follow a person
fails, it just go to the next one. All three tests went well for what I was able
to say, all three accounts followed the gurus.

One of the first benefit about using HoneyComb is that out of the box they are
able to detect errors looking at the events you return and the graph is made by
them. Just clicking around to their UI I ended up with weirdness like this graph:

![Requests break down by HTTP Status](/img/o11y-guru-series/first-bug-http-status.png){:class="img-fluid"}

I noticed some `500` error page, and I do not like that. As you can see
there is an `Error` tab, built by Honeycomb again and this is what it showed to me:

![Span with an error](/img/o11y-guru-series/first-bug-span-with-error.png){:class="img-fluid"}

At this point it is clear to me where the problem is: "You can't follow
yourself". It sounds reasonable.

I changed the code and I added a simple `if` statement to skip the guru if it is
the person actually following all the other people.

```go
// me comes from above when I validate that the token behaves to a user.

if guru == me.ScreenName {
    continue
}
```

It does not sound trivial at all but when I tried the fix didn't work.

![](/img/o11y-guru-series/rambo.jpg){:class="img-fluid"}

I decided to face up the problem differently. *Spoiler alert: I didn't
write any unit test yet. Feel free to leave now.*

Looking at the trace I knew I had set for every **following request** the guru name
to follow and at the root span I had who required to follow the
gurus. In practice, I had in the root span `required_by=me.ScreenName`, and for every
guru its span with their name. The next image has those two
spans side by side:

* At the left the span `newFriendship` describes a single following action (a
  twitter create friendship api request). As you can see it has the `error="you
  can't follow yourself"` and the `follower_screenname=gianarb`.
* The one to the right is the root span, it has the `required_by=GianArb` field,
  it is the `me.ScreenName` variable.

![](/img/o11y-guru-series/first-bug-compare-spans.png){:class="img-fluid"}

Looking at this span the situation is clear, I was comparing `GianArb` the
`required_by` variable that you see in the right with `gianarb`, the
`follower_screenname` you see at the left span.

At the end of the story the check needs to be case-insensitive. And that's how
it is now:

```
if strings.EqualFold(guru, me.ScreenName) {
    continue
}
```

This is the history of the first but I randomly discovered and I had to fix
twice for `o11y.guru`.
