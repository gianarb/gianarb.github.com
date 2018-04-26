---
layout: post
title:  "Go how to cleanup HTTP request terminated."
date:   2018-04-25 10:38:27
img: /img/gopher.png
categories: [post]
tags: [golang]
summary: "Cleaning up HTTP request, the most expensive one can be a huge
performance improvement for your application. This short article shows how to
handle HTTP request termination in Go."
priority: 1
---
Expensive HTTP handler is everywhere, doesn't matter how good you are as a
developer. Business logic is what matters in our application, and it can be
pretty complicated. It can create large files, resources on AWS starts
thousands of containers on Kubernetes.

This kind of procedures have in common they can be very slow and they produce a
lot of garbage if the system/person who requires that stops prematurely by
mistake or not.

If your API requests create AWS resources and the client, terminate the call
you should clean what you created.

if you are generating a report and the customer changes are mind and refresh
you should stop the procedure.

You bet! Queues, background processes probably fit better but coming back on
the previous example, if you are computing something and who is waiting for the
result changed his mind, stop and release resources can be a massive
optimization.

```bash
package main

import (
    "fmt"
    "io/ioutil"
    "net/http"
    "os"
    "time"
)

func main() {
    http.HandleFunc("/a", func(w http.ResponseWriter, r *http.Request) {
        err := ioutil.WriteFile(os.TempDir()+"/txt", []byte("hello"), 0644)
        if err != nil {
            panic(err)
        }
        println("new file " + os.TempDir() + "/txt")
        notify := w.(http.CloseNotifier).CloseNotify()
        go func() {
            <-notify
            println("The client closed the connection prematurely. Cleaning up.")
            os.Remove(os.TempDir() + "/txt")
        }()
        time.Sleep(4 * time.Second)
        fmt.Fprintln(w, "File persisted.")
    })
    http.ListenAndServe(":8080", nil)
}
```

When you are building an HTTP server in Go, you can use a channel provided by
the Zhttp.ResponseWriter` to wait for the connection to be closed. And if it
happens, you can take action.  The prototype above is very simple, every
request stores a file but I would like, remove the file if the client closes
the connection.

```bash
$ run main.go
```

You can start the server, and from another terminal, you can start a `curl`, you
will see that after almost 4 seconds your request will succeed and the file
will be persisted on disk. Check it!

```
$ time curl http://localhost:8080/a
File persisted.

real    0m4.018s
user    0m0.008s
sys     0m0.006s
$ cat /tmp/txt
```

Now let's suppose that the client terminates the connection because it is too
slow or the person who made the request doesn't care anymore.
Are you going to leave that request going? Event if nobody cares and it is just
consuming resources?

As you can see I am using the Notifier to remove the file if the client
terminates the connection:

```go
notify := w.(http.CloseNotifier).CloseNotify()
go func() {
    <-notify
    println("The client closed the connection prematurely. Cleaning up.")
    os.Remove(os.TempDir() + "/txt")
}()
```

You can check it stopping a `curl` just after starting it:

```
$ time curl http://localhost:8080/a
^C

real    0m1.016s
user    0m0.008s
sys     0m0.005s
```
And the server reports

```
$ go run main.go
new file /tmp/txt
The client closed the connection prematurely. Cleaning up.
```

That's it! Build and clean after yourself!
