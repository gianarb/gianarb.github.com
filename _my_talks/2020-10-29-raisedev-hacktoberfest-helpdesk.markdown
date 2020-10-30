---
layout: talk
title: 'Hacktoberfest Helpdesk'
description: "Joe and the raise.dev community organized a set of interviews
during hacktoberfest to help developers to contribute to opensource. I had the
opportunity to share my experience as developer. How I started my journey and
why open source matters to me."
date: 2020-10-29T11:00:00Z
eventName: Raise.dev
eventLink: https://raise.dev
city: "Twitchland"

---

<div class="embed-responsive embed-responsive-16by9 mb-3">
    <iframe
        class="embed-responsive-item"
        src="https://player.twitch.tv/?video=v785377719&parent=gianarb.it&autoplay=false"
        frameborder="0"
        scrolling="no"
        allowfullscreen="true">
    </iframe>
</div>

[Joe](https://twitter.com/jna_sh) is a professional host. I didn't know about
the raise.dev's mission until his introduction but I think it is a very cool
one! I am glad I had the opportunity to be part of it.

The [raise.dev](https://raise.dev)'s goal is to help developers to find their
direction in the tech industry. I think it is much needed. The career letter is
getting crazy every day more. New job titles come and go daily, and it is hard
to figure where you want to go.

I didn't have the opportunity to share during the interview, but I think it is
essential to mention the advantage you get in life when you know what you DO NOT
like. I am far away from that. I am starting to feel about it, but I am not that
close to saying: "I don't like YAML." As soon as you know what you dislike, it
is time to move on and leave it out of your life very quickly. In tech, it can
be a technology or a way of thinking.

Anyway, I have to admit I find myself way more comfortable than what I thought
sharing how I got to development. I presume there is nothing unusual in there,
but Joe highlighted a bunch of points that I hope will help more people to step
up, learn in public, and contribute to opensource. Because I think it is a great
way to make friends, work, and grow up.

As you will hear in the interview, I want to highlight what I think we can call:
"my way to familiarize with a repository," obviously it works with private and
opensource one. I use the same techniques when I start a new job.

I think this list is project independent a well, and I am an actionable person.
So I like to make my hands dirty with the project as soon as possible. If you
are not like me, you probably need another list!

### Clone the repository

Nat Friedman, GitHub's CEO, shared during a GitHub event something that sounded
to me like: "to start contributing to opensource you have to clone the
repository locally" it looked obvious to me when I heard it, but now, it means
everything. As soon as you clone code to your local environment, it becomes like
every other code you write every day. It becomes your code. As you do with your
code, now it is time to run the code, find bugs, compile, and see if it works
better than before.

### Have a look at the README

Usually, in opensource, this is an excellent way to start familiarizing with the
developer who wrote it; very often, it is the creator. Please don't take it too
seriously; it usually represents how the project should look like in theory, or
let's say, the best scenario. It is the front door. It tends to be clean and
welcoming.

In a close source environment, it is hard to get a good readme for my
experience.

### CI/CD configuration does not lie.

When you have the code locally, I want to compile it. From my experience,
open-source tools have CI/CD. Very often, they have Makefile, but it tends to
become unreadable pretty quickly. So if you can't find what you are looking for
in there, look for GitHub actions workflows description, Travis CI files, Drone,
Jenkinsfile or similar.

Based on the maturity of the project, the CI/CD pipeline gets executed way more
often, and it tends to be very noisy when it fails compared with commands listed
in a README, so they are more likely to work. In there, you will find the build
command!

### Where is the entry point?

When you know the language of the codebase, this task is more accessible. You
should look for the program's entry point. Every application has one. It is more
challenging for libraries or frameworks, but every project has at least one
entry point (if it is a mono repo, for example).

When it comes to a Golang application, you know that somewhere there is the main
function. It is often in a file called `main.go`, and this file is in the root of
the project or inside the cmd package. You will have to find the right pattern
for your language.

## Conclusion

That's it! I hope you will try to use the list I just shared and watch
raise.dev, it is a cool project. I want to thanks
[Rain](https://twitter.com/rainleander) for the opportunity as well!
