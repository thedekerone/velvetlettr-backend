package main

import (
	"log"
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

		rows, err := database.DB.Query("SELECT * FROM users")

		if err != nil {
			log.Fatal("failed to query users")
			return
		}
		columns, _ := rows.Columns()

		ctx.JSON(http.StatusOK, gin.H{
			"data": columns,
		})
	})
	r.Run(":8080")
}
