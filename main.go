package main

import (
	"go-redis/handlers"
	"go-redis/repositories"
	"go-redis/services"

	"github.com/go-redis/redis/v9"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func main() {
	db := initDatabase()
	redisClient := initRedis()
	_ = redisClient

	productRepository := repositories.NewProductRepositoryDB(db)
	productService := services.NewCatalogService(productRepository)
	productHandler := handlers.NewCatalogHandler(productService)

	app := fiber.New()
	app.Use(logger.New())

	app.Get("/products", productHandler.GetProducts)

	app.Listen(":9000")
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
