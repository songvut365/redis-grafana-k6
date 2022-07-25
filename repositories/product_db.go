package repositories

import (
	"fmt"
	"math/rand"
	"time"

	"gorm.io/gorm"
)

type productRepositoryDB struct {
	db *gorm.DB
}

func NewProductRepositoryDB(db *gorm.DB) ProductRepository {
	db.AutoMigrate(&product{})
	mockData(db)
	return productRepositoryDB{db: db}
}

func mockData(db *gorm.DB) error {
	// check data already exist
	var count int64
	db.Model(&product{}).Count(&count)
	if count > 0 {
		return nil
	}

	// random quantity
	seed := rand.NewSource(time.Now().UnixNano())
	random := rand.New(seed)

	products := []product{}
	for i := 0; i < 5000; i++ {
		products = append(products, product{
			Name:     fmt.Sprintf("Product%v", i+1),
			Quantity: random.Intn(100),
		})
	}

	return db.Create(&products).Error
}

func (repository productRepositoryDB) GetProducts() (products []product, err error) {
	err = repository.db.Order("quantity desc").Limit(30).Find(&products).Error
	return products, err
}
