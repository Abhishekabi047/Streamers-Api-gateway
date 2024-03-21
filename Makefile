proto:
	protoc --go_out=. --go-grpc_out=. ./pkg/proto/*.proto
	
run:
	go run cmd/main.go