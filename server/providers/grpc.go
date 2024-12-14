package providers

import (
	"log"
	"net"

	"google.golang.org/grpc"
)

const (
	protocol = "unix"
	sockAddr = "/tmp/echo.sock"
)

func GetGrpcServer() (net.Listener, *grpc.Server) {
	listener, err := net.Listen("tcp", "127.0.0.1:10000")
	if err != nil {
		log.Print("Error while listening for grpc server")
		panic(err)
	}
	grpcServer := grpc.NewServer()
	return listener, grpcServer
}

func StartGrpcServer(grpcServer *grpc.Server, grpcListener net.Listener) {
	log.Print("Starting Central GRPC server...")
	if err := grpcServer.Serve(grpcListener); err != nil {
		log.Print("err", err)
		panic(err)
	}
}
