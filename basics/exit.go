package basics

// stops execution of program abruptly. does not invoke deferred functions. nothing after exit fnc is executed.

import (
	"fmt"
	"os"
)

func main() {
	fmt.Println("Starting the main function")

	defer fmt.Println("Deferred from output")

	os.Exit(1)
}
