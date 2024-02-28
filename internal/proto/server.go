package proto

import "context"

type server struct {
	UnimplementedElectrocardiogramServer

	recordProto RecordProto
}

func (s server) ListRecords(_ context.Context, filter *Filter) (*ListRecordsResponse, error) {
	return s.recordProto.ListRecords(filter)
}

func (s server) StreamRecords(_ *Empty, streamer Electrocardiogram_StreamRecordsServer) error {
	return s.recordProto.StreamRecords(streamer)
}

func (s server) Analyze(_ context.Context, filter *Filter) (*Analysis, error) {
	return s.recordProto.Analyze(filter)
}
