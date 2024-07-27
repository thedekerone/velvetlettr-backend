package database

import (
	"log"

	"github.com/thedekerone/velvetlettr-backend/internal/models"
)

func getUserColFromDb(colName string, user *models.User) interface{} {
	switch colName {
	case "id":
		return &user.ID
	case "email":
		return &user.Email
	case "password_hash":
		return &user.PasswordHash
	case "first_name":
		return &user.FirstName
	case "last_name":
		return &user.LastName
	case "created_at":
		return &user.CreatedAt
	case "updated_at":
		return &user.UpdatedAt
	default:
		panic("not allowed column field")
	}
}

func GetUsersAll() ([]models.User, error) {
	query := "SELECT * FROM users"
	var users []models.User

	rows, err := DB.Query(query)

	if err != nil {
		log.Fatal("failed to query all users")
		return users, err
	}

	defer rows.Close()

	columns, _ := rows.Columns()
	columnsSize := len(columns)

	for rows.Next() {
		rowUser := models.User{}

		cols := make([]interface{}, columnsSize)

		for i := range columnsSize {
			cols[i] = getUserColFromDb(columns[i], &rowUser)
		}

		err := rows.Scan(cols...)

		if err != nil {
			log.Fatal(err)
			return nil, err
		}

		users = append(users, rowUser)
	}

	return users, nil
}
