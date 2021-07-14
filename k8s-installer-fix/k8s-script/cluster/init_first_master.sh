#!/bin/bash

#高可用
kubeadm init --config /root/k8s-installer/k8s-script/cluster/kubeadmin_init.yaml
#单master
#kubeadm init --kubernetes-version=KUBERNETES_VERSION --apiserver-advertise-address MASTER_INIT_IP  --pod-network-cidr=192.168.0.0/16

mkdir -p $HOME/.kube
cp -i /etc/kubernetes/admin.conf $HOME/.kube/config
chown $(id -u):$(id -g) $HOME/.kube/config

# 配置定时任务
cp ./clean_k8spod.sh /usr/local/bin
cp ./clean_docker.sh /usr/local/bin

cat <<EOF > /var/spool/cron/root
0 0 * * * /usr/local/bin/clean_docker.sh
0 0 1 * * /usr/local/bin/clean_k8spod.sh
EOF

kubectl apply -f /root/k8s-installer/k8s-script/cluster/NETWORK.yaml