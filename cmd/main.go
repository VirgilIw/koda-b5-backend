package main

import (
	"github.com/gin-gonic/gin"
	"github.com/virgilIw/koda-b5-backend/internal/router"
)

func main() {
	// buat gin engine
	app := gin.Default()
	// register
	router.Init(app)
	// login

	app.Run(":4000")
}
