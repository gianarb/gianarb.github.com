---
img: /img/gianarb.png
layout: post
title:  "From sequential to parallel with Go"
date:   2019-02-21 08:08:27
categories: [post]
tags: [go, golang, parallelization, trick, code]
summary: "From a sequence of action to parallelization in Go. Using channels and
wait groups from the sync package."
changefreq: daily
---
Everything starts as a sequence of events. You have a bunch of things to do and
you are not sure how long or hard to manage they will be.

As a pragmatic developer, you go over the list of things, and you make them one
by one. The script runs, it works, and everyone is happy.

```go
package main

import (
    "fmt"
    "log"
    "time"
)

func main() {
    list := []string{"a", "b", "c", "d", "e", "f", "g", "h", "i", "l"}
    for _, v := range list {
        v, err := do(v)
        if err != nil {
            log.Printf("nop")
        }
        fmt.Println(v)
    }

}

func do(s string) (string,error) {
    time.Sleep(1*time.Second)
    return fmt.Sprintf("%s-%d", s, time.Now().UnixNano()),nil
}
```

Let's execute it:

```
$ time go run c.go
a-1550742371537033061
b-1550742372537419148
c-1550742373537846015
d-1550742374538086031
e-1550742375538488129
f-1550742376538746707
g-1550742377539047837
h-1550742378539540979
i-1550742379539938404
l-1550742380540339887

real    0m10.174s
user    0m0.149s
sys     0m0.074s
```

Until something changes from the outside, the outside world is a terrible place.

![](https://media.giphy.com/media/124pc9nFq7ZScU/giphy.gif)

The list of things to do grows too much, and your program runs too slow to be
competitive, so you start to think about parallelization.

Luckily for you, every action doesn't depend on anything else, so you don't need
to stop if one of them fails or even worst you don't need to do nothing weird,
you skip that, and you log the failure.

There is an easy way to migrate the code about with something that safely runs
in parallel just using some built-in functions in Go like channels and
WaitGroups.

```go
package main

import (
    "fmt"
    "log"
    "sync"
    "time"
)

func main() {
    fmt.Println("Start")
    parallelization := 2
    list := []string{"a", "b", "c", "d", "e", "f", "g", "h", "i", "l"}
    c := make(chan string)

    var wg sync.WaitGroup
    wg.Add(parallelization)
    for ii := 0; ii < parallelization; ii++ {
        go func(c chan string) {
            for {
                v, more := <-c
                if more == false {
                    wg.Done()
                    return
                }

                v, err := do(v)
                if err != nil {
                    log.Printf("nop")
                }
                fmt.Println(v)
            }
        }(c)
    }
    for _, a := range list {
        c <- a
    }
    close(c)
    wg.Wait()
    fmt.Println("End")
}

func do(s string) (string, error) {
    time.Sleep(1 * time.Second)
    return fmt.Sprintf("%s-%d", s, time.Now().UnixNano()), nil
}
```

`parallelization` should be an external parameter that you can change to
parallelize more or less. With a parallelization factor of 2 the benchmark looks
like:

```bash
$ time go run c.go
Start
a-1550742531701829912
b-1550742531701820924
d-1550742532702088077
c-1550742532702180981
e-1550742533702473002
f-1550742533703389899
g-1550742534702714251
h-1550742534703981070
i-1550742535702992582
l-1550742535704308486
End

real    0m5.269s
user    0m0.249s
sys     0m0.078s
```

Almost half of the time. Let's try with 5.

```bash
$ time go run c.go
Start
e-1550742633337320607
b-1550742633337280491
c-1550742633337474112
d-1550742633337280481
a-1550742633337298154
h-1550742634338002235
i-1550742634338073772
f-1550742634338033897
g-1550742634338019639
l-1550742634338231670
End

real    0m2.145s
user    0m0.144s
sys     0m0.058s
```

I wrote this article because I like how easy it was for this use case to run in
parallel. Based on how complicated your `do` function is you need to be more
careful.

If your `do` function calls an external service it can fail, or it can rate
limit you because you are parallelizing too much. But these are all problem that
you can solve increasing the number of safeguards in your code.

Something I learned using this and calling AWS intensively to take snapshots is
the fact that EC2 snapshots happen in the background on AWS, so if you have
thousands of nodes and you call AWS it will rate limit you or you won't have a
good experience of what happens on the AWS side in reality.

A basic trick is to place a `batch delay` parameter that sleeps before every
execution

```go
v, more := <-c
if more == false {
    wg.Done()
    return
}

// Sleep here!

v, err := do(v)
if err != nil {
    log.Printf("nop")
}
fmt.Println(v)
```

This is a very crafty fix but if you catch this problem like me when everything
is failing this is a safe bullet you should try.

Parallelization is fun, but in reality, it increases complexity. Go servers
primitives that are solid foundations but it is not you to instrument your code
well enough to be confident about how it works.

I will write the next chapter about this where I will use opencensus or
opentracing to trace what is going on here!
