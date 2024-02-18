package repository

import "gorm.io/gorm"

type Repository interface {
	GetDB() *gorm.DB
}
