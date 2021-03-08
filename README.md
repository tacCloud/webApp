# Inventory webpage

This is a simple web application that allow users to view and "buy" items from an inventory.
The inventory backend is in the form of a mysql db.

The code is stupid simple, I just did this to explore:

* Golang
* Jenkins (to test this repo in isolation)
* ArgoCd (to test )


See this repo for a good example of using Jenkins and golang
https://levelup.gitconnected.com/automating-build-and-test-of-go-application-with-jenkins-9f96879b5479
https://github.com/wilsontwm/go_simple_rest/

Here are some instructions on installing Jenkins
https://www.digitalocean.com/community/tutorials/how-to-install-jenkins-on-kubernetes


Notes:

* Need to add the Go plugin to Jenkins
* I'm not sure how to build docker images from Jenkins ... Looks like you either have to [configure docker to expose its API](https://plugins.jenkins.io/docker-plugin/) or


git clone --depth 1 --branch v4.19.121-linuxkit https://github.com/linuxkit/linux 4.19.121-linuxkit
docker build --target dev . -t go
docker run --privileged --cap-add=ALL -it -v /lib/modules:/lib/modules \
   -v ${PWD}:/work \
   -v $PWD/4.19.121-linuxkit:/usr/src/linux-headers-4.19.121-linuxkit \
   -v /etc/localtime:/etc/localtime:ro \
   go sh
go version

  apk info bcc
  apk add bcc
  apk add bcc-dev
  apk add musl-dev
  apk add linux-headers
  apk info -L bcc
  apk add bcc-tools
  history | grep apk

----

Redis testing
docker run --name redis-test-instance -p 6379:6379 -d redis

docker run -it --rm redis redis-cli -h 172.17.0.2

