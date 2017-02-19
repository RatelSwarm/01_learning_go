package main

import "fmt"

// first way to define a constant
const a string = "Constant string"

func main() {
	// second way to define a constant
	const b = 40

	// third way to define constants
	const (
		c = 11
		d = "also a constant"
	)

	fmt.Printf("Const 'a': %v - %T\n", a, a)
	fmt.Printf("Const 'b': %v - %T\n", b, b)
	fmt.Printf("Const 'c': %v - %T\n", c, c)
	fmt.Printf("Const 'd': %v - %T\n", d, d)
}
