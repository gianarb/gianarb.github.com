---
layout: post
title:  "Asus universal dock station driver"
date:   2017-08-03 10:08:27
categories: [post]
tags: [setup]
summary: "Every developer loves to speak about its setup. I am here to share my
trouble with my new laptop. Asus Zenbook 3."
changefreq: yearly
---
Every developer loves to share things about it's setup. They also loves to make
it better and to spend time on it.

Lorenzo [(fntlnz)](https://twitter.com/fntlnz) is super on it! I am
not, plus I bought a Zenbook 3. Super slim, less than 1kg, I can use it to cut
ham probably but the unique USB-C is driving me crazy.

Probably more than the actual 40 degrees that I have in my home office now!
It is probably why I am writing this post btw.

When I bought this laptop 7 months ago the Universal Docker Station was not
available and I wasn't even able to install linux on this laptop.

Now I have an [Asus Universal Dock
station](https://www.asus.com/Laptops-Accessory/Universal-Dock/). I am feeling a
little bit better but to work it replace a normal charger, it means that without
a socket near you, I can not use a USB... Amazing experience.

I tried other adapter but I didn't find one good enough. Every one of them had
some input or output port unusable for some reason. Most of them because the
BIOS has a different watt limit and they can not charge the laptop. I never
received a response from ASUS about it. That's great.

Anyway I am writing this article just as note for myself about the driver that
Lorenzo discover to have the Asus Universal Dock Station's ethernet port
running.

[Realtek ethernet
driver](http://www.realtek.com/DOWNLOADS/downloadsView.aspx?Langid=1&PNid=13&PFid=5&Level=5&Conn=4&DownTypeID=3&GetDown=false).
It's super easy to install. Just compile it and it will work.
