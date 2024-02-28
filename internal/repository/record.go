package repository

import (
	"github.com/matisiekpl/electrocardiogram-server/internal/model"
	"gorm.io/gorm"
)

type RecordRepository interface {
	Filter(start, end int64) ([]model.Record, error)
	Insert(record *model.Record) error
}

type recordRepository struct {
	db *gorm.DB
}

func newRecordRepository(db *gorm.DB) RecordRepository {
	return &recordRepository{db: db}
}

func (r recordRepository) Filter(start, end int64) ([]model.Record, error) {
	var records []model.Record
	err := r.db.Where("timestamp > ? and timestamp < ?", start, end).Find(&records).Error
	if err != nil {
		return nil, err
	}
	return records, nil
}

func (r recordRepository) Insert(record *model.Record) error {
	return r.db.Create(&record).Error
}
