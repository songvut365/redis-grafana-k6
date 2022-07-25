package repositories

import (
	"context"
	"encoding/json"
	"time"

	"github.com/go-redis/redis/v9"
	"gorm.io/gorm"
)

type productRepositoryRedis struct {
	db          *gorm.DB
	redisClient *redis.Client
}

func NewProductRepositoryRedis(db *gorm.DB, redisClient *redis.Client) ProductRepository {
	db.AutoMigrate(&product{})
	mockData(db)

	return productRepositoryRedis{db, redisClient}
}

func (repository productRepositoryRedis) GetProducts() (products []product, err error) {
	key := "repository::GetProducts"

	// Get from Redis
	productsJson, err := repository.redisClient.Get(context.Background(), key).Result()
	if err == nil {
		err = json.Unmarshal([]byte(productsJson), &products)
		if err != nil {
			return products, nil
		}
	}

	// Get from database
	err = repository.db.Order("quantity desc").Limit(30).Find(&products).Error
	if err != nil {
		return nil, err
	}

	// Set to redis
	data, err := json.Marshal(products)
	if err != nil {
		return nil, err
	}

	err = repository.redisClient.Set(context.Background(), key, string(data), time.Second*10).Err()
	if err != nil {
		return nil, err
	}

	return products, nil
}
