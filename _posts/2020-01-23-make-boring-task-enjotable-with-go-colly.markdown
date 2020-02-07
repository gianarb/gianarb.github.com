---
img: /img/go.png
layout: post
title: "Make boring tasks enjoyable with go and colly"
date: 2020-01-23 09:08:27
categories: [post]
tags: [golang, colly, jekyll]
summary: "Recently I had the idea to update the conference page on my website with the end
goal to make it a bit more structured. Where structured means a bit more
reusable compared with the static HTML table I used to have. I mixed a bit of
hacky Go, colly for scraping and that's who I did it"
changefreq: daily
---
Recently I had the idea to update the [conference](/conferences.html) page on my website with the end
goal to make it a bit more structured. Where structured means a bit more
reusable compared with the static HTML table I used to have.

In the beginning, I decided to do an HTML table every year listing all the
conferences as a single row. It worked but I think at this point I can do
something even cooler with a single page for every conference talk with YouTube
and slides embedded, the abstract and few links to deep dive on the topic.

Jekyll has a cool feature called
[collections](https://jekyllrb.com/docs/collections/): “Collections are a great
way to group related content like members of a team or talks at a conference.” I
decided to do a “my_talks” collection.

I first added the right configuration in the `_config.yaml` and I added my first
conference in 2020, DevOps Pro in Vilnius (see you there!!).

```
collections:
  my_talks:
    output: true
```

I have created my first talk as a markdown file, just as I do for my posts:

```
---
title: Continuous Profiling Go Application Running in Kubernetes
date: 2020-03-24
slide:
embedSlide:
video:
embedVideo:
eventName: DevOps Pro Europe
eventLink: https://devopspro.lt/
city: Vilnius, Lithuania
---
Microservices and Kubernetes help our architecture to scale and to be
independent at the price of running many more applications. Golang provides a
powerful profiling tool called pprof, it is useful to collect information from a
running binary for future investigation. The problem is that you are not always
there to take a profile when needed, sometimes you do not even know when you
need to one, that's how a continuous profiling strategy helps. Profefe is an
open-source project that collects and organizes profiles. Gianluca wrote a
project called kube-profefe to integrate Kubernetes with Profefe. Kube-profefe
contains a kubectl plugin to capture locally or on profefe profiles from running
pods in Kubernetes. It also provides an operator to discover and continuously
profile applications running inside Pods.
```

As you can see I decided to set a bunch of variables that I will hope to re-use
where I will do the “single page” for each talk.

That’s it. All done: 2020 looks awesome and I added a for loop in the conference
page to print out the row as before:

```
<div class="row">
    <h3>{{ include.year }}</h3>
    <div class="col-md-12">
        <table class="table table-hover" id="{{ include.year}}">
          <thead>
            <tr>
              <th>Date</th>
              <th>Event</th>
              <th>Talk</th>
              <th>Slide</th>
            </tr>
          </thead>
          <tbody>
            {% for talk in include.talks %}
            <tr>
              <td>{{ talk.date | date_to_string }}</td>
              <td><a href="{{ talk.eventLink }}">{{ talk.eventName }} {{ talk.city }}</a></td>
              <td>{{ talk.title }}</td>
              <td>
                {% if talk.slide and talk.slide != ""  %}
                <a href="{{ talk.slide }}" target="_blank">Slide</a>
                {% endif %}
                {% if talk.video and talk.video != ""%}
                <a href="{{ talk.video }}" target="_blank">Video</a>
                {% endif %}
              </td>
            </tr>
            {% endfor %}
          </tbody>
      </table>
    </div>
</div>
```

In order to make everything a bit more reusable and organized this piece of code
is what Jekyll call [include](https://jekyllrb.com/docs/includes/). The way I
use it inside the conference page looks like:

```
{ assign talks2020 = site.my_talks | where:'date', "2020" }
{ include talks_per_year.html year="2020" talks=talks2020 }
```

Everything is working fine, and I am pretty happy but I have over 6 years of
talks to convert in this new format, it means over 50 conferences to convert one
by one in the new format made of files and YAML.

{:refdef: class="text-center"}
![https://media.giphy.com/media/KFz5cubdh5eskezQ6d/giphy.gif](https://media.giphy.com/media/KFz5cubdh5eskezQ6d/giphy.gif){:.img-fluid}
{: refdef}

## Scraping is my superpower

I am not a fan of scraping things around and I never did that before, but hey!
This solution looks less boring that me doing it manually. I deep dive looking
for scraping libraries in new languages (yes you always have to learn new
languages when doing a new side project), but at the end I discovered
[colly](https://github.com/gocolly/colly): “Elegant Scraper and Crawler
Framework for Golang”. I decided to be elegant and effective.

## A bit about Colly

I have to say that it took me less than 2 hours to hack a script in Go using
Colly that converted all my tables year by year from HTML to files with the
format you saw above. I also added some sweet sugar like:

* Be able to convert YouTube links when detected to their embeddable version
* I converted and standardized the end/start date for the talks because it
  changed year by year (I am lazy and unconsistent! Don’t tell anybody)

It was soo easy that I didn’t write any test… yep, that’s it. The file name is a
bit weird but at the end it works, so who cares!


```
$ tree ./_my_talks/
./_my_talks/
├── 2013-09-12-what-is-vagrant.markdown
├── 2014-02-c'è-un-modulo-zf2-per-tutto!---there-is-a-module-for-all.markdown
├── 2014-03-zend-queue.markdown
├── 2014-05-getting-start-chromecast-developer.markdown
├── 2014-05-vagrant,-riutilizzo-dell'infrastruttura---vagrant,-reuse-architecture.markdown
├── 2014-10-sviluppo-di-api-rest-con-zf2-&-mongodb.markdown
├── 2014-10-time-series-database,php-&-influx-db.markdown
├── 2015-01-angularjs-advanced-startup.markdown
├── 2015-06-delorean-made-in-home---reaspberry,-gobot-and-mqtt.markdown
├── 2015-07-joomla-and-scalability-with-aws-beanstalk.markdown
├── 2015-09-penny-php-middleware-framework.markdown
├── 2015-10-angularjs-in-cloud.markdown
├── 2015-10-doctrine-orm-cache-layer---it-is-not-a-boomerang.markdown
├── 2015-11-wordpress-and-scalability-with-docker.markdown
├── 2016-02-slimmer---poc-born-after-a-revolt-instant-vs-jenkins.markdown
├── 2016-03-a-zf-story:-parallel-made-easy.markdown
├── 2016-04-listen-your-infrastructure-and-please-sleep.markdown
├── 2016-05-continuous-delivery-with-jenkins-in-the-real-world.markdown
├── 2016-06-aws-under-the-hood.markdown
├── 2016-06-listen-your-infrastructure-and-please-sleep.markdown
├── 2016-06-parallel-made-easy.markdown
├── 2016-07-docker-1.12-and-orchestration-built-in.markdown
```

{:refdef: class="text-center"}
![https://i.kym-cdn.com/photos/images/newsfeed/000/345/534/4a2.jpg](https://i.kym-cdn.com/photos/images/newsfeed/000/345/534/4a2.jpg){:.img-fluid}
{: refdef}

Anyway, let’s get to some snippets!

```
type Talk struct {
	Title      string            `yaml:"title"`
	Date       time.Time         `yaml:"date"`
	Slide      string            `yaml:"slide"`
	EmbedSlide string            `yaml:"embedSlide"`
	Video      string            `yaml:"video"`
	EmbedVideo string            `yaml:"embedVideo"`
	EventName  string            `yaml:"eventName"`
	EventLink  string            `yaml:"eventLink"`
	City       string            `yaml:"city"`
	Links      map[string]string `yaml:"links"`
}

var dateLayout = "_2 Jan 2006"
var year = "2020"
var outputDir = "/tmp"

var errorsToCheck = map[string]string{}
```

Those are the variables and struct I set. The Talk represent every single talk,
the dataLayout converts the way the end/start date is written into a time.Time
object. `year` is a parameter that tells which table to scrape, `outputDir`
tells where to place the files. Those 3 variables can be changed with cli flags:

```
flag.StringVar(&year, "year", "2020", "The year used to identify the table to parse")
flag.StringVar(&dateLayout, "date-layout", "_2 Jan 2006", "The golang format layour to parse the event date column")
flag.StringVar(&outputDir, "output-dir", "/tmp", "Where to place the generated files")

flag.Parse()
```

`errorsToCheck` is an easy way to collect all the errors for every run. I
printed them in a file, if the errors were easy to fix with a code change I did
that, if they were easier to change modifying the current conference page I did
that.

```
// Instantiate default collector
c := colly.NewCollector(
	// Visit only domains: coursera.org, www.coursera.org
	colly.AllowedDomains("gianarb.it", "www.gianarb.it"),

	// Cache responses to prevent multiple download of pages
	// even if the collector is restarted
	colly.CacheDir("./gianarb_cache"),
)
talks := []Talk{}
c.OnHTML("table[id=\""+year+"\"] tbody", func(e *colly.HTMLElement) {
	e.ForEach("tr", func(_ int, row *colly.HTMLElement) {
		talk := Talk{}
                        // for each line "tr" do amazing things
		talks = append(talks, talk)
	})
})

// Before making a request print "Visiting ..."
c.OnRequest(func(r *colly.Request) {
	log.Println("visiting", r.URL.String())
})
err := c.Visit("https://gianarb.it/conferences.html")
if err != nil {
	println(err)
}
```

This is how easy colly is to run. You have to configure the collector and with
the function `OnHTML` you can look for whatever you need to scrape. In this case
I was looking for the table identified with the `id` equals the year got from
the CLI. For each TR element I was creating a new talk to append in a slice. The
`talk` has to to be populated with the actual values scraped cell by cell. It
means that ForEach row we need to look for each td (cell in html) and based on
its index we can identify the content. In my case it looks like this:

```
c.OnHTML("table[id=\""+year+"\"] tbody", func(e *colly.HTMLElement) {
	e.ForEach("tr", func(_ int, row *colly.HTMLElement) {
		talk := Talk{}
		row.ForEach("td", func(_ int, el *colly.HTMLElement) {
			switch el.Index {
			case 0:
			    // Date
			case 1:
                      // Event Name and conference URL (task.EventLink)
			case 3:
                      // Video and slides link
			}
		})
		talks = append(talks, talk)
	})
})
```

I can show you how I coded the case 3, the one that looks for Video or Slides,
takes its link and in case of a YouTube Video it also converts the link into an
embeddable one:

```
links := map[string]string{}
el.ForEach("a", func(_ int, el *colly.HTMLElement) {
	switch el.Text {
	case "Video":
		talk.Video = el.Attr("href")
		if strings.Contains(talk.Video, "youtube.com") {
			u, err := url.Parse(talk.Video)
			if err == nil {
				talk.EmbedVideo = "https://www.youtube.com/embed/" + u.Query().Get("v")
			} else {
errorsToCheck[row.Text+"/youtube_video_without_id"] = el.Text
			}
		} else {
			errorsToCheck[row.Text+"/no_youtube_video"] = el.Attr("href")
		}
	case "Slides":
		talk.Slide = el.Attr("href")
	default:
		links[el.Text] = el.Attr("href")
	}
	talk.Links = links
})
```


This is how I made a boring task enjoyable! And now I have all the talks (minus
two that didn’t get converted but I will add manually) converted and ready to be
rendered as posts.


## Conclusion

This post should not start a useless war between static side generator,
Wordpress or whatever. If you follow me on
[Twitter](https://twitter.com/gianarb) you know that I tweeted recently about
changing Jekyll with something else, mainly because I was thinking how to make a
better use of the contents I create. Digging deeper with Jekyll I discovered
that for now I don’t need more than that and changing tool will end up as
useless and probably not that fun exercise. I am sure all other tools like
Wordpress, Hugo, Gatsby have something similar.
