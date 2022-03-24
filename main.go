package main

import (
	controller "framework/Controller"

	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
)

func main() {
	router := gin.Default()

	router.GET("/getSahams", controller.GetSahams)
	router.POST("/addSaham", controller.AddSahams)
	router.PUT("/updateSaham", controller.UpdateSaham)
	router.DELETE("/deleteSaham", controller.DeleteSaham)
	router.Run()
}
