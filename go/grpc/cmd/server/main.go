package main

import (
	"context"
	"fmt"
	"log"
	"net"
	pb "simple-grpc/pkg/note_v1"
	"sync"
	"time"

	"github.com/google/uuid"
	"google.golang.org/grpc"
)

// Network configuration constants
var (
	network = "tcp"
	port    = ":50051"
)

// noteServer implements the NoteServiceServer interface
// It provides CRUD operations for notes with thread-safe access
type noteServer struct {
	pb.UnimplementedNoteServiceServer
	mu    sync.RWMutex        // RWMutex for thread-safe access to notes map
	notes map[string]*pb.Note // In-memory storage for notes
}

// newNoteServer creates and returns a new instance of noteServer
// with initialized notes map
func newNoteServer() *noteServer {
	return &noteServer{
		notes: make(map[string]*pb.Note),
	}
}

// Create implements the Create RPC method
// Creates a new note with unique ID and current timestamp
func (s *noteServer) Create(ctx context.Context, req *pb.CreateIn) (*pb.CreateOut, error) {
	s.mu.Lock()
	defer s.mu.Unlock()

	now := time.Now().Format(time.RFC3339)
	note := &pb.Note{
		Id:        uuid.New().String(), // Generate unique ID
		Title:     req.Title,
		Content:   req.Content,
		CreatedAt: now,
		UpdatedAt: now,
	}

	s.notes[note.Id] = note
	return &pb.CreateOut{Note: note}, nil
}

// Get implements the Get RPC method
// Retrieves a note by its ID
func (s *noteServer) Get(ctx context.Context, req *pb.GetIn) (*pb.GetOut, error) {
	s.mu.RLock()
	defer s.mu.RUnlock()

	note, exists := s.notes[req.Id]
	if !exists {
		return nil, fmt.Errorf("note not found: %s", req.Id)
	}

	return &pb.GetOut{Note: note}, nil
}

// List implements the List RPC method
// Returns all stored notes
func (s *noteServer) List(ctx context.Context, req *pb.ListIn) (*pb.ListOut, error) {
	s.mu.RLock()
	defer s.mu.RUnlock()

	notes := make([]*pb.Note, 0, len(s.notes))
	for _, note := range s.notes {
		notes = append(notes, note)
	}

	return &pb.ListOut{Notes: notes}, nil
}

// Update implements the Update RPC method
// Updates an existing note's title and content
func (s *noteServer) Update(ctx context.Context, req *pb.UpdateIn) (*pb.UpdateOut, error) {
	s.mu.Lock()
	defer s.mu.Unlock()

	note, exists := s.notes[req.Id]
	if !exists {
		return nil, fmt.Errorf("note not found: %s", req.Id)
	}

	note.Title = req.Title
	note.Content = req.Content
	note.UpdatedAt = time.Now().Format(time.RFC3339)

	return &pb.UpdateOut{Note: note}, nil
}

// Delete implements the Delete RPC method
// Removes a note by its ID
func (s *noteServer) Delete(ctx context.Context, req *pb.DeleteIn) (*pb.DeleteOut, error) {
	s.mu.Lock()
	defer s.mu.Unlock()

	if _, exists := s.notes[req.Id]; !exists {
		return nil, fmt.Errorf("note not found: %s", req.Id)
	}

	delete(s.notes, req.Id)
	return &pb.DeleteOut{Success: true}, nil
}

func main() {
	// Create TCP listener
	lis, err := net.Listen(network, port)
	if err != nil {
		log.Fatal(err)
	}

	// Create new gRPC server
	grpcServer := grpc.NewServer()
	// Register our service implementation
	pb.RegisterNoteServiceServer(grpcServer, newNoteServer())

	log.Printf("server is running on port %s", port)
	// Start serving requests
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatal(err)
	}
}
