---
layout: post
title:  "Generale Jekyll sitemap without plugin"
date:   2013-08-09 09:38:27
img: /img/jekyll.png
categories: jekyll
tags: jekyll, sitemap, blog, seo
summary: Generate sitemap for Jekyll blog, without plugin
priority: 1
---

This blog is a static blog and uses GitHub pages, GitHub pages are generally deployed using Jekyll.

### How can you generate a sitemap without Jekyll plugin?
This [gist](https://gist.github.com/GianArb/6172377) answers your question.
I use some post values: changefreq, date and priority, if you don't set any specific values for them default values are used that are, 0.8 for priority and month for frequency.
In a single post you add this params for use correct params!
{% highlight php %}
---
layout: post
title:  "Why this blog?"
date:   2013-07-22 23:08:27
categories: me
tags: me, developer, presentation, gianarb
summary: Gianluca Arbezzano, developer, Italian, why open this blog?
changefreq: monthly
---
{% endhighlight %}

If you want to know more about the Sitemap Protocol read [this](http://www.sitemaps.org/protocol.html).

[Marco](https://github.com/MarcoDeBortoli) thanks for English! :)
