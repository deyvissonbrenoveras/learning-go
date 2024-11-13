package inventory

import "inventory-management/products"

func AddProduct(product products.Product) {
	storedProducts = append(storedProducts, product)
}

func GetAllProducts() []products.Product {
	return storedProducts
}
