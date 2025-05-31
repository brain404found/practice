PROTO_DIR=api	# папка с proto-файлами
PROTO_FILE=$(PROTO_DIR)/note_v1/note.proto	# путь до конкретного proto-файла
GEN_OUT_DIR=pkg	# куда положить сгенерированные Go файлы

# пути и имена бинарников для сервера и клиента
SERVER_BIN=bin/server
CLIENT_BIN=bin/client

.PHONY: check-tools generate build-server build-client build run-server run-client clean

# проверяем наличие protoc и плагинов для Go
check-tools:
	@command -v protoc >/dev/null || (echo "Install protoc"; exit 1)
	@command -v protoc-gen-go >/dev/null || (echo "Install protoc-gen-go"; exit 1)
	@command -v protoc-gen-go-grpc >/dev/null || (echo "Install protoc-gen-go-grpc"; exit 1)

# генерируем Go код из proto в конкретно указанный нами место, внутри pkg
generate: check-tools
	protoc \
		--proto_path=$(PROTO_DIR) \
		--go_out=paths=source_relative:$(GEN_OUT_DIR) \
		--go-grpc_out=paths=source_relative:$(GEN_OUT_DIR) \
		$(PROTO_FILE)

# собираем бинарник сервера из cmd/server
build-server:
	go build -o $(SERVER_BIN) ./cmd/server

# собираем бинарник клиента из cmd/client
build-client:
	go build -o $(CLIENT_BIN) ./cmd/client

# собираем и сервер, и клиент
build: build-server build-client

run-server:
	go run ./cmd/server

run-client:
	go run ./cmd/client

# чистим сгенерированные .pb.go файлы и бинарники
clean:
	rm -rf $(GEN_OUT_DIR)/*.pb.go bin/
