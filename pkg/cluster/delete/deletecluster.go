package cluster

import (
	"computedriver/pkg/k8s/kubelet"
)

func DeleteCluster() {
	kubelet.StopKubelet()
}
