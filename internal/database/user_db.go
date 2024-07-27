package database

import (
	"log"

	"github.com/thedekerone/velvetlettr-backend/internal/models"
)

func GetUsersAll() ([]models.User, error) {
	query := "SELECT * FROM users"
	var users []models.User

	rows, err := DB.Query(query)

	if err != nil {
		log.Fatal("failed to query all users")
		return users, err
	}

	for rows.Next() {
		rowUser := models.User{}
		columns, _ := rows.Columns()

		log.Print(columns)
		err := rows.Scan(&rowUser.ID, &rowUser.Email, &rowUser.PasswordHash, &rowUser.FirstName, &rowUser.LastName, &rowUser.CreatedAt, &rowUser.UpdatedAt)

		if err != nil {
			log.Fatal(err)
			return nil, err
		}

		users = append(users, rowUser)
	}

	return users, nil
}
