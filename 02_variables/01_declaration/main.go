package main

import "fmt"

func main() {
	// shorthand
	// usually this method is preferred
	// can be used only inside functions
	// when declaring global variables, only full declaration may be used
	a := 10
	b := "golang"
	c := 3.14
	d := true

	// %v is the value in default format
	// %T is used to print type
	fmt.Printf("%v - %T\n", a, a)
	fmt.Printf("%v - %T\n", b, b)
	fmt.Printf("%v - %T\n", c, c)
	fmt.Printf("%v - %T\n", d, d)

	// long declaration
	// declaration
	var e string
	// assignment
	e = "Hello"
	fmt.Printf("%v - %T\n", e, e)

	// another way to do the same
	// declaration and assignment at the same time
	// the type is recognised from the right part
	var f = "Hello2"
	fmt.Printf("%v - %T\n", f, f)

}
