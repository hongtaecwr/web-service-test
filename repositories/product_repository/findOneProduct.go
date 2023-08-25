package product_repository

import "web-service-test/entities/product_entity"

func (productRepo ProductRepository) FindOneProduct(id string) (bool, error, *product_entity.ResponseProducts) {
	products := new(product_entity.ResponseProducts)
	findAllProduct := productRepo.db.Where("id = ?", id).Find(&products)
	if findAllProduct.RowsAffected == 0 {
		return false, nil, nil
	}
	return true, nil, products
}
