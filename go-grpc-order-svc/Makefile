proto:
	protoc pkg/pb/*.proto --go_out=. --go-grpc_out=require_unimplemented_servers=false:.

server:
	go run cmd/main.go
