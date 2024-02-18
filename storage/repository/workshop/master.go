package workshop

import (
	"errors"
	"kursarbeit/storage/models/workshop"
	common_repo "kursarbeit/storage/repository/common"
	"strconv"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type MasterRepo struct {
	common_repo.CommonRepo
}

func (repo *MasterRepo) GetByID(ID int) *workshop.Master {
	var master workshop.Master

	err := repo.DB.Preload(clause.Associations).First(&master, ID).Error
	if err != nil {
		return nil
	}

	return &master
}

func (repo *MasterRepo) GetAll(ctx *gin.Context) []workshop.Master {
	var masters []workshop.Master
	query := repo.DB.InnerJoins("PersonalInfo")

	search := ctx.Query("search")
	if search != "" {
		search = "%" + search + "%"
		query = query.Where("\"PersonalInfo\".name LIKE ?", search).
			Or("\"PersonalInfo\".phone_number LIKE ?", search).
			Or("\"PersonalInfo\".email LIKE ?", search)
	}

	sort_by := ctx.Query("sort_by")
	if sort_by == "name" {
		order := "\"PersonalInfo\"." + sort_by
		if ctx.Query("desc") == "true" {
			order += " DESC"
		}
		query = query.Order(order)
	}

	offset, _ := strconv.Atoi(ctx.DefaultQuery("offset", "0"))
	limit, _ := strconv.Atoi(ctx.DefaultQuery("limit", "10"))

	query.Offset(offset).Limit(limit).Find(&masters)

	return masters
}

func (repo *MasterRepo) DeleteByID(ID int) error {
	db := repo.DB.Delete(&workshop.Master{}, ID)
	if db.Error != nil {
		return db.Error
	} else if db.RowsAffected < 1 {
		return errors.New("entries to delete not found")
	}

	return nil
}

func (repo *MasterRepo) GetDB() *gorm.DB {
	return repo.DB
}
