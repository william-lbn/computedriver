package app

import (
	pb "computedriver/driver"
	"computedriver/pkg/sandbox"
	"flag"
	"fmt"
	"google.golang.org/grpc"
	"k8s.io/klog/v2"
	"net"
)

func NewAPIServerCommand() {
	var port int
	flag.IntVar(&port, "port", 50002, "listen port")
	flag.Parse()

	klog.Infof("port : %v", port)

	listenPort, err := net.Listen("tcp", fmt.Sprintf("0.0.0.0:%d", port))
	if err != nil {
		klog.Errorf("Failed to listen: %v", err)
	}

	grpcServer := grpc.NewServer()
	pb.RegisterRunSandboxServer(grpcServer, &sandbox.Server{})

	klog.Infof("server listening at %v", listenPort.Addr())
	if err := grpcServer.Serve(listenPort); err != nil {
		klog.Errorf("Failed to serve: %v", err)
	}
}
