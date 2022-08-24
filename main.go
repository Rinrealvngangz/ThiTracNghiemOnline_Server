package main

import (
	"log"

	utils "github.com/Rinrealvngangz/ThiTracNghiemOnline_Server/utils"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func main() {
	router := SetupRouter()
	router.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
			"title":   "GoLang",
		})
	})
	log.Println("Server running...")
	log.Println("OK")
	router.Run(":" + utils.GodotEnv("GO_PORT"))

}

func SetupRouter() *gin.Engine {
	router := gin.Default()
	router.Use(cors.New(cors.Config{
		AllowOrigins:  []string{"*"},
		AllowMethods:  []string{"*"},
		AllowHeaders:  []string{"*"},
		AllowWildcard: true,
	}))
	return router
}
