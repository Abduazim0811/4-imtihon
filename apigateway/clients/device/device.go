package device

import (
	"Api-GateWay/protos/deviceproto"
	"log"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func DialGrpcDevice() deviceproto.DeviceServiceClient {
	conn, err := grpc.NewClient("device-service:8888", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatal("failed to dial grpc client:", err)
	}

	return deviceproto.NewDeviceServiceClient(conn)

}
