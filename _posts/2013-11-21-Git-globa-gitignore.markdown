---
layout: post
title:  "Git global gitignore"
date:   2013-11-21 02:38:27
categories: versioning
tags: git, vcs, versioning
summary: Git, manage global gitignote files or direcories
priority: 1
---

 ``` .gitignore  ``` help me to manage my commit. Becouse I can set wicth files or directory doesn't add in my repository, but I know two good practices if you work for example into open source project:

* You don't commit your IDE configurations
* Not use .gitignore file for exclude IDE configuration, becouse this is personal problem

I follow this practices for all my projects, if you are Mac user you have a DS_STORE files, there is a method for exclude this file of default.

 ``` ~./.gitconfig ``` is yout configuration file, eatch users have her, if you execute this command
{% highlight bash %}
$. git config --global core.excludesfile ~/.gitignore_global
{% endhighlight %}
into file it write this lines
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
