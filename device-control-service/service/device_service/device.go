package deviceservice

import (
	"Device-service/internal/handlers/userhandler"
	"Device-service/internal/mongodb"
	"Device-service/internal/rabbitmq"
	pb "Device-service/pb/deviceproto"
	"context"
	"encoding/json"
	"fmt"
)

type DeviseService struct {
	pb.UnimplementedDeviceServiceServer
	mongodb *mongodb.DeviceMongoDb
}

func NewDeviceService(mongodb *mongodb.DeviceMongoDb) *DeviseService {
	return &DeviseService{mongodb: mongodb}
}

func (d *DeviseService) CreateDevice(ctx context.Context, req *pb.DeviceRequest) (*pb.DeviceResponse, error) {
	_, err := userhandler.Connect(req.UserId)
	if err != nil {
		return nil, fmt.Errorf("user not found")
	}
	id, err := d.mongodb.AddDevice(ctx, req)
	if err != nil {
		return nil, err
	}

	return &pb.DeviceResponse{DeviceId: id}, nil
}

func (d *DeviseService) GetDevicebyId(ctx context.Context, req *pb.DeviceResponse) (*pb.Device, error) {
	device, err := d.mongodb.GetbyIdDevice(ctx, req.DeviceId)
	if err != nil {
		return nil, err
	}

	return device, err
}

func (d *DeviseService) UpdateDevice(ctx context.Context, req *pb.Device) (*pb.DeviceOperationResponse, error) {
	err := d.mongodb.UpdateDevice(ctx, req)

	if err != nil {
		return nil, err
	}

	return &pb.DeviceOperationResponse{Message: "Update successfully"}, nil
}

func (d *DeviseService) DeleteDevice(ctx context.Context, req *pb.DeviceResponse) (*pb.DeviceOperationResponse, error) {
	err := d.mongodb.DeleteDevice(ctx, req.DeviceId)
	if err != nil {
		return nil, err
	}

	return &pb.DeviceOperationResponse{Message: "Device deleted"}, nil
}

func (d *DeviseService) CreateCommand(ctx context.Context, req *pb.CommandRequest) (*pb.CommandResponse, error) {
	command := &pb.Command{
		DeviceId:       req.GetDeviceId(),
		UserId:         req.GetUserId(),
		Type:           req.GetType(),
		CommandPayload: req.GetCommandPayload(),
		TimeStamp:      req.GetTimeStamp(),
		Status:         req.GetStatus(),
	}
	_, err := d.mongodb.GetbyIdDevice(ctx, command.DeviceId)
	if err != nil {
		return nil, fmt.Errorf("device not found")
	}

	_, err = userhandler.Connect(req.UserId)
	if err != nil {
		return nil, fmt.Errorf("user not found")
	}

	commandData, err := json.Marshal(command)
	if err != nil {
		return nil, err
	}

	err = rabbitmq.PublishOrder(commandData)
	if err != nil {
		return nil, err
	}

	return &pb.CommandResponse{Id: command.GetId()}, nil
}

func (d *DeviseService)GetCommandById(ctx context.Context, req *pb.CommandResponse) (*pb.Command, error) {
	command, err := d.mongodb.GetCommandbyId(ctx, req.Id)
	if err != nil {
		return nil, err
	}
	return command, nil
}
