// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.34.2
// 	protoc        v5.27.1
// source: pb/deviceproto/device.proto

package deviceproto

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type DeviceOperationResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Message string `protobuf:"bytes,1,opt,name=message,proto3" json:"message,omitempty"`
}

func (x *DeviceOperationResponse) Reset() {
	*x = DeviceOperationResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pb_deviceproto_device_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeviceOperationResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeviceOperationResponse) ProtoMessage() {}

func (x *DeviceOperationResponse) ProtoReflect() protoreflect.Message {
	mi := &file_pb_deviceproto_device_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeviceOperationResponse.ProtoReflect.Descriptor instead.
func (*DeviceOperationResponse) Descriptor() ([]byte, []int) {
	return file_pb_deviceproto_device_proto_rawDescGZIP(), []int{0}
}

func (x *DeviceOperationResponse) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

type Device struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	DeviceId              string                  `protobuf:"bytes,1,opt,name=device_id,json=deviceId,proto3" json:"device_id,omitempty"`
	UserId                string                  `protobuf:"bytes,2,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	DeviceType            string                  `protobuf:"bytes,3,opt,name=device_type,json=deviceType,proto3" json:"device_type,omitempty"`
	DeviceName            string                  `protobuf:"bytes,4,opt,name=device_name,json=deviceName,proto3" json:"device_name,omitempty"`
	DeviceStatus          string                  `protobuf:"bytes,5,opt,name=device_status,json=deviceStatus,proto3" json:"device_status,omitempty"`
	ConfigurationSettings *Configuration_Settings `protobuf:"bytes,6,opt,name=configuration_settings,json=configurationSettings,proto3" json:"configuration_settings,omitempty"`
	LastUpdated           string                  `protobuf:"bytes,7,opt,name=last_updated,json=lastUpdated,proto3" json:"last_updated,omitempty"`
	Location              string                  `protobuf:"bytes,8,opt,name=location,proto3" json:"location,omitempty"`
	FirmwareVersion       string                  `protobuf:"bytes,9,opt,name=firmware_version,json=firmwareVersion,proto3" json:"firmware_version,omitempty"`
	ConnectivityStatus    string                  `protobuf:"bytes,10,opt,name=connectivity_status,json=connectivityStatus,proto3" json:"connectivity_status,omitempty"`
}

func (x *Device) Reset() {
	*x = Device{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pb_deviceproto_device_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Device) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Device) ProtoMessage() {}

func (x *Device) ProtoReflect() protoreflect.Message {
	mi := &file_pb_deviceproto_device_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Device.ProtoReflect.Descriptor instead.
func (*Device) Descriptor() ([]byte, []int) {
	return file_pb_deviceproto_device_proto_rawDescGZIP(), []int{1}
}

func (x *Device) GetDeviceId() string {
	if x != nil {
		return x.DeviceId
	}
	return ""
}

func (x *Device) GetUserId() string {
	if x != nil {
		return x.UserId
	}
	return ""
}

func (x *Device) GetDeviceType() string {
	if x != nil {
		return x.DeviceType
	}
	return ""
}

func (x *Device) GetDeviceName() string {
	if x != nil {
		return x.DeviceName
	}
	return ""
}

func (x *Device) GetDeviceStatus() string {
	if x != nil {
		return x.DeviceStatus
	}
	return ""
}

func (x *Device) GetConfigurationSettings() *Configuration_Settings {
	if x != nil {
		return x.ConfigurationSettings
	}
	return nil
}

func (x *Device) GetLastUpdated() string {
	if x != nil {
		return x.LastUpdated
	}
	return ""
}

func (x *Device) GetLocation() string {
	if x != nil {
		return x.Location
	}
	return ""
}

func (x *Device) GetFirmwareVersion() string {
	if x != nil {
		return x.FirmwareVersion
	}
	return ""
}

func (x *Device) GetConnectivityStatus() string {
	if x != nil {
		return x.ConnectivityStatus
	}
	return ""
}

type Configuration_Settings struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Brightness int32  `protobuf:"varint,1,opt,name=brightness,proto3" json:"brightness,omitempty"`
	Color      string `protobuf:"bytes,2,opt,name=color,proto3" json:"color,omitempty"`
}

func (x *Configuration_Settings) Reset() {
	*x = Configuration_Settings{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pb_deviceproto_device_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Configuration_Settings) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Configuration_Settings) ProtoMessage() {}

func (x *Configuration_Settings) ProtoReflect() protoreflect.Message {
	mi := &file_pb_deviceproto_device_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Configuration_Settings.ProtoReflect.Descriptor instead.
func (*Configuration_Settings) Descriptor() ([]byte, []int) {
	return file_pb_deviceproto_device_proto_rawDescGZIP(), []int{2}
}

func (x *Configuration_Settings) GetBrightness() int32 {
	if x != nil {
		return x.Brightness
	}
	return 0
}

func (x *Configuration_Settings) GetColor() string {
	if x != nil {
		return x.Color
	}
	return ""
}

type Command struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id             string           `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	DeviceId       string           `protobuf:"bytes,2,opt,name=device_id,json=deviceId,proto3" json:"device_id,omitempty"`
	UserId         string           `protobuf:"bytes,3,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	Type           string           `protobuf:"bytes,4,opt,name=type,proto3" json:"type,omitempty"`
	CommandPayload *Command_Payload `protobuf:"bytes,5,opt,name=command_payload,json=commandPayload,proto3" json:"command_payload,omitempty"`
	TimeStamp      string           `protobuf:"bytes,6,opt,name=timeStamp,proto3" json:"timeStamp,omitempty"`
	Status         string           `protobuf:"bytes,7,opt,name=status,proto3" json:"status,omitempty"`
}

func (x *Command) Reset() {
	*x = Command{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pb_deviceproto_device_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Command) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Command) ProtoMessage() {}

func (x *Command) ProtoReflect() protoreflect.Message {
	mi := &file_pb_deviceproto_device_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Command.ProtoReflect.Descriptor instead.
func (*Command) Descriptor() ([]byte, []int) {
	return file_pb_deviceproto_device_proto_rawDescGZIP(), []int{3}
}

func (x *Command) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *Command) GetDeviceId() string {
	if x != nil {
		return x.DeviceId
	}
	return ""
}

func (x *Command) GetUserId() string {
	if x != nil {
		return x.UserId
	}
	return ""
}

func (x *Command) GetType() string {
	if x != nil {
		return x.Type
	}
	return ""
}

func (x *Command) GetCommandPayload() *Command_Payload {
	if x != nil {
		return x.CommandPayload
	}
	return nil
}

func (x *Command) GetTimeStamp() string {
	if x != nil {
		return x.TimeStamp
	}
	return ""
}

func (x *Command) GetStatus() string {
	if x != nil {
		return x.Status
	}
	return ""
}

type Command_Payload struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Brightness int32 `protobuf:"varint,1,opt,name=brightness,proto3" json:"brightness,omitempty"`
}

func (x *Command_Payload) Reset() {
	*x = Command_Payload{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pb_deviceproto_device_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Command_Payload) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Command_Payload) ProtoMessage() {}

func (x *Command_Payload) ProtoReflect() protoreflect.Message {
	mi := &file_pb_deviceproto_device_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Command_Payload.ProtoReflect.Descriptor instead.
func (*Command_Payload) Descriptor() ([]byte, []int) {
	return file_pb_deviceproto_device_proto_rawDescGZIP(), []int{4}
}

func (x *Command_Payload) GetBrightness() int32 {
	if x != nil {
		return x.Brightness
	}
	return 0
}

type DeviceRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserId                string                  `protobuf:"bytes,1,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	DeviceType            string                  `protobuf:"bytes,2,opt,name=device_type,json=deviceType,proto3" json:"device_type,omitempty"`
	DeviceName            string                  `protobuf:"bytes,3,opt,name=device_name,json=deviceName,proto3" json:"device_name,omitempty"`
	DeviceStatus          string                  `protobuf:"bytes,4,opt,name=device_status,json=deviceStatus,proto3" json:"device_status,omitempty"`
	ConfigurationSettings *Configuration_Settings `protobuf:"bytes,5,opt,name=configuration_settings,json=configurationSettings,proto3" json:"configuration_settings,omitempty"`
	LastUpdated           string                  `protobuf:"bytes,6,opt,name=last_updated,json=lastUpdated,proto3" json:"last_updated,omitempty"`
	Location              string                  `protobuf:"bytes,7,opt,name=location,proto3" json:"location,omitempty"`
	FirmwareVersion       string                  `protobuf:"bytes,8,opt,name=firmware_version,json=firmwareVersion,proto3" json:"firmware_version,omitempty"`
	ConnectivityStatus    string                  `protobuf:"bytes,9,opt,name=connectivity_status,json=connectivityStatus,proto3" json:"connectivity_status,omitempty"`
}

