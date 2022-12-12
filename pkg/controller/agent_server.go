package controller

import (
	"fmt"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"log"
	"net"
)

func (grpcAgnt *GRPCAgent) StartGRPCServer() {
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", grpcAgnt.Port))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	reflection.Register(s)

	RegisterEndPointServiceServer(s, grpcAgnt.EndPointServices)
	log.Printf("server listening at %v", lis.Addr())
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}

}
