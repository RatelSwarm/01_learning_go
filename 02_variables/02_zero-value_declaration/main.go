package main

import "fmt"

func main() {
	// 0 by default
	var a int
	// empty string is printed as nothing ("")
	var b string
	// 0 by default
	var c float64
	// false by default
	var d bool

	fmt.Printf("|%v| - %T\n", a, a)
	fmt.Printf("|%v| - %T\n", b, b)
	fmt.Printf("|%v| - %T\n", c, c)
	fmt.Printf("|%v| - %T\n", d, d)

}
