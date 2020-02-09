---
img: /img/golang-mockmania.png
layout: post
title: "Golang MockMania InfluxDB Client v2"
date: 2020-02-09 09:08:27
categories: [post]
tags: [golang, mockmania]
summary: ""
changefreq: daily
---

Recently I had to develop an integration with the [InfluxDB Client v2 Golang
SDK](https://github.com/influxdata/influxdb-client-go).

This SDK is useful to interact with InfluxDB v2, create organizations and users,
write new points, and submit queries; it accepts the Golang http.Client.

```golang
influx, err := influxdb.New(myHTTPInfluxAddress, myToken, influxdb.WithHTTPClient(myHTTPClient))
if err != nil {
	panic(err)
}
```

Having the ability to pass the HTTP client from the outside
`influxdb.WithHTTPClient(myHTTPClient)` improves the familiarity golang
developers have with the library; they know how to configure Transporters or how
to inject logging, tracing, debugging.  For what concerns `Golang MockMania`, it
gives to use the possibility to pass the
[httptest](https://golang.org/pkg/net/http/httptest/#example_Server) client.

```golang
influxDBServer := httptest.NewServer(http.HandlerFunc(func(rw http.ResponseWriter, req *http.Request) {

}))
influxClient, err := influxdb.New(myHTTPInfluxAddress, myToken, influxdb.WithHTTPClient(influxDBServer.Client()))
```

At this point you can write the response you expect from the influxdb server
using the `http.ResponseWriter`.

Either way, even if you have to check what influxdb receives from the sdk or if
you have to obtain a specific answer from InfluxDB to validate what your
business logic will do, nothing will stop you from using checking the
http.Request or utilizing the http.ResponseWriter to get what you expect.
