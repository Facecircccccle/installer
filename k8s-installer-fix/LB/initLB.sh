#!/bin/bash
chmod 644 /etc/keepalived/keepalived.conf
chmod 644 /etc/haproxy/haproxy.cfg
systemctl enable keepalived.service
systemctl start keepalived.service
systemctl enable haproxy.service
systemctl start haproxy.service 

