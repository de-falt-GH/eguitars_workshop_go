package config

import (
	repo "kursarbeit/storage/repository"
)

var Repo *repo.Repo

func init() {
	Repo = repo.InitRepo()
}
