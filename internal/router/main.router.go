package router

import (
	"github.com/gin-gonic/gin"
	"github.com/virgilIw/koda-b5-backend/internal/controller"
	"github.com/virgilIw/koda-b5-backend/internal/service"
)

// pakai gin engine
func Init(app *gin.Engine) {
	// service
	authService := service.NewAuthService()
	// controller
	authController := controller.NewAuthController(authService)
	// method
	app.POST("/register", authController.GetRegister)

	app.POST("/login", authController.GetLogin)
}
