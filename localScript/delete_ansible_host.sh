#!/bin/bash
#input 1=role
ansibleHostPath="/etc/ansible/hosts"
roleInit="$1"
role="[$roleInit]"
roleWithBackslash="\[$roleInit\]"

if [ -e $ansibleHostPath ]; then
  if cat $ansibleHostPath | grep $roleWithBackslash > /dev/null; then
    sed -i '/'"${roleWithBackslash}"'/,+1d' $ansibleHostPath
  fi
fi
