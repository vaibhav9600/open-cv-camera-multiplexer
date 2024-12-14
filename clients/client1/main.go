package main

import (
	"context"
	"fmt"
	"image"
	"io"
	"log"
	"time"

	"gocv.io/x/gocv"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/protobuf/types/known/emptypb"

	pb "client1/camera_stream"
)

func main() {
	// Connect to gRPC server
	conn, err := grpc.Dial("localhost:10000", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("Error connecting to gRPC server: %v", err)
	}
	defer conn.Close()

	// Create VideoWriter
	fps := 30.0
	frameSize := image.Point{X: 640, Y: 480} // Desired output size
	fileName := fmt.Sprintf("output_%s.mp4", time.Now().Format("2006-01-02_15-04-05"))

	writer, err := gocv.VideoWriterFile(
		fileName,
		"mp4v", // codec
		fps,
		frameSize.X,
		frameSize.Y,
		true, // isColor
	)
	if err != nil {
		log.Fatalf("Error creating video writer: %v", err)
	}
	defer writer.Close()

	// Create gRPC client and stream
	client := pb.NewStreamingServiceClient(conn)
	stream, err := client.GetDataStreaming(context.Background(), &emptypb.Empty{})
	if err != nil {
		log.Fatalf("Error creating stream: %v", err)
	}

	frameCount := 0
	startTime := time.Now()

	fmt.Println("Started receiving frames...")

	for {
		resp, err := stream.Recv()
		if err == io.EOF {
			log.Println("Stream ended")
			break
		}
		if err != nil {
			log.Printf("Error receiving from stream: %v", err)
			break
		}

		// Convert received bytes to Mat
		imgBytes := resp.GetImage()
		mat, err := gocv.IMDecode(imgBytes, gocv.IMReadColor)
		if err != nil {
			log.Printf("Error decoding image: %v", err)
			continue
		}

		// Check if frame is empty
		if mat.Empty() {
			log.Println("Received empty frame")
			mat.Close()
			continue
		}

		// Resize frame to desired size
		resized := gocv.NewMat()
		gocv.Resize(mat, &resized, frameSize, 0, 0, gocv.InterpolationLinear)

		// Write frame to video file
		err = writer.Write(resized)
		if err != nil {
			log.Printf("Error writing frame: %v", err)
			mat.Close()
			resized.Close()
			continue
		}

		frameCount++
		if frameCount%100 == 0 {
			elapsed := time.Since(startTime)
			actualFPS := float64(frameCount) / elapsed.Seconds()
			fmt.Printf("Processed %d frames (%.2f FPS)\n", frameCount, actualFPS)
		}

		mat.Close()
		resized.Close()
	}

	elapsed := time.Since(startTime)
	actualFPS := float64(frameCount) / elapsed.Seconds()
	fmt.Printf("\nFinished processing %d frames in %.2f seconds (%.2f FPS)\n",
		frameCount, elapsed.Seconds(), actualFPS)
	fmt.Printf("Video saved to: %s\n", fileName)
}
