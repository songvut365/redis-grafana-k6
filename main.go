package main

import (
	"fmt"
	"go-redis/repositories"

	"github.com/go-redis/redis/v9"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func main() {
	db := initDatabase()
	redisClient := initRedis()

	productRepository := repositories.NewProductRepositoryRedis(db, redisClient)
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

func initRedis() *redis.Client {
	return redis.NewClient(&redis.Options{
		Addr:     "localhost:6380",
		Password: "",
		DB:       0,
	})
}
