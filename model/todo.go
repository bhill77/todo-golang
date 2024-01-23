package model

import "gorm.io/gorm"

type Todo struct {
	gorm.Model
	Label  string
	Finish bool
}
