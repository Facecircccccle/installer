#!/bin/bash

cat > /etc/sysconfig/kubelet << EOF
KUBELET_EXTRA_ARGS=--enforce-node-allocatable=pods --cgroup-driver=systemd --kube-reserved=cpu=1,memory=500Mi,ephemeral-storage=10Gi --system-reserved=cpu=1,memory=500Mi,ephemeral-storage=10Gi --eviction-hard=memory.available<500Mi,nodefs.available<10%
EOF
systemctl restart kubelet
systemctl status kubelet
