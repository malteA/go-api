package main

import (
	"context"
	"fmt"
	"log"

	"google.golang.org/grpc"

	pb "github.com/maltea/go-api/carpb"
	_ "github.com/maltea/go-api/docs"

	_ "github.com/jnewmano/grpc-json-proxy/codec"
)

type carServiceServer struct {
	pb.UnimplementedCarServiceServer
}

// @title Swagger Car API
// @version 1.0
// @description This is a sample golang server with fiber and gorm
// @host localhost:3000
// @basePath /api
func main() {
	opts := grpc.WithInsecure()
	cc, err := grpc.Dial("localhost:3000", opts)
	if err != nil {
		log.Fatal(err)
	}
	defer cc.Close()

	client := pb.NewCarServiceClient(cc)
	// request := &google.protobuf.Empty{}
	request := &pb.GetCarRequest{
		Id: "1",
	}

	resp, _ := client.GetCar(context.Background(), request)
	// if resp != nil {
	fmt.Printf("Receive response => [%v]", resp.Brand)
	// }
}
