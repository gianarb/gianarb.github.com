---
layout: post
title:  "Chef Server startup notes"
date:   2016-11-10 10:08:27
categories: [post]
img: /img/chef.png
tags: [chef, devops]
summary: "Docker The fundamental is the second chapter of my book Scale Docker.
Drive your boat like a captain. I decided to share free the second chapter of
the book. It covers getting started with Docker. It's a good tutorial for
people that are no idea about how container means and how docker works."
changefreq: yearly
---
I worked with different provisioning tools and configuration managers in the
last couple of years: Chef, Saltstack, Puppet, Shell, Python, Terraform.
Everything that was allowing me to make automation and describe my
infrastructure as a code.

I really think that this is the correct street and every companies need to stop
to persist random commands in a server:

* The code used to describe your infrastructure is reausable.
* The code is a good backup and you can put it in your repository to study of
  it changed and manage rollbacks.
* Your servers become collaborative and your team can review what you do.

Chef is my first configuration manager, I started to use it with Vagrant few
years ago but I never had change to deep drive into it and into the full
chef-server configuration from scratch.

I had this change few days ago and I am here to share some notes. I used
digitalocean to start 1 Chef Server and two nodes, during this post I am not
focused about the recipe and cookbook syntax but I will share some commands and
notes that I took during my test to start and configure a Chef Server.

First of all `doctl` is the command line application provided by digitalocean
to manage droplets and everything, I used that tool to start my droplets.

The Chef Server doesn't run in little box, we need 2gb RAM, I tried with small
size but nothing was working, the installation process gone out of memory very
soon. Thanks Ruby.

```sh
$ doctl compute droplet create chef-server \
  --region ams2 --size 2gb --image 20385558 \
  --access-token $DO --ssh-keys $DO_SSH

$ doctl compute droplet create n1 \
  --region ams2 --size 512mb --image 20385558 \
  --access-token $DO --ssh-keys $DO_SSH
```
`$DO` contains my digitalocean access key and `$DO_SSH` the id of the ssh key
to log into the servers. You can leave the last one empty and you will receive
an email with the password.

When the process is gone you will be able to copy the ip of the chef-server and go into it.

```bash
$ doctl compute droplet ls

ID              Name            Public IPv4     Public IPv6     Memory  VCPUs   Disk    Region  Image           Status  Tags
30cw4230        chef-server                                     2gb     1       20      ams2    Debian 8.6 x64  new
q0514230        n1                                              512     1       20      ams2    Debian 8.6 x64  new
```

This provisioning script installs Chef-Server from the official deb package and
also install chef-manage.  Chef-manage provides a nice web interface to manage
users, cookbooks and everything is stored into the server.

```bash
cd /tmp
sudo apt-get update
sudo apt-get install -y unzip curl
curl -LS https://packages.chef.io/stable/ubuntu/16.04/chef-server-core_12.9.1-1_amd64.deb -o chef-server-core_12.9.1-1_amd64.deb
dpkg -i chef-server-core_12.9.1-1_amd64.deb
chef-server-ctl reconfigure
curl -SL https://packages.chef.io/stable/ubuntu/16.04/chef-manage_2.4.3-1_amd64.deb -o chef-manage.deb
dpkg -i chef-manage.deb
chef-manage-ctl reconfigure --accept-license
```

After this you are enable to see the interface over https in your browser, you
can use the IP of your server. I need to remember you that this is just an
example, the server is public and this is not a good practice AT ALL! You need
to configure this service in your VPN to make it totally private.

```bash
chef-server-ctl user-create gianarb Gianluca Arbezzano ga@thumpflow.com 'hellociaobye' --filename /root/gianarb_test.pem
chef-server-ctl org-create tf 'ThumpFlow' --association_user gianarb --filename /root/tf-validator.pem
chef-server-ctl org-user-add tf gianarb
```

Chef-Server works with the concept of Organization and User. The organization
is a group of users that share cookbooks, rules and so on.  Users can update
cookbooks and there is also a set of permission to manage access on particular
resources like:

* Add a new node
* Syncronize cookbook with the server
* add new users

At this point we have one user with its own key and credential. You can come
back into  the UI and use username (gianarb) and password (hellociaobye) to
login in.  The key (--filename) is used to configure knife and encrypt
communication between client and server.  There are 3 main actors and this
point that we need to know:

* Chef-Server contains all our recipes, cookbooks and it's the brain of the cluster.
* Nodes are all servers configurated by Chef.
* Workstation are usually enable to syncronize, update cookbooks. For example
Jenkins or your Continuous Integration System after every new commit can push
  every changes into the server.

