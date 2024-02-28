package proto

import (
	"github.com/matisiekpl/electrocardiogram-server/internal/dto"
	"github.com/matisiekpl/electrocardiogram-server/internal/model"
	"github.com/matisiekpl/electrocardiogram-server/internal/service"
	"github.com/sirupsen/logrus"
)

type RecordProto interface {
	ListRecords(filter *Filter) (*ListRecordsResponse, error)
	StreamRecords(streamer Electrocardiogram_StreamRecordsServer) error
	Analyze(filter *Filter) (*Analysis, error)
}

type recordProto struct {
	recordService service.RecordService
}

func newRecordProto(recordService service.RecordService) RecordProto {
	return &recordProto{
		recordService: recordService,
	}
}

func (r recordProto) ListRecords(filter *Filter) (*ListRecordsResponse, error) {
	records, err := r.recordService.ListRecords(dto.Filter{Start: filter.Start, End: filter.End})
	if err != nil {
		return nil, err
	}

	var elements []*Record
	for _, record := range records {
		elements = append(elements, r.convert(record))
	}
	return &ListRecordsResponse{Records: elements}, nil
}

func (r recordProto) StreamRecords(streamer Electrocardiogram_StreamRecordsServer) error {
	for record := range r.recordService.Channel() {
		err := streamer.Send(r.convert(record))
		if err != nil {
			logrus.Error(err)
			return err
		}
	}
	return nil
}

func (r recordProto) Analyze(filter *Filter) (*Analysis, error) {
	//TODO implement me
	panic("implement me")
}

func (r recordProto) convert(record model.Record) *Record {
	return &Record{
		Id:        int64(record.ID),
		Value:     record.Value,
		Timestamp: record.Timestamp,
	}
}
