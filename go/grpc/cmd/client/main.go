package main

import (
	"context"
	"log"
	"time"

	pb "simple-grpc/pkg/note_v1"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

const (
	address = "localhost:50051"
)

// createNote creates a new note with given title and content
func createNote(client pb.NoteServiceClient, ctx context.Context, title, content string) (*pb.Note, error) {
	resp, err := client.Create(ctx, &pb.CreateIn{
		Title:   title,
		Content: content,
	})
	if err != nil {
		return nil, err
	}
	log.Printf("Created note: %v", resp.Note)
	return resp.Note, nil
}

// getNote retrieves a note by its ID
func getNote(client pb.NoteServiceClient, ctx context.Context, id string) (*pb.Note, error) {
	resp, err := client.Get(ctx, &pb.GetIn{
		Id: id,
	})
	if err != nil {
		return nil, err
	}
	log.Printf("Got note: %v", resp.Note)
	return resp.Note, nil
}

// updateNote updates an existing note
func updateNote(client pb.NoteServiceClient, ctx context.Context, id, title, content string) (*pb.Note, error) {
	resp, err := client.Update(ctx, &pb.UpdateIn{
		Id:      id,
		Title:   title,
		Content: content,
	})
	if err != nil {
		return nil, err
	}
	log.Printf("Updated note: %v", resp.Note)
	return resp.Note, nil
}

// listNotes retrieves all notes
func listNotes(client pb.NoteServiceClient, ctx context.Context) ([]*pb.Note, error) {
	resp, err := client.List(ctx, &pb.ListIn{})
	if err != nil {
		return nil, err
	}
	log.Printf("List of notes: %v", resp.Notes)
	return resp.Notes, nil
}

// deleteNote removes a note by its ID
func deleteNote(client pb.NoteServiceClient, ctx context.Context, id string) (bool, error) {
	resp, err := client.Delete(ctx, &pb.DeleteIn{
		Id: id,
	})
	if err != nil {
		return false, err
	}
	log.Printf("Delete note result: %v", resp.Success)
	return resp.Success, nil
}

func main() {
	// Set up connection to the server
	conn, err := grpc.Dial(address, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()

	// Create client
	client := pb.NewNoteServiceClient(conn)
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	// Create a new note
	note, err := createNote(client, ctx, "My First Note", "This is the content of my first note")
	if err != nil {
		log.Fatalf("could not create note: %v", err)
	}

	// Get the created note
	_, err = getNote(client, ctx, note.Id)
	if err != nil {
		log.Fatalf("could not get note: %v", err)
	}

	// Update the note
	_, err = updateNote(client, ctx, note.Id, "Updated Note Title", "Updated note content")
	if err != nil {
		log.Fatalf("could not update note: %v", err)
	}

	// List all notes
	_, err = listNotes(client, ctx)
	if err != nil {
		log.Fatalf("could not list notes: %v", err)
	}

	// Delete the note
	_, err = deleteNote(client, ctx, note.Id)
	if err != nil {
		log.Fatalf("could not delete note: %v", err)
	}
}
