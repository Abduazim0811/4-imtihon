package userhandler

import (
	"Device-service/pb/userproto"
	"context"
	"time"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)


func Connect (id string) (*userproto.User, error){
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	conn, err := grpc.NewClient(":8888", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil{
		return nil, err
	}

	client := userproto.NewUserServiceClient(conn)

	res, err := client.GetByIdUser(ctx, &userproto.GetUserRequest{UserId: id})
	if err != nil {
		return nil, err
	}
	return res, err
}