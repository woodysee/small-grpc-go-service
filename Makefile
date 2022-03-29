default: proto

proto:
	protoc --go_out=. --go-grpc_out=. --go_opt=paths=source_relative --go-grpc_opt=paths=source_relative **/*.proto
client:
	go run ./client
server:
	go run ./server