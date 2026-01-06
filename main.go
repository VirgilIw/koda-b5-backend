package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	// initialization
	app := gin.Default()
	// routing

	app.POST("/auth", func(c *gin.Context) {
		// data binding
		// gin.H type alias
		var authen Authentication
		if err := c.ShouldBindJSON(&authen); err != nil {
			log.Println(err.Error())
			c.JSON(http.StatusInternalServerError, gin.H{
				"message": "internal server error",
			})
			return
		}

		if authen.Email == "" {
			fmt.Println("error email kosong")
			return
		}

		if len(authen.Password) <= 6 {
			fmt.Println("error password kurang dari 6")
			return
		}
		//
		// validasi
	})
	app.Run(":4000")
}

type Authentication struct {
	Email    string ``
	Password string
}
