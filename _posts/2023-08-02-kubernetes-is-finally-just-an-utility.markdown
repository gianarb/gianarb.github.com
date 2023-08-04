---
layout: post
title: "Kubernetes is finally just an utility"
date: 2023-08-02 10:08:27
categories: [post]
tags: [kubernetes]
summary: "Kubernetes is a tool, not a religion. It can tech you a lot about
about scalability, resiliency but if your buisness is not driven by this
specific tool you should not overlook it."
---
Kubernetes and cloud-native is a topic I spent a good time of my professional
life contributing to. I wrote some code, operators, articles, talks, libraries,
and I have been a member of the awesome Release Team for a few iterations.
Kubernetes helped me a lot to grow as a developer and improved my ability to
collaborate with people all over the world.

I know how to operate it, how it is written, the components it is made of, and
so on, but I do not work for a company that makes money out of it anymore.

I worked for cloud providers making money building software that integrates
with Kubernetes. I worked for companies that made money and took investments
saying: "Our solution runs natively on Kubernetes". They didn't age well and I
am not surprised at all.

If you are like me and your company does not profit from Kubernetes do not look
at it as something important. It is good to learn about it, to operate it but
that's it, like you do with Linux, systemd, git, or everything else, because
that's what it is. I spoke with many people who told me that Kubernetes is
complex, I got it, so it can't be that complicated, can't be more complicated
than systemd. There are cloud providers or companies that you can pay to get a
fully functioning Kubernetes endpoint to interact with up and running. I read
articles related to how the EC2 service works, and how and why they build it,
but that's it. I don't feel bad about using tools or services without knowing
all the details about how they are made.

I don't want to discourage you from using Kubernetes or contributing to it. I
advocate for the good practices that Kubernetes enforces and teaches but that's
the best it does. Today after two years of not touching it I installed kind,
the kubectl and I wrote 352 lines of YAML that I successfully applied to a
Kubernetes cluster because I am working with a potential customer that runs on
Azure and we picked Kubernetes as the common language, I think this is its
superpower. A technology capable of improving collaboration, and breaking
barriers is a gift that we should protect. And it does not require me to know
about CNI, CRI, CTO, and kubelet (can you guess the wrong one?).

Last week we expanded our solution from AWS, where I built the infrastructure
out of autoscaling groups, Launch Templates, EC2, load balancers, and duct take
to GCP where I decided to use GKE Autopilot driven all by Terraform. Not Flux
or Helm, two technologies that I never used but Terraform because the "IoC"
solution we have is 100% based on that and I didn't feel the need for something
else. GKE because it makes sense, it is quick and I don't have the experience
    on GPC as I have on AWS to operate at a "lower level" in a reasonable time.

The solution we have on AWS is too expensive because there are a million tiny
details when building using simple components like the one I mentioned that are
painful to figure out at least for me, or at least not interesting from a
business point of view. This is why I am probably moving to ECS pretty soon.
Not EKS, because EKS does not look at simple as GKE Autopilot.

The developer I want to be and the one I like to work with put effort into
finding the right solution based on their current context, building all the
surroundings that will play a difference in the game we all have to fight with:
evolution and time. It requires knowledge and skills exceeding a specific
technology or trend.  There are similarities with woodworking where complex
cuts or repetitive tasks require jigs, and those need to get built with
accuracy because they can drive the success of the primary project.
