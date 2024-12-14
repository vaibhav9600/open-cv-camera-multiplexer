package grpcserver

import (
	"log"

	pb "streaming-server/camera_stream"
	"streaming-server/global"

	"google.golang.org/grpc"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
)

type Server struct {
	GrpcServer *grpc.Server
	pb.UnimplementedStreamingServiceServer
}

func (g Server) Init() error {
	log.Println("Initializing grpc server for Product Search domain")
	pb.RegisterStreamingServiceServer(g.GrpcServer, &g)
	return nil
}

// *DataRequest, StreamingService_GetDataStreamingServer) error
func (s Server) GetDataStreaming(req *emptypb.Empty, srv pb.StreamingService_GetDataStreamingServer) error {
	log.Println("Fetch data streaming")

	for {
		data, ok := <-global.Channel1
		if !ok {
			log.Println("some issue occurred while streaming, reading from channel 1")
		}

		resp := pb.DataResponse{
			Image: data.([]byte),
		}

		if err := srv.Send(&resp); err != nil {
			log.Println("error generating response")
			return err
		}
	}

	return nil
}

func (s Server) GetDataStreamingStream2(req *emptypb.Empty, srv pb.StreamingService_GetDataStreamingStream2Server) error {
	log.Println("Fetch data streaming")

	for {
		data, ok := <-global.Channel2
		if !ok {
			log.Println("some issue occurred while streaming, reading from channel 1")
		}

		resp := pb.DataResponse{
			Image: data.([]byte),
		}

		if err := srv.Send(&resp); err != nil {
			log.Println("error generating response")
			return err
		}
	}

	return nil
}
