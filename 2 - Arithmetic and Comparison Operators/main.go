package main

import "fmt"

func main() {
	a, b := 5, 2

	fmt.Println(a + b)
	fmt.Println(a - b)
	fmt.Println(a * b)
	fmt.Println(a / b) //intergers truncate float
	fmt.Println(a % b)

	fmt.Println(float32(a) / float32(b)) //convert to float
	fmt.Println(a == b)
	fmt.Println(a < b)
	fmt.Println(a > b)
}
