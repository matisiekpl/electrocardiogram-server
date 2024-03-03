package proto

import (
	"crypto/tls"
	"net"

	"github.com/matisiekpl/electrocardiogram-server/internal/service"
	"github.com/sirupsen/logrus"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
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
	grpcCredentials, err := tls.LoadX509KeyPair("server.crt", "server.key")
	if err != nil {
		logrus.Panic(err)
	}
	_ = grpcCredentials
	//grpcServer := grpc.NewServer(grpc.Creds(credentials.NewServerTLSFromCert(&grpcCredentials)))
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
