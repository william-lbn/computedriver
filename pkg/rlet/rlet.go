package rlet

import (
	"computedriver/pkg/execcmd"
	pb "computedriver/pkg/rlet/proto"
	"context"
	"google.golang.org/grpc"
	"k8s.io/klog/v2"
)

const (
	// address = "localhost:6006"
	 address = "192.168.2.79:6006"
)

func NewClusterIP() string {
	// Set up a connection to the server.
	conn, err := grpc.Dial(address, grpc.WithInsecure())
	if err != nil {
		klog.Errorf("Failed to connect: %v", err)
		return ""
	}
	defer func(conn *grpc.ClientConn) {
		err := conn.Close()
		if err != nil {
			klog.Errorf("conn.Close() error : %v", err)
		}
	}(conn)
	c := pb.NewRletServiceClient(conn)

	// 获取机器标识
	command := "cat /etc/machine-id"
	machineId := execcmd.ExecCommand(command)

	r, err := c.GetChessmap(context.Background(), &pb.ChessmapRequest{NodeId: machineId})

	if err != nil {
		klog.Errorf("Failed to GetChessmap: %v", err)
		return ""
	}
	klog.Infof("####### get server Greeting response: %s", r.ComponentList)

	// 异常考虑
	if len(r.ComponentList) == 0 {
		klog.Errorf("Failed to GetChessmap, and the response is empty.")
		return ""
	}
	if len(r.ComponentList[0].SupervisorList) == 0 {
		klog.Errorf("Failed to GetChessmap, and the response `r.ComponentList[0].SupervisorList` is empty.")
		return ""
	}

	return r.ComponentList[0].SupervisorList[0]

}
