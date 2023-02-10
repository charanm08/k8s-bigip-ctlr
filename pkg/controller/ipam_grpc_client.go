package controller

import (
	"context"
	ipamsvc "github.com/F5Networks/k8s-bigip-ctlr/pkg/ipam-grpc/pb"
	log "github.com/F5Networks/k8s-bigip-ctlr/pkg/vlogger"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"time"
)

func NewGRPCClient(uri string) *grpc.ClientConn {
	// Set up a connection to the server.
	conn, err := grpc.Dial(uri, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	return conn
}

func (ctlr *Controller) releaseGrpcIP(ipamLabel, key string) string {
	grpc := ipamsvc.NewIpamGRPCServiceClient(ctlr.ipamGrpcCli)
	// Contact the server and print out its response.
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	if key != "" {
		out, err := grpc.ReleaseIP(ctx, &ipamsvc.ReleaseIPRequest{Label: ipamLabel, Hostname: key})
		if err != nil {
			log.Errorf("error allocating ip address: %v", err)
			return ""
		}
		return out.Ipaddress
	} else {
		log.Errorf("[IPAM] Invalid host and key.")
		return ""
	}
}

func (ctlr *Controller) allocateGrpcIP(ipamLabel, key string) (string, int) {
	grpc := ipamsvc.NewIpamGRPCServiceClient(ctlr.ipamGrpcCli)
	// Contact the server and print out its response.
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	ip, err := grpc.AllocateIP(ctx, &ipamsvc.AllocateIPRequest{Label: ipamLabel, Hostname: key})
	if err != nil {
		log.Errorf("error allocating ip address: %v", err)
		return "", InvalidInput
	}
	return ip.Ipaddress, Allocated
}
