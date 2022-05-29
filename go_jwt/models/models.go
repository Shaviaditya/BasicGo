package models

import (
	"github.com/jinzhu/gorm"
);

type Details struct {
	gorm.Model
	Username	string	`json:"username"`
	Email		string	`json:"email"`
	Password	string	`json:"password"`
}
