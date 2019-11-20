---
img: /img/gianarb.png
layout: post
title: "o11y.guru introduction and first set of iterations"
date: 2019-11-09 08:08:27
categories: [post]
tags: [o11y.guru, introduction, side project, twitter]
summary: "Part of the o11y.guru series this post is an introduction for this
side project and it describes the first architecture designed for the website."
changefreq: daily
---

{% include o11y-series-intro.markdown %}

![](/img/o11y-guru-series/index.png){:class="img-fluid"}

olly.guru is a website that I wrote in Go. It lists a group of people active on
twitter that I like to follow around monitoring, reliability, observability. It
allows you to follow them all in once as well.

A few years ago when I was developing almost in PHP somebody from the community,
I don't remember who was, I am getting old, did a similar website and I thought
it was a great idea.

Since then, the project was in the back, in my mind. I am a lazy person when it
comes to writing code. I think there is enough useless code around, and I don't
want to incentive that practice. That's why I tend to write as less code I can.

There are a bunch of reasons why I changed my mind, and I started to do it:

Better mood and I had to try Honeycomb.io, but I never had the right
opportunity. I didn't want to try it with a demo running on my laptop.  I
have a couple of new friends from CherryServer that supports my crazy ideas, and
I was looking for a reason to glue a bunch of reliable infrastructure as code that I
actually like. Even if it is usually a task I hate, mainly because there is no code
involved.

As Docker Captain and CNCF Ambassador, I have the feeling than a project like
that can be re-used.

I made the mistake that everyone does; I started to think about cool technologies
and not the problem I was going to solve or the project I was going to write.

I made a react application, just the foldering, and I quickly realized that I do not
know how to React, and I was wasting my time. But for this project, I got lucky
enough to keep going. I started to think about the problem again, and I decided
to make it as simple as possible. In practice almost everything gets generated
by an html template and a piloted by a list of names in a `txt` file. Very easy!

```
.
├── cmd
│   ├── generate
│   └── www
├── Dockerfile
├── go.mod
├── go.sum
├── index.tmpl
├── Makefile
├── people
│   └── people.go
├── people.txt
├── style
│   ├── css
│   ├── fonts
│   ├── img
│   ├── index.html
│   ├── js
│   ├── node_modules
│   ├── package.json
│   ├── package-lock.json
│   └── scss
├── vendor
└── www
```

You can see the shape of the project; it has a minimal amount of technologies
involved: Go, HTML, Bootstrap 4, and a bit of Javascript.  I started from the
`style` directory. I use it for prototyping the HTML, CSS part. I am far away to
be good with colors and CSS, and I am cool with that, we do not like each other.
So I do all my tests there, and when I am ready, I port the `style/index.html`
into `index.tmpl.` I started from a Bootstrap 4 layout already done as you can
see. It is in their documentation.

![](/img/o11y-guru-series/sheldon.jpeg)

`index.tmpl` is the template I use to render the actual homepage.  `www` is the
target destination for all the static files and the generated index page. I use
Make to copy files from `style` into `www,` and I wrote a CLI that generates the
HTML and it populates it with Twitter informations. It is inside `./cmd/generate.`

`./people.txt` is the list of twitter gurus. It is just a list:

```
gianarb
rakyll
```

The `cmd/generate` reads that file and, it gets the information it needs from
the Twitter API like user bio, avatar and it renders the `./index.tmpl`
into the actual index inside the `www` folder.

`./cmd/www` is an HTTP server written Go that serves the content of the `www`
directory. Plus it uses:

```
github.com/dghubble/go-twitter
github.com/dghubble/gologin/v2
```

To manage the Twitter authentication flow.

I am sure you are wondering, "is he gonna open-source that!?". I am. Not now.
The project needs to refactoring and some code needs to get stronger around
instrumentation and logging. As you see in the introduction, I am using this
experience as a use case to write down a bunch of practices I like or that I
would like to investigate.
So stay tuned! It will be available very soon.

## tldr lesson learned

Some of lessons I learned comes from how I am, but hey, this is my blog, I can
do whatever I like!
It is refreshing to start a project, but **it is way cooler to have something to
show**. So be careful when you start it, get it right so you won't get tired.

**Set clear goals**, and see point 1, they need to be very easy to
achieve, at least at the beginning.

**Do not type on your terminal, but write bash scripts.** I started to do this
months ago at work. Bash scripts are way better than random commands in a
terminal because you can move them around composing way more powerful workflows.
You won't lose them. That's how I built my Makefile, just from the terminal
history or from the shortcut I made along the way.
Often a well-done **dotenv file is enough to manage everything you need**.

## let's get to some code

