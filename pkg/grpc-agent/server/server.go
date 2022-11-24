package server

import (
	"flag"
	"fmt"
	pb "github.com/F5Networks/k8s-bigip-ctlr/pkg/grpc-agent/proto"
	"google.golang.org/grpc"
	"log"
	"net"
)

func (grpcAgnt *GRPCAgent) StartGRPCServer() {
	flag.Parse()
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", *grpcAgnt.Server.Port))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()

	pb.RegisterEndPointServiceServer(s, grpcAgnt.Server)
	log.Printf("server listening at %v", lis.Addr())
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
