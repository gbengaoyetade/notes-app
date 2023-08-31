package main

import (
	"fmt"
	"log"
	"net/http"
	"notes-app/note"
	"notes-app/user"
	"os"

	"notes-app/utils"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func loadEnv () {
	err := godotenv.Load(".env")
	if err != nil {
			log.Fatal("Error loading .env file")
	}
}


func loadDatabase () {
	utils.ConnectDB()
	utils.Database.AutoMigrate(&user.User{})
	utils.Database.AutoMigrate(&note.Note{})
}

func serveApplication() {
	router := gin.Default()

	port := os.Getenv("APP_PORT")

	publicRoutes := router.Group("/api")
	router.GET("/", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{"user": "Welcome aboard"})
	})

	publicRoutes.POST("/user", user.SignUp)
	

	router.Run(port)
	fmt.Println("Server running on port 8080")
}

func main () {
	loadEnv()
	loadDatabase()
	serveApplication()
}
