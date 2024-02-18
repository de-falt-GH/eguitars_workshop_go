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

type CustomerRepo struct {
	common_repo.CommonRepo
}

func (repo *CustomerRepo) GetByID(ID int) *workshop.Customer {
	var customer workshop.Customer

	err := repo.DB.Preload(clause.Associations).First(&customer, ID).Error
	if err != nil {
		return nil
	}

	return &customer
}

func (repo *CustomerRepo) GetAll(ctx *gin.Context) []workshop.Customer {
	var customers []workshop.Customer
	query := repo.DB.InnerJoins("PersonalInfo").InnerJoins("CustomerRank")

	search := ctx.Query("search")
	if search != "" {
		search = "%" + search + "%"
		query = query.Where("\"customer\".notes LIKE ?", search).
			Or("\"PersonalInfo\".name LIKE ?", search).
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

	query.Offset(offset).Limit(limit).Find(&customers)

	return customers
}

func (repo *CustomerRepo) DeleteByID(ID int) error {
	db := repo.DB.Delete(&workshop.Customer{}, ID)
	if db.Error != nil {
		return db.Error
	} else if db.RowsAffected < 1 {
		return errors.New("entries to delete not found")
	}

	return nil
}

func (repo *CustomerRepo) GetDB() *gorm.DB {
	return repo.DB
}
