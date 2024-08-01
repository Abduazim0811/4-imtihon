package mongodb

import (
	"Device-service/internal/models"
	"Device-service/pb/deviceproto"
	pb "Device-service/pb/deviceproto"
	"context"
	"fmt"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type DeviceMongoDb struct {
	client     *mongo.Client
	collection *mongo.Collection
}

func NewDeviceMongoDb() (*DeviceMongoDb, error) {
	clientOptions := options.Client().ApplyURI("mongodb://device-mongodb:27017")

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	client, err := mongo.Connect(ctx, clientOptions)
	if err != nil {
		return nil, err
	}

	err = client.Ping(ctx, nil)
	if err != nil {
		return nil, err
	}

	collection := client.Database("Device").Collection("device")
	return &DeviceMongoDb{client: client, collection: collection}, nil
}

func (d *DeviceMongoDb) AddDevice(ctx context.Context, req *pb.DeviceRequest) (string, error) {
	device := bson.M{
		"user_id":       req.GetUserId(),
		"device_type":   req.GetDeviceType(),
		"device_name":   req.GetDeviceName(),
		"device_status": req.GetDeviceStatus(),
		"configuration_settings": bson.M{
			"brightness": req.GetConfigurationSettings().GetBrightness(),
			"color":      req.GetConfigurationSettings().GetColor(),
		},
		"last_updated":        req.GetLastUpdated(),
		"location":            req.GetLocation(),
		"firmware_version":    req.GetFirmwareVersion(),
		"connectivity_status": req.GetConnectivityStatus(),
	}

	result, err := d.collection.InsertOne(ctx, device)
	if err != nil {
		return "", err
	}

	deviceID := result.InsertedID.(primitive.ObjectID).Hex()
	return deviceID, nil
}

func (d *DeviceMongoDb) GetbyIdDevice(ctx context.Context, deviceID string) (*pb.Device, error) {
	objID, err := primitive.ObjectIDFromHex(deviceID)
	if err != nil {
		return nil, fmt.Errorf("primitive id error: %v", err)
	}
	var device pb.Device
	var devicemongo models.Device

	res := d.collection.FindOne(ctx, bson.M{"_id": objID})

	err = res.Decode(&devicemongo)
	if err != nil {
		return nil, fmt.Errorf("error decode:%v", err)
	}
	device.DeviceId = devicemongo.DeviceID
	device.UserId = devicemongo.UserID
	device.DeviceType = devicemongo.DeviceType
	device.DeviceName = devicemongo.DeviceName
	device.DeviceStatus = devicemongo.DeviceStatus
	if devicemongo.ConfigurationSettings != (models.ConfigurationSettings{}) {
		device.ConfigurationSettings = &pb.Configuration_Settings{
			Brightness: devicemongo.ConfigurationSettings.Brightness,
			Color:      devicemongo.ConfigurationSettings.Color,
		}
	}
	device.LastUpdated = devicemongo.LastUpdated
	device.Location = devicemongo.Location
	device.FirmwareVersion = devicemongo.FirmwareVersion
	device.ConnectivityStatus = devicemongo.ConnectivityStatus

	return &device, nil
}

func (d *DeviceMongoDb) UpdateDevice(ctx context.Context, req *pb.Device) error {
	objID, err := primitive.ObjectIDFromHex(req.GetDeviceId())
	if err != nil {
		return err
	}

	update := bson.M{
		"$set": bson.M{
			"user_id":       req.GetUserId(),
			"device_type":   req.GetDeviceType(),
			"device_name":   req.GetDeviceName(),
			"device_status": req.GetDeviceStatus(),
			"configuration_settings": bson.M{
				"brightness": req.GetConfigurationSettings().GetBrightness(),
				"color":      req.GetConfigurationSettings().GetColor(),
			},
			"last_updated":        req.GetLastUpdated(),
			"location":            req.GetLocation(),
			"firmware_version":    req.GetFirmwareVersion(),
			"connectivity_status": req.GetConnectivityStatus(),
		},
	}

	_, err = d.collection.UpdateOne(ctx, bson.M{"_id": objID}, update)
	return err
}

func (d *DeviceMongoDb) DeleteDevice(ctx context.Context, id string) error {
	objID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return err
	}

	_, err = d.collection.DeleteOne(ctx, bson.M{"_id": objID})
	return err
}

func (d *DeviceMongoDb) AddCommand(ctx context.Context, req *deviceproto.CommandRequest) (string, error) {
	command := bson.M{
		"device_id": req.GetDeviceId(),
		"user_id":   req.GetUserId(),
		"type":      req.GetType(),
		"command_payload": bson.M{
			"brightness": req.GetCommandPayload().GetBrightness(),
		},
		"timeStamp": req.GetTimeStamp(),
		"status":    req.GetStatus(),
	}

	res, err := d.collection.InsertOne(ctx, command)
	if err != nil {
		return "", err
	}

	commandID := res.InsertedID.(primitive.ObjectID).Hex()
	return commandID, nil

}

func (d *DeviceMongoDb) GetCommandbyId(ctx context.Context, commandID string) (*deviceproto.Command, error) {
	objID, err := primitive.ObjectIDFromHex(commandID)
	if err != nil {
		return nil, err
	}

	var command deviceproto.Command
	res := d.collection.FindOne(ctx, bson.M{"_id": objID})

	err = res.Decode(&command)
	if err != nil {
		return nil, err
	}

	return &command, nil
}
