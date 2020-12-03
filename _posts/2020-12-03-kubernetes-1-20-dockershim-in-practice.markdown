---
layout: post
title:  "Kubenetes v1.20 the docker deprecation dilemma in practice"
date:   2020-12-03 10:08:27
categories: [post]
tags: [golang]
summary: "there are many discussions going on Twitter about why Kubernetes v1.20
deprecated Docker and dockershim as default runtime. But it was a well thought
an planned effort. Nothing to really worry about and here I will go over the
process of updating Kubenetes from 1.19 to 1.20"
---

Kubernetes v1.20 is not yet out but there is already a lot going on behind the
scene. The main reason is the deprecation of Docker as default runtime.

I won't go too deep in the theory because at this point I think it is a well
covered part. But a few things:

1. If you run Docker, you run containerd. That's it. Even if you didn't know, or
   you don't like the idea.
2. The container runtime interface is there since a good amount of time and the
   goal for it was to decouple the orchestrator (kubernetes), from other
   business like running containers. Who cares about containers at the end.
3. Deprecating dockershim or at least removing it from the kubelet itself is the
   right thing to do.

More about this topic:

* [Dockershim Deprecation FAQ](https://kubernetes.io/blog/2020/12/02/dockershim-faq/)
* [Don't Panic: Kubernetes and Docker](https://kubernetes.io/blog/2020/12/02/dont-panic-kubernetes-and-docker/)

I want to tell you how it works in practice. And this article contains my
experience updating a Kubernetes cluster from v1.19 to v1.20.

So I created a two node clusters on [Equinix
Metal](https://console.equinix.com/) running Ubuntu using this simple script and
kubeadm.

```bash
#!bin/bash

apt-get update
apt-get install -y vim git

apt-get install -y \
    apt-transport-https \
    ca-certificates \
    curl \
    gnupg-agent \
    vim \
    git \
    software-properties-common

releaseName=$(lsb_release -cs)
if [ $releaseName == "groovy" ]
then
    releaseName="focal"
fi

curl -fsSL https://download.docker.com/linux/ubuntu/gpg | sudo apt-key add -
apt-key fingerprint 0EBFCD88
add-apt-repository \
   "deb [arch=amd64] https://download.docker.com/linux/ubuntu \
   ${releaseName} \
   test"

apt-get update
apt-get install -y docker-ce docker-ce-cli containerd.io

apt-get update && sudo apt-get install -y apt-transport-https curl
curl -s https://packages.cloud.google.com/apt/doc/apt-key.gpg | apt-key add -
cat <<EOF | sudo tee /etc/apt/sources.list.d/kubernetes.list
deb https://apt.kubernetes.io/ kubernetes-xenial main
EOF
sudo apt-get update
sudo apt-get install -y kubelet kubeadm kubectl
sudo apt-mark hold kubelet kubeadm kubectl
```

I placed that script as CloudInit for my two nodes and I ran `kubeadm init/join`
to get my cluster.

```terminal
# kubectl get node
NAME            STATUS   ROLES    AGE     VERSION
gianarb-k8s     Ready    master   2m11s   v1.19.4
gianarb-k8s01   Ready    <none>   91s     v1.19.4
```

I have installed Flannel and now the nodes are ready. That's it. That's how I
measure success here.

It is not time to update to v1.20, so I downloaded the binaries from the
registry.

```terminal
# wget https://dl.k8s.io/v1.20.0-rc.0/kubernetes-server-linux-amd64.tar.gz
# tar xzvf ./kubernetes-server-linux-amd64.tar.gz
# ./kubernetes/server/bin/kubeadm version
kubeadm version: &version.Info{Major:"1", Minor:"20+", GitVersion:"v1.20.0-rc.0", GitCommit:"3321f00ed14e07f774b84d3198ede545c1dee697", GitTreeState:"clean", BuildDate:"2020-12-01T10:36:46Z", GoVersion:"go1.15.5", Compiler:"gc", Platform:"linux/amd64"}
```

I checked available upgrade plan from `kubeadm` with the flag
`--allow-experimental-upgrades`, if you do this process when v1.20 will be
officially released, you won't need that flag.

```terminal
./kubernetes/server/bin/kubeadm upgrade plan --allow-experimental-upgrades
[upgrade/config] Making sure the configuration is correct:
[upgrade/config] Reading configuration from the cluster...
[upgrade/config] FYI: You can look at this config file with 'kubectl -n kube-system get cm kubeadm-config -o yaml'
[preflight] Running pre-flight checks.
[upgrade] Running cluster health checks
[upgrade] Fetching available versions to upgrade to
[upgrade/versions] Cluster version: v1.19.4
[upgrade/versions] kubeadm version: v1.20.0-rc.0
[upgrade/versions] Latest stable version: v1.19.4
[upgrade/versions] Latest stable version: v1.19.4
[upgrade/versions] Latest version in the v1.19 series: v1.19.4
[upgrade/versions] Latest version in the v1.19 series: v1.19.4
I1203 21:50:13.860850   59152 version.go:251] remote version is much newer: v1.21.0-alpha.0; falling back to: stable-1.20
W1203 21:50:14.100325   59152 version.go:101] could not fetch a Kubernetes version from the internet: unable to fetch file. URL: "https://dl.k8s.io/release/stable-1.20.txt", status: 404 Not Found
W1203 21:50:14.100362   59152 version.go:102] falling back to the local client version: v1.20.0-rc.0
[upgrade/versions] Latest experimental version: v1.20.0-rc.0
[upgrade/versions] Latest experimental version: v1.20.0-rc.0
[upgrade/versions] Latest : v1.19.5-rc.0
[upgrade/versions] Latest : v1.19.5-rc.0

Components that must be upgraded manually after you have upgraded the control plane with 'kubeadm upgrade apply':
COMPONENT   CURRENT       AVAILABLE
kubelet     2 x v1.19.4   v1.20.0-rc.0

Upgrade to the latest experimental version:

COMPONENT                 CURRENT    AVAILABLE
kube-apiserver            v1.19.4    v1.20.0-rc.0
kube-controller-manager   v1.19.4    v1.20.0-rc.0
kube-scheduler            v1.19.4    v1.20.0-rc.0
kube-proxy                v1.19.4    v1.20.0-rc.0
CoreDNS                   1.7.0      1.7.0
etcd                      3.4.13-0   3.4.13-0

You can now apply the upgrade by executing the following command:

        kubeadm upgrade apply v1.20.0-rc.0 --allow-release-candidate-upgrades

_____________________________________________________________________


The table below shows the current state of component configs as understood by this version of kubeadm.
Configs that have a "yes" mark in the "MANUAL UPGRADE REQUIRED" column require manual config upgrade or
resetting to kubeadm defaults before a successful upgrade can be performed. The version to manually
upgrade to is denoted in the "PREFERRED VERSION" column.

API GROUP                 CURRENT VERSION   PREFERRED VERSION   MANUAL UPGRADE REQUIRED
kubeproxy.config.k8s.io   v1alpha1          v1alpha1            no
kubelet.config.k8s.io     v1beta1           v1beta1             no
_____________________________________________________________________
```

As you can see from the output there is `v1.20.0-rc.0` available so it is time
to apply that plan and rollout the upgrade.

```terminal
# ./kubernetes/server/bin/kubeadm upgrade apply v1.20.0-rc.0 --allow-release-candidate-upgrades
[upgrade/config] Making sure the configuration is correct:
[upgrade/config] Reading configuration from the cluster...
[upgrade/config] FYI: You can look at this config file with 'kubectl -n kube-system get cm kubeadm-config -o yaml'
[preflight] Running pre-flight checks.
[upgrade] Running cluster health checks
[upgrade/version] You have chosen to change the cluster version to "v1.20.0-rc.0"
[upgrade/versions] Cluster version: v1.19.4
[upgrade/versions] kubeadm version: v1.20.0-rc.0
[upgrade/confirm] Are you sure you want to proceed with the upgrade? [y/N]: y
...
```

Time to check kubelet with journalctl

```terminal
#journalctl -xe -u kubeletDec 03 21:58:49 gianarb-k8s kubelet[108749]: I1203 21:58:49.648539  108749 server.go:416] Version: v1.20.0-rc.0

Dec 03 21:58:49 gianarb-k8s kubelet[108749]: I1203 21:58:49.649168  108749 server.go:837] Client rotation is on, will bootstrap in background
Dec 03 21:58:49 gianarb-k8s kubelet[108749]: I1203 21:58:49.651975  108749 certificate_store.go:130] Loading cert/key pair from "/var/lib/kubelet/pki/kubelet-client-current.pem".
Dec 03 21:58:49 gianarb-k8s kubelet[108749]: I1203 21:58:49.653428  108749 dynamic_cafile_content.go:167] Starting client-ca-bundle::/etc/kubernetes/pki/ca.crt
Dec 03 21:58:49 gianarb-k8s kubelet[108749]: I1203 21:58:49.764200  108749 server.go:645] --cgroups-per-qos enabled, but --cgroup-root was not specified.  defaulting to /
Dec 03 21:58:49 gianarb-k8s kubelet[108749]: I1203 21:58:49.764717  108749 container_manager_linux.go:274] container manager verified user specified cgroup-root exists: []
Dec 03 21:58:49 gianarb-k8s kubelet[108749]: I1203 21:58:49.764743  108749 container_manager_linux.go:279] Creating Container Manager object based on Node Config: {RuntimeCgroupsName: SystemCgroupsName: KubeletCgroupsName: ContainerRuntime>
Dec 03 21:58:49 gianarb-k8s kubelet[108749]: I1203 21:58:49.764862  108749 topology_manager.go:120] [topologymanager] Creating topology manager with none policy per container scope
Dec 03 21:58:49 gianarb-k8s kubelet[108749]: I1203 21:58:49.764874  108749 container_manager_linux.go:310] [topologymanager] Initializing Topology Manager with none policy and container-level scope
Dec 03 21:58:49 gianarb-k8s kubelet[108749]: I1203 21:58:49.764881  108749 container_manager_linux.go:315] Creating device plugin manager: true
Dec 03 21:58:49 gianarb-k8s kubelet[108749]: W1203 21:58:49.765010  108749 kubelet.go:297] Using dockershim is deprecated, please consider using a full-fledged CRI implementation
```

We finally got it `Using dockershim is deprecated, please consider using a
full-fledged CRI implementation`. But everything still works.

```terminal
# kubectl get node
NAME            STATUS   ROLES                  AGE   VERSION
gianarb-k8s     Ready    control-plane,master   17m   v1.20.0-rc.0
gianarb-k8s01   Ready    <none>                 16m   v1.19.4
```

There is something cool as well, now the role is `control-plane` and `master`, I
am sure at some point we will deprecate `master` as well and this is just a
transition phase, same as it happened for Docker Shim.

We don't want that working, so it is not time to change the default
configuration of `containerd`, because as you know, it is there sitting behind
docker since forever almost.

```terminal
cat /etc/containerd/

# cat /etc/containerd/config.toml
... 
# comment this line because we need cri enable
# disabled_plugins = ["cri"]
...
```

We need to enable the plugin `cri`, by default it is disabled when installing
docker-ce via the Docker registry because `dockerd` does not need a CRI, but
Kubernetes needs it obviously. Now you can restart the service with `systemctl`
and we have to tell the kubelet that now it has to use `containerd`.

```terminal
# mkdir -p /etc/systemd/system/kubelet.service.d/
cat << EOF | sudo tee  /etc/systemd/system/kubelet.service.d/0-containerd.conf
[Service]
Environment="KUBELET_EXTRA_ARGS=--container-runtime=remote --runtime-request-timeout=15m --container-runtime-endpoint=unix:///run/containerd/containerd.sock"
EOF
```

After a systemd daemon reload and a kubelet.service restart we are back again.

```
# kubectl get node
NAME            STATUS   ROLES                  AGE   VERSION
gianarb-k8s     Ready    control-plane,master   23m   v1.20.0-rc.0
gianarb-k8s01   Ready    <none>                 23m   v1.19.4
```

Exercise for you, have a look at the kubelet logs, the warning is not there
anymore and you are good to go.
