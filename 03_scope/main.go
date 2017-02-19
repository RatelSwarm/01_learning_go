package main

// importing is effective only for certain file, not for whole package
import "fmt"

// this variable is available inside whole package
// it means, that if package consists of multiple files, then each function
// in each file will have access to this varuable
// this is called "package level scope"
// this variable can be defined before or after function, it doesn't matter
var a = 45

// also, variables starting from lower case letter, can be accessed only within
// the package. E.g. if I import my "main" package into another package, I won't
// be able to use "a" variable I used in "main". But if I had a variable "Abc",
// then I could've used it in another package

func main() {
	fmt.Printf("'a': %v\n", a)

	// we can declare function later, than "main" and still use it
	// not "from top to bottom style" as in Python
	foo()

	// this variable can be used only inside this function
	// it is inside "block-level scope"
	b := 11
	fmt.Printf("'b': %v\n", b)

	// here is another inner scope.
	// everything outside this inner scope is accessible by it
	{
		fmt.Printf("Inner scope 'a': %v\n", a)
		fmt.Printf("Inner scope 'b': %v\n", b)
		c := 26
		fmt.Printf("Inner scope 'c': %v\n", c)

		// function as variable result
		fmt.Printf("'d' itself: %v\n", d)
		fmt.Printf("'Increment d' result1: %v\n", increment())
		fmt.Printf("'Increment d' result2: %v\n", increment())
		// incrementE is considered function result, so it should be presented
		// as 'function' in printf
		// it processes operations needed before result (as in Python)
		fmt.Printf("'IncrementE' result1: %v\n", incrementE())
		fmt.Printf("'IncrementE' result2: %v\n", incrementE())
	}
}

func foo() {
	fmt.Printf("foo 'a': %v\n", a)
}

// variable 'x' is package-scoped, so every increment increases 'x' value
// if it was block scoped, it would've result in identical lines of
// 'result1' and 'result2'
var d int

// named function
func increment() int {
	d++
	return d
}

var e int

// anonymous function - result is assigned to variable
// function inside of a function is defined like that
var incrementE = func() int {
	e++
	fmt.Printf("Test 'e': %v\n", e)
	return e
}
