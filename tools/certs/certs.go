package certs

import (
	"computedriver/tools/file"
)

func MasterCheck() bool {

	// 判断证书是否存在
	return file.PathExists("/etc/kubernetes/pki/ca.crt") &&
		file.PathExists("/etc/kubernetes/pki/ca.key") &&
		file.PathExists("/etc/kubernetes/pki/apiserver-etcd-client.crt") &&
		file.PathExists("/etc/kubernetes/pki/apiserver-etcd-client.key") &&
		file.PathExists("/etc/kubernetes/pki/apiserver-kubelet-client.crt") &&
		file.PathExists("/etc/kubernetes/pki/apiserver-kubelet-client.key") &&
		file.PathExists("/etc/kubernetes/pki/apiserver.crt") &&
		file.PathExists("/etc/kubernetes/pki/apiserver.key") &&
		file.PathExists("/etc/kubernetes/pki/front-proxy-ca.crt") &&
		file.PathExists("/etc/kubernetes/pki/front-proxy-ca.key") &&
		file.PathExists("/etc/kubernetes/pki/front-proxy-client.crt") &&
		file.PathExists("/etc/kubernetes/pki/front-proxy-client.key") &&
		file.PathExists("/etc/kubernetes/pki/sa.key") &&
		file.PathExists("/etc/kubernetes/pki/sa.pub") &&
		file.PathExists("/etc/kubernetes/pki/kube-controller-manager.pem") &&
		file.PathExists("/etc/kubernetes/pki/kubelet-client-current.pem") &&
		file.PathExists("/etc/kubernetes/pki/kube-scheduler.pem")

}

func NodeCheck(certPath string) bool {
	return file.PathExists(certPath)
}
