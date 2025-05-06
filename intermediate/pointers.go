package intermediate

import "fmt"

func main() {
	var a int = 10
	var ptr *int = &a // referencing

	fmt.Println(a)
	fmt.Println(ptr)
	fmt.Println(*ptr) // dereferencing

	modifyValue(ptr)
	fmt.Println(a)

}

func modifyValue(ptr *int) {
	*ptr++
}
