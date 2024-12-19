package main

import (
	"fmt"
	"inventory-management/inventory"
	"inventory-management/menu"
	"strings"
)

var separator = strings.Repeat("-", 20)

func main() {
	fmt.Println("Inventory management system\n")
loop:
	for {
		selectedOption, _ := menu.PrintMenuAndSelectOption()

		switch selectedOption {
		case 1:
			addProduct()
		case 2:
			updateInventory()
		case 3:
			deleteProduct()
		case 4:
			showInventory()
		case 5:
			getInventorySummary()
		case 6:
			startSales()
		case 7:
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
	err = inventory.AddProduct(product)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Printf("Product registered: %+v\n\n", product)
}

func updateInventory() {
	product, err := menu.GetProductInputFromUser()
	if err != nil {
		fmt.Println(err)
		return
	}
	err = inventory.UpdateProduct(product)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Printf("Product updated: %+v\n\n", product)
}

func deleteProduct() {
	id, err := menu.ReadProductId()
	if err != nil {
		fmt.Println(err)
		return
	}
	err = inventory.DeleteProduct(id)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Printf("Product with id %s was deleted successfully\n\n", id)
}

func showInventory() {
	products := inventory.GetAllProducts()

	if len(products) > 0 {
		menu.ClearConsole()
		fmt.Println("Inventory list:\n\n")
		for index, product := range products {
			fmt.Println(separator)
			fmt.Printf("Product %d\n", index+1)
			fmt.Printf("Id: %s\n", product.Id)
			fmt.Printf("Name: %s\n", product.Name)
			fmt.Printf("Price: %.2f\n", product.Price)
			fmt.Printf("Inventory quantity: %d\n", product.InventoryQuantity)
		}
		fmt.Println(separator)
	} else {
		fmt.Println("No prodcuts found...\n")
	}

}

func getInventorySummary() {
	totalValue, productCount := inventory.GetInventorySummary()
	fmt.Println(separator)
	fmt.Println("\nInventory summary: \n")
	fmt.Printf("Product count: %d\nProducts total price: R$ %.2f\n\n", productCount, totalValue)
	fmt.Println(separator)
}

func startSales() {
	menu.ClearConsole()

}
