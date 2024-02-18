package workshop

import "gorm.io/gorm"

type MasterRank struct {
	ID      uint
	Name    string
	Salary  int
	Masters []Master
}

type Master struct {
	ID             uint
	PersonalInfoID uint
	PersonalInfo   *PersonalInfo
	MasterRankID   uint `gorm:"default:1"`
	Orders         []Order
}

func (master *Master) Insert() error {
	db := repo.GetDB()
	if err := db.Create(master).Error; err != nil {
		return err
	}

	return nil
}

func (master *Master) Update() error {
	db := repo.GetDB()

	if err := db.Session(&gorm.Session{FullSaveAssociations: true}).Save(master).Error; err != nil {
		return err
	}

	return nil
}

func (master *Master) Delete() error {
	db := repo.GetDB()
	if err := db.Delete(master).Error; err != nil {
		return err
	}

	return nil
}
