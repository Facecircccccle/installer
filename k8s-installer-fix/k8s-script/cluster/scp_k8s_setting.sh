#!/bin/bash

/bin/rpm -qa|/bin/grep -q expect
if [ $? -ne 0 ];then
  yum -y install expect
  exit
fi

dest=$1
passwd=$2
scp -r root@$dest:/root/k8s-setting /root/k8s-playbook/

expect -c "
    spawn scp -r /root/k8s-playbook/ root@$dest:/root/k8s-setting/ 
    expect {
        \"*assword\" {set timeout 20; send \"$passwd\r\"; exp_continue;} 
        \"yes/no\" {send \"yes\r\"; exp_continue;}
    }
expect eof"
