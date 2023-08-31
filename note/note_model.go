package note

import "gorm.io/gorm"


type Note struct {
	gorm.Model
	Note string `gorm:"type:text" json:"note"`
	UserId string `gorm:"type:varchar" `
}