package models

import "github.com/jinzhu/gorm"

type Category struct {
	gorm.Model
	categoryName string
	status       int
}
