package main

import (
	. "kursarbeit/config"
	"kursarbeit/storage/models/user"
	"kursarbeit/storage/models/workshop"
)

func main() {
	db := Repo.User.DB

	// automigrate workshop models
	db.AutoMigrate(
		&workshop.Order{},
		&workshop.PersonalInfo{},
		&workshop.CustomerRank{},
		&workshop.MasterRank{},
		&workshop.Customer{},
		&workshop.Master{},
		&workshop.OrderStatus{},
		&workshop.OrderType{},
		&workshop.Guitar{},
		&workshop.RequiredComponents{},
		&workshop.Component{},
	)

	// automigrate user models
	db.AutoMigrate(
		&user.User{},
		&user.Credentials{},
	)
}
