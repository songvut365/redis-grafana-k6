package handlers

import (
	"context"
	"encoding/json"
	"go-redis/services"
	"time"

	"github.com/go-redis/redis/v9"
	"github.com/gofiber/fiber/v2"
)

type catalogHandlerRedis struct {
	catalogService services.CatalogService
	redisClient    *redis.Client
}

func NewCatalogHandlerRedis(catalogService services.CatalogService, redisClient *redis.Client) CatalogHandler {
	return catalogHandlerRedis{catalogService, redisClient}
}

func (h catalogHandlerRedis) GetProducts(c *fiber.Ctx) error {
	key := "handler::GetProducts"

	// Get from Redis
	responseJson, err := h.redisClient.Get(context.Background(), key).Result()
	if err == nil {
		c.Set("Content-Type", "application/json")
		return c.SendString(responseJson)
	}

	// Get from Service
	products, err := h.catalogService.GetProducts()
	if err != nil {
		return err
	}

	response := fiber.Map{
		"status":   "ok",
		"products": products,
	}

	// Set to Redis
	data, err := json.Marshal(response)
	if err == nil {
		h.redisClient.Set(context.Background(), key, string(data), time.Second*10)
	}

	return c.JSON(response)
}
