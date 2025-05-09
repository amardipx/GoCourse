package intermediate

import "fmt"

func main() {

	// s := fmt.Sprint("Hello", "World!", 123, 456)
	// fmt.Println(s)

	// s1 := fmt.Sprintln("Hello", "World!", 123, 456)
	// fmt.Print(s1)

	// similarly sprintf

	// var name string
	// var age int

	// fmt.Println("Enter your name and age:")
	// // fmt.Scan(&name, &age)
	// fmt.Scanln(&name, &age)
	// fmt.Printf("Name: %s, Age: %d\n", name, age)

	// Error formatting functions

	err := checkAge(15)
	if err != nil {
		fmt.Println("Error:", err)
	}

}

func checkAge(age int) error {
	if age < 18 {
		return fmt.Errorf("Age %d is too young to drive", age)
	} else {
		return nil
	}
}
