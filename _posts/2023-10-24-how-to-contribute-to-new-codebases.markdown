---
layout: post
title:  "How I discover new codebases"
date: 2023-10-24 09:10:10
categories: [post]
tags: [codebase]
summary: "Strategies I use when I want to contribute to new codebases"
---

![](/img/you-gen-a-new-job-no-docs-read-the-code-meme.jpg)

Today is my time to be honest with you. I think this meme describes a lot of
places and codebases that I had to deal with or that I contributed to.

I don't want to tell you why because there are many reasons about how you can
end up with a blob of undocumented code and I can write an article on its own
but I want to share the strategy I use to figure it out.

Why am I the right person to do so? Because I change job frequently, not
because of undocumented code obviously and I like to contribute to opensource
software that I use, and sometime I end up contributing to small undocumented
libraries, or to overly documented massive projects.

1. Take a look at the CI/CD system

Many applications have a basic CI/CD system, to run tests, to check code
formatting and sometime to build the software itself. When I don't event know
the language because you I am contributing to a codebase developed in a
language I am not familiar with the CI/CD teaches a lot about the toolchain
that I need to have in order to be effective. Moving forward I tend to replace
tools I don't like with alternatives I am more familiar with but at the
beginning CI/CD, makefiles, npm packages or equivalent files are gold. Worst
case if they don't have unit tests or they don't build their code in CI/CD I
end up knowing about what they use to format their code, it does not look much
but usually it drives to the required tech stack.

2. Dockerfile

Dockerfile are useful to figure out dependency tree and system dependencies
that can teach you a lot about the codebase you are dealing with. It is also
useful to figure out if my teammate are familiar with containers, or more old
style cmake and `./configure` kind of people.

3. The entrypoint, look for that!

`fn main`, `func main`, `index.php` look for the entrypoint! If I see more than
one entrypoint I am in a monorepo, if there is only one it is a single
application. If I can't find one maybe is a library but libraries should have
an entrypoint as well, so look for `Option` or `Configure` classes or objects.
If you find one the class using it is often the library endpoint.

4. Run the test suite

I like to run code locally when I can because it makes things a bit more real,
It validates that I figured out the right toolchain and that I am starting from a
trusted checkpoint.

5. I need an easy win

Why are you looking at such codebase? Do not miss the why! If it is your first
day at work and you got assigned to an apparently easy bug to fix this is your
goal, so try to figure out the right path, you know how to build the software
now, you know how to run the testsuite so leverage that when running the entire
application is still an unknown or is not possible.

I am not saying that tests should be the end goal for you, because there are
codebase with zero or unless tests, but they can be helpful at the beginning as a
north start, if they are unreliable or absent point 5 is still valid but there
is no escape the entrypoint needs to be discovered and used to validate that
your change has the right impact.

When I feel brave, I figured out the part of the code I want to contribute and
the are no tests I write a small scripts that imports such path so I can run
the subset of the code in isolation, quickly and repeatedly without too much
noise. At some point it ca even be turned into a unit test.
