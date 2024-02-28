package repository

import (
	"gorm.io/gorm"
)

type Repositories interface {
}

type repositories struct {
}

func NewRepositories(db *gorm.DB) Repositories {
	//err := db.AutoMigrate(&model.User{})
	//if err != nil {
	//	logrus.Panic(err)
	//}

	return &repositories{}
}
