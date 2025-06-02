# Simple gRPC Notes Service

A simple gRPC-based notes service that demonstrates CRUD operations using gRPC in Go. This project serves as a practical example of building a microservice using gRPC.

## Architecture

```

├── api/
│   └── note_v1/
│       └── note.proto      # Protocol Buffers definition
├── cmd/
│   ├── client/
│   │   └── main.go        # gRPC client implementation
│   └── server/
│       └── main.go        # gRPC server implementation
├── pkg/
│   └── note_v1/           # Generated Go code from proto
└── readme.md
```

### Components

- **Protocol Buffers Definition** (`note.proto`): Defines the service interface and message types
- **Server** (`cmd/server/main.go`): Implements the gRPC service with in-memory storage
- **Client** (`cmd/client/main.go`): Demonstrates how to interact with the service

## Features

- Create, Read, Update, Delete (CRUD) operations for notes
- Thread-safe in-memory storage
- Unique ID generation for notes
- Timestamp tracking for creation and updates
- Simple and clean API

## Prerequisites

- Go 1.16 or higher
- Protocol Buffers compiler (protoc)
- Go plugins for protoc:
  - protoc-gen-go
  - protoc-gen-go-grpc

## Getting Started

1. **Install Protocol Buffers compiler and Go plugins**

```bash
# Install protoc (macOS)
brew install protobuf

# Install Go plugins
go install google.golang.org/protobuf/cmd/protoc-gen-go@latest
go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest
```

2. **Initialize the project**

```bash
# Initialize Go module
go mod init simple-grpc

# Install dependencies
go get google.golang.org/grpc
go get github.com/google/uuid
```

3. **Generate Go code from proto**

```bash
protoc --go_out=. --go_opt=paths=source_relative \
    --go-grpc_out=. --go-grpc_opt=paths=source_relative \
    api/note_v1/note.proto
```

4. **Run the server**

```bash
go run cmd/server/main.go
```

5. **Run the client**

In a new terminal:
```bash
go run cmd/client/main.go
```

## Testing the Service

The client demonstrates all CRUD operations in sequence:

1. Creates a new note
2. Retrieves the created note
3. Updates the note
4. Lists all notes
5. Deletes the note

You should see log output for each operation in both the server and client terminals.

## API Reference

### Note Service

```protobuf
service NoteService {
    rpc Create(CreateIn) returns (CreateOut);
    rpc Get(GetIn) returns (GetOut);
    rpc List(ListIn) returns (ListOut);
    rpc Update(UpdateIn) returns (UpdateOut);
    rpc Delete(DeleteIn) returns (DeleteOut);
}
```

### Message Types

- `Note`: Represents a note with id, title, content, and timestamps
- `CreateIn`: Input for creating a note (title, content)
- `GetIn`: Input for getting a note (id)
- `ListIn`: Empty input for listing all notes
- `UpdateIn`: Input for updating a note (id, title, content)
- `DeleteIn`: Input for deleting a note (id)
