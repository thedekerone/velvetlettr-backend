package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/thedekerone/velvetlettr-backend/internal/database"
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
	r.GET("/ping", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})

	r.GET("/users", func(ctx *gin.Context) {
		if database.DB == nil {
			database.InitDB(config)
		}

		users, err := database.GetUsersAll()

		if err != nil {
			ctx.Error(err)
			return
		}

		ctx.JSON(http.StatusOK, gin.H{
			"data": users,
		})

	})
	r.Run(":8080")
}
