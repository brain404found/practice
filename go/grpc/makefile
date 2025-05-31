PROTO_DIR=api/note_v1
PROTO_FILE=$(PROTO_DIR)/note.proto
GEN_OUT_DIR=pkg
SERVER_BIN=bin/server
CLIENT_BIN=bin/client

check-tools:
	@command -v protoc-gen-go >/dev/null || (echo "Install protoc-gen-go"; exit 1)
	@command -v protoc-gen-go-grpc >/dev/null || (echo "Install protoc-gen-go-grpc"; exit 1)

generate: check-tools
	protoc --go_out=$(GEN_OUT_DIR) --go-grpc_out=$(GEN_OUT_DIR) $(PROTO_FILE)

build-server:
	go build -o $(SERVER_BIN) ./cmd/server

build-client:
	go build -o $(CLIENT_BIN) ./cmd/client

build: build-server build-client

run-server:
	go run ./cmd/server

run-client:
	go run ./cmd/client

clean:
	rm -rf $(GEN_OUT_DIR)/note_v1/*.pb.go bin/
