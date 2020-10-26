---
layout: post
title:  "Reactive planing in Golang. Reach a desired number adding and subtracting random numbers"
date:   2020-10-26 10:08:27
categories: [post]
tags: [software design, golang, go, reconciliation loop, reactive planning]
summary: "An example about how to write reactive planning in Go. Code and step
by step solution for an exercise I developed to learn planner"
heroimg: /img/steel-iron-metal.jpg
---

Ciao! A few months ago, probably a year, I wrote a small library called
[planner](https://github.com/gianarb/planner). It comes from my experience using
reactive planning and Kubernetes. I am really in love with this way of writing
code because it sounds very reliable to me.

Over the last couple of days, I decided to write documentation for it! So now it
is presentable; I streamed that with [Twitch](https://twitch.tv/gianarb) if you
like to watch people coding!

As part of the library's readme, I wrote a small program, and I left a couple of
exercises to the reader. With this article, I want to solve them.

You can follow this article and try it yourself, starting from
[play.golang.com](https://play.golang.com/p/0LuIoMtp10f).

```golang
package main

import (
	"context"
	"time"

	"github.com/gianarb/planner"
	"go.uber.org/zap"
)

func main() {
	ctx, done := context.WithTimeout(context.Background(), 10*time.Second)
	defer done()

	countPlan := &CountPlan{
		Target: 20,
	}
	scheduler := planner.NewScheduler()
	scheduler.WithLogger(initLogger())

	scheduler.Execute(ctx, countPlan)
}

type CountPlan struct {
	Target  int
	current int
}

func (p *CountPlan) Create(ctx context.Context) ([]planner.Procedure, error) {
	if p.current < p.Target {
		return []planner.Procedure{&AddNumber{plan: p}}, nil
	}
	return nil, nil
}

func (p *CountPlan) Name() string {
	return "count_plan"
}

type AddNumber struct {
	plan *CountPlan
}

func (a *AddNumber) Name() string {
	return "add_number"
}

func (a *AddNumber) Do(ctx context.Context) ([]planner.Procedure, error) {
	a.plan.current = a.plan.current + 1
	return nil, nil
}

func initLogger() *zap.Logger {
	cfg := zap.NewProductionConfig()
	cfg.Encoding = "console"
	l, _ := cfg.Build()
	return l
}
```

This program tries to reach the `Target` (20 in this case) adding numbers to the
current state. If you execute this program as it is you will get the following
logs:

```console
1.257894e+09	info	planner@v0.0.1/scheduer.go:41	Started execution plan count_plan	{"execution_id": "98d28eed-9b3b-4ad8-bfbd-1b5338d1a649"}
1.257894e+09	info	planner@v0.0.1/scheduer.go:59	Plan executed without errors.	{"execution_id": "98d28eed-9b3b-4ad8-bfbd-1b5338d1a649", "execution_time": "0s", "step_executed": 20}
```

As you can see, the scheduler executed the plan `count_plan` successfully, and
it took 20 steps to get there (`step_executed: 20`).

Reasonable because, as you can see, the `CounterPlan.Create` function returns an
`AddNumber` procedure and that procedure only adds 1 to the current state. It is
just a counter; let's make it a bit more fun. I want to add or substract random
number until the target is reached. The program adds when the current state is above the
target, when above it subtracts. If it's equal we are done. This is a simple way
to simulate something that has to adapt, too simple to sound cool but still
something understandable.

### Change the AddNumber to use a randomly generated number.

We need to change the AddNumber in order to add not 1 but a random number. Let's
do it:

```go
var random *rand.Rand
func initRandom() {
    s1 := rand.NewSource(time.Now().UnixNano())
    random = rand.New(s1)
}
```

At this point, we can use `random` as part of the `AddNumber.Do` function.

```go
type AddNumber struct {
	plan *CountPlan
}

func (a *AddNumber) Name() string {
	return "add_number"
}

func (a *AddNumber) Do(ctx context.Context) ([]planner.Procedure, error) {
	a.plan.current = a.plan.current + random.Intn(10)
	return nil, nil
}
```

For simplicity, I am taking a random number between 0 and 10. What happens now?
The problem now is that we can go above the target, so we have to make our
`CounterPlan.Create` function and our logic a bit more complicated.

## Evolve the Create function to subtract numbers from the current state

```go
func (p *CountPlan) Create(ctx context.Context) ([]planner.Procedure, error) {
	if p.current < p.Target {
		return []planner.Procedure{&AddNumber{plan: p}}, nil
	} else if p.current > p.Target {
		return []planner.Procedure{&SubtractNumber{plan: p}}, nil
	}
	return nil, nil
}
```

When we go above the target, the Plan subtracts a random number, and it keeps
going until we get to it. `SubtractNumber` does the opposite of what
`AddNumber` does, it subtracts a random number between 0 an 10.

```go
type SubtractNumber struct {
	plan *CountPlan
}

func (a *SubtractNumber) Name() string {
	return "subtract_number"
}

func (a *SubtractNumber) Do(ctx context.Context) ([]planner.Procedure, error) {
	a.plan.current = a.plan.current - random.Intn(10)
	return nil, nil
}
```

You can run the result [here](https://play.golang.com/p/JDuizzUI86M), and you
will see that based on the random numbers, it adds or subtracts the number of
executed steps changes.

NOTE: the golang playground always starts from the same time; in my example, I
use time as Seed; for this reason, to see a variation in the number of steps,
you will have to run the code locally.

## Conclusion

This is probably a too straightforward example, but let's imagine that
your Target is not fixed and varies based on external factors. Your house's
temperature and this program is a thermostat that has to keep your room at the
desired temperature. Or the number of instances running in your cloud provider,
and you have to keep them balanced. This last use case is the exact problem I
solved writing [keepit](https://github.com/gianarb/keepit) a replica set for
[Equinix Metal](https://metal.equinix.com) servers. I used planner, so check it out.

I didn't highlight this example because this pattern gives you an excellent way
to measure how reliable your program is. Think about it in this way; you can
programmatically handle errors returning a procedure or more than one that can
mitigate the error itself. It can be a "sleep for 5 minutes and retry", or you
can do something more complicated, and until the Plan keeps returning work to
do, you will have the opportunity to succeed. I extracted an highlight from the
[Twitch stream](https://www.twitch.tv/videos/780401570) rambling about this.

Have a nice week!
