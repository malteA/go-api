# Go Car API Example

Example Service to manage Cars.

gRPC is used for the interaction, GORM to interact with the Database.

## Usage

You can build this programm on your platform by yourself or simply start it with "`go run server/server.go.`". The application opens a socket at the port 3000.

To run Client `go run client/client.go`

## GRPC

### Setup

- install `protobuf`
- install `go get -u github.com/golang/protobuf/protoc-gen-go`
- install `go get -u google.golang.org/grpc`

- Install `go get -u github.com/jnewmano/grpc-json-proxy` [https://github.com/jnewmano/grpc-json-proxy#grpc-json-proxy](grpc-json-proxy)