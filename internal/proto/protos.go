package proto

import (
	"github.com/matisiekpl/electrocardiogram-server/internal/service"
	"google.golang.org/grpc"
	"net"
)

type Protos interface {
	Serve(listener net.Listener) error
}
type protos struct {
	grpcServer *grpc.Server
}

func (p protos) Serve(listener net.Listener) error {
	return p.grpcServer.Serve(listener)
}

func NewProtos(services service.Services) Protos {
	grpcServer := grpc.NewServer()
	s := &server{}
	RegisterElectrocardiogramServer(grpcServer, s)
	return &protos{
		grpcServer: grpcServer,
	}
}
