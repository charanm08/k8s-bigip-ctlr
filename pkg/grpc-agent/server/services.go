package server

import (
	"context"
	"github.com/F5Networks/k8s-bigip-ctlr/pkg/controller"
	pb "github.com/F5Networks/k8s-bigip-ctlr/pkg/grpc-agent/proto"
)

func (ser EndPointServices) ProcessEndpoints(ctx context.Context, in *pb.ClusterAgentRequest) (*pb.ServerResponse, error) {
	var eps controller.EndPoints
	eps.SvcName = in.SvcName
	eps.ClusterName = in.ClusterName
	var rcd *controller.Record

	for _, record := range in.Records {

		rcd.TargetPort = record.TargetPort
		rcd.SvcPort = record.SvcPort

		for _, info := range record.NwInfo {
			rcd.NetworkInfos = append(rcd.NetworkInfos, controller.NetworkInfo{
				EndPoint: info.EndPoint,
				Mac:      info.EndPoint,
			})
		}
		eps.Records = append(eps.Records, rcd)
	}

	ser.EndPointChan <- eps
	return &pb.ServerResponse{Message: "Hello " + in.ClusterName}, nil
}
