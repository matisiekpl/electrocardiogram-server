package service

import (
	"github.com/matisiekpl/electrocardiogram-server/internal/client"
	"github.com/matisiekpl/electrocardiogram-server/internal/dto"
	"github.com/matisiekpl/electrocardiogram-server/internal/repository"
)

type Services interface {
	Record() RecordService
}

type services struct {
	recordService RecordService
}

func (s services) Record() RecordService {
	return s.recordService
}

func NewServices(repositories repository.Repositories, config dto.Config, clients client.Clients) Services {
	recordService := newRecordService(repositories.Record(), config)
	return &services{
		recordService: recordService,
	}
}
