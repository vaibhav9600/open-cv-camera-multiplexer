package core

import (
	"context"
	"fmt"
	"log"
	"streaming-server/global"

	"gocv.io/x/gocv"
)

// TeeChannel splits messages from a single input channel to multiple output channels.
func TeeChannel(ctx context.Context, input <-chan interface{}, outputs ...chan interface{}) {
	go func() {
		defer func() {
			// Close all output channels when returning
			for _, ch := range outputs {
				close(ch)
			}
		}()

		for {
			select {
			case <-ctx.Done():
				return
			case msg, ok := <-input:
				if !ok {
					return
				}

				// Send the message to all output channels with timeout
				for _, ch := range outputs {
					select {
					case ch <- msg:
						// Successfully sent
					case <-ctx.Done():
						return
					default:
						// Skip if channel is full (non-blocking)
						fmt.Println("Channel full, skipping message")
					}
				}
			}
		}
	}()
}

func StartTee() {
	// Create a context with cancellation
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	// Create buffered data channel
	dataChan := make(chan interface{})
	defer close(dataChan)

	// Start the TeeChannel before sending data
	TeeChannel(ctx, dataChan, global.Channel1, global.Channel2)

	camera, err := gocv.OpenVideoCapture(0)
	if err != nil {
		log.Fatalf("Error opening video capture device: %v", err)
	}
	defer camera.Close()

	// Create a Mat to hold the video frame
	frame := gocv.NewMat()
	defer frame.Close()

	for {
		fmt.Println("inside core loop")

		// Read a frame from the camera
		if ok := camera.Read(&frame); !ok || frame.Empty() {
			log.Println("Cannot read frame from camera")
			break
		}

		// Encode the frame to JPEG
		buf, err := gocv.IMEncode(".jpg", frame)
		if err != nil {
			log.Printf("Error encoding frame: %v", err)
			break
		}

		// Send the bytes through the channel
		select {
		case dataChan <- buf.GetBytes():
			// Successfully sent
		default:
			// Skip if channel is full
			fmt.Println("Channel full, skipping frame")
		}

		buf.Close()
	}
}
