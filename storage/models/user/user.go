package user

import (
	"kursarbeit/storage/interfaces/repository"
	// . "kursarbeit/storage/models/workshop"

	"gorm.io/gorm"
)

var (
	repo repository.Repository
)

// 1 - customer,
// 2 - master,
// 3 - manager,
// 4 - admin
type UserRole int32

const (
	ROLE_CUSTOMER UserRole = 1
	ROLE_MASTER   UserRole = 2
	ROLE_MANAGER  UserRole = 3
	ROLE_ADMIN    UserRole = 4
)

type Credentials struct {
	ID       uint
	Login    string `gorm:"unique"`
	Password string
}

type User struct {
	ID            uint
	Role          UserRole     `gorm:"default:4; check: role between 1 and 4"`
	CredentialsID uint         `gorm:"foreignKey:ID"`
	Credentials   *Credentials `gorm:"foreignKey:CredentialsID"`
	MasterID      *uint        `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
	// Master        *Master
	CustomerID *uint `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
	// Customer      *Customer
	Name string `gorm:"default:Anonymous"`
}

func (user *User) Insert() error {
	db := repo.GetDB()
	if err := db.Create(user).Error; err != nil {
		return err
	}

	return nil
}

func (user *User) Update() error {
	db := repo.GetDB().Session(&gorm.Session{FullSaveAssociations: true})
	if err := db.Save(user).Error; err != nil {
		return err
	}

	return nil
}

func (user *User) Delete() error {
	db := repo.GetDB()
	if err := db.Delete(user).Error; err != nil {
		return err
	}

	return nil
}

func SetUserRepo(newRepo repository.Repository) {
	repo = newRepo
}
