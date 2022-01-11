#!/bin/bash
echo "file binary path : $1"
echo "version : $2"

# 1、创建目录，在目录下操作
mkdir computedriver-v0.1
cd ./computedriver-v0.1

# 2、安装containerd
# 2.1 下载
wget http://172.17.20.4/releases/containerd/cri-containerd-cni-1.5.8-linux-amd64.tar.gz
# 2.2 解压
tar -C / -xzf cri-containerd-cni-1.5.8-linux-amd64.tar.gz

# 2.3 加入环境变量
export PATH=$PATH:/usr/local/bin:/usr/local/sbin
source ~/.bashrc

# 2.4 下载runc配置文件
mkdir -p /etc/containerd ; wget  http://172.17.20.4/releases/containerd/runc-config.toml -O /etc/containerd/config.toml

# 2.5 配置代理
mkdir -p /etc/systemd/system/containerd.service.d

cat > /etc/systemd/system/containerd.service.d/http_proxy.conf << EOF
[Service]
Environment="HTTP_PROXY=http://172.17.9.143:7890/"
Environment="HTTPS_PROXY=http://172.17.9.143:7890/"
EOF

# 2.5 重启container
systemctl restart containerd

# 安装 crictl
wget http://172.17.20.4/releases/crictl-v1.21.0-linux-amd64.tar.gz

sudo tar zxvf crictl-v1.21.0-linux-amd64.tar.gz -C /usr/local/bin

# 3、安装kubelet、kubeadm首节点使用、或者首节点需要启动api-server、controller等
apt-get install -y kubelet=1.23.1-00
apt-get install -y kubeadm=1.23.1-00

rm -rf /etc/kubernetes/*

mkdir -p /etc/kubernetes/pki
mkdir -p /etc/kubernetes/manifests


sudo systemctl stop computedriver

rm -rf /usr/local/bin/computedriver
rm -rf /etc/systemd/system/computedriver.service

# 4、计算固件下载
wget -c $1/computedriver/computedriver-$2 -O /usr/local/bin/computedriver ; chmod +x /usr/local/bin/computedriver

# computedriver.service 到   etc/systemd/system/
wget -c $1/computedriver/computedriver.service -O /etc/systemd/system/computedriver.service


# 启动
sudo systemctl daemon-reload
sudo systemctl restart computedriver

