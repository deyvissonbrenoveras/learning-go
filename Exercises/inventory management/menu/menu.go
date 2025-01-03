package menu

import (
	"bufio"
	"errors"
	"fmt"
	"inventory-management/inventory"
	"inventory-management/products"
	sales "inventory-management/sale"
	"os"
	"strconv"
	"strings"
)

func PrintMenuAndSelectOption() (int, error) {
	fmt.Println(menuText)
	return getMenuOptionFromUser()
}

func GetProductInputFromUser() (products.Product, error) {
	var product products.Product
	ClearConsole()
	fmt.Print("Type the product Id: ")
	input, err := readUserInput()
	if err != nil {
		return product, fmt.Errorf("an error has occurred while reading product id: %v", err)
	}
	product.Id = input

	ClearConsole()
	fmt.Print("Type the product name: ")
	input, err = readUserInput()
	if err != nil {
		return product, fmt.Errorf("an error has occurred while reading product name: %v", err)
	}
	product.Name = input

	ClearConsole()
	fmt.Print("Type the product price: ")
	input, err = readUserInput()
	if err != nil {
		return product, fmt.Errorf("an error has occurred while reading product price: %v", err)
	}
	price, err := strconv.ParseFloat(input, 32)
	if err != nil {
		return product, errors.New("an error has occurred while reading product price: invalid price")
	}
	product.Price = float32(price)

	ClearConsole()
	fmt.Print("Type the product available quantity: ")
	input, err = readUserInput()
	if err != nil {
		return product, fmt.Errorf("an error has occurred while reading product available quantity: %v", err)
	}
	inventoryQuantity, err := strconv.ParseInt(input, 10, 32)
	if err != nil {
		return product, errors.New("an error has occurred while reading product price: invalid price")
	}
	if inventoryQuantity < 1 {
		return product, errors.New("inventory quantity must be greater than 0")
	}
	product.InventoryQuantity = int32(inventoryQuantity)
	return product, nil

}

var reader = bufio.NewReader(os.Stdin)

func getMenuOptionFromUser() (int, error) {
	input, err := readUserInput()

	if err != nil {
		errToReturn := errors.New("an error has occurred while reading user input")
		fmt.Println(errToReturn)
		return 0, errToReturn
	}

	selectedNumber, err := strconv.Atoi(input)

	if err != nil {
		fmt.Println(err)
		errToReturn := errors.New("user input is not valid")
		fmt.Println(errToReturn)
		return 0, errToReturn
	}

	return selectedNumber, nil
}

func ReadProductId() (id string, err error) {
	ClearConsole()
	fmt.Print("Please, type the product ID to delete: ")
	id, err = readUserInput()
	if err != nil {
		fmt.Println(err)
		return "", err
	}
	ClearConsole()
	return id, nil
}

func CreateSale() (sales.Sale, error) {
	PrintSeparator()
	newSale := sales.Sale{}
create_sale_loop:
	for {

		fmt.Println(saleText)
		selectedOption, _ := getMenuOptionFromUser()

		switch selectedOption {
		case 1:
			product, quantitySelected, _ := addProduct()
			fmt.Printf("%v+ %d", product, quantitySelected)
		case 2:
			// encerrar venda
		case 3:
			// cancelar
			ClearConsole()
			break create_sale_loop
		}

		break
	}
	return newSale, nil
}

func addProduct() (products.Product, int, error) {
	fmt.Println("type the product Id:")
	input, err := readUserInput()

	if err != nil {
		fmt.Println(err)
		return products.Product{}, 0, err
	}

	for _, prod := range inventory.GetAllProducts() {
		if prod.Id == input {
			fmt.Printf("Product selected: %+v\ntype product quantity:", prod)
			input, err = readUserInput()

			if err != nil {
				fmt.Println(err)
				return products.Product{}, 0, err
			}
			quantitySelected, err := strconv.ParseInt(input, 10, 32)

			if err != nil {
				fmt.Println(err)
				return products.Product{}, 0, err
			}

			return prod, int(quantitySelected), nil
		}
	}

	return products.Product{}, 0, fmt.Errorf("product with id %s not found", input)
}

func PrintSeparator() {
	fmt.Println(strings.Repeat("-", 20))
}

func readUserInput() (string, error) {
	input, err := reader.ReadString('\n')
	input = strings.TrimSpace(input)

	if err != nil {
		return "", err
	}

	return input, nil
}

func ClearConsole() {
	fmt.Println("\033[2J")
}
