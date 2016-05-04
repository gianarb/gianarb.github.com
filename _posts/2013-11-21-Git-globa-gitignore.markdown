---
layout: post
title:  "Git global gitignore"
date:   2013-11-21 12:38:27
img: /img/git.png
categories: [post]
tags: git, vcs, versioning
summary: Git, manage global gitignote files or direcories
priority: 1
---

 ``` .gitignore  ``` helps me manage my commits by setting which files or directory don't end in my repository. I know two good practices if you work for example on an open source project:

* You don't commit your IDE configurations
* Not use .gitignore file for exclude IDE configuration, because this is personal problem. There are differents IDE, if all devs exclude this files on a repository level the lists is very long.

I follow this practices for all my projects, if you are Mac user you have a DS_STORE files, there is a method for exclude this file of default.

 ``` ~./.gitconfig ``` is your configuration file, every user has it. If you execute this command
{% highlight bash %}
$. git config --global core.excludesfile ~/.gitignore_global
{% endhighlight %}
into this file it write thiese lines
{% highlight bash %}
[core]
	excludesfile = /Users/gianarb/.gitignore_global
{% endhighlight %}

 ``` /Users/gianarb/.gitignore_global ``` is my global gitignore file!
{% highlight bash %}
# IDE #
#######
.idea

# COMPOSER #
############
composer.phar

# OS generated files #
######################
.DS_Store
.DS_Store?
._*
.Spotlight-V100
.Trashes
ehthumbs.db
Thumbs.db
{% endhighlight %}

