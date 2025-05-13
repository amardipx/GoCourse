package intermediate

import (
	"fmt"
)

// Optional generic swap function (uncomment if needed)
// func swap[T any](a, b T) (T, T) {
// 	return b, a
// }

type Stack[T any] struct {
	elements []T
}

func (s *Stack[T]) push(element T) {
	s.elements = append(s.elements, element)
}

func (s *Stack[T]) pop() (T, bool) {
	if len(s.elements) == 0 {
		var zero T // Fixed: declared type for zero value
		return zero, false
	}

	element := s.elements[len(s.elements)-1]
	s.elements = s.elements[:len(s.elements)-1]
	return element, true
}

func (s *Stack[T]) isEmpty() bool { // Fixed: removed space in method name
	return len(s.elements) == 0
}

func (s Stack[T]) printAll() {
	if len(s.elements) == 0 {
		fmt.Println("No elements to be printed")
		return
	}
	fmt.Println("Stack elements:")
	for _, element := range s.elements {
		fmt.Printf("%v\n", element) // Fixed: fmt.Printf and correct variable
	}
}

func main() {
	// Optional: swap demo
	// x, y := 1, 2
	// fmt.Println("x:", x, "y:", y)
	// x, y = swap(x, y)
	// fmt.Println("x:", x, "y:", y)

	// intStack := Stack[int]{}
	// intStack.push(1) // Fixed: typo "insStack" → "intStack"
	// intStack.push(2)
	// intStack.push(3)
	// intStack.push(4)
	// intStack.push(5)
	// intStack.printAll()
	// fmt.Println(intStack.pop())
	// fmt.Println(intStack.pop())
	// fmt.Println(intStack.pop())
	// intStack.printAll()

	stringStack := Stack[string]{}
	stringStack.push("Amardip")
	stringStack.push("Karn")
	stringStack.push("Arunima")
	stringStack.push("Karn")
	stringStack.printAll()

}
