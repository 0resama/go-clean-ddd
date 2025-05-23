package handler

import (
	"net/http"
	"strconv"

	"github.com/0resama/go-clean-ddd/internal/domain/model"
	useruc "github.com/0resama/go-clean-ddd/internal/domain/usecase/user"
	"github.com/gin-gonic/gin"
)

type UserHandler struct {
	createUser  *useruc.CreateUser
	getUserById *useruc.GetUserByIDUseCase
}

func NewUserHandler(createUC *useruc.CreateUser, getUC *useruc.GetUserByIDUseCase) *UserHandler {
	return &UserHandler{createUser: createUC, getUserById: getUC}
}

func (h *UserHandler) CreateUser(c *gin.Context) {
	var user model.User
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if err := h.createUser.Execute(&user); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, user)
}

func (h *UserHandler) GetUserByID(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid id"})
		return
	}
	userEntity, err := h.getUserById.Execute(uint(id))
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, userEntity)
}
