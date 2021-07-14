#!/bin/bash

mkdir  -p /root/k8s-setting/etcd
cp /etc/kubernetes/admin.conf     ~/k8s-setting
cp /etc/kubernetes/pki/{ca.*,sa.*,front-proxy-ca.*}   ~/k8s-setting
cp /etc/kubernetes/pki/etcd/ca.*   ~/k8s-setting/etcd
#cp /root/k8s-script/kube_init.log    ~/k8s-setting

