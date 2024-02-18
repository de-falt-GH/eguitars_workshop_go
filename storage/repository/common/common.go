package common

import "gorm.io/gorm"

type CommonRepo struct {
	DB *gorm.DB
}

func (repo CommonRepo) GetDB() *gorm.DB {
	return repo.DB
}
