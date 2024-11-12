package main

import "fmt"

func main() {
	const a = 42
	var i int = a
	var f float64 = a
	fmt.Println(i, f)

	const b float32 = 3
	var f32 float32 = b
	var f64 float64 = float64(b)
	fmt.Println(f32, f64)

	const c = iota
	fmt.Println(c)

	const (
		d = 2 * 5
		e
		f1 = iota
		g
		h = 10 * iota
	)

	fmt.Println(d, e, f1, g, h)
}
