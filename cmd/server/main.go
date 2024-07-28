package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/thedekerone/velvetlettr-backend/internal/database"
	"github.com/thedekerone/velvetlettr-backend/internal/handlers"
	"github.com/thedekerone/velvetlettr-backend/internal/services"
)

var config database.Config = database.Config{
	User:     "postgres",
	Password: "postgres",
	DBName:   "velvetlettr-db",
	Host:     "localhost",
	Port:     5432,
	SSLMode:  "disable",
}

func main() {
	r := gin.Default()

	if database.InitDB(config) != nil {
		panic("error when trying to initialize db")
	}

	r.GET("/ping", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})

	usersService := services.UserService{}
	usersHandler := handlers.NewUserHandler(&usersService)

	userRoutes := r.Group("/users")
	{
		userRoutes.GET("", usersHandler.GetUsersHandler)
		userRoutes.GET("/:id", usersHandler.GetUserHandler)
		userRoutes.POST("/", usersHandler.CreateUserHandler)
	}

	r.Run(":8080")
}
