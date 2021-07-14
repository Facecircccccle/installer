#!/bin/bash

./k8s-installer/preSet/preSet.sh
ansible-playbook ./k8s-installer/addFirstMaster/HaMasterPreSet.yaml
ansible-playbook ./k8s-installer/LB/configLB.yaml
./k8s-installer/addFirstMaster/addFirstMaster.sh
rm -rf /home/cafile
mkdir -p /home/cafile/etcd
ansible-playbook ./k8s-installer/addOtherMaster/scpfile.yaml
ansible-playbook ./k8s-installer/addOtherMaster/addothermaster.yaml
ansible-playbook ./k8s-installer/addNode/add_node.yaml
