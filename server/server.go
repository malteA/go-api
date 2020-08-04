package main

import (
	"context"
	"log"
	"net"

	"google.golang.org/grpc"
	"google.golang.org/protobuf/types/known/emptypb"

	"github.com/jinzhu/gorm"
	"github.com/maltea/go-api/car"
	pb "github.com/maltea/go-api/carpb"
	"github.com/maltea/go-api/database"
	_ "github.com/maltea/go-api/docs"

	_ "github.com/jnewmano/grpc-json-proxy/codec"
)

type carServiceServer struct{}

func (*carServiceServer) ListCars(ctx context.Context, request *emptypb.Empty) (*pb.ListCarsResponse, error) {
	db := database.DBConn

	var cars []car.Car
	db.Find(&cars)
	var responseCars []*pb.GetCarResponse

	for _, car := range cars {
		responseCars = append(responseCars, &pb.GetCarResponse{
			Brand: car.Brand,
			Make:  car.Make,
		})
	}

	response := &pb.ListCarsResponse{
		Cars: responseCars,
	}
	return response, nil
}

func (*carServiceServer) GetCar(ctx context.Context, request *pb.GetCarRequest) (*pb.GetCarResponse, error) {
	id := request.Id
	db := database.DBConn

	var car car.Car
	db.Find(&car, id)

	response := &pb.GetCarResponse{
		Brand: car.Brand,
		Make:  car.Make,
	}
	return response, nil
}

func (*carServiceServer) CreateCar(ctx context.Context, request *pb.CreateCarRequest) (*pb.CreateCarResponse, error) {
	car := car.Car{
		Brand: request.Brand,
		Make:  request.Make,
	}

	db := database.DBConn
	db.Create(&car)

	response := &pb.CreateCarResponse{
		Brand: car.Brand,
		Make:  car.Make,
	}

	return response, nil
}

func (*carServiceServer) DeleteCar(ctx context.Context, request *pb.DeleteCarRequest) (*pb.DeleteCarResponse, error) {
	return nil, nil
}

func initDatabase() {
	var err error
	database.DBConn, err = gorm.Open("sqlite3", "cars.db")
	if err != nil {
		panic("Failed to connect to database")
	}
	log.Println("Database connection opened")

	database.DBConn.AutoMigrate(&car.Car{})
	log.Println("Automigration completed")
}

// @title Swagger Car API
// @version 1.0
// @description This is a sample golang server with fiber and gorm
// @host localhost:3000
// @basePath /api
func main() {
	lis, err := net.Listen("tcp", "0.0.0.0:3000")
	if err != nil {
		log.Fatalf("Failed to listen on port 3000: %v", err)
	}

	initDatabase()
	defer database.DBConn.Close()

	grpcServer := grpc.NewServer()
	pb.RegisterCarServiceServer(grpcServer, &carServiceServer{})

	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("Failed to serve gRPC server over port 3000 %v", err)
	}
}
