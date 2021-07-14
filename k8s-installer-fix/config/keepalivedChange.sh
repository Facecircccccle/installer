#!/bin/bash
card="$1"
sed -i 's/NET_CARD/'"${card}"'/g' /root/k8s-installer/config/keepalived.conf