---
heroimg: /img/hero/action.jpg
layout: post
title:  "GitHub actions to deliver on kubernetes"
date:   2019-01-22 08:08:27
categories: [post]
tags: [kubernetes, github action, serverless, ci, automation]
summary: "GitHub recently released a new feature called GitHub Actions. They are
a serverless approach to allow developers to run their own code based on what
happens to a particular repository. They are amazing for continuous integration
and delivery. I used them to deploy and validate kubernetes code."
changefreq: daily
---
Recently GitHub released a new feature called Actions. To me, it looks like the
best implementation I can think of for serverless.  I used AWS Lambda and API
Gateway for some basic API, and I wrote a prototype of an application capable of
running functions using containers called
[gourmet](https://github.com/gianarb/gourmet) I don't buy the fact that it will
make my code easy to manage. At least not to write API or web applications.

<blockquote class="twitter-tweet tw-align-center"><p lang="en" dir="ltr">I used the <a
href="https://twitter.com/hashtag/GitHubActions?src=hash&amp;ref_src=twsrc%5Etfw">#GitHubActions</a>
to verify and deploy code to a <a
href="https://twitter.com/hashtag/kubernetes?src=hash&amp;ref_src=twsrc%5Etfw">#kubernetes</a>
cluster <a href="https://t.co/nfkjmYKPKs">https://t.co/nfkjmYKPKs</a> I am
impressed about how wonderful this feature is designed and implemented! <a
href="https://twitter.com/github?ref_src=twsrc%5Etfw">@Github</a> you
🤘!</p>&mdash; :w !sudo tee % (@GianArb) <a
href="https://twitter.com/GianArb/status/1087640589838008321?ref_src=twsrc%5Etfw">January
22, 2019</a></blockquote> <script async
src="https://platform.twitter.com/widgets.js" charset="utf-8"></script>


That's why I like what GitHub did because they used serverless for what I think
it is designed for, extensibility.

GitHub Actions just like Lambda functions on AWS are a powerful and managed way
to extend their product straightforwardly.

With AWS Lambda you can hook your code to almost whatever event happens: EC2
creations, termination, route53 DNS record change and a lot more. You don't need
to run a server, you load your code, and it just works.

Jess Frazelle wrote a blog post about ["The Life of a GitHub
Action](https://blog.jessfraz.com/post/the-life-of-a-github-action/), and I
decided to try something I had my mind since a couple of weeks but it required a
CI server, and it was already too much for me.

Time to time I like the idea to have a kubernetes cluster that I can use for the
testing purpose, so I created a private repository that it is not ready to be
open source because it is a mess with secrets inside and so on.

![](/img/sorry.jpg)

In any case, to give you an idea, this is the project's folder:

```
├── .github
│   ├── actions
│   │   ├── deploy
│   │   │   ├── deploy
│   │   │   └── Dockerfile
│   │   └── dryrun
│   │       ├── Dockerfile
│   │       └── dryrun
│   └── main.workflow
└── kubernetes
    ├── digitalocean.yaml
    ├── external-dns.yaml
    ├── micro.yaml
    ├── namespaces.yaml
    ├── nginx.yaml
    └── openvpn.yaml
```
The `kubernetes` directory contains all the things I would like to install in my
cluster.  For every new push on this repository, I would like to check if it can
be applied to the kubernetes cluster with the command `kubectl apply -f
./kubernetes --dryrun` and when the PR is merged the changes should get applied.

So I created my workflow in `.github/main.workflow`: ( I left some comment to
make it understandable)

```
## Workflow defines what we want to call a set of actions.

## For every new push check if the changes can be applied to kubernetes ## using the action called: kubectl dryrun
workflow "after a push check if they apply to kubernetes" {
  on = "push"
  resolves = ["kubectl dryrun"]
}

## When a PR is merged trigger the action: kubectl deploy. To apply the new code to master.
workflow "on merge to master deploy on kubernetes" {
  on = "pull_request"
  resolves = ["kubectl deploy"]
}

## This is the action that checks if the push can be applied to kubernetes
action "kubectl dryrun" {
  uses = "./.github/actions/dryrun"
  secrets = ["KUBECONFIG"]
}

## This is the action that applies the change to kubernetes
action "kubectl deploy" {
  uses = "./.github/actions/deploy"
  secrets = ["KUBECONFIG"]
}
```
The `secrets` are an array of environment variables that you can use to set
values from the outside. If your account has GitHub Action enabled there is a
new Tag inside the Settings in every repository called "Secrets."

You can set key-value pairs usable as you see in my workflow. For this example,
I set the `KUBECONFIG` as the base64 of a kubeconfig file that allows the GitHub
Action to authorize itself to my Kubernetes cluster.

Both actions are similar the first one is in the directory
`.github/actions/dryrun`

```
├── .github
    ├── actions
        └── dryrun
            ├── Dockerfile
            └── dryrun
```
It contains a Dockerfile

```
FROM alpine:latest

## The action name displayed by GitHub
LABEL "com.github.actions.name"="kubectl dryrun"
## The description for the action
LABEL "com.github.actions.description"="Check the kubernetes change to apply."
## https://developer.github.com/actions/creating-github-actions/creating-a-docker-container/#supported-feather-icons
LABEL "com.github.actions.icon"="check"
## The color of the action icon
LABEL "com.github.actions.color"="blue"

RUN     apk add --no-cache \
        bash \
        ca-certificates \
        curl \
        git \
        jq

RUN curl -L -o /usr/bin/kubectl https://storage.googleapis.com/kubernetes-release/release/v1.13.0/bin/linux/amd64/kubectl && \
  chmod +x /usr/bin/kubectl && \
  kubectl version --client

COPY dryrun /usr/bin/dryrun
CMD ["dryrun"]
```

As you can see to describe an action, you need just a Dockerfile, and it works
the same as in docker. The CMD `dryrun` is the bash script I copied here:

```bash
#!/bin/bash

main(){
    echo ">>>> Action started"
    # Decode the secret passed by the action and paste the config in a file.
    echo $KUBECONFIG | base64 -d > ./kubeconfig.yaml
    echo ">>>> kubeconfig created"
    # Check if the kubernetes directory has change
    diff=$(git diff --exit-code HEAD~1 HEAD ./kubernetes)
    if [ $? -eq 1 ]; then
        echo ">>>> Detected a change inside the kubernetes directory"
        # Apply the changes with --dryrun just to validate them
        kubectl apply --kubeconfig ./kubeconfig.yaml --dry-run -f ./kubernetes
    else
        echo ">>>> No changed detected inside the ./kubernetes folder. Nothing to do."
    fi
}

main "$@"
```
The second action is almost the same as this one, the Dockerfile is THE same, so
I am not posting it here, but the CMD looks like this:

```bash
#!/bin/bash

main(){
    # Decode the secret passed by the action and paste the config in a file.
    echo $KUBECONFIG | base64 -d > ./kubeconfig.yaml
     # Check if it is an event generated by the PR is a merge
    merged=$(jq --raw-output .pull_request.merged "$GITHUB_EVENT_PATH")
    # Retrieve the base branch for the PR because I would like to apply only PR merged to master
    baseRef=$(jq --raw-output .pull_request.base.ref "$GITHUB_EVENT_PATH")

    if [[ "$merged" == "true" ]] && [[ "$baseRef" == "master" ]]; then
        echo ">>>> PR merged into master. Shipping to k8s!"
        kubectl apply --kubeconfig ./kubeconfig.yaml -f ./kubernetes
    else
        echo ">>>> Nothing to do here!"
    fi
}

main "$@"
```
That's everything, and I am thrilled!

![](/img/party.jpg)

There is nothing more to say other than "GitHub actions are amazing!". They look
well designed since day! The workflow file has a generator that even if I didn't
use it because I don't like colors, it seems amazing. The secrets allow us to do
integration with third-party services out of the box and you can use bash to do
whatever you like! Let me know what you use them for on
[Twitter](https://twitter.com/gianarb).
