package repository

import (
	"github.com/matisiekpl/electrocardiogram-server/internal/model"
	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
	"time"
)

type RecordRepository interface {
	Filter(start, end int64) ([]model.Record, error)
	Insert(record *model.Record) error
	PurgeOlderThan(time time.Time) error
}

type recordRepository struct {
	db *gorm.DB
}

func newRecordRepository(db *gorm.DB) RecordRepository {
	return &recordRepository{db: db}
}

func (r recordRepository) Filter(start, end int64) ([]model.Record, error) {
	var records []model.Record
	err := r.db.Model(model.Record{}).Where("timestamp > ? and timestamp < ?", start, end).Order("timestamp").Find(&records).Error
	if err != nil {
		return nil, err
	}
	return records, nil
}

func (r recordRepository) Insert(record *model.Record) error {
	return r.db.Create(&record).Error
}

func (r recordRepository) PurgeOlderThan(time time.Time) error {
	tx := r.db.Model(&model.Record{}).Unscoped().Where("timestamp < ?", time.UnixMilli()).Delete(&model.Record{})
	if tx.Error != nil {
		return tx.Error
	}
	logrus.Infof("deleted %d rows", tx.RowsAffected)
	return r.db.Exec("VACUUM;").Error
}
