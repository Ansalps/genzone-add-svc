package main

import (
	"fmt"
	"log"
	"net"

	"github.com/Ansalps/genzone-add-svc/pkg/config"
	"github.com/Ansalps/genzone-add-svc/pkg/db"
	"google.golang.org/grpc"
	"github.com/Ansalps/genzone-add-svc/pkg/services"
	"github.com/Ansalps/genzone-add-svc/pkg/pb"
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
	fmt.Println("Product Svc on", c.Port)
	s := services.Server{
		H: h,
	}
	grpcServer := grpc.NewServer()
	pb.RegisterAddressServiceServer(grpcServer, &s)
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalln("Failed to serve:", err)
	}
}
