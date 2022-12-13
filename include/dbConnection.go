package include

import (
	"fmt"

	"github.com/VeeramaniRajkumar32/webApi/models"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

var DB *gorm.DB

const (
	USER     = "root"
	PASSWORD = ""
	DATABASE = "blogpagego"
)

func Conn() {
	dsn := fmt.Sprintf(
		"%s:%s@tcp(127.0.0.1:3306)/%s?charset=utf8&parseTime=True&loc=Local",
		USER, PASSWORD, DATABASE,
	)

	DB, err := gorm.Open("mysql", dsn)

	if err != nil {
		fmt.Println(err.Error())
		fmt.Println("Could not connect the database")
	}

	fmt.Println("Database %s connected successfully \n", DATABASE)

	// defer DB.Close()
	DB.AutoMigrate(&models.Category{})
}
