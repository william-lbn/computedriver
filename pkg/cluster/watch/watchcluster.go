package cluster

import (
	"computedriver/pkg/hyperNode"
	"computedriver/pkg/k8s/kubelet"
	"computedriver/tools/file"
	"time"
)

func Watch(clusterIp string) {
	for {
		if file.PathExists("/etc/kubernetes/pki/ca.crt") && file.PathExists(kubelet.Cert()) {

			kubelet.WriteKubeletConf(clusterIp)

			hyperNode.Run()
			// 采用了runtime.Goexit() 此处是防止异常导致死循环
			break
		}
		time.Sleep(10 * time.Second)
	}
}
