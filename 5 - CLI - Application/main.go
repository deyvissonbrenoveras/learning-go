package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	fmt.Printf("What would you like me to scream?")
	in := bufio.NewReader(os.Stdin)
	s, _ := in.ReadString('\n')
	s = strings.TrimSpace(s)
	s = strings.ToUpper(s)
	fmt.Println(s + "!")

	var name string
	fmt.Print("Type your name:")
	fmt.Scan(&name)
	fmt.Printf("Hello, %s!\n", name)
}
