package main

import (
	"context"
	"log"
	"net"
	"simple-grpc/pkg/note_v1"
	"time"

	"google.golang.org/grpc"
)

var (
	network = "tcp"
	port    = ":50051"
)

type noteServer struct {
	note_v1.UnimplementedNoteServiceServer
}

func (n *noteServer) Craete(ctx context.Context, in *note_v1.CreateIn) (*note_v1.CreateOut, error) {
	return &note_v1.CreateOut{Note: &note_v1.Note{
		Id:        "123",
		Title:     "test-note",
		Content:   "test-content",
		CreatedAt: time.Now().String(),
		UpdatedAt: time.Now().String(),
	}}, nil
}

func main() {
	lis, err := net.Listen(network, port)
	if err != nil {
		log.Fatal(err)
	}

	grpcServer := grpc.NewServer()
	note_v1.RegisterNoteServiceServer(grpcServer, &noteServer{})

	log.Printf("server is running on port %s", port)
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatal(err)
	}
}
