package basics

import "fmt"

func main() {

	//...Ellipsis

	fmt.Println(sum("1", 2, 3, 4, 5, 6, 7, 8))

	numbers := []int{1, 2, 3, 4}
	fmt.Println(sum("2", numbers...))

}

func sum(sequence string, nums ...int) (string, int) { // variadic parameters must come last
	total := 0
	for _, v := range nums {
		total += v
	}
	return sequence, total
}
