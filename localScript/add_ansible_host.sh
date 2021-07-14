#!/bin/bash
#input 1=role 2=IP
ansibleHostPath="/etc/ansible/hosts"
roleInit="$1"
role="[$roleInit]"
roleWithBackslash="\[$roleInit\]"
ip=`echo ${@:2}`
if [ -e $ansibleHostPath ]; then
  if cat $ansibleHostPath | grep $roleWithBackslash > /dev/null; then
    sed -i '/'"${roleWithBackslash}"'/a\'"$ip"'' $ansibleHostPath
  else
    echo -e "\n$role\n$ip" >> $ansibleHostPath
  fi
else
  echo -e "\n$role\n$ip" >> $ansibleHostPath
fi
