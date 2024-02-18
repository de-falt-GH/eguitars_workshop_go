package workshop

import (
	"kursarbeit/storage/interfaces/repository"
)

var (
	repo repository.Repository
)

type PersonalInfo struct {
	ID          uint
	Name        string
	PhoneNumber string
	Email       string
}
