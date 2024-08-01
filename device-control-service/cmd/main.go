package main

import (
	"Device-service/internal/mongodb"
	"Device-service/internal/rabbitmq"
	pb "Device-service/pb/deviceproto"
	deviceservice "Device-service/service/device_service"
	"log"
	"net"

	"google.golang.org/grpc"
)

func main() {
	mongoclient, err := mongodb.NewDeviceMongoDb()
	if err != nil {
		log.Fatal(err)
	}

	deviceservice := deviceservice.NewDeviceService(mongoclient)
	server := grpc.NewServer()
	pb.RegisterDeviceServiceServer(server, deviceservice)

	go rabbitmq.ConsumeOrders()

	lis, err := net.Listen("tcp", ":8888")
	if err != nil {
		log.Fatalf("Tarmoqli ulanishda xatolik: %v", err)
	}

	log.Println("Server 8888 portda ishga tushdi")

	if err := server.Serve(lis); err != nil {
		log.Fatalf("Serverni ishga tushirishda xatolik: %v", err)
	}
}
