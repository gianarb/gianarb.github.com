---
img: /img/kubernetes.png
layout: post
title: "Unit test kubernetes client in Go"
date: 2020-01-10 09:08:27
categories: [post]
tags: [golang, kubernetes]
summary: "A flexible an easy to use testing framework makes all the difference.
Kubernetes provides a fake client in Go that works like a charm."
changefreq: daily
---

I write a lot of operations and integrations with Kubernetes those days. You can
follow my journey in its dedicated section on this blog ["Building
Kubernetes"](/planet/assemble-kubernetes.html).

I had to write a function recently capable of filtering pods based on assigned
annotations.

```go

const (
	ProfefeEnabledAnnotation = "profefe.com/enable"
)

// GetSelectedPods returns all the pods with the profefe annotation enabled
// filtered by the selected labels
func GetSelectedPods(clientset kubernetes.Interface,
	namespace string,
	listOpt metav1.ListOptions) ([]v1.Pod, error) {

	target := []v1.Pod{}
	pods, err := clientset.CoreV1().Pods(namespace).List(listOpt)
	if err != nil {
		return target, err
	}
	for _, pod := range pods.Items {
		enabled, ok := pod.Annotations[ProfefeEnabledAnnotation]
		if ok && enabled == "true" && pod.Status.Phase == v1.PodRunning {
			target = append(target, pod)
		}
	}
	return target, nil
}
```
This function is pretty easy, but it has a good amount of assertions that we can
check. Even more when we have a so well scoped functions writing tests should be
almost mandatory.

* The returned list of pods should only contains pods with the
  `ProfefeEnabledAnnotation` set
* The returned list of pods should only returns pods from the specified
  `namespace`
* The returned list of pods should observe the filtering and label selection
  criteria specified by `metav1.ListOptions`

Covering those use cases will give us a solid foundation to avoid regression
when this function will get more complicated (usually that's the evolution for
successful piece of code).

## Kubernetes Client Mock

Kubernetes offers a simple and powerful `fake` client that has a very efficient
mechanism to simulate the desired output from a specific request, in our case
`clientset.CoreV1().Pods(namespace).List(listOpt)`. You have to pass the slice
of `runtime.Object` you desire when you create a new fake client. Awesome and
easy.

```go
clientset: fake.NewSimpleClientset(&v1.Pod{
    ObjectMeta: metav1.ObjectMeta{
        Name:        "influxdb-v2",
        Namespace:   "default",
        Annotations: map[string]string{},
    },
}, &v1.Pod{
    ObjectMeta: metav1.ObjectMeta{
        Name:        "chronograf",
        Namespace:   "default",
        Annotations: map[string]string{},
    },
}),
```
For example this `clientset` will return two pods, one called `influxdb-v2` and
one called `chronograf`, but you can return what ever you need: Services,
Deployments, Ingress, Custom Resource Definition or even a mix of everything.

## In practice

I wrote a bunch of tests for
[kube-profefe](https://github.com/profefe/kube-profefe/blob/master/pkg/kubeutil/kube_test.go)
that are using a fake client. You can get inspiration over there.

## Conclusion

`fake` client is easy to use, so easy that since I added it in my tool chain for
some functions like the one I described here I efficiently do `TDD` because it
makes the iteration over my code way faster.
