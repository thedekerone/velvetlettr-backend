package handlers

import (
	"errors"
	"log"
	"net/http"
	"strconv"

	"github.com/thedekerone/velvetlettr-backend/internal/services"

	"github.com/gin-gonic/gin"
)

type UserHandler struct {
	service *services.UserService
}

func NewUserHandler(userService *services.UserService) *UserHandler {
	h := UserHandler{
		service: userService,
	}

	return &h
}

func (h *UserHandler) GetUsersHandler(ctx *gin.Context) {
	users, err := h.service.GetUsers()

	if err != nil {
		ctx.Error(err)
	}

	ctx.JSON(http.StatusOK, gin.H{
		"data": users,
	})
}

func (h *UserHandler) GetUserHandler(ctx *gin.Context) {
	id, exists := ctx.Params.Get("id")

	if !exists {
		ctx.Error(errors.New("You need to provide a user id"))
	}

	user, err := h.service.GetUserById(id)

	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"error": "user not found"})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"data": user,
	})

}

type createUserBody struct {
	Email     string `json:"email" binding:"required"`
	Password  string `json:"password" binding:"required"`
	FirstName string `json:"firstName" binding:"required"`
	LastName  string `json:"lastName" binding:"required"`
}

func (h *UserHandler) CreateUserHandler(ctx *gin.Context) {
	log.Printf("creating..")
	var body createUserBody

	if err := ctx.ShouldBindJSON(&body); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request"})
		return
	}
	log.Printf("correct body")

	id, err := h.service.CreateUser(body.Email, body.Password, body.FirstName, body.LastName)

	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err})
	}

	ctx.JSON(http.StatusOK, id)
}

type deleteUserBody struct {
	ID string `json:"id" binding:"required"`
}

func (h *UserHandler) DeleteUserHandler(ctx *gin.Context) {
	idStr, exists := ctx.Params.Get("id")

	if !exists {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request"})
		return
	}

	id, err := strconv.Atoi(idStr)

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "unexpected error"})
		return
	}

	err = h.service.DeleteUser(id)

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "failed to delete user"})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "deleted correctly yeeyyyy"})
}
