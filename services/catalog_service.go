package services

import "go-redis/repositories"

type catalogService struct {
	productRepository repositories.ProductRepository
}

func NewCatalogService(productRepository repositories.ProductRepository) CatalogService {
	return catalogService{productRepository}
}

func (service catalogService) GetProducts() (products []Product, err error) {
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

	return products, nil
}
