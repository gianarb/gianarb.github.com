---
img: /img/gianarb.png
layout: post
title:  "Hack your Google Calendar with gcalcli"
date:   2019-08-26 08:08:27
categories: [post]
tags: [oss, gcalcli, calendar, google, linux]
summary: "Everybody uses google calendar in a way or another and if you are a
Linux with a light desktop manager such as i3 you lack on
some commodities like reminders and notifications for your events. I find gcalcli a
very good solution for my pain."
changefreq: daily
---
I am pretty bad with meetings. I forget about them for a lot of different
reasons, sometime I do not show up even if few minutes earlier my mind briefly
remembered it.

Meetings are not my daily job, and I do not have them with a lot of different
people: IPM with my team, one-to-one with my manager, various stand up. I can
remember the recurrent one pretty well but it is still an annoying a useless
exercise.

When they are not recurrent they are usually out of my small circle of friends
and it gets even worst because I do not like to be late or to miss it! I swear I
am not like that in real life! I am on time and I prefer to be there earlier.

Anyway! Ryan Betts VP Engineer at InfluxData shared a very nice CLI tools called
[gcalcli](https://github.com/insanum/gcalcli). I love CLI tools as much as I
love API! Probably a bit more because they are the perfect glue between server
side and the best UX ever (also known as **my terminal**).

![A good gin tonic is great as close as my terminal](/img/gintonic.jpg)

**gcalcli** is a lovely CLI tool that uses the Google Calendar API to help you
to manage your Google Calendar.

You can do a lot of things: list, search, edit, add events and even more.
The [authentication is well
documented](https://github.com/insanum/gcalcli#login-information) you need to
create a project on Google Development Platform with Calendar API access. After
that you get your credential and you follow the link I just posted! Super easy.

When you are logged I wrote this system unit and a timer in order to check every
10 minutes if there are upcoming events:

```
[Service]
SyslogIdentifier=gcalcli-notification
ExecStart=/usr/bin/gcalcli remind

[Install]
WantedBy=multi-user.target
```

```
[Unit]
Description="Send notification for every meetings set for xxxxx@gmail.com"

[Timer]
OnBootSec=0min
OnCalendar=*:0/10

[Install]
WantedBy=timers.target
```

The timer runs every 10 minutes this command `/usr/bin/gcalcli remind`.
`remind` uses `notify-send` to show a lovely notification.

I set it up for my working calendar and let me tell you it works great!
For that reason I was looking for a way to support multiple Google account,
because I would like to use it for my personal Google Calendar as well.

There is a global flag for `gcalcli` called `--config-folder`, by default it set
te none it creates a config file with credentials and preferences in your home
directory.  If you run `gcalcli` with that parameter set with a different
location:

```bash
$ gcalcli --config-folder ~/.gcalclirc-anotheraccount list
```

The CLI won't find the configuration file and it will proceed with a brand-new
authentication and it will create a new file located where specified. Sweet! I
did that trick in order to have the second Google Account configured and I have
created a new unit and timer with the right flags and now I get notification
from everywhere! So far so good!

Ryan allowed me to share a script he hacked called `next`, I have it in my
`bashrc`

```bash
next() {
    datetime=$(date "+%Y-%m-%dT%H:%M")
    whatwhere=$(gcalcli --calendar name-your-calendar agenda --tsv --details location $datetime 8pm | head -n 1 | awk 'BEGIN {FS = "\t+"} ; {print $5 " " $6}')

    re="([[:digit:]]+)"
    if [[ $whatwhere =~ $re ]]; then
       room="zoommtg://zoom.us/join?confno=${BASH_REMATCH[1]}"
    fi

    echo "What: '$whatwhere'"
    echo "xdg-open $room"
    echo "xdg-open $room" | clipc
}
```

I use Linux, he uses MacOS, so I changed the script a bit.

`xdg-open` to make it to work with `X`, `next` gets the next closer meeting you have in one
particular calendar (`name-your-calendar` in my case) and it stores on my
clipboard (via `clicp`) the command to join a zoom channel. It is super when you
are in a hurry, you will join zoom meetings in a second.

If you use `gcalcli` and you have other tricks let me know via twitter
[@gianarb](https://twitter.com/gianarb) because I would like to try them as well!