Chef Server has a HTTP api and `knife` is a CLI that provide an easy
integration for your node and workstation.  With this command we are installing
knife. You can do it in your local environment, to become a workstation and into
the server. (it's usually a god practice create a user, we are doing everything
as root right know but it's BAD! don't be bad!).

We have two certificate one is `gianarb_test.pem` and it's identify a specific
user, we need to generate our for every workstation/member of the team and the
`validation_client` represent the organization, it could be the same across
multiple users.

```bash
curl -O -L http://www.opscode.com/chef/install.sh
bash ./install.sh
```

You can copy paste the 2 keys into the local machine and run this command that
will drive your into the process to create a `~/.chef/knife.rb` file that your
cli uses to communicate with the chef server.

```bash
knife configure
```

This is an example of generated knife configuration file that I did in my
server.  I lose times to understand the `chef_server_url` it contains the
hostname of the server but also the `/organization/<organization_short_name>'
be careful about this or knife will come back with an HTML response in your terminal.

```ruby
log_level                :info
log_location             STDOUT
node_name                'gianarb'
client_key               '/root/gianarb_test.pem'
validation_client_name   'tf-validator'
validation_key           '/root/tf-validator.pem'
chef_server_url          'https://chef-server:443/organizations/tf'
syntax_check_cache_path  '/root/.chef/syntax_check_cache'
cookbook_path []
ssl_verify_mode          :verify_none
```

#BAAAAAAAAAAAAA
The last 2 commands download and validate the SSH certificate because in the
default configuration the CA is unofficial and we need to force our client to
trust the cert.

```bash
knife ssh fetch
knife ssh check
```

Know that we did that in our server and also in our local environment we can
clone chef-pluto a repository that contains recipes, rules and cookbooks to
configure our node, we need to syncronize it into the server.

```bash
git clone git@github.com:gianarb/chef-pluto.git ~/chef-pluto
cd ~/chef-pluto
knife update .
```
The last command update all our repository into the chef server. You can log in
into the web ui and see the `micro` cookbook and the `power` rule.

[micro](https://github.com/gianarb/micro) is an application that I wrote in go and it just expose the ip of the
machine. It's a binary and the cookbook downloads and starts it, pretty
straightforward.

At this point we need to make a provisioning of our first node, usually is the
server that install and start the Chef Client into the node, what we can do
it's store a private key into the server to allow chef to connect to the node.
I copied the digitalocean private key into the server (~/do), from security
point of view you can create a dedicate one. You can also use the -P option if
you are not using an ssh-key to run this example.

```bash
knife bootstrap <ip-node> -N node1 --ssh-user root -r 'role[power]' -i ~/do
```

If everything it's good you can reach the application from port `8000` into the
browser. The log is something like:

```bash
$ knife bootstrap 95.85.52.211 -N testNode --ssh-user root -r 'role[power]' -i ~/do
Doing old-style registration with the validation key at /root/tf-validator.pem...
Delete your validation key in order to use your user credentials instead

Connecting to 95.85.52.211
95.85.52.211 -----> Existing Chef installation detected
95.85.52.211 Starting the first Chef Client run...
95.85.52.211 Starting Chef Client, version 12.15.19
95.85.52.211 resolving cookbooks for run list: ["micro"]
95.85.52.211 Synchronizing Cookbooks:
95.85.52.211   - micro (0.1.0)
95.85.52.211 Installing Cookbook Gems:
95.85.52.211 Compiling Cookbooks...
95.85.52.211 Converging 2 resources
95.85.52.211 Recipe: micro::default
95.85.52.211   * remote_file[Download micro] action create_if_missing (up to date)
95.85.52.211   * service[Start micro] action start
95.85.52.211     - start service service[Start micro]
95.85.52.211
95.85.52.211 Running handlers:
95.85.52.211 Running handlers complete
95.85.52.211 Chef Client finished, 1/2 resources updated in 02 seconds
```

knife started the client, syncronized cookbooks, it assigned the `power` role
at the node and run the correct recipes.  Your server is ready and you can
create and delete nodes to make your infrastructure complex how much you like.

Chef is quite old and it's in ruby (the first one could be a plus but the
second one no really) but it continue to be a good way to make a provisioning o
your infrastructure. Lots of people moved to Ansible but the agent that they
reject offer a very good orchestration feature that it's something that I
usually search.

I worked with StalStack and it's very nice, the syntax is easy
and it seems less expensive in terms of configuration, resources and setup but
I am not really sure about the YAML specification. I am not a ruby developer
and I don't love the ruby syntax but in the end is a programming languages and
I am doing infrastructure as a code.
