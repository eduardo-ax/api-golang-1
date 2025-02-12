package schemas

import (
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	ID       int64
	Name     string
	Email    string
	Password string
	Year     int
	Plan     int8
}
