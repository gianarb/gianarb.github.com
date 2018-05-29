---
layout: post
title:  "Cloud Native Intranet with Kubernetes, CoreDNS and OpenVPN"
date:   2018-05-29 10:38:27
img: /img/kubernetes.png
categories: [post]
tags: [devops, kubernetes, security, network, intranet, dns, vpn, openvpn, coredns]
summary: "Designing an architecture the network should be a top priority because
it is very hard to change moving forward. Even in a cloud environment running on
Kubernetes the situation doesn't change. Security and networking are hard
pattern hard to inject in old projects. In this talk I will share a practical
idea about how to start in the best way with OpenVPN and private DNS in a
Kubernetes cluster in order to build your own intranet."
priority: 1
---
This article has a marketing and buzzword oriented title. I know.

Let me introduce you to what I am going to speak with you here with better
worlds: VPN, private, DNS, kubernetes, security.

I hope we all agree that VPN should be a must-have when you set up an
infrastructure. It doesn't matter what you are doing, how many people are
working with you.

When you design a new system usually you need to expose to the public only some
service over HTTP and HTTPS all the rest: Jenkins, monitoring tools,
dashboards, log management should be locked-down and accessible just in a
private network. An intranet.

> An intranet is a private network accessible only to an organization's staff.
Often, a wide range of information and services are available on an
organization's internal intranet that is unavailable to the public, unlike the
Internet.

All these concepts apply to "Cloud Native" ecosystem as well.

Kubernetes has a powerful dashboard and CTL that you can use to interact with
the API. That API doesn't need to be publicly exposed, and to use the CLI from
your laptop, you should set up a VPN.

## OpenVPN
Usually, I configure an OpenVPN using the image
[kylemanna/openvpn](https://hub.docker.com/r/kylemanna/openvpn/) available on
Docker Hub. It is straightforward to apply, and it offers a set of utilities
around user creation and certification management.

```yml
apiVersion: v1
kind: Service
metadata:
  name: openvpn
  namespace: openvpn
  labels:
    app: openvpn
spec:
  ports:
  - name: openvpn
    nodePort: 1194
    port: 1194
    protocol: UDP
    targetPort: 1194
  selector:
    app: openvpn
  type: NodePort
---
apiVersion: apps/v1
kind: Deployment
metadata:
  name: openvpn
  namespace: "openvpn"
  labels:
    app: openvpn
spec:
  replicas: 1
  strategy:
    type: Recreate
  selector:
    matchLabels:
      app: openvpn
  template:
    metadata:
      labels:
        app: openvpn
    spec:
      nodeSelector:
        role: vpn
      containers:
        - name: openvpn
          image: docker.io/kylemanna/openvpn
          command: ["/etc/openvpn/setup/configure.sh"]
          env:
            - name: VPN_HOSTNAME
              valueFrom:
                configMapKeyRef:
                  name: vpn-hostname
                  key: hostname
            - name: VPN_DNS
              valueFrom:
                configMapKeyRef:
                  name: vpn-hostname
                  key: dns
          ports:
            - containerPort: 1194
              name: openvpn
          securityContext:
            capabilities:
              add:
                - NET_ADMIN
          volumeMounts:
            - mountPath: /etc/openvpn/setup
              name: openvpn
              readOnly: false
            - mountPath: /etc/openvpn/certs
              name: certs
              readOnly: false
      volumes:
        - name: openvpn
          configMap:
            name: openvpn
            defaultMode: 0755
        - name: certs
          persistentVolumeClaim:
            claimName: openvpncerts
```
I put the persistentVolumeClaim to remember you to store in a persisted and
safe place (and you should backup them too) the certificates used and generated
by the VPN `/etc/openvpn/certs`.

I won't write more about this topic; we are all excellent yaml developers!

How to create users, configuration and so on is a well knows topic that you can
easily [find in the OpenVPN's
documentation](https://openvpn.net/index.php/open-source/documentation.html).

I don't know if you realized that, but this VPN runs inside a Kubernetes
Cluster, so well configurated allow us to reach pods via a private network and
a bonus point also via kubedns to ping services, pods and all the other
resources registered to it.

To do that OpenVPN server can be configured to push kubedns to the client:
```
dhcp-option DNS <kube-dns-ip>
```
Something with learned is that if you are using Linux the
NetworkManager-OpenVPN plugin pushes the DNS correctly, but the OpenVPN cli
tool doesn't if you are using the last one you need to set it up in another
way.

Tips: You can take the `<kube-dns-ip>` doing `cat /etc/resolv.conf` from inside a pod.

## DNS
Push the KubeDNS or the DNS used by kubernetes is not enough to have a complete
intranet. You should be able to set up a custom domain to have friendly or
short URL.

You can take two different directions. KubeDNS can have static
record configured, but some person is not happy to touch or customize too much
the KubeDNS because Kubernetes itself use it and if you mess it up all it can
be a problem.

A possible solution is to deploy another DNS like CoreDNS and
configures it to resolve KubeDNS as a fallback. In this way, you will be free
to register custom LTDs and records. Kubernetes is going to use KubeDNS as
usual, and if you mess up CoreDNS, only a fraction of your system will blow
out.

Naturally to resolve your custom domains from the VPN you need to push
the CoreDNS ip and not the one used by Kubernetes.

If two DNSs are too much take the option one or from Kubernetes 1.10 you can
use CoreDNS as kubernetes DNS so it is a bit more flexible and you can use only
that one if you are brave enough.

I suggested CoreDNS because it supports records configuration via
[etcd](https://github.com/coredns/coredns/tree/master/plugin/etcd). Here an
example of Corefile:

```
. {
      errors
      etcd *.myinternal {
          stubzones
          path /skydns
          entrypoint  http://etcd-1:2379,http://etcd-2:2379,http://etcd-3:2379
          upstream /etc/resolv.conf
      }
      proxy . /etc/resolv.conf
}
```
Running this configuration inside a pod automatically fallback to kubedns (that
automatically fallback to the one configured to reach internet). Because of
`upstream` point to `resolv.conf` that inside a pod contains kubedns.

## Benefits
Resolve Kubernetes DNS record from your local environment is very comfortable
to build a shared or dynamic development environment for you and your
colleagues.

You can set up per-developer namespaces that they can use to
deploy services reachable from the program that they are writing. Or you can
deploy your application, and another person connected to the VPN will be able
to use it.
