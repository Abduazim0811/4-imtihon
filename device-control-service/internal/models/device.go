package models


type ConfigurationSettings struct {
	Brightness int32  `bson:"brightness"`
	Color      string `bson:"color"`
}

type Device struct {
	DeviceID            string                `bson:"_id,omitempty"`
	UserID              string                `bson:"user_id"`
	DeviceType          string                `bson:"device_type"`
	DeviceName          string                `bson:"device_name"`
	DeviceStatus        string                `bson:"device_status"`
	ConfigurationSettings ConfigurationSettings `bson:"configuration_settings"`
	LastUpdated         string                `bson:"last_updated"`
	Location            string                `bson:"location"`
	FirmwareVersion     string                `bson:"firmware_version"`
	ConnectivityStatus  string                `bson:"connectivity_status"`
}

type CommandPayload struct {
	Brightness int32 `bson:"brightness"`
}

type Command struct {
	ID              string         `bson:"_id,omitempty"`
	DeviceID        string         `bson:"device_id"`
	UserID          string         `bson:"user_id"`
	Type            string         `bson:"type"`
	CommandPayload  CommandPayload `bson:"command_payload"`
	TimeStamp       string         `bson:"timeStamp"`
	Status          string         `bson:"status"`
}