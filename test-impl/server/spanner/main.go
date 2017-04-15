package main

import (
	"fmt"
	"net"
	_"github.com/lib/pq"
	pb "github.com/tcncloud/protoc-gen-persist/examples"
	"google.golang.org/grpc"
)


func main() {
	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		panic(err)
	}
	grpcServer := grpc.NewServer()
	connString := "projects/algebraic-ratio-149721/instances/test-instance/databases/test-project"
	s, err := pb.NewMySpannerImpl(connString, nil)
	if err != nil {
		panic(err)
	}
	pb.RegisterMySpannerServer(grpcServer, s)
	fmt.Printf("server listening on 50051\n")
	grpcServer.Serve(lis)
}
