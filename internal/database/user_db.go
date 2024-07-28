package database

import (
	"log"

	"github.com/thedekerone/velvetlettr-backend/internal/models"
	"golang.org/x/crypto/bcrypt"
)

func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)

	return string(bytes), err
}

func CheckPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}

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
	rows, err := DB.Query("SELECT * FROM users")

	if err != nil {
		log.Fatal("failed to query all users")
		return nil, err
	}

	defer rows.Close()

	var users []models.User
	var columns []string
	columns, err = rows.Columns()
	columnsSize := len(columns)

	if err != nil {
		log.Fatal("failed to retrieve columns")
		return nil, err
	}

	for rows.Next() {
		rowUser := models.User{}

		cols := make([]interface{}, columnsSize)

		for i := range columnsSize {
			cols[i] = getUserColFromDb(columns[i], &rowUser)
		}

		err = rows.Scan(cols...)

		if err != nil {
			log.Fatal(err)
			return nil, err
		}

		users = append(users, rowUser)
	}
	err = rows.Err()
	if err != nil {
		log.Fatal(err)
		return nil, err
	}

	return users, nil
}

func GetUserByID(id string) (models.User, error) {
	user := models.User{}

	row := DB.QueryRow("SELECT * FROM users WHERE id=$1", id)

	err := row.Scan(&user.ID, &user.Email, &user.FirstName, &user.LastName, &user.PasswordHash, &user.CreatedAt, &user.UpdatedAt)

	if err != nil {
		log.Printf("error when querying user")
		return user, err
	}

	return user, nil
}

func CreateUser(email string, password string, firstName string, lastName string) (int, error) {
	var userId int
	passwordHash, _ := HashPassword(password)

	err := DB.QueryRow("INSERT INTO users (email, password_hash, first_name, last_name) VALUES ($1,$2,$3,$4) RETURNING id", email, passwordHash, firstName, lastName).Scan(&userId)

	if err != nil {
		return userId, err
	}

	return userId, err
}
