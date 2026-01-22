package main

import (
	"context"
	"fmt"
	"log"
	"net"
	"time"

	pb "github.com/khbdev/amaliy-proto/proto/video"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)


func unaryInterceptor(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (interface{}, error) {
	log.Printf("[UNARY ] Incoming request: %s", info.FullMethod)
	resp, err := handler(ctx, req)
	if err != nil {
		log.Printf("[UNARY] Error: %v", err)
	}
	return resp, err
}

func streamInterceptor(
	srv interface{}, 
	ss grpc.ServerStream,
	info *grpc.StreamServerInfo,
	handler grpc.StreamHandler,
) error {
	   log.Printf("[STREAM] %s started", info.FullMethod)
    err := handler(srv, ss)
    if err != nil {
        log.Printf("[STREAM] Error: %v", err)
    } else {
        log.Printf("[STREAM] %s finished", info.FullMethod)
    }
    return err
}

type  server struct {
	pb.UnimplementedVideoServiceServer
}

func (s *server) StreamVideo(req *pb.VideoRequest, stream pb.VideoService_StreamVideoServer) error {
	log.Printf("Streaming video: %s", req.VideoId)

	 videoChunks := [][]byte{
        []byte("chunk1 data"),
        []byte("chunk2 data"),
        []byte("chunk3 data"),
    }

	for i, chunk := range videoChunks {
		if i == 1 && req.VideoId == "bad" {
			return status.Errorf(codes.NotFound, "Video %s not found", req.VideoId)
		}
		   if err := stream.Send(&pb.VideoChunk{
            Index: int32(i),
            Data:  chunk,
        }); err != nil {
            return status.Errorf(codes.Internal, "Failed to send chunk %d: %v", i, err)
        }
		time.Sleep(time.Second)
	}
	return  nil

	
}


func main(){
	 lis, err := net.Listen("tcp", ":50051")
    if err != nil {
        log.Fatalf("failed to listen: %v", err)
    }

    grpcServer := grpc.NewServer(
        grpc.ChainUnaryInterceptor(unaryInterceptor),
        grpc.ChainStreamInterceptor(streamInterceptor),
    )

    pb.RegisterVideoServiceServer(grpcServer, &server{})
    fmt.Println("Server listening on :50051")
    if err := grpcServer.Serve(lis); err != nil {
        log.Fatalf("failed to serve: %v", err)
    }
}