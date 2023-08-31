package user

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func SignUp(context *gin.Context) {
	var input AuthenticationInput

	if err := context.ShouldBindJSON(&input); err != nil {
			context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
	}

	user := User{
			Email: input.Email,
			Password: input.Password,
	}

	savedUser, err := user.Save()

	if err != nil {
			context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
	}

	context.JSON(http.StatusCreated, gin.H{"user": savedUser})
}