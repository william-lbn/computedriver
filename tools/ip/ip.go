package ip

import (
	"k8s.io/klog/v2"
	"net"
)

var LocalIP string

func init() {
	LocalIP = GetLocalIp()
}

func GetLocalIp() string {
	addrs, err := net.InterfaceAddrs()

	if err != nil {
		klog.Errorf("%v", err)
		return ""
	}

	for _, address := range addrs {
		// 检查ip地址判断是否回环地址
		if ipnet, ok := address.(*net.IPNet); ok && !ipnet.IP.IsLoopback() {
			if ipnet.IP.To4() != nil {
				return ipnet.IP.String()
			}
		}
	}
	klog.Errorf("Can not find the local ip address!")
	return ""
}
