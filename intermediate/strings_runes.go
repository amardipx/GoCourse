package intermediate

import (
	"fmt"
	"unicode/utf8"
)

func main() {

	message := "Hello, Go!"
	rawMessage := `Hello\nGo` // raw string literal, escape sequences will also be printed

	fmt.Println(message)
	fmt.Println(rawMessage)

	fmt.Println("The first character in message var is:", message[0]) // ASCII value of H

	for i, char := range message {
		fmt.Printf("Character at index %d is %c\n", i, char)
		// fmt.Printf("%x\n", char) // Hexadecimal value
	}

	fmt.Println("No of runes in message: ", utf8.RuneCountInString(message))

	// Runes
	var ch rune = 'a'
	fmt.Println(ch)
	fmt.Printf("%c\n", ch)
}
