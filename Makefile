LOCAL_BIN := $(shell pwd)/bin

proto: install-go-deps generate-go-clients_service_v1-api generate-go-books_service_v1-api

install-go-deps:
	GOBIN=$(LOCAL_BIN) go install google.golang.org/protobuf/cmd/protoc-gen-go@v1.32.0
	GOBIN=$(LOCAL_BIN) go install -mod=mod google.golang.org/grpc/cmd/protoc-gen-go-grpc@v1.2
	GOBIN=$(LOCAL_BIN) go install github.com/pseudomuto/protoc-gen-doc/cmd/protoc-gen-doc@latest
	GOBIN=$(LOCAL_BIN) go install github.com/envoyproxy/protoc-gen-validate@v0.10.1

generate-go-clients_service_v1-api:
	mkdir -p src/go/pkg/grpc/clients/v1
	GOBIN=$(LOCAL_BIN) protoc \
	--proto_path . \
	--proto_path api/grpc/protobuf/clients_v1 \
	--go_out=src/go/pkg/grpc/clients/v1 \
	--go_opt=paths=source_relative \
	--go-grpc_out=src/go/pkg/grpc/clients/v1 \
	--go-grpc_opt=paths=source_relative \
	--plugin=protoc-gen-go=bin/protoc-gen-go \
	--plugin=protoc-gen-go-grpc=bin/protoc-gen-go-grpc \
	--plugin=protoc-gen-doc=bin/protoc-gen-doc \
	--doc_out=. \
	--doc_opt=markdown,clients_readme.md,source_relative \
	api/grpc/protobuf/clients_v1/*.proto

generate-go-books_service_v1-api:
	mkdir -p src/go/pkg/grpc/books/v1
	GOBIN=$(LOCAL_BIN) protoc \
	--proto_path . \
	--proto_path api/grpc/protobuf/books_v1 \
	--go_out=src/go/pkg/grpc/books/v1 \
	--go_opt=paths=source_relative \
	--go-grpc_out=src/go/pkg/grpc/books/v1 \
	--go-grpc_opt=paths=source_relative \
	--plugin=protoc-gen-go=bin/protoc-gen-go \
	--plugin=protoc-gen-go-grpc=bin/protoc-gen-go-grpc \
	--plugin=protoc-gen-doc=bin/protoc-gen-doc \
	--doc_out=. \
	--doc_opt=markdown,books_readme.md,source_relative \
	api/grpc/protobuf/books_v1/*.proto
