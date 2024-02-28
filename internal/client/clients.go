package client

import "github.com/matisiekpl/electrocardiogram-server/internal/dto"

type Clients interface{}

type clients struct{}

func NewClients(_ dto.Config) Clients {
	return &clients{}
}
