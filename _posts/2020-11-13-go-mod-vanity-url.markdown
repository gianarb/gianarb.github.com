---
layout: post
title:  "Vanity URL for Go mod with zero infrastructure"
date:   2020-11-13 10:08:27
categories: [post]
tags: [golang]
summary: "A lot of Go modules today are hosted on GitHub. But you can setup
vanity URL using a custom domain. This is good to decouple your library from
GitHub. The best part is that it does not require any infrastructure, if you
don't want to."
heroimg: /img/jamie-street-vanity-unsplash.jpg
---
This post is about how I renamed a go module from
github.com/something/somethingelse to go.gianarb.it/somethingelse. It requires
zero infrastructure, just a static site that can run on GitHub.

I used one of the many projects I have that nobody cares about
[gianarb/go-irc](https://github.com/gianarb/go-irc).

### Why

You know, is one of those ideas you have in your mind for ages, but who cares?
At least for me that I don't have any cool open source project under my name.

I am not one of those people who suffer when realizing that it is not escale. If
you end up having a project that gets traction and is tight to github.com
because you didn't think about another way to go, you are stuck. And even if
GitHub today is cool, it won't stay cool forever.

Filippo Valsorda [@FiloSottile](https://twitter.com/FiloSottile) today
[tweeted](https://twitter.com/FiloSottile/status/1327240411266641920) about this
topic, and I looked at how he set up filippo.io/age to solve this little
dilemma.

### Goals

This is not about how to escape from GitHub but is about setting up a "vanity
URL" that won't lock you or your project to GitHub. It does not require any
infrastructure, just a domain that you can point to a GitHub Pages.

### Prerequisite

1.  Create a DNS record that points as CNAME to `<github-handle>.github.io`. I
    used go.gianarb.it
2.  Create a repository; it will be the home for your static site. Mine is
    [gianarb/go-libraries](https://github.com/gianarb/go-libraries)
3.  Set the repository up to be a [GitHub page](https://pages.github.com/) and
    enable HTTPS. You can enable it via the repository Settings; we will push
    HTML files to it directly, so I used the master branch as the GitHub page's
    source.

### Add your first library.

If your library is already using go mod, you have to change the module name to
the new one. In my case, from was github.com/gianarb/go-irc to
go.gianarb.it/irc. I just searched and replaced with my editor in all the
project. Renaming a module is a bc break; I am not sure how to avoid or mitigate
that; if you know, let me know!

You can push a new file to your static site repository; I called my irc:

```html
<html>
    <head>
        <meta name="go-import" content="go.gianarb.it/irc git https://github.com/gianarb/go-irc">
        <meta http-equiv="refresh" content="0;URL='https://github.com/gianarb/go-irc'">
    </head>
    <body>
        Redirecting you to the <a href="https://github.com/gianarb/go-irc">project page</a>...
    </body>
</html>
```

Replace your URL accordingly, but as soon as you push this file and GitHub will
release it to your page, you will be able to import: `go.gianarb.it/irc`.

### Conclusion

This methods works as a safeguard if you decide to move your code out from
GitHub. The static site can be deployed to Netlify, S3 or served by Nginx. It
does not need to stay on GitHub.

Same for your code, if you decide to move from GitHub to GitLab you can do it
transparently.
