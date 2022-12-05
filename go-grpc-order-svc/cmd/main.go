package main

import (
	"fmt"
	"github.com/sercheo87/go-grpc-order-svc/pkg/client"
	"github.com/sercheo87/go-grpc-order-svc/pkg/config"
	"github.com/sercheo87/go-grpc-order-svc/pkg/db"
	"github.com/sercheo87/go-grpc-order-svc/pkg/pb"
	service "github.com/sercheo87/go-grpc-order-svc/pkg/services"
	"google.golang.org/grpc"
	"log"
	"net"
)

func main() {
	c, err := config.LoadConfig()

	if err != nil {
		log.Fatalln("Failed at config", err)
	}

	h := db.Init(c.DBUrl)

	lis, err := net.Listen("tcp", c.Port)

	if err != nil {
		log.Fatalln("Failed to listing:", err)
	}

	productSvc := client.InitProductServiceClient(c.ProductSvcUrl)

	if err != nil {
		log.Fatalln("Failed to listing:", err)
	}

	fmt.Println("Order Svc on", c.Port)

	s := service.Server{
		H:          h,
		ProductSvc: productSvc,
	}

	grpcServer := grpc.NewServer()

	pb.RegisterOrderServiceServer(grpcServer, &s)

	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalln("Failed to serve:", err)
	}
}
