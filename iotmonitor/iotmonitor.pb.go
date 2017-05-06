// Code generated by protoc-gen-go.
// source: iotmonitor.proto
// DO NOT EDIT!

/*
Package iotmonitor is a generated protocol buffer package.

It is generated from these files:
	iotmonitor.proto

It has these top-level messages:
	RegisterDeviceRequest
	RegisterDeviceReply
	StatusUpdateRequest
	StatusUpdateReply
	TelemetrySubmitRequest
	TelemetrySubmitReply
	Location
*/
package iotmonitor

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

import (
	context "golang.org/x/net/context"
	grpc "google.golang.org/grpc"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type DeviceType int32

const (
	DeviceType_DRONE  DeviceType = 0
	DeviceType_SENSOR DeviceType = 1
)

var DeviceType_name = map[int32]string{
	0: "DRONE",
	1: "SENSOR",
}
var DeviceType_value = map[string]int32{
	"DRONE":  0,
	"SENSOR": 1,
}

func (x DeviceType) String() string {
	return proto.EnumName(DeviceType_name, int32(x))
}
func (DeviceType) EnumDescriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

type RegisterDeviceRequest struct {
	Name         string     `protobuf:"bytes,1,opt,name=name" json:"name,omitempty"`
	Serialnumber string     `protobuf:"bytes,2,opt,name=serialnumber" json:"serialnumber,omitempty"`
	Owner        string     `protobuf:"bytes,3,opt,name=owner" json:"owner,omitempty"`
	Devicetype   DeviceType `protobuf:"varint,4,opt,name=devicetype,enum=iotmonitor.DeviceType" json:"devicetype,omitempty"`
}

func (m *RegisterDeviceRequest) Reset()                    { *m = RegisterDeviceRequest{} }
func (m *RegisterDeviceRequest) String() string            { return proto.CompactTextString(m) }
func (*RegisterDeviceRequest) ProtoMessage()               {}
func (*RegisterDeviceRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *RegisterDeviceRequest) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *RegisterDeviceRequest) GetSerialnumber() string {
	if m != nil {
		return m.Serialnumber
	}
	return ""
}

func (m *RegisterDeviceRequest) GetOwner() string {
	if m != nil {
		return m.Owner
	}
	return ""
}

func (m *RegisterDeviceRequest) GetDevicetype() DeviceType {
	if m != nil {
		return m.Devicetype
	}
	return DeviceType_DRONE
}

type RegisterDeviceReply struct {
	Registered bool   `protobuf:"varint,1,opt,name=registered" json:"registered,omitempty"`
	Deviceid   uint64 `protobuf:"varint,2,opt,name=deviceid" json:"deviceid,omitempty"`
}

func (m *RegisterDeviceReply) Reset()                    { *m = RegisterDeviceReply{} }
func (m *RegisterDeviceReply) String() string            { return proto.CompactTextString(m) }
func (*RegisterDeviceReply) ProtoMessage()               {}
func (*RegisterDeviceReply) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *RegisterDeviceReply) GetRegistered() bool {
	if m != nil {
		return m.Registered
	}
	return false
}

func (m *RegisterDeviceReply) GetDeviceid() uint64 {
	if m != nil {
		return m.Deviceid
	}
	return 0
}

type StatusUpdateRequest struct {
	Deviceid         uint64    `protobuf:"varint,1,opt,name=deviceid" json:"deviceid,omitempty"`
	Location         *Location `protobuf:"bytes,2,opt,name=location" json:"location,omitempty"`
	Batteryremaining uint32    `protobuf:"varint,3,opt,name=batteryremaining" json:"batteryremaining,omitempty"`
}

func (m *StatusUpdateRequest) Reset()                    { *m = StatusUpdateRequest{} }
func (m *StatusUpdateRequest) String() string            { return proto.CompactTextString(m) }
func (*StatusUpdateRequest) ProtoMessage()               {}
func (*StatusUpdateRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

func (m *StatusUpdateRequest) GetDeviceid() uint64 {
	if m != nil {
		return m.Deviceid
	}
	return 0
}

func (m *StatusUpdateRequest) GetLocation() *Location {
	if m != nil {
		return m.Location
	}
	return nil
}

func (m *StatusUpdateRequest) GetBatteryremaining() uint32 {
	if m != nil {
		return m.Batteryremaining
	}
	return 0
}

type StatusUpdateReply struct {
	Acknowledged bool `protobuf:"varint,1,opt,name=acknowledged" json:"acknowledged,omitempty"`
}

func (m *StatusUpdateReply) Reset()                    { *m = StatusUpdateReply{} }
func (m *StatusUpdateReply) String() string            { return proto.CompactTextString(m) }
func (*StatusUpdateReply) ProtoMessage()               {}
func (*StatusUpdateReply) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{3} }

