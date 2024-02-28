package repository

import (
	"github.com/matisiekpl/electrocardiogram-server/internal/model"
	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

type Repositories interface {
	Record() RecordRepository
}

type repositories struct {
	recordRepository RecordRepository
}

func (r repositories) Record() RecordRepository {
	return r.recordRepository
}

func NewRepositories(db *gorm.DB) Repositories {
	err := db.AutoMigrate(&model.Record{})
	if err != nil {
		logrus.Panic(err)
	}

	recordRepository := newRecordRepository(db)

	return &repositories{
		recordRepository: recordRepository,
	}
}
