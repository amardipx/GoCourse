package basics

import "fmt"

func main() {
	message := "Hello world"

	for i, v := range message {
		// fmt.Println(i, v)
		fmt.Printf("%d %c\n", i, v)
	}
}
