#!/bin/bash

# 下载
wget https://github.com/containerd/containerd/releases/download/v1.5.8/cri-containerd-cni-1.5.8-linux-amd64.tar.gz
# 解压
tar -C / -xzf cri-containerd-cni-1.5.8-linux-amd64.tar.gz

# 加入环境变量
export PATH=$PATH:/usr/local/bin:/usr/local/sbin
source ~/.bashrc

# 下载runc配置文件
mkdir -p /etc/containerd ; wget  http://172.17.20.4/releases/containerd/runc-config.toml -O /etc/containerd/config.toml


# 重启containerd
systemctl daemon-reload
systemctl restart containerd

# 安装 kubelet
apt-get install -y kubelet=1.22.4-00
apt-get install -y kubeadm=1.22.4-00