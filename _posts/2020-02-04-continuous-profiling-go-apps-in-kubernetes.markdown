---
img: /img/profefe.png
layout: post
title: "Continuous Profiling Go applications running in Kubernetes"
date: 2020-02-04 09:08:27
categories: [post]
tags: [golang, pprof, profefe, kubernetes]
summary: "Kube-Profefe is an open source project that acts like a bridge between
Kubernetes and Profefe. It helps you to implement continuous profiling for Go
applications running in Kubernetes."
changefreq: daily
---
Recently I wrote [“Continuous profiling in Go with
Profefe”](https://gianarb.it/blog/go-continuous-profiling-profefe), an article
about the new shiny open source project I am contributing to.

**TLDR:** Profefe is a registry for pprof profiles. You can push them embedding
an SDK in your application or you can write a collector (cronjob) that gets
profiles and push the tar via the Profefe API. Side by side with the profile you
have to send other information like:

* Type: represents the profile type such as mutex, goroutines, CPU and so on
* Service: identifies the source for this profile, for example, the binary name
* InstanceID: identifies where it comes from, for example, pod name or server
  hostname
* Labels: are optional key/value pairs that you can use at query time to filter
  profiles. If you are building the same service with two different Go versions
  to check for performance degradation you can label the profiles with
  `go=1.13.4` for example.

The article has way more content but that’s enough. You can keep reading with
only this information.

## Kubernetes

As you know at InfluxData we use Kubernetes, our services already expose the
[pprof HTTP handler](https://golang.org/pkg/net/http/pprof/) and we can not
instrument all the services with the Profefe SDK, for those reasons we had to
write our own collectors capable of getting pprof profiles via the Kubernetes
API and to push them into Profefe. That’s why we decided to go with a different
approach. I wrote a project called
[kube-profefe](https://github.com/profefe/kube-profefe). It acts as a bridge
between the Profefe API and Kubernetes. The repository provides two different
binaries:

* A kubectl plugin that you can install (even via krew) that servers useful
  utilities to interact with the profefe API (profefe at the moment does not
  have a CLI) and to capture profiles from running pod.
* A collector that can run as a cronjob, it goes pod by pod looking for profiles
  to collect and it will push them to Profefe.


## Architecture

In order to configure the collector or to capture profiles from a running
container, it leverages pod annotations. Only the pods with the annotation
`pprof.com/enable=true` will be taken into consideration from kube-profefe.
Other annotations are optional or they have default values. This one is the
unique one that has to be set to make kube-profefe aware of your pod.

The example above shows a Pod spec that enables profefe capabilities:

```yaml
apiVersion: v1
kind: Pod
metadata:
  name: influxdb-v2
  annotations:
    "profefe.com/enable": "true"
    "profefe.com/port": "9999"
spec:
  containers:
  - name: influxdb
    image: quay.io/influxdb/influxdb:2.0.0-alpha
    ports:
    - containerPort: 9999
```

As you can see there are other annotations such as `profefe.com/port` by default
is 6060. In this case it is pointed to 9999 because that's where the pprof HTTP
handler runs in InfluxDB v2.  A full list of annotations is maintained in the
project's README.md.

There is not a lot more to know about the underling mechanism that enpowers
kube-profefe, we are gonna deep dive on both components: the kubectl plugin and
the collector.

## Kubectl-profefe: the kubectl plugin

A kubectl plugin is nothing more than a binary located in your $PATH with the
prefix name “kubectl-”. In my case the binary is released with the name
kubectl-profefe, when located in your $PATH you will be able to run a command
like:

```bash
$ kubectl profefe --help
It is a kubectl plugin that you can use to retrieve and manage profiles in Go.

Usage:
  kubectl-profefe [flags]
  kubectl-profefe [command]

Available Commands:
  capture     Capture gathers profiles for a pod or a set of them. If can filter by namespace and via label selector.
  get         Display one or many resources
  help        Help about any command
  load        Load a profile you have locally to profefe

Flags:
  -A, --all-namespaces                 If present, list the requested object(s) across all namespaces. Namespace in current context is ignored even if specified with --namespace.
      --as string                      Username to impersonate for the operation
      --as-group stringArray           Group to impersonate for the operation, this flag can be repeated to specify multiple groups.
      --cache-dir string               Default HTTP cache directory (default "/home/gianarb/.kube/http-cache")
      --certificate-authority string   Path to a cert file for the certificate authority
      --client-certificate string      Path to a client certificate file for TLS
      --client-key string              Path to a client key file for TLS
      --cluster string                 The name of the kubeconfig cluster to use
      --context string                 The name of the kubeconfig context to use
  -f, --filename strings               identifying the resource.
  -h, --help                           help for kubectl-profefe
      --insecure-skip-tls-verify       If true, the server's certificate will not be checked for validity. This will make your HTTPS connections insecure
      --kubeconfig string              Path to the kubeconfig file to use for CLI requests.
  -n, --namespace string               If present, the namespace scope for this CLI request
  -R, --recursive                      Process the directory used in -f, --filename recursively. Useful when you want to manage related manifests organized within the same directory. (default true)
      --request-timeout string         The length of time to wait before giving up on a single server request. Non-zero values should contain a corresponding time unit (e.g. 1s, 2m, 3h). A value of zero means don't timeout requests. (default "0")
  -l, --selector string                Selector (label query) to filter on, supports '=', '==', and '!='.(e.g. -l key1=value1,key2=value2)
  -s, --server string                  The address and port of the Kubernetes API server
      --token string                   Bearer token for authentication to the API server
      --user string                    The name of the kubeconfig user to use

Use "kubectl-profefe [command] --help" for more information about a command.
```

This output should look very familiar to you, there are a lot of options usable
with any other kubectl native command. Mainly around authentication: --user,
--server, --kubeconfig, --client-certificate… Or around pod selection: -l,
--selector, -n, --namespace, --all-namespaces. If you are curious about how to
write a friendly kubectl plugin I wrote [“kubectl flags in your
plugin”](https://gianarb.it/blog/kubectl-flags-in-your-plugin) check it out.

This plugin, even if it is not native, uses the same authentication mechanism in
use from the kubectl so, where ever the kubectl works, this plugin should work
as well.

The pod selectors -l, -n, for example, are useful when running the command:

```
$ kubectl profefe capture
```

Capture, as the name suggests, goes straight to one or more pods and it
downloads or pushes to profefe various profiles. It is very flexible, you can
capture pprof profiles from a specific pod (or multiple pods) by ID:

```
$ kubectl profefe capture <pod-id>,<pod-id>...
```

_NB: just remember to use the namespace where the pods are running with the flag
-n or --namespace._

You can use the pod selectors to collect multiple profiles:

```
$ kubectl profefe capture -n web
```

Captures profiles from all the pod with the pprof.com/enable=true annotation
running in the pod namespace and it will store them under the `/tmp` directory.
You can change the output directory with `--output-dir`. If you do not want to
store them locally you can push them to profefe specifying its location via
`--profefe-hostport`.

The are other combinations for the capture command and you can get profiles from
profefe, I will leave the rest to you!

{:refdef:.text-center}
![](/img/stopwatch.jpg){:.img-fluild}
{:refdef}

{:.small .text-center}
Hero image via [Pixabay](https://pixabay.com/illustrations/time-time-management-stopwatch-3216244/)

## Kprofefe: the collector

The main responsability for the collector is to make the continuous profiling
magic to happen! It uses the same mechanism we already saw for the capture
kubectl plugin but it is a single binary and it can run as a cronjob.

```
apiVersion: batch/v1beta1
kind: CronJob
metadata:
  name: kprofefe-allnamespaces
  namespace: profefe
spec:
  concurrencyPolicy: Replace
  jobTemplate:
    metadata:
    spec:
      template:
        spec:
          containers:
          - args:
            - --all-namespaces
            - --profefe-hostport
            - http://profefe-collector:10100
            image: profefe/kprofefe:v0.0.8
            imagePullPolicy: IfNotPresent
            name: kprofefe
          restartPolicy: Never
          serviceAccount: kprofefe-all-namespaces
          serviceAccountName: kprofefe-all-namespaces
  schedule: '*/10 * * * *'
  successfulJobsHistoryLimit: 3
```

You can run a single cronjob that will over all the pods across all the
namespaces or you can deploy multiple cronjobs, playing with the label selector
(-l) and the namespace selector (-n) you can configure the ownership for every
running cronjob. The reasons to split in multiple cronjobs can be:

*   Scalability: one cronjob is not enough, so you can have one per namespace
    for example
*   Time segmentation: if you have a single cronjob it means that all the pods
    profiles will get captured with the same frequency, but you will may want to
    get high frequent profiles for a specific subset of applications and less
    dencity for others.

Documentation about [“Label and
Selector”](https://kubernetes.io/docs/concepts/overview/working-with-objects/labels/)
for your reference.

_Note: serviceAccount is required only if you have RBAC enabled (you should)
because the collector needs access to Kubernetes API to list/view pods across
all namespaces in this case._


## Conclusion

There is a lot to do in both the collector and kubectl plugins. I would like to
add logs and monitoring to the collector for example. The kubectl plugin get
profiles command needs some love, ideally using the same format that `kubectl
get` has via
[kubernetes/cli-runtime/pkg/printers](https://github.com/kubernetes/cli-runtime/tree/master/pkg/printers).
Try, contribute and [let me know](https://twitter.com/gianarb)!
