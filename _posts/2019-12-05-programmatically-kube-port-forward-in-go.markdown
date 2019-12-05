---
img: /img/gianarb.png
layout: post
title: "Programmatically Kubernetes port forward in Go"
date: 2019-12-05 06:08:27
categories: [post]
tags: [kubernetes, k8s, golang]
summary: "Depending on your networking configuration port forwarding will may be
the unique way for you to reach pods or services running in Kubernetes. When you
develop a CLI integration that has to interact with pods running inside the
cluster you can programmatically do a port forwarding in golang."
changefreq: daily
---
Along the way I saw at least two different ways to manage Kubernetes clusters
from a networking prospective. Some companies configure a VPN inside the
Kubernetes Cluster, in this way a developer connected in the VPN can reach pods
and services.

It is not mandatory but suggested having a good network segmentation in order
to be able to manage what a person connected in the VPN can touch and see.
Achieving this level of control is not easy in Kubernetes a lot of the open
source CNI plugin does not have this feature at all and I understand why in
operations this is evaluated as a safe approach. It is very convenient if close
an eye because pods and services are just IPs that you can reach from your
laptop and if you configure the VPN to push the Kubernetes DNS you can also
resolve them as DNS lookup.

The alternative I saw is to lock everybody out leaves as unique way to interact
with a service or a pod the command `kubectl port-forward`. In this way the
authentication and authorization method in Kubernetes allows you to decide who
can do port-forwarding on what based on namespace for example. Or at least you
can use Kubernetes Audit logs to figure out who did port forwarding if something
bad happens.

We tried both ways, I was to one pushing for the first, but we never achieved a
good segmentation and at some point I got locked down, sadly as it sound. Anyway
I like to automate things and I had to figure out a way to make my scripts to
work with this new approach.

![](/img/sub.jpg){:class="img-fluid"}

I started to dig in the `kubectl` code because we all know that it is capable of
doing the port forwarding. I had some trouble figuring out the right parameters
and to make them to work but at the end I did it! So here we are! If I can do it
you can do it as well!

The main repository with the code and an example is in
[github.com/gianarb/kube-port-forward](https://github.com/gianarb/kube-port-forward),
you can run it there. I am gonna explain it a bit here.

It is a simple CLI that mocks what `kubectl port-forward` already does but I
extrapolated the code needed to do and control a port forwarding. I will write
here as soon as the reason about why I did that is open source, I am telling it
to you right now STAY TUNED! It will be great!

First of all I used the `k8s.io/cli-runtime/pkg/genericclioptions` library to
configure a stream, we already used in the [blog post about writing a CLI that
uses the same flags as the kubectl](/blog/kubectl-flags-in-your-plugin). A
stream is a `struct` used by different `kubernetes` service when they need to
get or print information from a stream, in this case I am using `os.Stdout`,
`os.Stdin`, `os.Stderr` for simplicity, but where I do not need to print out the
output I use a `bytes.Stream` like this:

```go
var berr, bout bytes.Buffer
buffErr := bufio.NewWriter(&berr)
buffOut := bufio.NewWriter(&bout)
```

In order to make this code easy to read I had a structure to request the port
forwarding for a pod:

```go
type PortForwardAPodRequest struct {
	// RestConfig is the kubernetes config
	RestConfig *rest.Config
	// Pod is the selected pod for this port forwarding
	Pod v1.Pod
	// LocalPort is the local port that will be selected to expose the PodPort
	LocalPort int
	// PodPort is the target port for the pod
	PodPort int
	// Steams configures where to write or read input from
	Streams genericclioptions.IOStreams
	// StopCh is the channel used to manage the port forward lifecycle
	StopCh <-chan struct{}
	// ReadyCh communicates when the tunnel is ready to receive traffic
	ReadyCh chan struct{}
}
```

And I wrote the function that actually does the port forward:

```go
func PortForwardAPod(req PortForwardAPodRequest) error {
	path := fmt.Sprintf("/api/v1/namespaces/%s/pods/%s/portforward",
		req.Pod.Namespace, req.Pod.Name)
	hostIP := strings.TrimLeft(req.RestConfig.Host, "htps:/")

	transport, upgrader, err := spdy.RoundTripperFor(req.RestConfig)
	if err != nil {
		return err
	}

	dialer := spdy.NewDialer(upgrader, &http.Client{Transport: transport}, http.MethodPost, &url.URL{Scheme: "https", Path: path, Host: hostIP})
	fw, err := portforward.New(dialer, []string{fmt.Sprintf("%d:%d", req.LocalPort, req.PodPort)}, req.StopCh, req.ReadyCh, req.Streams.Out, req.Streams.ErrOut)
	if err != nil {
		return err
	}
	return fw.ForwardPorts()
}
```
An exercise that I can leave for you is to add Service support to this function,
you can open a PR if you like on
[github.com/gianarb/kube-port-forward](https://github.com/gianarb/kube-port-forward).

The `Stop` and `Ready` channels are crucial to manage the port forward because
as you see in the example it is a blocking operation it means that it will
luckily always run inside a goroutine. Those two channels gives you what you
need to understand when the port forward is ready to get traffic `ReadyCh` and
you have the capabilities to stop it `StopCh`.

My example is basic, I am closing the port forwarding when the `SIGTERM` signal
gets notified:

```go
sigs := make(chan os.Signal, 1)
signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM)
go func() {
    <-sigs
    fmt.Println("Bye...")
    close(stopCh)
    wg.Done()
}()
```

I just wait until the readyCh tells me that the connection is
up and running

```go
select {
case <-readyCh:
    break
}
println("Port forwarding is ready to get traffic. have fun!")
```

As soon as I coded this feature I saw that it was gonna be an easy but useful
post. I wrote a [report with O'Reilly](/blog/extending-kubernetes-oreilly) about
how to extend Kubernetes, you can find more about Go and Kube there. It is a
free PDF.

I hope you enjoyed it and [let me know](https://twitter.com/gianarb) what cool
things you are gonna do port-forwarding the universe!

