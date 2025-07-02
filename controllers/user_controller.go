package controllers

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/mufferio/concave-tech/services"
)

type UserController struct{ svc services.UserService }

func NewUserController(s services.UserService) *UserController { return &UserController{s} }

func (uc *UserController) Register(r *gin.Engine) { r.GET("/users", uc.allUsers) }

func (uc *UserController) allUsers(c *gin.Context) {
	users, err := uc.svc.GetUsers()
	if err != nil {
		log.Printf("allUsers: %v", err) // ← add this line
		c.JSON(http.StatusInternalServerError, gin.H{"error": "internal"})
		return
	}
	c.JSON(http.StatusOK, users)
}
