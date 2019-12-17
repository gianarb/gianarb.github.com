---
img: /img/gianarb.png
layout: post
title: "kubectl flags in your plugin"
date: 2019-10-07 08:08:27
categories: [post]
tags: [kubernetes, flags, kubectl, plugin, kubeassemble]
summary: "Develop cure custom kubectl plugins with friendly flags from the kubectl"
changefreq: daily
---
This is not at all a new topic, no hacking involved, but it is something
everybody needs to know where we design kubectl plugin.

I was recently working at one and I had to make the user experience as friendly
as possible compared with the `kubectl`, because that's what a good developer does!
Tricks other developers to make their life comfortable, if you are used to do:

```bash
$ kubectl get pod -n your-namespace -L app=http
```

To get pods from a particular namespace `your-namemespace` filtered by label
`app=http` and your plugin does something similar or it will benefit from an
interaction that remembers the classic `get` you should re-use those flags.

Example: design a `kubectl-plugin` capable of running a `pprof` profile against a
set of containers.

My expectation will be to do something like:

```bash
$ kubectl pprof -n your-namespace -n pod-name-go-app
```

The Kubernetes community writes a lot of their code in Go, it means that there
are a lot of libraries that you can re-use.

[kubernetes/cli-runtime](https://github.com/kubernetes/cli-runtime/tree/master/pkg/genericclioptions)
is a library that provides utilities to create kubectl plugins. One of their
packages is called `genericclioptions` and as you can get from its name the goal
is obvious.

```go

// import "github.com/spf13/cobra"
// import "github.com/spf13/pflag"
// import "k8s.io/cli-runtime/pkg/genericclioptions"

// Create the set of flags for your kubect-plugin
flags := pflag.NewFlagSet("kubectl-plugin", pflag.ExitOnError)
pflag.CommandLine = flags

// Configure the genericclioptions
streams := genericclioptions.IOStreams{
    In:     os.Stdin,
    Out:    os.Stdout,
    ErrOut: os.Stderr,
}

// This set of flags is the one used for the kubectl configuration such as:
// cache-dir, cluster-name, namespace, kube-config, insecure, timeout, impersonate,
// ca-file and so on
kubeConfigFlags := genericclioptions.NewConfigFlags(false)

// It is a set of flags related to a specific resource such as: label selector
(-L), --all-namespaces, --schema and so on.
kubeResouceBuilderFlags := genericclioptions.NewResourceBuilderFlags()

var rootCmd = &cobra.Command{
    Use:   "kubectl-plugin",
    Short: "My root command",
    Run: func(cmd *cobra.Command, args []string) {
		cmd.SetOutput(streams.ErrOut)
    }
}

// You can join all this flags to your root command
flags.AddFlagSet(rootCmd.PersistentFlags())
kubeConfigFlags.AddFlags(flags)
kubeResouceBuilderFlags.AddFlags(flags)
```

This is the output:

```bash
$ kubectl-plugin --help
My root command

Usage:
  kubectl-plugin [flags]

Flags:
      --as string                      Username to impersonate for the operation
      --as-group stringArray           Group to impersonate for the operation, this flag can be repeated to specify multiple groups.
      --cache-dir string               Default HTTP cache directory (default "/home/gianarb/.kube/http-cache")
      --certificate-authority string   Path to a cert file for the certificate authority
      --client-certificate string      Path to a client certificate file for TLS
      --client-key string              Path to a client key file for TLS
      --cluster string                 The name of the kubeconfig cluster to use
      --context string                 The name of the kubeconfig context to use
  -f, --filename strings               identifying the resource.
  -h, --help                           help for kubectl-pprof
      --insecure-skip-tls-verify       If true, the server's certificate will not be checked for validity. This will make your HTTPS connections insecure
      --kubeconfig string              Path to the kubeconfig file to use for CLI requests.
  -n, --namespace string               If present, the namespace scope for this CLI request
  -R, --recursive                      Process the directory used in -f, --filename recursively. Useful when you want to manage related manifests organized within the same directory. (default true)
      --request-timeout string         The length of time to wait before giving up on a single server request. Non-zero values should contain a corresponding time unit (e.g. 1s, 2m, 3h). A value of zero means don't timeout requests. (default "0")
  -l, --selector string                Selector (label query) to filter on, supports '=', '==', and '!='.(e.g. -l key1=value1,key2=value2)
  -s, --server string                  The address and port of the Kubernetes API server
      --token string                   Bearer token for authentication to the API server
      --user string                    The name of the kubeconfig user to use
```
