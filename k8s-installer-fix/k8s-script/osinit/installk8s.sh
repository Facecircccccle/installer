#!/bin/bash

#==================================================================================================================
# Setup k8s yum repo
#------------------------------------------------------------------------------------------------------------------
cat <<EOF > /etc/yum.repos.d/kubernetes.repo
[kubernetes]
name=Kubernetes
baseurl=https://mirrors.aliyun.com/kubernetes/yum/repos/kubernetes-el7-x86_64/
enabled=1
gpgcheck=0
repo_gpgcheck=0
gpgkey=https://mirrors.aliyun.com/kubernetes/yum/doc/yum-key.gpg https://mirrors.aliyun.com/kubernetes/yum/doc/rpm-package-key.gpg
EOF

yum makecache

#==================================================================================================================
# Install k8s
#------------------------------------------------------------------------------------------------------------------
yum install -y kubeadm-KUBEADM_VERSION kubelet-KUBELET_VERSION kubectl-KUBECTL_VERSION --disableexcludes=kubernetes
systemctl enable kubelet
systemctl restart kubelet

kubeadm version

