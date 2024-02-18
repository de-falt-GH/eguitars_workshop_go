package repo

import (
	user_model "kursarbeit/storage/models/user"
	workshop_models "kursarbeit/storage/models/workshop"

	common_repo "kursarbeit/storage/repository/common"
	user_repo "kursarbeit/storage/repository/user"
	workshop_repo "kursarbeit/storage/repository/workshop"
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
)

type Repo struct {
	User     *user_repo.UserRepo
	Customer *workshop_repo.CustomerRepo
	Master   *workshop_repo.MasterRepo
	Order    *workshop_repo.OrderRepo
}

func InitRepo() *Repo {
	dsn := "host=127.0.0.1 user=postgres password=postgres dbname=eguitars_workshop port=5432 sslmode=disable TimeZone=Europe/Moscow"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
		NamingStrategy: schema.NamingStrategy{
			SingularTable: true, // struct User -> table user (not users)
		},
	})
	if err != nil {
		log.Fatal(err)
	}

	commonRepo := common_repo.CommonRepo{DB: db}
	repo := Repo{
		User:     &user_repo.UserRepo{CommonRepo: commonRepo},
		Customer: &workshop_repo.CustomerRepo{CommonRepo: commonRepo},
		Master:   &workshop_repo.MasterRepo{CommonRepo: commonRepo},
		Order:    &workshop_repo.OrderRepo{CommonRepo: commonRepo},
	}

	user_model.SetUserRepo(repo.User)
	workshop_models.SetWorkshopRepo(repo.Order)

	return &repo
}
