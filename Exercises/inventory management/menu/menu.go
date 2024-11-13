package menu

import (
	"bufio"
	"errors"
	"fmt"
	"inventory-management/products"
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