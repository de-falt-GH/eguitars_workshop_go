package user

import (
	"errors"
	"kursarbeit/storage/models/user"
	common_repo "kursarbeit/storage/repository/common"
	"strconv"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type UserRepo struct {
	common_repo.CommonRepo
}

func (repo *UserRepo) GetByID(ID int) *user.User {
	var user user.User

	err := repo.DB.Preload(clause.Associations).First(&user, ID).Error
	if err != nil {
		return nil
	}

	return &user
}

func (repo *UserRepo) IDtoMasterID(ID int) *uint {
	var user user.User

	err := repo.DB.First(&user, ID).Error
	if err != nil {
		return nil
	}

	return user.MasterID
}

func (repo *UserRepo) IDtoCustomerID(ID int) *uint {
	var user user.User

	err := repo.DB.First(&user, ID).Error
	if err != nil {
		return nil
	}

	return user.CustomerID
}

func (repo *UserRepo) GetByLogin(login string) *user.User {
	var user user.User
	err := repo.DB.InnerJoins("Credentials").Where(`"Credentials".login = ?`, login).First(&user).Error
	if err != nil {
		return nil
	}

	return &user
}

func (repo *UserRepo) GetAll(ctx *gin.Context) []user.User {
	var users []user.User
	// query := repo.DB.Preload(clause.Associations)
	query := repo.DB.InnerJoins("Credentials")

	search := ctx.Query("search")
	if search != "" {
		search = "%" + search + "%"
		query = query.Where("\"user\".name LIKE ? OR \"Credentials\".login LIKE ?", search, search)
	}

	sort_by := ctx.Query("sort_by")
	if sort_by == "name" {
		order := sort_by
		if ctx.Query("desc") == "true" {
			order += " DESC"
		}
		query = query.Order(order)
	}

	offset, _ := strconv.Atoi(ctx.DefaultQuery("offset", "0"))
	limit, _ := strconv.Atoi(ctx.DefaultQuery("limit", "10"))

	query.Offset(offset).Limit(limit).Find(&users)

	return users
}

func (repo *UserRepo) DeleteByID(ID int) error {
	db := repo.DB.Delete(&user.User{}, ID)
	if db.Error != nil {
		return db.Error
	} else if db.RowsAffected < 1 {
		return errors.New("entries to delete not found")
	}

	return nil
}

func (repo *UserRepo) GetDB() *gorm.DB {
	return repo.DB
}
