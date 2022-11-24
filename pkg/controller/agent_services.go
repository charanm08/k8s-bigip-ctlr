package controller

import (
	"context"
	"strconv"
)

func (ser EndPointServices) ProcessEndpoints(ctx context.Context, in *ClusterAgentRequest) (*ServerResponse, error) {
	var eps EndPoints
	eps.SvcName = in.SvcName
	eps.ClusterName = in.ClusterName
	var rcd *AgentRecord

	for _, record := range in.Records {

		targetPort, _ := strconv.ParseInt(record.TargetPort, 0, 8)
		rcd.TargetPort.IntVal = int32(targetPort)
		svcPort, _ := strconv.ParseInt(record.SvcPort, 0, 8)
		rcd.SvcPort.IntVal = int32(svcPort)

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
