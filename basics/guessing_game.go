package basics

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	source := rand.NewSource(time.Now().UnixNano())
	random := rand.New(source)

	// Generate a random number bw 1 and 100
	target := random.Intn(100) + 1

	// Welcome message
	fmt.Println("Welcome to the game!")
	fmt.Println("guess the number bw 1 and 100")

	var guess int
	guesses := 5
	for guesses > 0 {

		fmt.Println("Enter your guess: ")
		fmt.Scanln(&guess)

		// Check if guess is correct
		if guess == target {
			fmt.Println("Congratulations you guessed correct!")
			break
		} else if guess < target {
			guesses -= 1
			fmt.Println("Too low. Guesses remaining: ", guesses)
		} else {
			guesses -= 1
			fmt.Println("Too high. Guesses remaining: ", guesses)
		}
	}
}
