---
layout: post
title:  "Docker The fundamentals"
date:   2016-08-25 12:08:27
categories: [post]
img: /img/docker.png
tags: [docker, scaledocker]
summary: "Docker The fundamental is the second chapter of my book Scale Docker.
Drive your boat like a captain. I decided to share free the second chapter of
the book. It covers getting started with Docker. It's a good tutorial for
people that are no idea about how container means and how docker works."
changefreq: yearly
---
I am writing a book about Docker SwarmKit and how manage a production
environment for your containers.

The second chapter of the book is a Getting Started about Docker, it covers
basic concepts about what container means and it's a started point to
understand the concepts expressed into the book.

<h2>Drive your boat like a Captain.
<small>Docker in production</small></h2>

The book is work in progress but you can find more information into the site
<a href="http://scaledocker.com">scaledocker.com</a>.

To receive the first chapter free leave your email and if you like your twitter account:

<div class="row">
	<div class="col-md-6">
        <img src="/img/the-fundamentals.jpg" class="img-responsive">
    </div>
	<div class="col-md-4">
		<form id="get-chapter">
		  <div class="form-group">
			<label for="exampleInputEmail1">Email address *</label>
			<input type="email" class="form-control" required="required" id="email" placeholder="Email">
		  </div>
		  <div class="form-group">
			<label for="exampleInputPassword1">Twitter</label>
			<input type="title" class="form-control" id="twitter" placeholder="@gianarb" pattern="^@.*">
			<p class="help-block">The first letter needs to be a @</p>
		  </div>
          <p class="text-success get-chapter-thanks">Check your email! Thanks!</p>
          <p class="text-warning get-chapter-sorry"><span class="err-text"></span>.
          Please notify the error with a comment or with an email</p>
		  <button class="btn btn-default">Get your free copy</button>
		</form>
	</div>
</div>

<h2>Contents</h2>
1. Introduction
2. Install Docker on Ubuntu 16.04
3. Install Docker on Mac
4. Install Docker on Windows
5. Run your first HTTP application
6. Docker engine architect
7. Image and Registry
8. Docker Command Line Tool
9. Volumes and File Systems 20
10. Network and Links
11. Conclusion

Enjoy your reading and leave me a feedback about the chapter!

<script>
    (function() {
        $(".get-chapter-thanks").hide();
        $(".get-chapter-sorry").hide();
        var api = "https://1lkdtyxdx4.execute-api.eu-west-1.amazonaws.com/prod";
        $("#get-chapter button").click(function(eve) {
            eve.preventDefault()
            $(".get-chapter-thanks").hide();
            $(".get-chapter-sorry").hide();
            var requestChapter = $.ajax({
                "url": api+"/the-fundamentals",
                "type": 'post',
                "data": {
                    email: $("#email").val(),
                    twitter: $("#twitter").val()
                },
                "dataType": 'json',
                "contentType": "application/json"
            });
            requestChapter.done(function() {
                $(".get-chapter-thanks").show();
            });
            requestChapter.fail(function(data) {
                $('.err-text').html("["+data.responseJSON.code+"]"+ data.responseJSON.text);
                $(".get-chapter-sorry").show();
            });
        });
    })();
</script>