func (m *StatusUpdateReply) GetAcknowledged() bool {
	if m != nil {
		return m.Acknowledged
	}
	return false
}

type TelemetrySubmitRequest struct {
	Deviceid uint64             `protobuf:"varint,1,opt,name=deviceid" json:"deviceid,omitempty"`
	Readings map[string]float32 `protobuf:"bytes,2,rep,name=readings" json:"readings,omitempty" protobuf_key:"bytes,1,opt,name=key" protobuf_val:"fixed32,2,opt,name=value"`
}

func (m *TelemetrySubmitRequest) Reset()                    { *m = TelemetrySubmitRequest{} }
func (m *TelemetrySubmitRequest) String() string            { return proto.CompactTextString(m) }
func (*TelemetrySubmitRequest) ProtoMessage()               {}
func (*TelemetrySubmitRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{4} }

func (m *TelemetrySubmitRequest) GetDeviceid() uint64 {
	if m != nil {
		return m.Deviceid
	}
	return 0
}

func (m *TelemetrySubmitRequest) GetReadings() map[string]float32 {
	if m != nil {
		return m.Readings
	}
	return nil
}

type TelemetrySubmitReply struct {
	Acknowledged bool `protobuf:"varint,1,opt,name=acknowledged" json:"acknowledged,omitempty"`
}

func (m *TelemetrySubmitReply) Reset()                    { *m = TelemetrySubmitReply{} }
func (m *TelemetrySubmitReply) String() string            { return proto.CompactTextString(m) }
func (*TelemetrySubmitReply) ProtoMessage()               {}
func (*TelemetrySubmitReply) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{5} }

func (m *TelemetrySubmitReply) GetAcknowledged() bool {
	if m != nil {
		return m.Acknowledged
	}
	return false
}

type Location struct {
	Longitude float32 `protobuf:"fixed32,1,opt,name=longitude" json:"longitude,omitempty"`
	Latitude  float32 `protobuf:"fixed32,2,opt,name=latitude" json:"latitude,omitempty"`
	Altitude  float32 `protobuf:"fixed32,3,opt,name=altitude" json:"altitude,omitempty"`
}

func (m *Location) Reset()                    { *m = Location{} }
func (m *Location) String() string            { return proto.CompactTextString(m) }
func (*Location) ProtoMessage()               {}
func (*Location) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{6} }

func (m *Location) GetLongitude() float32 {
	if m != nil {
		return m.Longitude
	}
	return 0
}

func (m *Location) GetLatitude() float32 {
	if m != nil {
		return m.Latitude
	}
	return 0
}

func (m *Location) GetAltitude() float32 {
	if m != nil {
		return m.Altitude
	}
	return 0
}

