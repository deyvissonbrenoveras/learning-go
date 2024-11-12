package main

import (
	"fmt"
	"inventory-management/menu"
)

func main() {
loop:
	for {
		selectedOption, _ := menu.PrintMenuAndSelectOption()

		switch selectedOption {
		case 1:
			fmt.Println("add product selected")
		case 5:
			fmt.Println("Quitting Inventory Management System")
			break loop
		}

		// fmt.Println(selectedOption)
	}
}
