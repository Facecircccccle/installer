#!/bin/bash

yum install -y deltarpm yum-utils
yum-config-manager --add-repo https://mirrors.aliyun.com/docker-ce/linux/centos/docker-ce.repo

yum makecache

yum install -y DOCKER_VERSION --nogpgcheck

mkdir -p /etc/docker

cat <<EOF > /etc/docker/daemon.json
{
  "exec-opts": ["native.cgroupdriver=systemd"],
  "log-driver": "json-file",
  "log-opts": {
    "max-size": "100m"
  },
  "storage-driver": "overlay2",
  "storage-opts": [
    "overlay2.override_kernel_check=true"
  ],
  "registry-mirrors": ["https://ll77psmm.mirror.aliyuncs.com"],
  "insecure-registries" : ["DOCKER_REGISTRIES"]
}
EOF

tee  /etc/sysctl.d/docker.conf <<EOF
net.bridge.bridge-nf-call-ip6tables=1
net.bridge.bridge-nf-call-iptables=1
net.ipv4.ip_forward=1
EOF

sysctl --system

systemctl daemon-reload
systemctl restart docker
systemctl enable docker

