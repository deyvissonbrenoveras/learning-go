package main

import (
	"fmt"
	"inventory-management/inventory"
	"inventory-management/menu"
	"strings"
)

func main() {
	fmt.Println("Inventory management system\n")
loop:
	for {
		selectedOption, _ := menu.PrintMenuAndSelectOption()

		switch selectedOption {
		case 1:
			addProduct()
		case 3:
			showInventory()
		case 5:
			menu.ClearConsole()
			fmt.Println("Quitting Inventory Management System")
			break loop
		}
	}
}

func addProduct() {
	product, err := menu.GetProductInputFromUser()
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Printf("Product registered: %+v\n\n", product)
	inventory.AddProduct(product)
}

func showInventory() {
	products := inventory.GetAllProducts()
	separator := strings.Repeat("-", 20)

	for index, product := range products {
		fmt.Println(separator)
		fmt.Printf("Product %d\n", index+1)
		fmt.Printf("Id: %s\n", product.Name)
		fmt.Printf("Name: %s\n", product.Name)
		fmt.Printf("Price: %.2f\n", product.Price)
	}
	fmt.Println(separator)
}
