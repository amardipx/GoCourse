package basics

import "fmt"

// Panic is a built in function that stops execution of a function immediately. When a function encounters a panic, it stops executing
// current activities, unwinds the stack, and then executes any deferred functions.

func main() {
	process(10)
	process(-2)
}

func process(i int) {
	defer fmt.Println("Deferred 1")
	defer fmt.Println("Deferred 2")

	if i < 0 {
		fmt.Println("Before panic")
		panic("Input must be a non negative number")
	} else {
		fmt.Println("Processing input", i)
	}
}
