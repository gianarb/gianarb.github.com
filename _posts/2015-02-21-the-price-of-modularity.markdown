---
layout: post
title:  "The price of modularity"
date:   2015-02-21 23:08:27
tags: zf2, code, develop, modularity, reusable
categories: [post]
summary: The modularity is not only a beautiful word, It has roles and a price.
priority: 0.6
changefreq: yearly
---
Today all frameworks are **modulable** but it isn't just a beautiful word, behind it there are a lot of concepts and ideas:

* The modularity helps you to reuse parts of code in different projects
* Every component is indipendent so you work on single part of code
* Every **component** solves a specific problem... it's a beautiful concept that helps you with maintainance!
* other stuffs..

As you can imagine there is a drawback, all this requires a big effort.
Ideally every component requires personal circle of release, repository, commits, pull requests, travis conf, documentation etc. etc.

Anyway several shorcuts are available. For instance, *git subtree* could help you in this war but the key is this: you need an agreement to win.

Zend Framwork Community choose another street, `Zend\Mvc` in this moment required:
{% highlight json %}
{
    "name": "zendframework/zend-mvc",
    "...": "...",
    "target-dir": "Zend/Mvc",
    "require": {
        "php": ">=5.3.23",
        "zendframework/zend-eventmanager": "self.version",
        "zendframework/zend-servicemanager": "self.version",
        "zendframework/zend-form": "self.version",
        "zendframework/zend-stdlib": "self.version"
    },
    "require-dev": {
        "zendframework/zend-authentication": "self.version",
        "zendframework/zend-console": "self.version",
        "zendframework/zend-di": "self.version",
        "zendframework/zend-filter": "self.version",
        "zendframework/zend-http": "self.version",
        "zendframework/zend-i18n": "self.version",
        "zendframework/zend-inputfilter": "self.version",
        "zendframework/zend-json": "self.version",
        "zendframework/zend-log": "self.version",
        "zendframework/zend-modulemanager": "self.version",
        "zendframework/zend-session": "self.version",
        "zendframework/zend-serializer": "self.version",
        "zendframework/zend-text": "self.version",
        "zendframework/zend-uri": "self.version",
        "zendframework/zend-validator": "self.version",
        "zendframework/zend-version": "self.version",
        "zendframework/zend-view": "self.version"
    },
    "suggest": {
        "zendframework/zend-authentication": "Zend\\Authentication component for Identity plugin",
        "zendframework/zend-config": "Zend\\Config component",
        "zendframework/zend-console": "Zend\\Console component",
        "zendframework/zend-di": "Zend\\Di component",
        "zendframework/zend-filter": "Zend\\Filter component",
        "...": "..."
    },
    "...": "..."
}
{% endhighlight %}

A few `require-dev` dependencies are used into the component to run some features, why? This force me to think *"Dependencies of this feature are included or not?"*!!
Composer was born to solve it! In my opinion the cost of the question is highest than download a few unused classes.
There are a lot of unused classes? Maybe too much?

Even if the right answer donsn't exist I think thant some indicators may help you to understand when is the moment to split the component:

* List of dependencies
* Complexity of component
* Features
* ..

No shortcuts.


