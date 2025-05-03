package basics

// awaits execution until surrounding function or statement has finished executing.
// Kinda like await and async in Node.js

// Just because defer gets executed at the end, doesn't mean that it gets evaluated at the end. It is evaluated normally. Keep in mind when
// changing values.

import "fmt"

func main() {
	process(10)
}

func process(i int) {
	defer fmt.Println("Deferred i value:", i)
	defer fmt.Println("First deferred statement executed")
	defer fmt.Println("Second deferred statement executed")
	defer fmt.Println("Third deferred statement executed") // If multiple defer statements, LIFO basis
	i++

	fmt.Println("Normal statement executed")
	fmt.Println("Value of i:", i)
}
