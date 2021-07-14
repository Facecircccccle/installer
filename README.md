
## EB-xxxxx
[![Build Status](https://travis-ci.com/Facecircccccle/installer.svg?branch=main)](https://travis-ci.com/Facecircccccle/installer)

EB-xxxx is a one-click installer for beginners that can install the basic Kubernetes cluster. 

Kubeadm helps users can quickly create Kubernetes clusters in line with best practices. However, they still need some configuration (IP table, swap, etc.) and installation (runtime, kubectl, etc.) as prerequisites. This project integrates these conditions with Kubeadm to help novices quickly build basic Kubernetes clusters.

This project selects Golang as the main logic and GUI, Ansible and Shell as the actual control of the remote installation process, and Kubeadm as the actual installation core.

Since it has just been developed, the current version only supports some basic features and versions, and will have more features in the future.


### Function

#### Install new cluster

Homepage -> setup -> one master setup / HA setup -> follow the prompts to input -> start -> wait for the logs -> Done.

#### Manage exist cluster

Copy the config file of the target cluster to the `~/.kube/config` directory of the deploy server, the cluster information can then be received. (For now, user can only add nodes, later more functions will be implemented.)

### Requirements

#### Deploy server

* Linux Centos 7.4+
* Golang 1.13+
* Ansible ready
> Install Ansible
```
    yum install epel-release
    yum install ansible -y
```
* SSH connect with others Kubernetes cluster servers and nodes
> Generate new SSH key in this server
```
    ssh-keygen
```
> Copy the SSH key in deploy server to all cluster server and node for password-free login
```
    ssh-copy-id <target IP>
```

#### Kubernetes cluster server and node

* Linux Centos 7.4+
* Available IP address and unique hostname

### Quick start
```
    go build
    ./installer
```

### Roadmap

* Support for more operating systems

* Support offline installation

* Support for more container runtimes

* Support for more versions of Kubernetes installation

* Support for more cluster management operations

* Better interface
