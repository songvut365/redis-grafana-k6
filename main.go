package main

import (
	"fmt"
	"go-redis/repositories"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func main() {
	db := initDatabase()

	productRepository := repositories.NewProductRepositoryDB(db)
	products, err := productRepository.GetProducts()
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(products)
}

func initDatabase() *gorm.DB {
	dial := mysql.Open("root:1234@tcp(localhost:3306)/store")
	db, err := gorm.Open(dial, &gorm.Config{})
	if err != nil {
		panic(err)
	}

	return db
}
