#!/bin/bash

#input 1=Name 2=IP

ansible-playbook sethostname.yaml -e "name=$1"
