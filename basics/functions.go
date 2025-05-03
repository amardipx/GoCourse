package basics

import (
	"errors"
	"fmt"
)

func main() {
	fmt.Printf("%d\n", add(5, 10))

	greet := func() {
		fmt.Println("Hello anonymous function")
	}

	greet()
	fmt.Println(divide(10, 3))

	result, err := compare(10, 20)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(result)
	}

	divide2(12, 3)
}

func add(a, b int) int {
	return a + b
}

// multiple return values
func divide(a, b int) (int, int) {
	quotient := a / b
	remainder := a % b
	return quotient, remainder
}

// sending errors
func compare(a, b int) (string, error) {
	if a > b {
		return "a is greater than b", nil
	} else if b > a {
		return "b is greater than a", nil
	} else {
		return "", errors.New("unable to compare which is greater")
	}
}

// named return values
func divide2(a, b int) (quotient int, remainder int) {
	quotient = a / b
	remainder = a % b
	return
}
