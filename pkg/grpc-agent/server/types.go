package server

import (
	"github.com/F5Networks/k8s-bigip-ctlr/pkg/controller"
	pb "github.com/F5Networks/k8s-bigip-ctlr/pkg/grpc-agent/proto"
	"sync"
)

type EndPointServices struct {
	pb.UnimplementedEndPointServiceServer
	EndPointChan chan controller.EndPoints
}

type GRPCAgent struct {
	EndPointServices *EndPointServices
	Port             *int
	declUpdate       sync.Mutex
}
