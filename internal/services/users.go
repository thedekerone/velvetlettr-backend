package services

import (
	"log"

	"github.com/thedekerone/velvetlettr-backend/internal/database"
	"github.com/thedekerone/velvetlettr-backend/internal/models"
)

type Service struct {
}

func GetUsers() ([]models.User, error) {
	users, err := database.GetUsersAll()

	if err != nil {
		log.Fatal(err)
		return nil, err
	}

	return users, nil
}
