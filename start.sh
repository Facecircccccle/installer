#!/bin/bash
#1、#传入 ansible的 host各个参数


#2、开始执行预处理
chmod -R +x . # 赋予所有文件执行权限
#tar czvf  k8s-installer.tar.gz . # 打包
ansible-playbook start.yaml #上传文件
#ansible-playbook ./localScript/start/start.yaml