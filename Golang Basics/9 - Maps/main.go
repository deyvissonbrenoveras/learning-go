package main

import "fmt"

func main() {
	var m map[string][]string

	m = map[string][]string{
		"coffee": {"Coffee", "Espresso", "Capuccino"},
		"tea":    {"Hot Tea", "Chai Tea", "Chai Latte"},
	}

	fmt.Println(m)
	fmt.Println(m["coffee"])

	m["other"] = []string{"Hot Chocollate"}
	fmt.Println(m)

	delete(m, "tea")
	fmt.Println(m)

	fmt.Println(m["tea"])
	v, ok := m["tea"]
	fmt.Println(v, ok)

	m2 := m
	m2["coffee"] = []string{"Coffee"}
	m["tea"] = []string{"Hot Tea"}

	fmt.Println(m)
	fmt.Println(m2)
}