func (x *DeviceRequest) Reset() {
	*x = DeviceRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pb_deviceproto_device_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeviceRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeviceRequest) ProtoMessage() {}

func (x *DeviceRequest) ProtoReflect() protoreflect.Message {
	mi := &file_pb_deviceproto_device_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeviceRequest.ProtoReflect.Descriptor instead.
func (*DeviceRequest) Descriptor() ([]byte, []int) {
	return file_pb_deviceproto_device_proto_rawDescGZIP(), []int{5}
}

func (x *DeviceRequest) GetUserId() string {
	if x != nil {
		return x.UserId
	}
	return ""
}

func (x *DeviceRequest) GetDeviceType() string {
	if x != nil {
		return x.DeviceType
	}
	return ""
}

func (x *DeviceRequest) GetDeviceName() string {
	if x != nil {
		return x.DeviceName
	}
	return ""
}

func (x *DeviceRequest) GetDeviceStatus() string {
	if x != nil {
		return x.DeviceStatus
	}
	return ""
}

func (x *DeviceRequest) GetConfigurationSettings() *Configuration_Settings {
	if x != nil {
		return x.ConfigurationSettings
	}
	return nil
}

func (x *DeviceRequest) GetLastUpdated() string {
	if x != nil {
		return x.LastUpdated
	}
	return ""
}

func (x *DeviceRequest) GetLocation() string {
	if x != nil {
		return x.Location
	}
	return ""
}

func (x *DeviceRequest) GetFirmwareVersion() string {
	if x != nil {
		return x.FirmwareVersion
	}
	return ""
}

func (x *DeviceRequest) GetConnectivityStatus() string {
	if x != nil {
		return x.ConnectivityStatus
	}
	return ""
}

type DeviceResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	DeviceId string `protobuf:"bytes,1,opt,name=device_id,json=deviceId,proto3" json:"device_id,omitempty"`
}

