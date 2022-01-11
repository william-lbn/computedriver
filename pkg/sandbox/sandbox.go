package sandbox

import (
	_ "computedriver/config/yaml"
	pb "computedriver/driver"
	"context"
	"k8s.io/klog/v2"
	"os"
	"os/exec"
	"strings"
)

const (
	OK           = 000000
	PartialError = 100001

	HyperOsPath = "/etc/hyperos/"
)

type Server struct {
	pb.UnimplementedRunSandboxServer
}

func (driver *Server) RunSandbox(ctx context.Context, in *pb.SandboxList) (*pb.SandboxReply, error) {
	klog.Infof("Begin to run sandbox")

	var resCode = false

	for i := 0; i < len(in.Sandbox); i++ {
		url := in.Sandbox[i].Url
		if !strings.HasSuffix(in.Sandbox[i].Url, "/") {
			url += "/"
		}
		downloadUrl := url + in.Sandbox[i].Name + "-" + in.Sandbox[i].Version + ".zip"
		path := "./driver/" + in.Sandbox[i].Name + "-" + in.Sandbox[i].Version
		filePath := path + "/" + in.Sandbox[i].Name + "-" + in.Sandbox[i].Version + ".zip"
		fileName := in.Sandbox[i].Name + "-" + in.Sandbox[i].Version + ".zip"
		err := os.MkdirAll(path, os.ModePerm)
		if err != nil {
			return nil, err
		}

		configPath := HyperOsPath + in.Sandbox[i].Name
		err = os.MkdirAll(configPath, os.ModePerm)
		if err != nil {
			return nil, err
		}

		// 防止重复调用产生很多重复文件
		deleteExistFile(path)

		command := "wget -P " + path + " " + downloadUrl + "; unzip " + filePath + " -d " + path +
			"; mv -f " + path + "/*config.yaml  " + configPath +
			"; mv -f " + path + "/*yaml " + " /etc/kubernetes/manifests"

		klog.Infof("command: %s", command)

		cmd := exec.Command("/bin/bash", "-c", command)

		err = cmd.Run()
		if err != nil {
			// vm可能不存在unzip,安装unzip
			if exec.Command("/bin/bash", "-c", "apt-get install -y unzip").Run() != nil {
				klog.Errorf("apt-get install -y unzip error!")
				klog.Errorf("%v", err)
				resCode = true
				continue
			}
			// 重新执行一次下载动作
			deleteExistFile(path)
			if exec.Command("/bin/bash", "-c", command).Run() != nil {
				klog.Errorf("download %s is error!", fileName)
				klog.Errorf("%v", err)
				resCode = true
				continue
			}
		}

	}

	if resCode {
		return &pb.SandboxReply{Code: PartialError, Message: "Partial error."}, nil
	}
	return &pb.SandboxReply{Code: OK, Message: "succeed."}, nil
}

func deleteExistFile(path string) {
	commandDelete := "rm -rf " + path
	cmd := exec.Command("/bin/bash", "-c", commandDelete)
	err := cmd.Run()
	if err != nil {
		klog.Errorf("rm -rf  %s is error!", path)
	}
}
