package menu

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func PrintMenuAndSelectOption() (int, error) {
	fmt.Println(menuText)
	return readUserInput()
}

var reader = bufio.NewReader(os.Stdin)

func readUserInput() (int, error) {
	input, err := reader.ReadString('\n')
	input = strings.TrimSpace(input)

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

func ClearConsole() {
	fmt.Println("\033[2J")
}
