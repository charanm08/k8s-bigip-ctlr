package controller

import (
	"context"
	log "github.com/F5Networks/k8s-bigip-ctlr/pkg/vlogger"
	ipamsvc "github.com/arzzon/ipam-as/api"
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
	grpc := ipamsvc.NewIPManagementClient(ctlr.ipamGrpcCli)
	// Contact the server and print out its response.
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	if key != "" {
		out, err := grpc.ReleaseIP(ctx, &ipamsvc.ReleaseIPRequest{Label: ipamLabel, Hostname: key})
		if err != nil {
			return ""
		}
		return out.IP
	} else {
		log.Debugf("[IPAM] Invalid host and key.")
		return ""
	}
}

func (ctlr *Controller) allocateGrpcIP(ipamLabel, key string) (string, int) {
	grpc := ipamsvc.NewIPManagementClient(ctlr.ipamGrpcCli)
	// Contact the server and print out its response.
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	ip, err := grpc.AllocateIP(ctx, &ipamsvc.AllocateIPRequest{Label: ipamLabel, Hostname: key})
	if err != nil {
		return "", InvalidInput
	}
	return ip.IP, Allocated
}
