#!/bin/bash
cat >> /etc/sysctl.conf << EOF
net.ipv4.ip_forward = 1
EOF
yum remove -y keepalived
rm -rf /etc/keepalived
yum install -y keepalived
cat >> /etc/sysctl.conf << EOF
net.ipv4.ip_nonlocal_bind = 1
EOF
yum remove -y haproxy
rm -rf /etc/haproxy
yum install -y haproxy

