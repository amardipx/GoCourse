package basics

import "fmt"

func main() {
	arr := [3]int{1, 2, 3}
	fmt.Println(arr)

	for i, v := range arr {
		fmt.Printf("Index: %d Value %d\n", i, v)
	}

	var arr1 [3]int = [3]int{10, 20, 30}
	fmt.Println(arr1)

	matrix := [3][3]int{
		{1, 2, 3},
		{4, 5, 6},
		{7, 8, 9},
	}

	fmt.Println(matrix)
	for i, v := range matrix {
		fmt.Printf("Index: %d Value: %d", i, v)
	}
}
