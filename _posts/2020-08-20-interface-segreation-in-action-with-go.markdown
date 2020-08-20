---
img: /img/me.jpg
heroimg: /img/wood_growth.jpg
layout: post
title: "Interface segregation in action with Go"
date:   2020-08-20 09:08:27
categories: [post]
tags: [golang]
summary: "It takes a couple of hours to get an hello world up and running in a
new language but it takes ages to learn it deeply. Even if Go has a learning
curve that is affordable some concepts take time to stick in mind. Interface are
everywhere and this flexibility makes them crucial to write maintainable Go code."
changefreq: daily
---

Everybody should write an article about Golang interface! I don't know why I
waited so long for mine!

Golang interfaces are your best friends when it comes to mocking an object or to
specify a well scoped set of functionalities required by a function to interact
with an object.

Yep! That's how they work, you have an entire object that does a lot of cool
things, but when you pass it to a function only a subset of it get used, that's
when you can replace the structure itself with an interface that only requires
what it is needed by the structure.

In this way you will have a smaller piece of code to mock in your test and to
deal with (this is a good way to hide functions you don't want other people or
yourself in a rush to use).

Even more when you remember to keep the interface small via composition.

For example let's suppose you have to build an interface that describes a generic
resource that you can Create, Update and Delete. This is useful to standardize
something that can be persisted in a database. I am setting this up
so.

You should not use `interface{}` because it is too generic. I used it for
simplicity but Kubernetes for examples uses an object called
[`runtime.Object`](https://godoc.org/k8s.io/apimachinery/pkg/runtime) and it way
better. Go 2 will have generics that will make this situation even easier. Or
you can use code generation as well. But the idea to use a serializable object
like Kubernetes is good.

```golang
type Resource interface {
    Create(ctx context.Context) error
    Update(ctx context.Context, updated interface{}) error
    Delete(ctx context.Context) error
}
```

This is a reasonably small interface, it is easy to satisfy but I do not like
the name. I think it does not give me the ability to figure out what's its
purpose. It represents, a resource but I prefer to call interface as actions or
a adjective. In this case the structure who implements this interface can be
stored in a database. I think a better name for it is:
["Persistable"](https://en.wiktionary.org/wiki/persistable) because it makes
clear its purpose.

A strategy to make an interface smaller in this case is to break it in actions:

```golang
type Creatable interface {
    Create(ctx context.Context) error
}

type Updatable interface {
    Update(ctx context.Context, updated interface{}) error
}

type Deletable interface {
    Delete(ctx context.Context) error
}
```

And you can use composition to create an interface that requires all the three
actions to work if you need it:

```golang
type Persistable interface {
    Deletable
    Updatable
    Creatable
}
```

This is useful when a function uses more than one of those actions, if you have
an interface that contains also `Get` or `View` you can think about a different
split `ReadOnly` contains `Get`, `View` and `Modifiable` that will require only
the functions `Update`, `Create`, `Delete`.

Imagine you are writing a set of http handlers to expose a CRUD API around your
resources:

```
Create
Update
Delete
List
GetByID
```

Usually it looks like this, you can create an interface for every function, all
your resources will implement the functions and you will be able to write a
single "Create" handle for all the resources:

```golang
func CreateHandle(c Creatable) func(w http.ResponseWriter, r *http.Request) {
    return http.HandleFunc("/resource", func(w http.ResponseWriter, r *http.Request) {
        if err := c.Create(r.Context); if err != nil {
            w.WriteHeader(http.StatusInternalServerError)
            return
        }
        w.WriteHeader(http.StatusCreated)
    })
}
```

If you have to write a test for the handler it does not matter how complicated
the resource is, you just have to mock the `Creatable` interface, one single
function. This is a very basic example, if you need to add validation the
`Creatable` function can require a `func Valid() error` that you can
add incrementally in all your resources.

```golang
func CreateHandle(c Creatable) func(w http.ResponseWriter, r *http.Request) {
    return http.HandleFunc("/resource", func(w http.ResponseWriter, r *http.Request) {
        if err := c.Valid(); err != nil {
            w.WriteHeader(http.StatusBadRequest)
            return
        }
        if err := c.Create(r.Context); if err != nil {
            w.WriteHeader(http.StatusInternalServerError)
            return
        }
        w.WriteHeader(http.StatusCreated)
    })
}
```