func init() {
	proto.RegisterType((*RegisterDeviceRequest)(nil), "iotmonitor.RegisterDeviceRequest")
	proto.RegisterType((*RegisterDeviceReply)(nil), "iotmonitor.RegisterDeviceReply")
	proto.RegisterType((*StatusUpdateRequest)(nil), "iotmonitor.StatusUpdateRequest")
	proto.RegisterType((*StatusUpdateReply)(nil), "iotmonitor.StatusUpdateReply")
	proto.RegisterType((*TelemetrySubmitRequest)(nil), "iotmonitor.TelemetrySubmitRequest")
	proto.RegisterType((*TelemetrySubmitReply)(nil), "iotmonitor.TelemetrySubmitReply")
	proto.RegisterType((*Location)(nil), "iotmonitor.Location")
	proto.RegisterEnum("iotmonitor.DeviceType", DeviceType_name, DeviceType_value)
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// Client API for Monitor service

type MonitorClient interface {
	RegisterDevice(ctx context.Context, in *RegisterDeviceRequest, opts ...grpc.CallOption) (*RegisterDeviceReply, error)
	UpdateDeviceStatus(ctx context.Context, in *StatusUpdateRequest, opts ...grpc.CallOption) (*StatusUpdateReply, error)
	SubmitTelemetry(ctx context.Context, in *TelemetrySubmitRequest, opts ...grpc.CallOption) (*TelemetrySubmitReply, error)
}

type monitorClient struct {
	cc *grpc.ClientConn
}

func NewMonitorClient(cc *grpc.ClientConn) MonitorClient {
	return &monitorClient{cc}
}

func (c *monitorClient) RegisterDevice(ctx context.Context, in *RegisterDeviceRequest, opts ...grpc.CallOption) (*RegisterDeviceReply, error) {
	out := new(RegisterDeviceReply)
	err := grpc.Invoke(ctx, "/iotmonitor.Monitor/RegisterDevice", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *monitorClient) UpdateDeviceStatus(ctx context.Context, in *StatusUpdateRequest, opts ...grpc.CallOption) (*StatusUpdateReply, error) {
	out := new(StatusUpdateReply)
	err := grpc.Invoke(ctx, "/iotmonitor.Monitor/UpdateDeviceStatus", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *monitorClient) SubmitTelemetry(ctx context.Context, in *TelemetrySubmitRequest, opts ...grpc.CallOption) (*TelemetrySubmitReply, error) {
	out := new(TelemetrySubmitReply)
	err := grpc.Invoke(ctx, "/iotmonitor.Monitor/SubmitTelemetry", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for Monitor service

type MonitorServer interface {
	RegisterDevice(context.Context, *RegisterDeviceRequest) (*RegisterDeviceReply, error)
	UpdateDeviceStatus(context.Context, *StatusUpdateRequest) (*StatusUpdateReply, error)
	SubmitTelemetry(context.Context, *TelemetrySubmitRequest) (*TelemetrySubmitReply, error)
}

func RegisterMonitorServer(s *grpc.Server, srv MonitorServer) {
	s.RegisterService(&_Monitor_serviceDesc, srv)
}

func _Monitor_RegisterDevice_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RegisterDeviceRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MonitorServer).RegisterDevice(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/iotmonitor.Monitor/RegisterDevice",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MonitorServer).RegisterDevice(ctx, req.(*RegisterDeviceRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Monitor_UpdateDeviceStatus_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(StatusUpdateRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MonitorServer).UpdateDeviceStatus(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/iotmonitor.Monitor/UpdateDeviceStatus",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MonitorServer).UpdateDeviceStatus(ctx, req.(*StatusUpdateRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Monitor_SubmitTelemetry_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(TelemetrySubmitRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MonitorServer).SubmitTelemetry(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/iotmonitor.Monitor/SubmitTelemetry",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MonitorServer).SubmitTelemetry(ctx, req.(*TelemetrySubmitRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _Monitor_serviceDesc = grpc.ServiceDesc{
	ServiceName: "iotmonitor.Monitor",
	HandlerType: (*MonitorServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "RegisterDevice",
			Handler:    _Monitor_RegisterDevice_Handler,
		},
		{
			MethodName: "UpdateDeviceStatus",
			Handler:    _Monitor_UpdateDeviceStatus_Handler,
		},
		{
			MethodName: "SubmitTelemetry",
			Handler:    _Monitor_SubmitTelemetry_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "iotmonitor.proto",
}

func init() { proto.RegisterFile("iotmonitor.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 504 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x54, 0x4d, 0x6f, 0xd3, 0x40,
	0x10, 0xc5, 0x4e, 0x5b, 0x92, 0xe9, 0x07, 0x61, 0x1b, 0xaa, 0xc8, 0xe2, 0x23, 0x98, 0x4b, 0xd4,
	0x43, 0x54, 0x05, 0x09, 0x50, 0xb9, 0x36, 0xb7, 0xd2, 0x8a, 0x4d, 0x10, 0x57, 0x36, 0xf1, 0xc8,
	0x5a, 0x75, 0xbd, 0x6b, 0xd6, 0xeb, 0x56, 0xfe, 0x13, 0xf0, 0x0b, 0xf8, 0x2d, 0xfc, 0x35, 0xe4,
	0x5d, 0xd7, 0xb1, 0x69, 0x28, 0xbd, 0x79, 0xde, 0x9b, 0x19, 0xbf, 0x79, 0x79, 0x31, 0xf4, 0xb9,
	0x32, 0x89, 0x92, 0xdc, 0x28, 0x3d, 0x49, 0xb5, 0x32, 0x8a, 0xc0, 0x1a, 0x09, 0x7f, 0x79, 0xf0,
	0x8c, 0x62, 0xcc, 0x33, 0x83, 0xfa, 0x0c, 0xaf, 0xf9, 0x0a, 0x29, 0x7e, 0xcf, 0x31, 0x33, 0x84,
	0xc0, 0x96, 0x64, 0x09, 0x0e, 0xbd, 0x91, 0x37, 0xee, 0x51, 0xfb, 0x4c, 0x42, 0xd8, 0xcb, 0x50,
	0x73, 0x26, 0x64, 0x9e, 0x2c, 0x51, 0x0f, 0x7d, 0xcb, 0xb5, 0x30, 0x32, 0x80, 0x6d, 0x75, 0x23,
	0x51, 0x0f, 0x3b, 0x96, 0x74, 0x05, 0x79, 0x07, 0x10, 0xd9, 0xf5, 0xa6, 0x48, 0x71, 0xb8, 0x35,
	0xf2, 0xc6, 0x07, 0xd3, 0xa3, 0x49, 0x43, 0x9a, 0x7b, 0xf9, 0xa2, 0x48, 0x91, 0x36, 0x3a, 0xc3,
	0xcf, 0x70, 0xf8, 0xb7, 0xbc, 0x54, 0x14, 0xe4, 0x25, 0x80, 0xae, 0x60, 0x8c, 0xac, 0xc4, 0x2e,
	0x6d, 0x20, 0x24, 0x80, 0xae, 0x5b, 0xc2, 0x23, 0x2b, 0x72, 0x8b, 0xd6, 0x75, 0xf8, 0xd3, 0x83,
	0xc3, 0xb9, 0x61, 0x26, 0xcf, 0xbe, 0xa4, 0x11, 0x33, 0xf5, 0xc1, 0xcd, 0x19, 0xaf, 0x3d, 0x43,
	0x4e, 0xa0, 0x2b, 0xd4, 0x8a, 0x19, 0xae, 0xa4, 0xdd, 0xb7, 0x3b, 0x1d, 0x34, 0xc5, 0x9f, 0x57,
	0x1c, 0xad, 0xbb, 0xc8, 0x31, 0xf4, 0x97, 0xcc, 0x18, 0xd4, 0x85, 0xc6, 0x84, 0x71, 0xc9, 0x65,
	0x6c, 0x1d, 0xd9, 0xa7, 0x77, 0xf0, 0xf0, 0x3d, 0x3c, 0x6d, 0x0b, 0x2a, 0x4f, 0x0c, 0x61, 0x8f,
	0xad, 0xae, 0xa4, 0xba, 0x11, 0x18, 0xc5, 0xf5, 0x91, 0x2d, 0x2c, 0xfc, 0xed, 0xc1, 0xd1, 0x02,
	0x05, 0x26, 0x68, 0x74, 0x31, 0xcf, 0x97, 0x09, 0x37, 0x0f, 0xb9, 0xe6, 0x1c, 0xba, 0x1a, 0x59,
	0xc4, 0x65, 0x9c, 0x0d, 0xfd, 0x51, 0x67, 0xbc, 0x3b, 0x3d, 0x69, 0x5e, 0xb3, 0x79, 0xe3, 0x84,
	0x56, 0x23, 0x33, 0x69, 0x74, 0x41, 0xeb, 0x0d, 0xc1, 0x47, 0xd8, 0x6f, 0x51, 0xa4, 0x0f, 0x9d,
	0x2b, 0x2c, 0xaa, 0xe0, 0x94, 0x8f, 0x65, 0x26, 0xae, 0x99, 0xc8, 0xd1, 0x7a, 0xe7, 0x53, 0x57,
	0x9c, 0xfa, 0x1f, 0xbc, 0xf0, 0x14, 0x06, 0x77, 0x5e, 0xf7, 0xd0, 0xeb, 0xbf, 0x41, 0xf7, 0xd6,
	0x78, 0xf2, 0x1c, 0x7a, 0x42, 0xc9, 0x98, 0x9b, 0x3c, 0x72, 0x91, 0xf5, 0xe9, 0x1a, 0x28, 0xcd,
	0x10, 0xcc, 0x38, 0xd2, 0x49, 0xa8, 0xeb, 0x92, 0x63, 0xa2, 0xe2, 0x3a, 0x8e, 0xbb, 0xad, 0x8f,
	0xdf, 0x00, 0xac, 0x73, 0x49, 0x7a, 0xb0, 0x7d, 0x46, 0x2f, 0x2f, 0x66, 0xfd, 0x47, 0x04, 0x60,
	0x67, 0x3e, 0xbb, 0x98, 0x5f, 0xd2, 0xbe, 0x37, 0xfd, 0xe1, 0xc3, 0xe3, 0x4f, 0xce, 0x3a, 0xb2,
	0x80, 0x83, 0x76, 0x5c, 0xc9, 0xeb, 0xa6, 0xb3, 0x1b, 0xff, 0x69, 0xc1, 0xab, 0xfb, 0x5a, 0x4a,
	0x33, 0x16, 0x40, 0x5c, 0x32, 0x1c, 0xe8, 0xb2, 0x42, 0x5a, 0x63, 0x1b, 0x02, 0x1d, 0xbc, 0xf8,
	0x77, 0x43, 0xb9, 0xf5, 0x2b, 0x3c, 0x71, 0x8e, 0xd7, 0x3f, 0x00, 0x09, 0xff, 0x1f, 0x83, 0x60,
	0x74, 0x6f, 0x4f, 0x2a, 0x8a, 0xe5, 0x8e, 0xfd, 0xcc, 0xbc, 0xfd, 0x13, 0x00, 0x00, 0xff, 0xff,
	0x62, 0xb6, 0xf1, 0x1a, 0x7a, 0x04, 0x00, 0x00,
}