I told you about bash scripts and Makefile, I will write a post about automation
for a small project, but this is part of my Makefile:

```
style/build:
    cd ./style && npm install
    cp -r ./style/node_modules/jquery/dist/jquery.js ./style/js/jquery.js
    cp ./style/node_modules/@fortawesome/fontawesome-free/js/all.js ./style/js/
    cp -r ./style/node_modules/@fortawesome/fontawesome-free/webfonts ./style/fonts
    cd ./style && npm run scss

style/start: style/build
    cd ./style && npm start

style/compile: style/build
    rm -rf ./www
    mkdir ./www
    cp -r ./style/img ./www
    cp -r ./style/fonts ./www
    cp -r ./style/css ./www
    cp -r ./style/js ./www
```

People can do the same with npm and, a hundred node packages, I like to keep
things simple at this point to avoid unnecessary blockers that will get my
tired. This is how I manage the `style` directory and how I build the `www` one.

> With blockers I mean: googling around for things that should be easy.

```go
flag.StringVar(&flags.consumerKey, "consumer-key", "", "Twitter Consumer Key")
flag.StringVar(&flags.consumerSecret, "consumer-secret", "", "Twitter Consumer Secret")
flag.StringVar(&flags.accessToken, "access-token", "", "Twitter access key")
flag.StringVar(&flags.accessSecret, "access-secret", "", "Twitter access secret")
flag.StringVar(&flags.guruFile, "guru-file", "", "File that contains the guru's name")
flag.StringVar(&flags.indexTemplate, "index-template", "", "File that contains the guru's name")
flag.Parse()
flagutil.SetFlagsFromEnv(flag.CommandLine, "TWITTER")

config := oauth1.NewConfig(flags.consumerKey, flags.consumerSecret)
token := oauth1.NewToken(flags.accessToken, flags.accessSecret)
httpClient := config.Client(oauth1.NoContext, token)

// Twitter client
client := twitter.NewClient(httpClient)

// Verify Credentials
verifyParams := &twitter.AccountVerifyParams{
    SkipStatus:   twitter.Bool(true),
    IncludeEmail: twitter.Bool(true),
}
_, _, err := client.Accounts.VerifyCredentials(verifyParams)
if err != nil {
    println(err.Error())
    os.Exit(1)
}

gurus := []*twitter.User{}

lines, err := people.ReadLineByLine(flags.guruFile)
if err != nil {
    println(err.Error())
    os.Exit(1)
}
for _, eachline := range lines {
    user, _, err := client.Users.Show(&twitter.UserShowParams{
        ScreenName: eachline,
    })
```

The generate command is straightforward, I get over the `people.txt` file line
by line, and for every record, I get information about the user. When I have the
slice of gurus populated I render the template:

```go
t, err := template.ParseFiles(flags.indexTemplate)
if err != nil {
    panic(err)
}
err = t.Execute(os.Stdout, Render{
    Gurus: gurus,
})
```
I decided to print the HTML into the stdout because it is way easier to use `>`
other than accepting another parameter to specify the target output.

LDD: laziness driven development.

The `cmd/www` uses the same people.txt file to know who to follow when the user
presses the `Follow` button and authorize the twitter application:

```go
for _, eachline := range lines {
    if strings.EqualFold(eachline, me.ScreenName) {
        continue
    }
    time.Sleep(5 * time.Second)
    err = newFriendship(ctx, twitterClient, eachline)
    if err != nil {
        logger.Warn(err.Error(), zap.String("follower_screenname", eachline), zap.Error(err))
    }
}
```

## The project in the project

This series of posts I am writing is a side project in the side project. As I
wrote earlier I like to share what and why I do things. I hope to keep having
practical experiences to write down.

The high level expectation I set are:

1. Have fun
2. Create a good network of followers on twitter that likes to speak about
   observability
3. Learning how Honeycomb works and why everybody says that it sounds like magic
4. Writing down something about code instrumentation, infrastructure as a code and
   automation
5. Exercise my experience as a decision maker driven by simplicity and
   efficiency.
6. I hope to work with a couple of friends from Docker, HashiCorp,
   CherryServier, InfluxData, HoneyComb to help me out with secret management,
   monitoring, terraform and, automation in order to build the coolest project
   ever. You will get an email from me (or reach out if you have suggestions).

## That's it

I am sure it gets out clearly from this article the friction between the
excitement about having an idea and the effort to make it real, even when it is
as simple as a single html page.. I struggle with
that all the time, and the laziness usually wins. Will this time be different?!
Well, I have a domain that is not a blank page. I think it is a good starting
point.

Time matter and have fun!

![](/img/o11y-guru-series/sleep.jpg){:class="img-fluid"}
