package user

import (
	"Api-GateWay/protos/userproto"
	"log"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func DialGRPCUser() userproto.UserServiceClient {
	conn, err := grpc.NewClient("user-service:8888", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatal("failed to dial grpc client:", err)
	}

	return userproto.NewUserServiceClient(conn)
}