func (x *DeviceResponse) Reset() {
	*x = DeviceResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pb_deviceproto_device_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeviceResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeviceResponse) ProtoMessage() {}

func (x *DeviceResponse) ProtoReflect() protoreflect.Message {
	mi := &file_pb_deviceproto_device_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeviceResponse.ProtoReflect.Descriptor instead.
func (*DeviceResponse) Descriptor() ([]byte, []int) {
	return file_pb_deviceproto_device_proto_rawDescGZIP(), []int{6}
}

func (x *DeviceResponse) GetDeviceId() string {
	if x != nil {
		return x.DeviceId
	}
	return ""
}

type CommandRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	DeviceId       string           `protobuf:"bytes,1,opt,name=device_id,json=deviceId,proto3" json:"device_id,omitempty"`
	UserId         string           `protobuf:"bytes,2,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	Type           string           `protobuf:"bytes,3,opt,name=type,proto3" json:"type,omitempty"`
	CommandPayload *Command_Payload `protobuf:"bytes,4,opt,name=command_payload,json=commandPayload,proto3" json:"command_payload,omitempty"`
	TimeStamp      string           `protobuf:"bytes,5,opt,name=timeStamp,proto3" json:"timeStamp,omitempty"`
	Status         string           `protobuf:"bytes,6,opt,name=status,proto3" json:"status,omitempty"`
}

func (x *CommandRequest) Reset() {
	*x = CommandRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pb_deviceproto_device_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CommandRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CommandRequest) ProtoMessage() {}

func (x *CommandRequest) ProtoReflect() protoreflect.Message {
	mi := &file_pb_deviceproto_device_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CommandRequest.ProtoReflect.Descriptor instead.
func (*CommandRequest) Descriptor() ([]byte, []int) {
	return file_pb_deviceproto_device_proto_rawDescGZIP(), []int{7}
}

func (x *CommandRequest) GetDeviceId() string {
	if x != nil {
		return x.DeviceId
	}
	return ""
}

func (x *CommandRequest) GetUserId() string {
	if x != nil {
		return x.UserId
	}
	return ""
}

func (x *CommandRequest) GetType() string {
	if x != nil {
		return x.Type
	}
	return ""
}

func (x *CommandRequest) GetCommandPayload() *Command_Payload {
	if x != nil {
		return x.CommandPayload
	}
	return nil
}

func (x *CommandRequest) GetTimeStamp() string {
	if x != nil {
		return x.TimeStamp
	}
	return ""
}

func (x *CommandRequest) GetStatus() string {
	if x != nil {
		return x.Status
	}
	return ""
}

type CommandResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *CommandResponse) Reset() {
	*x = CommandResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pb_deviceproto_device_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CommandResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CommandResponse) ProtoMessage() {}

func (x *CommandResponse) ProtoReflect() protoreflect.Message {
	mi := &file_pb_deviceproto_device_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CommandResponse.ProtoReflect.Descriptor instead.
func (*CommandResponse) Descriptor() ([]byte, []int) {
	return file_pb_deviceproto_device_proto_rawDescGZIP(), []int{8}
}

func (x *CommandResponse) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

var File_pb_deviceproto_device_proto protoreflect.FileDescriptor

var file_pb_deviceproto_device_proto_rawDesc = []byte{
	0x0a, 0x1b, 0x70, 0x62, 0x2f, 0x64, 0x65, 0x76, 0x69, 0x63, 0x65, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x2f, 0x64, 0x65, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x33, 0x0a,
	0x17, 0x44, 0x65, 0x76, 0x69, 0x63, 0x65, 0x4f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x6d, 0x65, 0x73, 0x73,
	0x61, 0x67, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61,
	0x67, 0x65, 0x22, 0x90, 0x03, 0x0a, 0x06, 0x44, 0x65, 0x76, 0x69, 0x63, 0x65, 0x12, 0x1b, 0x0a,
	0x09, 0x64, 0x65, 0x76, 0x69, 0x63, 0x65, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x08, 0x64, 0x65, 0x76, 0x69, 0x63, 0x65, 0x49, 0x64, 0x12, 0x17, 0x0a, 0x07, 0x75, 0x73,
	0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x75, 0x73, 0x65,
	0x72, 0x49, 0x64, 0x12, 0x1f, 0x0a, 0x0b, 0x64, 0x65, 0x76, 0x69, 0x63, 0x65, 0x5f, 0x74, 0x79,
	0x70, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x64, 0x65, 0x76, 0x69, 0x63, 0x65,
	0x54, 0x79, 0x70, 0x65, 0x12, 0x1f, 0x0a, 0x0b, 0x64, 0x65, 0x76, 0x69, 0x63, 0x65, 0x5f, 0x6e,
	0x61, 0x6d, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x64, 0x65, 0x76, 0x69, 0x63,
	0x65, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x23, 0x0a, 0x0d, 0x64, 0x65, 0x76, 0x69, 0x63, 0x65, 0x5f,
	0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x64, 0x65,
	0x76, 0x69, 0x63, 0x65, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x4e, 0x0a, 0x16, 0x63, 0x6f,
	0x6e, 0x66, 0x69, 0x67, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x73, 0x65, 0x74, 0x74,
	0x69, 0x6e, 0x67, 0x73, 0x18, 0x06, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x17, 0x2e, 0x43, 0x6f, 0x6e,
	0x66, 0x69, 0x67, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x53, 0x65, 0x74, 0x74, 0x69,
	0x6e, 0x67, 0x73, 0x52, 0x15, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x75, 0x72, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x53, 0x65, 0x74, 0x74, 0x69, 0x6e, 0x67, 0x73, 0x12, 0x21, 0x0a, 0x0c, 0x6c, 0x61,
	0x73, 0x74, 0x5f, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x0b, 0x6c, 0x61, 0x73, 0x74, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x12, 0x1a, 0x0a,
	0x08, 0x6c, 0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x08, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x08, 0x6c, 0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x29, 0x0a, 0x10, 0x66, 0x69, 0x72,
	0x6d, 0x77, 0x61, 0x72, 0x65, 0x5f, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x18, 0x09, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x0f, 0x66, 0x69, 0x72, 0x6d, 0x77, 0x61, 0x72, 0x65, 0x56, 0x65, 0x72,
	0x73, 0x69, 0x6f, 0x6e, 0x12, 0x2f, 0x0a, 0x13, 0x63, 0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74, 0x69,
	0x76, 0x69, 0x74, 0x79, 0x5f, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x0a, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x12, 0x63, 0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74, 0x69, 0x76, 0x69, 0x74, 0x79, 0x53,
	0x74, 0x61, 0x74, 0x75, 0x73, 0x22, 0x4e, 0x0a, 0x16, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x75,
	0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x53, 0x65, 0x74, 0x74, 0x69, 0x6e, 0x67, 0x73, 0x12,
	0x1e, 0x0a, 0x0a, 0x62, 0x72, 0x69, 0x67, 0x68, 0x74, 0x6e, 0x65, 0x73, 0x73, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x05, 0x52, 0x0a, 0x62, 0x72, 0x69, 0x67, 0x68, 0x74, 0x6e, 0x65, 0x73, 0x73, 0x12,
	0x14, 0x0a, 0x05, 0x63, 0x6f, 0x6c, 0x6f, 0x72, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05,
	0x63, 0x6f, 0x6c, 0x6f, 0x72, 0x22, 0xd4, 0x01, 0x0a, 0x07, 0x43, 0x6f, 0x6d, 0x6d, 0x61, 0x6e,
	0x64, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69,
	0x64, 0x12, 0x1b, 0x0a, 0x09, 0x64, 0x65, 0x76, 0x69, 0x63, 0x65, 0x5f, 0x69, 0x64, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x64, 0x65, 0x76, 0x69, 0x63, 0x65, 0x49, 0x64, 0x12, 0x17,
	0x0a, 0x07, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x74, 0x79, 0x70, 0x65, 0x18,
	0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x74, 0x79, 0x70, 0x65, 0x12, 0x39, 0x0a, 0x0f, 0x63,
	0x6f, 0x6d, 0x6d, 0x61, 0x6e, 0x64, 0x5f, 0x70, 0x61, 0x79, 0x6c, 0x6f, 0x61, 0x64, 0x18, 0x05,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x10, 0x2e, 0x43, 0x6f, 0x6d, 0x6d, 0x61, 0x6e, 0x64, 0x5f, 0x50,
	0x61, 0x79, 0x6c, 0x6f, 0x61, 0x64, 0x52, 0x0e, 0x63, 0x6f, 0x6d, 0x6d, 0x61, 0x6e, 0x64, 0x50,
	0x61, 0x79, 0x6c, 0x6f, 0x61, 0x64, 0x12, 0x1c, 0x0a, 0x09, 0x74, 0x69, 0x6d, 0x65, 0x53, 0x74,
	0x61, 0x6d, 0x70, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x74, 0x69, 0x6d, 0x65, 0x53,
	0x74, 0x61, 0x6d, 0x70, 0x12, 0x16, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x07,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x22, 0x31, 0x0a, 0x0f,
	0x43, 0x6f, 0x6d, 0x6d, 0x61, 0x6e, 0x64, 0x5f, 0x50, 0x61, 0x79, 0x6c, 0x6f, 0x61, 0x64, 0x12,
	0x1e, 0x0a, 0x0a, 0x62, 0x72, 0x69, 0x67, 0x68, 0x74, 0x6e, 0x65, 0x73, 0x73, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x05, 0x52, 0x0a, 0x62, 0x72, 0x69, 0x67, 0x68, 0x74, 0x6e, 0x65, 0x73, 0x73, 0x22,
	0xfa, 0x02, 0x0a, 0x0d, 0x44, 0x65, 0x76, 0x69, 0x63, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x12, 0x17, 0x0a, 0x07, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x12, 0x1f, 0x0a, 0x0b, 0x64, 0x65,
	0x76, 0x69, 0x63, 0x65, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x0a, 0x64, 0x65, 0x76, 0x69, 0x63, 0x65, 0x54, 0x79, 0x70, 0x65, 0x12, 0x1f, 0x0a, 0x0b, 0x64,
	0x65, 0x76, 0x69, 0x63, 0x65, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x0a, 0x64, 0x65, 0x76, 0x69, 0x63, 0x65, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x23, 0x0a, 0x0d,
	0x64, 0x65, 0x76, 0x69, 0x63, 0x65, 0x5f, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x04, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x0c, 0x64, 0x65, 0x76, 0x69, 0x63, 0x65, 0x53, 0x74, 0x61, 0x74, 0x75,
	0x73, 0x12, 0x4e, 0x0a, 0x16, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x75, 0x72, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x5f, 0x73, 0x65, 0x74, 0x74, 0x69, 0x6e, 0x67, 0x73, 0x18, 0x05, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x17, 0x2e, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x5f, 0x53, 0x65, 0x74, 0x74, 0x69, 0x6e, 0x67, 0x73, 0x52, 0x15, 0x63, 0x6f, 0x6e, 0x66,
	0x69, 0x67, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x53, 0x65, 0x74, 0x74, 0x69, 0x6e, 0x67,
	0x73, 0x12, 0x21, 0x0a, 0x0c, 0x6c, 0x61, 0x73, 0x74, 0x5f, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65,
	0x64, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x6c, 0x61, 0x73, 0x74, 0x55, 0x70, 0x64,
	0x61, 0x74, 0x65, 0x64, 0x12, 0x1a, 0x0a, 0x08, 0x6c, 0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x6c, 0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x12, 0x29, 0x0a, 0x10, 0x66, 0x69, 0x72, 0x6d, 0x77, 0x61, 0x72, 0x65, 0x5f, 0x76, 0x65, 0x72,
	0x73, 0x69, 0x6f, 0x6e, 0x18, 0x08, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0f, 0x66, 0x69, 0x72, 0x6d,
	0x77, 0x61, 0x72, 0x65, 0x56, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x12, 0x2f, 0x0a, 0x13, 0x63,
	0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74, 0x69, 0x76, 0x69, 0x74, 0x79, 0x5f, 0x73, 0x74, 0x61, 0x74,
	0x75, 0x73, 0x18, 0x09, 0x20, 0x01, 0x28, 0x09, 0x52, 0x12, 0x63, 0x6f, 0x6e, 0x6e, 0x65, 0x63,
	0x74, 0x69, 0x76, 0x69, 0x74, 0x79, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x22, 0x2d, 0x0a, 0x0e,
	0x44, 0x65, 0x76, 0x69, 0x63, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x1b,
	0x0a, 0x09, 0x64, 0x65, 0x76, 0x69, 0x63, 0x65, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x08, 0x64, 0x65, 0x76, 0x69, 0x63, 0x65, 0x49, 0x64, 0x22, 0xcb, 0x01, 0x0a, 0x0e,
	0x43, 0x6f, 0x6d, 0x6d, 0x61, 0x6e, 0x64, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1b,
	0x0a, 0x09, 0x64, 0x65, 0x76, 0x69, 0x63, 0x65, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x08, 0x64, 0x65, 0x76, 0x69, 0x63, 0x65, 0x49, 0x64, 0x12, 0x17, 0x0a, 0x07, 0x75,
	0x73, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x75, 0x73,
	0x65, 0x72, 0x49, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x74, 0x79, 0x70, 0x65, 0x18, 0x03, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x04, 0x74, 0x79, 0x70, 0x65, 0x12, 0x39, 0x0a, 0x0f, 0x63, 0x6f, 0x6d, 0x6d,
	0x61, 0x6e, 0x64, 0x5f, 0x70, 0x61, 0x79, 0x6c, 0x6f, 0x61, 0x64, 0x18, 0x04, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x10, 0x2e, 0x43, 0x6f, 0x6d, 0x6d, 0x61, 0x6e, 0x64, 0x5f, 0x50, 0x61, 0x79, 0x6c,
	0x6f, 0x61, 0x64, 0x52, 0x0e, 0x63, 0x6f, 0x6d, 0x6d, 0x61, 0x6e, 0x64, 0x50, 0x61, 0x79, 0x6c,
	0x6f, 0x61, 0x64, 0x12, 0x1c, 0x0a, 0x09, 0x74, 0x69, 0x6d, 0x65, 0x53, 0x74, 0x61, 0x6d, 0x70,
	0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x74, 0x69, 0x6d, 0x65, 0x53, 0x74, 0x61, 0x6d,
	0x70, 0x12, 0x16, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x06, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x22, 0x21, 0x0a, 0x0f, 0x43, 0x6f, 0x6d,
	0x6d, 0x61, 0x6e, 0x64, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x0e, 0x0a, 0x02,
	0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x32, 0xbb, 0x02, 0x0a,
	0x0d, 0x44, 0x65, 0x76, 0x69, 0x63, 0x65, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x2f,
	0x0a, 0x0c, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x44, 0x65, 0x76, 0x69, 0x63, 0x65, 0x12, 0x0e,
	0x2e, 0x44, 0x65, 0x76, 0x69, 0x63, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x0f,
	0x2e, 0x44, 0x65, 0x76, 0x69, 0x63, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12,
	0x29, 0x0a, 0x0d, 0x47, 0x65, 0x74, 0x44, 0x65, 0x76, 0x69, 0x63, 0x65, 0x42, 0x79, 0x49, 0x64,
	0x12, 0x0f, 0x2e, 0x44, 0x65, 0x76, 0x69, 0x63, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x1a, 0x07, 0x2e, 0x44, 0x65, 0x76, 0x69, 0x63, 0x65, 0x12, 0x31, 0x0a, 0x0c, 0x55, 0x70,
	0x64, 0x61, 0x74, 0x65, 0x44, 0x65, 0x76, 0x69, 0x63, 0x65, 0x12, 0x07, 0x2e, 0x44, 0x65, 0x76,
	0x69, 0x63, 0x65, 0x1a, 0x18, 0x2e, 0x44, 0x65, 0x76, 0x69, 0x63, 0x65, 0x4f, 0x70, 0x65, 0x72,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x39, 0x0a,
	0x0c, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x44, 0x65, 0x76, 0x69, 0x63, 0x65, 0x12, 0x0f, 0x2e,
	0x44, 0x65, 0x76, 0x69, 0x63, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x1a, 0x18,
	0x2e, 0x44, 0x65, 0x76, 0x69, 0x63, 0x65, 0x4f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x32, 0x0a, 0x0d, 0x43, 0x72, 0x65, 0x61,
	0x74, 0x65, 0x43, 0x6f, 0x6d, 0x6d, 0x61, 0x6e, 0x64, 0x12, 0x0f, 0x2e, 0x43, 0x6f, 0x6d, 0x6d,
	0x61, 0x6e, 0x64, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x10, 0x2e, 0x43, 0x6f, 0x6d,
	0x6d, 0x61, 0x6e, 0x64, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x2c, 0x0a, 0x0e,
	0x47, 0x65, 0x74, 0x43, 0x6f, 0x6d, 0x6d, 0x61, 0x6e, 0x64, 0x42, 0x79, 0x49, 0x64, 0x12, 0x10,
	0x2e, 0x43, 0x6f, 0x6d, 0x6d, 0x61, 0x6e, 0x64, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x1a, 0x08, 0x2e, 0x43, 0x6f, 0x6d, 0x6d, 0x61, 0x6e, 0x64, 0x42, 0x0f, 0x5a, 0x0d, 0x2e, 0x2f,
	0x64, 0x65, 0x76, 0x69, 0x63, 0x65, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x06, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x33,
}

var (
	file_pb_deviceproto_device_proto_rawDescOnce sync.Once
	file_pb_deviceproto_device_proto_rawDescData = file_pb_deviceproto_device_proto_rawDesc
)

func file_pb_deviceproto_device_proto_rawDescGZIP() []byte {
	file_pb_deviceproto_device_proto_rawDescOnce.Do(func() {
		file_pb_deviceproto_device_proto_rawDescData = protoimpl.X.CompressGZIP(file_pb_deviceproto_device_proto_rawDescData)
	})
	return file_pb_deviceproto_device_proto_rawDescData
}

var file_pb_deviceproto_device_proto_msgTypes = make([]protoimpl.MessageInfo, 9)
var file_pb_deviceproto_device_proto_goTypes = []any{
	(*DeviceOperationResponse)(nil), // 0: DeviceOperationResponse
	(*Device)(nil),                  // 1: Device
	(*Configuration_Settings)(nil),  // 2: Configuration_Settings
	(*Command)(nil),                 // 3: Command
	(*Command_Payload)(nil),         // 4: Command_Payload
	(*DeviceRequest)(nil),           // 5: DeviceRequest
	(*DeviceResponse)(nil),          // 6: DeviceResponse
	(*CommandRequest)(nil),          // 7: CommandRequest
	(*CommandResponse)(nil),         // 8: CommandResponse
}
var file_pb_deviceproto_device_proto_depIdxs = []int32{
	2,  // 0: Device.configuration_settings:type_name -> Configuration_Settings
	4,  // 1: Command.command_payload:type_name -> Command_Payload
	2,  // 2: DeviceRequest.configuration_settings:type_name -> Configuration_Settings
	4,  // 3: CommandRequest.command_payload:type_name -> Command_Payload
	5,  // 4: DeviceService.CreateDevice:input_type -> DeviceRequest
	6,  // 5: DeviceService.GetDeviceById:input_type -> DeviceResponse
	1,  // 6: DeviceService.UpdateDevice:input_type -> Device
	6,  // 7: DeviceService.DeleteDevice:input_type -> DeviceResponse
	7,  // 8: DeviceService.CreateCommand:input_type -> CommandRequest
	8,  // 9: DeviceService.GetCommandById:input_type -> CommandResponse
	6,  // 10: DeviceService.CreateDevice:output_type -> DeviceResponse
	1,  // 11: DeviceService.GetDeviceById:output_type -> Device
	0,  // 12: DeviceService.UpdateDevice:output_type -> DeviceOperationResponse
	0,  // 13: DeviceService.DeleteDevice:output_type -> DeviceOperationResponse
	8,  // 14: DeviceService.CreateCommand:output_type -> CommandResponse
	3,  // 15: DeviceService.GetCommandById:output_type -> Command
	10, // [10:16] is the sub-list for method output_type
	4,  // [4:10] is the sub-list for method input_type
	4,  // [4:4] is the sub-list for extension type_name
	4,  // [4:4] is the sub-list for extension extendee
	0,  // [0:4] is the sub-list for field type_name
}

func init() { file_pb_deviceproto_device_proto_init() }
func file_pb_deviceproto_device_proto_init() {
	if File_pb_deviceproto_device_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_pb_deviceproto_device_proto_msgTypes[0].Exporter = func(v any, i int) any {
			switch v := v.(*DeviceOperationResponse); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_pb_deviceproto_device_proto_msgTypes[1].Exporter = func(v any, i int) any {
			switch v := v.(*Device); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_pb_deviceproto_device_proto_msgTypes[2].Exporter = func(v any, i int) any {
			switch v := v.(*Configuration_Settings); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_pb_deviceproto_device_proto_msgTypes[3].Exporter = func(v any, i int) any {
			switch v := v.(*Command); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_pb_deviceproto_device_proto_msgTypes[4].Exporter = func(v any, i int) any {
			switch v := v.(*Command_Payload); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_pb_deviceproto_device_proto_msgTypes[5].Exporter = func(v any, i int) any {
			switch v := v.(*DeviceRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_pb_deviceproto_device_proto_msgTypes[6].Exporter = func(v any, i int) any {
			switch v := v.(*DeviceResponse); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_pb_deviceproto_device_proto_msgTypes[7].Exporter = func(v any, i int) any {
			switch v := v.(*CommandRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_pb_deviceproto_device_proto_msgTypes[8].Exporter = func(v any, i int) any {
			switch v := v.(*CommandResponse); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_pb_deviceproto_device_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   9,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_pb_deviceproto_device_proto_goTypes,
		DependencyIndexes: file_pb_deviceproto_device_proto_depIdxs,
		MessageInfos:      file_pb_deviceproto_device_proto_msgTypes,
	}.Build()
	File_pb_deviceproto_device_proto = out.File
	file_pb_deviceproto_device_proto_rawDesc = nil
	file_pb_deviceproto_device_proto_goTypes = nil
	file_pb_deviceproto_device_proto_depIdxs = nil
}
