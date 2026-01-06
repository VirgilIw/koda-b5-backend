package controller

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/virgilIw/koda-b5-backend/internal/dto"
	"github.com/virgilIw/koda-b5-backend/internal/service"
)

type AuthController struct {
	authService *service.AuthService
}

func NewAuthController(authService *service.AuthService) *AuthController {
	return &AuthController{
		authService: authService,
	}
}

func (a *AuthController) GetRegister(c *gin.Context) {
	var register dto.Register

	if err := c.ShouldBindJSON(&register); err != nil {
		log.Println(err.Error())
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "terjadi kesalahan server",
		})
		return
	}

	userData, err := a.authService.ServiceRegister(register)
	if err != nil {
		statusCode := http.StatusBadRequest
		if err.Error() == "email sudah terdaftar" {
			statusCode = http.StatusConflict
		}

		c.JSON(statusCode, gin.H{
			"message": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "register sukses",
		"data":    userData,
	})
}

func (a *AuthController) GetLogin(c *gin.Context) {
	var authen dto.Authentication

	if err := c.ShouldBindJSON(&authen); err != nil {
		log.Println(err.Error())
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "terjadi kesalahan server",
		})
		return
	}

	err := a.authService.ServiceLogin(authen)
	if err != nil {
		statusCode := http.StatusBadRequest
		if err.Error() == "email atau password salah / belum terdaftar" {
			statusCode = http.StatusUnauthorized
		}

		c.JSON(statusCode, gin.H{
			"message": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "login valid",
	})
}
