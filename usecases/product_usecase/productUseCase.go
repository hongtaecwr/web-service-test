package product_usecase

import (
	"web-service-test/entities/product_entity"
	"web-service-test/repositories/product_repository"
)

type ProductUseCase struct {
	productRepo product_repository.ProductRepository
}

func NewProductUseCase(productRepo product_repository.ProductRepository) ProductUseCase {
	return ProductUseCase{productRepo: productRepo}
}

func (pu ProductUseCase) GetAllProduct() (bool, error, *[]product_entity.ResponseProducts) {
	ok, err, products := pu.productRepo.FindAllProduct()
	if !ok {
		return false, err, nil
	}
	return true, err, products
}

func (pu ProductUseCase) GetProductDetail(id string) (bool, error, *product_entity.ResponseProducts) {
	ok, err, products := pu.productRepo.FindOneProduct(id)
	if !ok {
		return false, err, nil
	}
	return true, err, products
}
