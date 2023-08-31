package user

import (
	"html"
	"notes-app/utils"
	"strings"

	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Email string  `gorm:"size:255;not null;unique" json:"email"`
	Password string `gorm:"size:255;not null;" json:"-"`
	ProfilePicUrl string ``
}

type AuthenticationInput struct {
	Email string `json:"email" binding:"required"`
	Password string `json:"password" binding:"required"`
}


func (user *User) Save() (*User, error) {
	err := utils.Database.Create(&user).Error
	if err != nil {
			return &User{}, err
	}
	return user, nil
}

func (user *User) BeforeSave(*gorm.DB) error {
	passwordHash, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
			return err
	}
	user.Password = string(passwordHash)
	user.Email = html.EscapeString(strings.TrimSpace(user.Email))
	return nil
}
