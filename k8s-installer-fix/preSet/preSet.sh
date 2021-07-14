#!/bin/bash

ansible-playbook ./k8s-installer/preSet/system-preSet.yaml
# 没有
sleep 60s
ansible-playbook ./k8s-installer/preSet/dockerInstall.yaml
#DOCKER_VERSION DOCKER_REGISTRIES
ansible-playbook ./k8s-installer/preSet/k8sInstall.yaml
#KUBEADM_VERSION KUBECTL_VERSION KUBELET_VERSION
