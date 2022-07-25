package services

import (
	"context"
	"encoding/json"
	"go-redis/repositories"
	"time"

	"github.com/go-redis/redis/v9"
)

type catalogServiceRedis struct {
	productRepository repositories.ProductRepository
	redisClient       *redis.Client
}

func NewCatalogServiceRedis(productRepository repositories.ProductRepository, redisClient *redis.Client) CatalogService {
	return catalogServiceRedis{productRepository, redisClient}
}

func (service catalogServiceRedis) GetProducts() (products []Product, err error) {
	key := "service::GetProducts"

	// Get from Redis
	productsJson, err := service.redisClient.Get(context.Background(), key).Result()
	if err == nil {
		err = json.Unmarshal([]byte(productsJson), &products)
		if err == nil {
			return products, nil
		}
	}

	// Get from database
	productDB, err := service.productRepository.GetProducts()
	if err != nil {
		return nil, err
	}

	for _, p := range productDB {
		products = append(products, Product{
			ID:       p.ID,
			Name:     p.Name,
			Quantity: p.Quantity,
		})
	}

	// Set to Redis
	data, err := json.Marshal(products)
	if err == nil {
		service.redisClient.Set(context.Background(), key, string(data), time.Second*10)
	}

	return products, nil
}
