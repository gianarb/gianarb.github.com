---
heroimg: /img/hero/infra-as-code.jpeg
layout: post
title:  "Infrastructure as (real) code"
date:   2018-12-31 08:08:27
categories: [post]
tags: [devops, infrastructure as code, kubernetes, golang, yaml, ansible, chef,
puppet, salt, helm]
summary: "Infrastructure as code today is wrong. Tools like Chef, Helm, Salt,
Ansible uses a template engine to make YAML or JSON way to smarter, but
comparing this solution with a proper coding language you always miss something.
GitOps forces you to stick your infrastructure code in a git repository this is
good. But infrastructure as code is way more."
changefreq: daily
---
I got different signals from the internet around the topic infrastructure as
code. I worked with a lot of configuration management tools: Chef, Ansible,
Salt. All of them are good and bad almost in the same way, for me it is mainly a
boring syntax switch between them. That's one of the reasons I have a repulsion
for these kind of tools.  This year at InfluxData we moved to Kubernetes, and I
had the chance to see how a migration like that works, and the unique
privileges to work with my collagues to design how the end result looks
like, even if it is a never ending work in progress based on the feedback
that we get from outself and other teams.  So I think at this point I can
try to explain why I think infrastructure as code today doesn't work.

<blockquote class="twitter-tweet tw-align-center" data-lang="en"><p lang="en" dir="ltr">I’m
starting to think the industry didn’t get the point of “infrastructure as code”.
That people believe codified infrastructure is checking YAMLs into a git repo is
troubling.</p>&mdash; Dan Woods (@danveloper) <a
href="https://twitter.com/danveloper/status/1078870433246662656?ref_src=twsrc%5Etfw">December
29, 2018</a></blockquote> <script async
src="https://platform.twitter.com/widgets.js" charset="utf-8"></script>

