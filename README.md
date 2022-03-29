# `small-grpc-go-service`

Tiny Golang service with a gRPC setup

## gRPC vs REST

- Payload data type
  - gRPC: [Protocol Buffers](https://developers.google.com/protocol-buffers)
  - REST: JSON
- Protocol type
  - gRPC: HTTP/2?
  - REST: HTTP/1.1
## Pre-requisites

- Go ^1.17
- GNU Make
- `make proto` requires installing `protoc-gen-go` and `protoc-gen-go-grpc` first.

```bash
go install google.golang.org/protobuf/cmd/protoc-gen-go@latest
go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest
```

## Development

```bash
make server
# In another Shell window
make client

# Generates or updates *.pb.go generated files. They should be not be directly edited.
. .envrc # Run these commands when starting a new shell window. Alternatively you can put this in your own global shell RC file.
make proto
```

## References

- [Protocol Buffers defined by Google](https://developers.google.com/protocol-buffers)
- [Protocol Buffers - What it is and Setting it up](https://dev.to/dsckiitdev/protocol-buffers-in-go-5bl7)
- [Setting up gRPC in an existing Go project](https://grpc.io/docs/languages/go/quickstart/)
- [Alternative setup up gRPC in Go](https://dev.to/dsckiitdev/build-a-grpc-server-in-go-1890)