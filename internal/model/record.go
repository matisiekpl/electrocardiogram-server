package model

import "gorm.io/gorm"

type Record struct {
	gorm.Model
	Value     int64
	Timestamp int64
}
