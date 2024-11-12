package inventory

import "inventory-management/products"

func AddProduct(product products.Product) {
	append(storedProducts, product)
}
