package file

import (
	"fmt"
	"k8s.io/klog/v2"
	"strings"
	"testing"
)

var kubeletService = `
[Unit]
Description=Kubernetes Kubelet Server
Documentation=https://github.com/GoogleCloudPlatform/k8s
After=containerd.service
Wants=containerd.service

[Service]
User=root
EnvironmentFile=-/etc/kubernetes/kubelet.env
ExecStart=/usr/bin/kubelet \
		$KUBE_LOGTOSTDERR \
		$KUBE_LOG_LEVEL \
		$KUBELET_API_SERVER \
		$KUBELET_ADDRESS \
		$KUBELET_PORT \
		$KUBELET_HOSTNAME \
		$KUBELET_ARGS \
		$DOCKER_SOCKET \
		$KUBELET_NETWORK_PLUGIN \
		$KUBELET_VOLUME_PLUGIN \
		$KUBELET_CLOUDPROVIDER
Restart=always
RestartSec=10s

[Install]
WantedBy=multi-user.target
`
var kubeletEnv = `
KUBE_LOGTOSTDERR="--logtostderr=true"
KUBE_LOG_LEVEL="--v=2"
KUBELET_ADDRESS="--node-ip=${localIp}"
KUBELET_HOSTNAME="--hostname-override=${hostname}"



KUBELET_ARGS="--kubeconfig=/etc/kubernetes/kubelet.conf \
--container-runtime=remote \
--container-runtime-endpoint=unix:///var/run/containerd/containerd.sock \
--runtime-cgroups=/systemd/system.slice \
--kubelet-cgroups=/systemd/system.slice \
  "
KUBELET_NETWORK_PLUGIN="--network-plugin=cni --cni-conf-dir=/etc/cni/net.d --cni-bin-dir=/opt/cni/bin"
KUBELET_CLOUDPROVIDER=""

PATH=/usr/local/bin:/usr/local/sbin:/usr/local/bin:/usr/sbin:/usr/bin:/sbin:/bin
`

var kubeletConf = `apiVersion: v1
clusters:
- cluster:
    certificate-authority-data: /etc/kubernetes/certificate/cer.pem
    server: https://${clusterIp}:6443
  name: default-cluster
contexts:
- context:
    cluster: default-cluster
    namespace: default
    user: default-auth
  name: default-context
current-context: default-context
kind: Config
preferences: {}
users:
- name: default-auth
  user:
    client-certificate: /etc/kubernetes/certificate/kubelet-client-current.pem
    client-key: /etc/kubernetes/certificate/kubelet-client-current.pem
`

func TestPathExists(t *testing.T) {
	if !PathExists("/etc/kubernetes/kubelet.conf") {
		fmt.Println("/etc/kubernetes/kubelet.conf not exist.")
	}

	if PathExists("./file.go") {
		fmt.Println("./file.go  exist.")
	}
}

func TestWriteToFile(t *testing.T) {
	newConf := strings.Replace(kubeletConf, "${clusterIp}", "1arfasdf", -1)

	err := WriteToFile("./kubelet.conf", newConf)
	if err != nil {
		klog.Errorf("WriteToFile error. : %v", err)
	}
}
