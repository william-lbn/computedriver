package cluster

import (
	clusterDelete "computedriver/pkg/cluster/delete"
	clusterJoin "computedriver/pkg/cluster/join"
	clusterWatch "computedriver/pkg/cluster/watch"
	"computedriver/pkg/execcmd"
	"computedriver/pkg/k8s"
	cluster "computedriver/pkg/k8s/kubelet"
	"computedriver/pkg/rlet"
	driverChan "computedriver/tools/chan"
	"computedriver/tools/file"
	"computedriver/tools/ip"
	"k8s.io/klog/v2"
	"time"
)

func Switch(clusterIp string) {
	for {
		// 模拟 clusterIp
		clusterChan := <-driverChan.ClusterChan

		if clusterChan == driverChan.ClusterChanJoin {
			klog.Info("receive join cluster signal")
			go func() {
				clusterJoin.JoinCluster(clusterIp)
			}()
		}

		if clusterChan == driverChan.ClusterChanDelete {
			klog.Info("receive delete cluster signal")

			go func() {
				clusterDelete.DeleteCluster()

				// 3s后重新加入集群
				time.Sleep(3 * time.Second)
				// 加入新集群
				joinWatchNewCluster()

				klog.Info("close Current cluster coroutine")
			}()

		}
		klog.Infof("ClusterChan signal is %s.", clusterChan)
	}
}

func joinWatchNewCluster() {
	for {
		// 获取新的集群IP地址
		clusterIp := rlet.NewClusterIP()
		klog.Infof("joinWatchNewCluster clusterIp  %s.", clusterIp)
		if clusterIp == "" {
			time.Sleep(10 * time.Second)
			continue
		} else if clusterIp == ip.LocalIP {
			// 首节点初始化
			klog.Infof("start to  init kubelet.")
			k8s.MaterInit(clusterIp)
			break
		} else {
			// 切换集群，证书需要重新分发，首先删掉之前的证书，以防止重新加入当前集群
			klog.Infof("start to  join the new cluster.")
			execcmd.ExecCommand("rm -rf /etc/kubernetes/pki/ca.crt")
			execcmd.ExecCommand("rm -rf  " + cluster.Cert())
			for {
				if file.PathExists("/etc/kubernetes/pki/ca.crt") && file.PathExists(cluster.Cert()) {
					break
				}
				klog.Warningf("Join new cluster need new certs.")
				klog.Warningf("If you provide correct certs, please check CR exist in new cluster? if not please create it.")
				time.Sleep(5 * time.Second)

			}
			if clusterIp != "" {
				cluster.WriteKubeletConf(clusterIp)
			}
			clusterWatch.Watch(clusterIp)
			break
		}
	}
}
