package app

import (
	clusterSwitch "computedriver/pkg/cluster/switch"
	clusterWatch "computedriver/pkg/cluster/watch"
	"computedriver/pkg/k8s"
	"computedriver/pkg/k8s/kubelet"
	"computedriver/pkg/rlet"
	"computedriver/pkg/sandbox/app"
	"computedriver/tools/ip"
	"k8s.io/klog/v2"
	"time"
)

func Run() {

	klog.Infof("start to  init kubelet.")
	kubelet.Init()
	klog.Infof("end to  init kubelet.")

	for {
		clusterIp := rlet.NewClusterIP()
		if clusterIp == "" {
			klog.Warningf("Retry after 10s to get cluster ip from rlet.")
			time.Sleep(10 * time.Second)
			continue
		} else if clusterIp == ip.LocalIP {
			// 首节点初始化
			k8s.MaterInit(clusterIp)
			// 首节点不进行watch
			break
		} else {
			klog.Infof("start to  switch cluster.")
			go func() {
				clusterSwitch.Switch(clusterIp)
			}()

			klog.Infof("start to  watch cluster.")
			go func() {
				clusterWatch.Watch(clusterIp)
			}()
			klog.Infof("clusterIp is %s , localIp is %s.", clusterIp, ip.LocalIP)
			break
		}
	}

	klog.Infof("start to grpc server.")
	// grpc server.
	app.NewAPIServerCommand()

}
