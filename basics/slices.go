package basics

import (
	"fmt"
	"slices"
)

func main() {
	// nil slice: var numbers []int
	var numbers1 = []int{1, 2, 3, 4, 5}

	slice := numbers1[1:4]
	fmt.Println(slice)

	slice = append(slice, 6, 7)
	fmt.Println("this is slice: ", slice)

	// copying a slice
	sliceCopy := make([]int, len(slice))
	copy(sliceCopy, slice)
	fmt.Println("this is slice copy: ", sliceCopy)

	for i, v := range slice {
		fmt.Println(i, v)
	}

	if slices.Equal(slice, sliceCopy) {
		fmt.Println("Equal")
	} else {
		fmt.Println("Not equal")
	}

	slice2 := slice[1:3]
	fmt.Println(slice2)

}
