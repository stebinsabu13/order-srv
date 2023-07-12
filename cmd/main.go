package main

import (
	"fmt"
	"log"
	"net"

	"github.com/stebin13/order-srv/pkg/client"
	"github.com/stebin13/order-srv/pkg/config"
	"github.com/stebin13/order-srv/pkg/db"
	"github.com/stebin13/order-srv/pkg/pb"
	"github.com/stebin13/order-srv/pkg/services"
	"google.golang.org/grpc"
)

func main() {
	c, err := config.LoadConfig()
	if err != nil {
		log.Fatalln("failed to load config", err)
	}
	h := db.InitDb(&c)

	lis, errr := net.Listen("tcp", c.Port)
	if errr != nil {
		log.Fatalln("failed to listen", errr)
	}

	prodCli := client.InitProdCli(&c)

	fmt.Println("Order Svc on", c.Port)

	s := services.Server{
		H:       h,
		ProdCli: prodCli,
	}
	grpcServer := grpc.NewServer()
	pb.RegisterOrderServiceServer(grpcServer, &s)

	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalln("Failed to serve:", err)
	}
}
