package main

import (
	"streaming-server/core"
	v1Grpc "streaming-server/grpcServer"
	"streaming-server/providers"
)

func main() {
	grpcListener, grpcServer := providers.GetGrpcServer()
	g := v1Grpc.Server{GrpcServer: grpcServer}

	g.Init()

	go providers.StartGrpcServer(grpcServer, grpcListener)

	core.StartTee()
}
