package controller

import (
	"context"
)

func (ser EndPointServices) ProcessEndpoints(ctx context.Context, in *ClusterAgentRequest) (*ServerResponse, error) {
	var eps EndPoints
	eps.SvcName = in.SvcName
	eps.ClusterName = in.ClusterName
	var rcd *AgentRecord

	for _, record := range in.Records {

		rcd.TargetPort = record.TargetPort
		rcd.SvcPort = record.SvcPort

		for _, info := range record.NwInfo {
			rcd.NetworkInfos = append(rcd.NetworkInfos, AgentNetworkInfo{
				EndPoint: info.EndPoint,
				Mac:      info.EndPoint,
			})
		}
		eps.Records = append(eps.Records, rcd)
	}

	ser.EndPointChan <- eps
	return &ServerResponse{Message: "Hello " + in.ClusterName}, nil
}
