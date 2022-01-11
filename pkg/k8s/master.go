package k8s

import (
	"computedriver/pkg/execcmd"
	"computedriver/pkg/k8s/apiserver"
	"computedriver/pkg/k8s/controllermanager"
	"computedriver/pkg/k8s/etcd"
	"computedriver/pkg/k8s/kubelet"
	"computedriver/pkg/k8s/scheduler"
	"k8s.io/klog/v2"
	"time"
)

func MaterInit(clusterIp string) {
	klog.Infof("master init begin.")
	// 初始化证书
	execcmd.ExecCommand("kubeadm init phase certs all")
	klog.Infof("kubeadm init phase certs all down.")
	// 初始化配置项
	execcmd.ExecCommand("kubeadm init phase kubeconfig all")
	klog.Infof("kubeadm init phase kubeconfig all down.")

	writeStaticYaml(clusterIp)

	// 判断 crictl ps
	time.Sleep(10 * time.Second)

	// 改env
	kubelet.WriteKubeletEnv()

	kubelet.StopKubelet()
	kubelet.StartKubelet()
	klog.Infof("master init end.")
}

func writeStaticYaml(clusterIp string) {

	etcd.EtcdInit(clusterIp)
	scheduler.SchedulerInit()
	controllermanager.ManagerInit()

	apiserver.ServerInit(clusterIp)
}
