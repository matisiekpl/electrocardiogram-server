package proto

import (
	"github.com/matisiekpl/electrocardiogram-server/internal/service"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"net"
)

type Protos interface {
	Serve(listener net.Listener) error
}
type protos struct {
	grpcServer *grpc.Server

	recordProto RecordProto
}

func (p protos) Serve(listener net.Listener) error {
	return p.grpcServer.Serve(listener)
}

func NewProtos(services service.Services) Protos {
	grpcServer := grpc.NewServer()
	recordProto := newRecordProto(services.Record())
	s := &server{
		recordProto: recordProto,
	}
	RegisterElectrocardiogramServer(grpcServer, s)
	reflection.Register(grpcServer)
	return &protos{
		grpcServer:  grpcServer,
		recordProto: recordProto,
	}
}
