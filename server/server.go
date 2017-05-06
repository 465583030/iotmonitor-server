package server

import (
	"fmt"

	"golang.org/x/net/context"

	"github.com/autodidaddict/iotmonitor-server/iotmonitor"
)

type Server struct{}

func (s *Server) RegisterDevice(ctx context.Context, in *iotmonitor.RegisterDeviceRequest) (*iotmonitor.RegisterDeviceReply, error) {
	fmt.Printf("Received request to register device name %s, type %v\n", in.Name, in.Devicetype)
	return &iotmonitor.RegisterDeviceReply{Registered: true, Deviceid: 99}, nil
}

func (s *Server) UpdateDeviceStatus(ctx context.Context, in *iotmonitor.StatusUpdateRequest) (*iotmonitor.StatusUpdateReply, error) {
	fmt.Printf("Received request to update device status %+v\n", in)
	return &iotmonitor.StatusUpdateReply{Acknowledged: true}, nil
}

func (s *Server) SubmitTelemetry(ctx context.Context, in *iotmonitor.TelemetrySubmitRequest) (*iotmonitor.TelemetrySubmitReply, error) {
	fmt.Printf("Received telemetry %+v\n", in)
	return &iotmonitor.TelemetrySubmitReply{Acknowledged: true}, nil
}
