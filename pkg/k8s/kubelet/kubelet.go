package kubelet

import (
	"computedriver/pkg/execcmd"
	"computedriver/tools/file"
	"computedriver/tools/ip"
	"encoding/base64"
	"io/ioutil"
	"k8s.io/klog/v2"
	"os"
	"os/exec"
	"sigs.k8s.io/controller-runtime/pkg/log"
	"strings"
)

var kubeletService = `[Unit]
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
var kubeletEnv = `KUBE_LOGTOSTDERR="--logtostderr=true"
KUBE_LOG_LEVEL="--v=2"
KUBELET_ADDRESS="--node-ip=${localIp}"
KUBELET_HOSTNAME="--hostname-override=${hostname}"



KUBELET_ARGS="--kubeconfig=/etc/kubernetes/kubelet.conf \
--cgroup-driver=systemd \
--pod-manifest-path=/etc/kubernetes/manifests/ \
--container-runtime=remote \
--container-runtime-endpoint=unix:///var/run/containerd/containerd.sock \
--runtime-cgroups=/systemd/system.slice \
--kubelet-cgroups=/systemd/system.slice \
  "
KUBELET_NETWORK_PLUGIN="--network-plugin=cni --cni-conf-dir=/etc/cni/net.d --cni-bin-dir=/opt/cni/bin"
KUBELET_CLOUDPROVIDER=""

PATH=/usr/local/bin:/usr/local/sbin:/usr/local/bin:/usr/sbin:/usr/bin:/sbin:/bin
`

var kubeletEnvInit = `KUBE_LOGTOSTDERR="--logtostderr=true"
KUBE_LOG_LEVEL="--v=2"
KUBELET_ADDRESS="--node-ip=${localIp}"
KUBELET_HOSTNAME="--hostname-override=${hostname}"



KUBELET_ARGS="--cgroup-driver=systemd \
--pod-manifest-path=/etc/kubernetes/manifests/ \
--container-runtime=remote \
--container-runtime-endpoint=unix:///var/run/containerd/containerd.sock \
"
#KUBELET_NETWORK_PLUGIN="--network-plugin=cni --cni-conf-dir=/etc/cni/net.d --cni-bin-dir=/opt/cni/bin"
KUBELET_CLOUDPROVIDER=""

PATH=/usr/local/bin:/usr/local/sbin:/usr/local/bin:/usr/sbin:/usr/bin:/sbin:/bin
`

var kubeletConf = `apiVersion: v1
clusters:
- cluster:
    certificate-authority-data: ${CA-BASE64}
    server: https://${clusterIp}:6443
  name: kubernetes
contexts:
- context:
    cluster: kubernetes
    namespace: default
    user: kubernetes-admin
  name: kubernetes-admin@kubernetes
current-context: kubernetes-admin@kubernetes
kind: Config
preferences: {}
users:
- name: kubernetes-admin
  user:
    client-certificate: /var/lib/kubelet/pki/kubelet-${hostname}.pem
    client-key: /var/lib/kubelet/pki/kubelet-${hostname}.pem
`

func StartKubelet() {

	// 修改kubelet配置文件 重新加入首节点集群，删掉kubeadm 生成的配置
	execcmd.ExecCommand("rm -rf /etc/systemd/system/kubelet.service.d")

	StopKubelet()
	klog.Infof("start kubelet.")
	commandStart := "systemctl daemon-reload; systemctl start  kubelet"
	cmd := exec.Command("/bin/bash", "-c", commandStart)

	err := cmd.Run()
	if err != nil {
		klog.Errorf("exec command %s is error!", commandStart)
	}
}

func StopKubelet() {
	klog.Infof("stop kubelet.")
	commandStop := "systemctl daemon-reload; systemctl stop  kubelet"
	cmd := exec.Command("/bin/bash", "-c", commandStop)

	err := cmd.Run()
	if err != nil {
		klog.Errorf("exec command %s is error!", commandStop)
	}
}

func writeKubeletService() {
	klog.Infof("write /lib/systemd/system/kubelet.service.")
	err := file.WriteToFile("/lib/systemd/system/kubelet.service", kubeletService)
	if err != nil {
		klog.Errorf("WriteToFile kubelet.service error. : %v", err)
	}
}

func writeKubeletEnv(config string) {

	newKubeletEnvLocalIP := strings.Replace(config, "${localIp}", ip.LocalIP, -1)
	hostName, err := os.Hostname()
	if err != nil {
		log.Log.Error(err, "get host name error.")
	}
	newKubeletEnv := strings.Replace(newKubeletEnvLocalIP, "${hostname}", hostName, -1)

	err = file.WriteToFile("/etc/kubernetes/kubelet.env", newKubeletEnv)
	if err != nil {
		klog.Errorf("WriteToFile kubelet.env error. : %v", err)
	}
	klog.Infof("write /etc/kubernetes/kubelet.env.")
}

func WriteKubeletConf(clusterIp string) {

	writeKubeletService()
	writeKubeletEnv(kubeletEnv)

	kubeletConfIP := strings.Replace(kubeletConf, "${clusterIp}", clusterIp, -1)

	data, err := ioutil.ReadFile("/etc/kubernetes/pki/ca.crt")
	if err != nil {
		klog.Errorf("read /etc/kubernetes/pki/ca.crt error. : %v", err)
	}

	caBase64 := base64.StdEncoding.EncodeToString(data)

	klog.Infof("ca.crt base64 : %s.", caBase64)
	kubeletConfNew := strings.Replace(kubeletConfIP, "${CA-BASE64}", caBase64, -1)

	hostName, err := os.Hostname()
	if err != nil {
		klog.Errorf("get host name error %v.", err)
	}

	kubeletConfCertNew := strings.Replace(kubeletConfNew, "${hostname}", hostName, -1)

	err = file.WriteToFile("/etc/kubernetes/kubelet.conf", kubeletConfCertNew)
	if err != nil {
		klog.Errorf("WriteToFile kubelet.conf error. : %v", err)
	}
}

func WriteKubeletEnv() {

	kubeletEnvIP := strings.Replace(kubeletEnv, "${localIp}", ip.LocalIP, -1)
	hostName, err := os.Hostname()
	if err != nil {
		klog.Errorf("get host name error %v.", err)
	}
	kubeletEnvNew := strings.Replace(kubeletEnvIP, "${hostname}", hostName, -1)

	err = file.WriteToFile("/etc/kubernetes/kubelet.env", kubeletEnvNew)
	if err != nil {
		klog.Errorf("WriteToFile kubelet.env error. : %v", err)
	}
}

func Cert() string {
	hostName, err := os.Hostname()
	if err != nil {
		klog.Errorf("get host name error %v.", err)
		return "/var/lib/kubelet/pki/kubelet-error-err.pem"
	}

	return "/var/lib/kubelet/pki/kubelet-" + hostName + ".pem"
}

func Init() {
	writeKubeletService()
	writeKubeletEnv(kubeletEnvInit)
	StartKubelet()
}
