package main

import (
	"github.com/VeeramaniRajkumar32/webApi/include"
	"github.com/VeeramaniRajkumar32/webApi/models"
)

func init() {
	include.Conn()
}

func main() {
	include.DB.AutoMigrate(&models.Category{})
}
