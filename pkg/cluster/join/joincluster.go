package cluster

import (
	"computedriver/pkg/k8s/kubelet"
	"computedriver/tools/certs"
	"time"
)

func JoinCluster(clusterIp string) {

	// 检查证书，证书不存在一致循环打告警日志
	for {
		if certs.NodeCheck(kubelet.Cert()) && certs.NodeCheck("/etc/kubernetes/pki/ca.crt"){
			break
		}

		time.Sleep(5 * time.Second)
	}

	// todo 判断kubelet 的conf 文件是否和clusterIP一致 并且是否正在运行，防止多次调谐 多次启停

	kubelet.WriteKubeletConf(clusterIp)
	kubelet.StartKubelet()

}
