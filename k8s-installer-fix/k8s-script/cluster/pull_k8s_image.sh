#!/bin/bash
# kubernetes
#docker pull registry.cn-hangzhou.aliyuncs.com/google_containers/kube-apiserver:v1.20.5
#docker pull registry.cn-hangzhou.aliyuncs.com/google_containers/kube-proxy:v1.20.5
#docker pull registry.cn-hangzhou.aliyuncs.com/google_containers/kube-scheduler:v1.20.5
#docker pull registry.cn-hangzhou.aliyuncs.com/google_containers/kube-controller-manager:v1.20.5
#docker pull registry.cn-hangzhou.aliyuncs.com/google_containers/etcd:3.4.13-0
#docker pull registry.cn-hangzhou.aliyuncs.com/google_containers/coredns:1.7.0
#docker pull registry.cn-hangzhou.aliyuncs.com/google_containers/pause:3.2
#
#docker tag registry.cn-hangzhou.aliyuncs.com/google_containers/kube-apiserver:v1.20.5 k8s.gcr.io/kube-apiserver:v1.20.5
#docker tag registry.cn-hangzhou.aliyuncs.com/google_containers/kube-proxy:v1.20.5 k8s.gcr.io/kube-proxy:v1.20.5
#docker tag registry.cn-hangzhou.aliyuncs.com/google_containers/kube-scheduler:v1.20.5 k8s.gcr.io/kube-scheduler:v1.20.5
#docker tag registry.cn-hangzhou.aliyuncs.com/google_containers/kube-controller-manager:v1.20.5 k8s.gcr.io/kube-controller-manager:v1.20.5
#docker tag registry.cn-hangzhou.aliyuncs.com/google_containers/etcd:3.4.13-0 k8s.gcr.io/etcd:3.4.13-0
#docker tag registry.cn-hangzhou.aliyuncs.com/google_containers/coredns:1.7.0 k8s.gcr.io/coredns:1.7.0
#docker tag registry.cn-hangzhou.aliyuncs.com/google_containers/pause:3.2 k8s.gcr.io/pause:3.2

# kubernetes
docker pull registry.cn-hangzhou.aliyuncs.com/google_containers/kube-apiserver:KUBE_APISERVER_VERSION
docker pull registry.cn-hangzhou.aliyuncs.com/google_containers/kube-proxy:KUBE_PROXY_VERSION
docker pull registry.cn-hangzhou.aliyuncs.com/google_containers/kube-scheduler:KUBE_SCHEDULER_VERSION
docker pull registry.cn-hangzhou.aliyuncs.com/google_containers/kube-controller-manager:KUBE_CONTROLLER_VERSION
docker pull registry.cn-hangzhou.aliyuncs.com/google_containers/etcd:ETCD_VERSION
docker pull registry.cn-hangzhou.aliyuncs.com/google_containers/coredns:COREDNS_VERSION
docker pull registry.cn-hangzhou.aliyuncs.com/google_containers/pause:PAUSE_VERSION

docker tag registry.cn-hangzhou.aliyuncs.com/google_containers/kube-apiserver:KUBE_APISERVER_VERSION k8s.gcr.io/kube-apiserver:KUBE_APISERVER_VERSION
docker tag registry.cn-hangzhou.aliyuncs.com/google_containers/kube-proxy:KUBE_PROXY_VERSION k8s.gcr.io/kube-proxy:KUBE_PROXY_VERSION
docker tag registry.cn-hangzhou.aliyuncs.com/google_containers/kube-scheduler:KUBE_SCHEDULER_VERSION k8s.gcr.io/kube-scheduler:KUBE_SCHEDULER_VERSION
docker tag registry.cn-hangzhou.aliyuncs.com/google_containers/kube-controller-manager:KUBE_CONTROLLER_VERSION k8s.gcr.io/kube-controller-manager:KUBE_CONTROLLER_VERSION
docker tag registry.cn-hangzhou.aliyuncs.com/google_containers/etcd:ETCD_VERSION k8s.gcr.io/etcd:ETCD_VERSION
docker tag registry.cn-hangzhou.aliyuncs.com/google_containers/coredns:COREDNS_VERSION k8s.gcr.io/coredns:COREDNS_VERSION
docker tag registry.cn-hangzhou.aliyuncs.com/google_containers/pause:PAUSE_VERSION k8s.gcr.io/pause:PAUSE_VERSION
