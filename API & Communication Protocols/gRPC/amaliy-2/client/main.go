package main

import (
	"context"
	"fmt"
	"io"
	"log"

	pb "github.com/khbdev/amaliy-proto/proto/video"
	"google.golang.org/grpc"
)


func main(){
	    conn, err := grpc.Dial("localhost:50051", grpc.WithInsecure())
		   if err != nil {
        log.Fatalf("Failed to connect: %v", err)
    }
    defer conn.Close()

	 client := pb.NewVideoServiceClient(conn)
	   stream, err := client.StreamVideo(context.Background(), &pb.VideoRequest{VideoId: "good"})
    if err != nil {
        log.Fatalf("StreamVideo error: %v", err)
    }
	for {
		  chunk, err := stream.Recv()
		    if err == io.EOF {
            break
        }
		  if err != nil {
            log.Fatalf("Error receiving chunk: %v", err)
        }
		fmt.Printf("Received chunk %d: %s\n", chunk.Index, string(chunk.Data))
		 fmt.Println("Video streaming finished")

	}

}