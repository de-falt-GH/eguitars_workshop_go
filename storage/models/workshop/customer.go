package workshop

import (
	"kursarbeit/storage/interfaces/repository"

	"gorm.io/gorm"
)

type CustomerRank struct {
	ID   uint
	Name string
}

type Customer struct {
	ID             uint
	PersonalInfoID uint
	PersonalInfo   *PersonalInfo
	CustomerRankID uint `gorm:"default:1"`
	CustomerRank   *CustomerRank
	TotalPurchase  int `gorm:"default:0"`
	Notes          string
}

func (customer *Customer) Insert() error {
	db := repo.GetDB()
	if err := db.Create(customer).Error; err != nil {
		return err
	}

	return nil
}

func (customer *Customer) Update() error {
	db := repo.GetDB().Session(&gorm.Session{FullSaveAssociations: true})
	if err := db.Save(customer).Error; err != nil {
		return err
	}

	return nil
}

func (customer *Customer) Delete() error {
	db := repo.GetDB()
	if err := db.Delete(customer).Error; err != nil {
		return err
	}

	return nil
}

func SetWorkshopRepo(newRepo repository.Repository) {
	repo = newRepo
}
