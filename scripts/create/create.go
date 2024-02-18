package main

import (
	. "kursarbeit/config"
	"kursarbeit/storage/models/user"
)

func main() {
	db := Repo.User.DB

	db.Create(&user.User{
		Role: 4,
		Credentials: &user.Credentials{
			Login:    "admin",
			Password: "$2a$12$wM0USgMpl9r4MCpfLQDzAeDa4CxbQj12srDYSPIpg/IkeQDgGPake",
		},
		Name: "Ivan",
	})

	db.Create(&user.User{
		Role: 3,
		Credentials: &user.Credentials{
			Login:    "manager",
			Password: "$2a$12$04iFU0QicS6918fDvhpbgutDOcfpTQmbKldiLpzbFdIVSh.VurA2G",
		},
		Name: "Vasily",
	})

	var masterID uint = 1
	db.Create(&user.User{
		Role: 2,
		Credentials: &user.Credentials{
			Login:    "master",
			Password: "$2a$12$6kq//zMUUk2l0fNsfbePPO1DLKrr8ktndzKnMIeiRT1q7JBSONrJ.",
		},
		Name:     "Dmitry",
		MasterID: &masterID,
	})

	var customerID uint = 1
	db.Create(&user.User{
		Role: 1,
		Credentials: &user.Credentials{
			Login:    "customer",
			Password: "$2a$12$OJP6xVxsXs6Fby5RVGrH7ecw1Q.SxJSKhZNoACkKUssmOhReUemWC",
		},
		Name:       "Alexey",
		CustomerID: &customerID,
	})
}
