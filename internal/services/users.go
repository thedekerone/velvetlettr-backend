package services

import (
	"errors"
	"log"

	"github.com/thedekerone/velvetlettr-backend/internal/database"
	"github.com/thedekerone/velvetlettr-backend/internal/models"
)

type UserService struct {
}

func (s *UserService) GetUsers() ([]models.User, error) {
	users, err := database.GetUsersAll()

	if err != nil {
		log.Fatal(err)
		return nil, err
	}

	return users, nil
}

func (s *UserService) GetUserById(id string) (models.User, error) {
	if id == "" {
		return models.User{}, errors.New("Missing id")
	}

	user, err := database.GetUserByID(id)

	if err != nil {
		return models.User{}, err
	}

	return user, nil
}

func (s *UserService) CreateUser(email string, password string, firstName string, lastName string) (int, error) {
	id, err := database.CreateUser(email, password, firstName, lastName)

	if err != nil || id == 0 {
		return id, err
	}

	return id, nil
}
