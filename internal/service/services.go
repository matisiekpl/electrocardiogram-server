package service

import (
	"github.com/matisiekpl/electrocardiogram-server/internal/client"
	"github.com/matisiekpl/electrocardiogram-server/internal/dto"
	"github.com/matisiekpl/electrocardiogram-server/internal/repository"
)

type Services interface {
}

type services struct {
}

func NewServices(repositories repository.Repositories, config dto.Config, clients client.Clients) Services {
	return &services{}
}
