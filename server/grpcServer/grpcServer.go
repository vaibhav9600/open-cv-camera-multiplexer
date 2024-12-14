package grpcserver

import (
	"log"
	"time"

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
	log.Println("Started data streaming")

	// Create timer for 10 second timeout
	timer := time.NewTimer(10 * time.Second)
	defer timer.Stop()

	for {
		select {
		case <-timer.C:
			log.Println("Stream timeout after 10 seconds")
			return nil

		case data, ok := <-global.Channel1:
			if !ok {
				log.Println("Channel closed, ending stream")
				return nil
			}

			// Type assert the data to []byte
			imageData, ok := data.([]byte)
			if !ok {
				log.Println("Invalid data type received from channel")
				continue
			}

			// Create and send response
			resp := &pb.DataResponse{
				Image: imageData,
			}

			if err := srv.Send(resp); err != nil {
				log.Printf("Error sending response: %v", err)
				return err
			}
		}
	}
}

func (s Server) GetDataStreamingStream2(req *emptypb.Empty, srv pb.StreamingService_GetDataStreamingStream2Server) error {
	log.Println("Started data streaming")

	// Create timer for 10 second timeout
	timer := time.NewTimer(10 * time.Second)
	defer timer.Stop()

	for {
		select {
		case <-timer.C:
			log.Println("Stream timeout after 10 seconds")
			return nil

		case data, ok := <-global.Channel2:
			if !ok {
				log.Println("Channel closed, ending stream")
				return nil
			}

			// Type assert the data to []byte
			imageData, ok := data.([]byte)
			if !ok {
				log.Println("Invalid data type received from channel")
				continue
			}

			// Create and send response
			resp := &pb.DataResponse{
				Image: imageData,
			}

			if err := srv.Send(resp); err != nil {
				log.Printf("Error sending response: %v", err)
				return err
			}
		}
	}
}
