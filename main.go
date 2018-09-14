package main

import (
	"github.com/gin-gonic/gin"
	"bookroom/models"
)

func main() {
	models.SetDB()
	models.AutoMigrate()

	r := gin.Default()
	r.Run(":8080")
}

