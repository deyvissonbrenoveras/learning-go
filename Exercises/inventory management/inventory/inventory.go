package inventory

import (
	"fmt"
	"inventory-management/products"
)

func AddProduct(product products.Product) error {
	for _, prod := range storedProducts {
		if prod.Id == product.Id {
			return fmt.Errorf(`product with id %s already exists`, product.Id)
		}
	}
	storedProducts = append(storedProducts, product)
	return nil
}

func UpdateProduct(product products.Product) error {
	for i, prod := range storedProducts {
		if prod.Id == product.Id {
			storedProducts[i] = product
			return nil
		}
	}
	return fmt.Errorf(`product with id %s not found`, product.Id)
}

func GetAllProducts() []products.Product {
	return storedProducts
}

func GetInventorySummary() (totalValue float32, productCount int) {

	for _, prod := range storedProducts {
		totalValue += float32(prod.Price)
		productCount += 1
	}

	return totalValue, productCount
}
