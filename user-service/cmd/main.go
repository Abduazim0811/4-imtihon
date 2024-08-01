package main

import (
	"UserService/internal/mongodb"
	redisdb "UserService/internal/redis"
	userservice "UserService/service/UserService"
	pb "UserService/userproto"
	"log"
	"net"

	"google.golang.org/grpc"
)

func main() {
	redisClient := redisdb.NewRedisClient("user-redis:6379", "", 0)
	mongoclient, err := mongodb.NewUser()
	if err != nil {
		log.Fatalf("MongoDB ulanishda xatolik: %v", err)
	}

	userService := userservice.NewUserService(mongoclient, redisClient)
	server := grpc.NewServer()
	pb.RegisterUserServiceServer(server, userService)

	lis, err := net.Listen("tcp", ":8888")
	if err != nil {
		log.Fatalf("Tarmoqli ulanishda xatolik: %v", err)
	}

	log.Println("Server 8888 portda ishga tushdi")

	if err := server.Serve(lis); err != nil {
		log.Fatalf("Serverni ishga tushirishda xatolik: %v", err)
	}

}