Configuration management are not entirely useless, but it is like [learning a
new framework](https://sizovs.net/2018/12/17/stop-learning-frameworks/), there
is always something good to learn but it is just a framework.  If you pick the
cooler Javascript one, you will probably get a well-paid job in a startup with
candies and a flexible workplace but I am always interested in learning the
underline architecture and patterns. The reconciliation loop that ReactJS built
to interact with the DOM is pretty nice, or the one that Kubernetes has to
manage all the resources.  Architecture, design patterns are well more useful
that syntactic sure that you can get from the framework itself even more when
the "sugar" looks like this:

```yaml
- name: "(Install: All OSs) Install NGINX Open Source Perl Module"
  package:
    name: nginx-module-perl
    state: present
  when: nginx_type == "opensource"
- name: "(Install: All OSs) Install NGINX Plus Perl Module"
  package:
    name: nginx-plus-module-perl
    state: present
  when: nginx_type == "plus"
- name: "(Setup: All NGINX) Load NGINX Perl Module"
  lineinfile:
    path: /etc/nginx/nginx.conf
    insertbefore: BOF
    line: load_module modules/ngx_http_perl.so;
  notify: "(Handler: All OSs) Reload NGINX"
```

The above code is an Ansible script that I took from a randomly from the [nginx
role](https://github.com/nginxinc/ansible-role-nginx/blob/master/tasks/modules/install-perl.yml)

```yaml
piVersion: extensions/v1beta1
kind: Deployment
metadata:
  name: { template "drone.fullname" . }}-agent
  labels:
    app: { template "drone.name" . }}
    chart: "{ .Chart.Name }}-{ .Chart.Version }}"
    release: "{ .Release.Name }}"
    heritage: "{ .Release.Service }}"
    component: agent
spec:
  replicas: { .Values.agent.replicas }}
  template:
    metadata:
      annotations:
        checksum/secrets: { include (print $.Template.BasePath "/secrets.yaml") . | sha256sum }}
{- if .Values.agent.annotations }
{ toYaml .Values.agent.annotations | indent 8 }
{- end }
      labels:
        app: { template "drone.name" . }}
        release: "{ .Release.Name }}"
        component: agent
```

This is an help chart I took from the [official GitHub
repository](https://github.com/helm/charts/blob/master/stable/drone/templates/deployment-agent.yaml).

To be clear, when I imagine a sweet dessert full of sugar it is way different
compared with what I have pasted above.

Both of them work with a template engine that is capable of rendering a template
that looks like YAML.  I will never buy that infrastructure as code doesn't use
the real code but a serialization language.

If you don't know why YAML or JSON or HCL these are a set of reasons that you
will hear:

* The curve to learn YAML, JSON, HCL is way more friendly than a proper language
  like Go, Javascript, PHP or whatever.
* You don't have all the utilities that a language provides but only what the
  template engine exposes. This should help you and your team to avoid terrible
  mistakes.

These concerns was reasonably at the beginning, when the DevOps culture started,
but now everyone has a good sense of how to code.  We do code review, and we
have a lot more experience around patterns and API to handle infrastructure
provisioning.

1. If you know Kubernetes, it has powerful API that you can leverage to write
   automation code, same for cloud provider like AWS, GCP or OpenStack.
2. Reconciliation loop, informer, Workqueue, Controller and CRDs are concepts
   from Kubernetes that you can reuse.
3. I wrote about [reactive
   planning](https://gianarb.it/blog/reactive-planning-is-a-cloud-native-pattern)
   and its application in cloud.

<blockquote class="twitter-tweet tw-align-center" data-lang="en"><p lang="en" dir="ltr">if people refuse to learn things, fire them.<br>if your management won&#39;t fire people for not pulling their weight, quit.<br><br>ENGINEERS: we live in a golden age of opportunity.  please use it while it lasts. <a href="https://t.co/OdB24UNl9X">https://t.co/OdB24UNl9X</a></p>&mdash; Charity Majors (@mipsytipsy) <a href="https://twitter.com/mipsytipsy/status/1078799382009470979?ref_src=twsrc%5Etfw">December 28, 2018</a></blockquote>
<script async src="https://platform.twitter.com/widgets.js" charset="utf-8"></script>

All the concerns that I raised in favor of `YAML, JSON vs. code` drives to the
risk of writing bad code, but I think there is no way to "remove bad code." Even
code that looks good today will look bad tomorrow.  Find a way to mitigate the
risk is admirable but I don't think YAML is the right solution, a code
architecture, the right patterns, testing, documentation and code review are the
way to go.

Today therea are people with the right skills to write good code even around
infrastructure, and if you use real code you will have:

* A reacher set of libraries and tools based on the language that you will pick.
* Unit, integration test frameworks.
* Compiling or interpreting an actual language will highlight more syntax errors
  that every template engine.
* Code is way more fun!
* You can import your code and you don't need to make trick things to join
  Kubernetes template together
* You can instantiate new objects, apply transformations of them from the same
  structure to reuse the code that describes your resources (aws autoscaling
  group, kubernetes ingress or whatever).

This discussion applied to a real word situation with Kubernetes used not via
YAML but with the Go struct provided by the
[kubernetes/client-go](https://github.com/kubernetes/client-go/tree/master/kubernetes/typed/core/v1)

```yaml
apiversion: apps/v1
kind: deployment
metadata:
  name: micro
  namespace: micro
  labels:
    app: micro
    component: micro
spec:
  replicas: 12
  selector:
    matchlabels:
      app: micro
  template:
    metadata:
      labels:
        app: micro
    spec:
      containers:
      - name: microapp
        image: gianarb/micro
        ports:
        - containerport: 8080
        env:
        - name: SLACK_TOKEN
          valuefrom:
            secretkeyref:
              name: slack
              key: token
        - name: SLACK_USERNAME
          value: "myuser'
        resources:
          limits:
            memory: 128mi
          requests:
            memory: 100mi
```

This YAML translated to Golang:

```golang
func newMicroDeployment() *appsv1.Deployment {
    return &appsv1.Deployment{
        TypeMeta: metav1.TypeMeta{
            Kind:       "Deployment",
            APIVersion: "apps/v1",
        },
        ObjectMeta: metav1.ObjectMeta{
            Name:      "micro",
            Namespace: twodotoh.Namespace,
            Labels: map[string]string{
                "app":       "micro",
                "component": "micro",
            },
        },
        Spec: appsv1.DeploymentSpec{
            Replicas: pointer.Int32Ptr(12),
            Selector: &metav1.LabelSelector{
                MatchLabels: map[string]string{
                    "app": "micro",
                },
            },
            Template: corev1.PodTemplateSpec{
                ObjectMeta: metav1.ObjectMeta{
                    Labels: map[string]string{
                        "app": "micro",
                    },
                },
                Spec: corev1.PodSpec{
                    Containers: []corev1.Container{
                        {
                            Name:  "microapp",
                            Image: "gianarb/micro",
                            Ports: []corev1.ContainerPort{
                                {
                                    ContainerPort: 8080,
                                },
                            },
                            Env: []corev1.EnvVar{
                                {
                                    Name: "SLACK_TOKEN",
                                    ValueFrom: &corev1.EnvVarSource{
                                        SecretKeyRef: &corev1.SecretKeySelector{
                                            LocalObjectReference: corev1.LocalObjectReference{
                                                Name: "slack",
                                            },
                                            Key: "token",
                                        },
                                    },
                                },
                                {
                                    Name:  "SLACK_USERNAME",
                                    Value: "myuser",
                                },
                            },
                        },
                    },
                },
            },
        },
    }
}
```

You can make the function more flexible passing variables like the number of
replicas for example, or you can write transformation function that looks like
`WithDifferentMemoryLimit` to apply transformation to your `runtime.Object`.

```golang
deployment := newMicroDeployment()

// You can transform them with utils like:
WithDifferentMemoryLimit("200mi", deployment)
```

If you play well will Go packages, and if you structure your code you can have
something like:

```golang
apps := []*runtime.Object{}
service := micro.NewKubernetesService()
deployment := micro.NewDeployment()
apps = append(apps, service)
apps = append(apps, deployment)
// Deploy via kubernetes api
```
I mean, you have the code now! So you can make all the mistakes you usually do
during your daily job!

{:.small}
Hero image via [Pixabay](https://pixabay.com/en/fractal-complexity-geometry-1758543/)
