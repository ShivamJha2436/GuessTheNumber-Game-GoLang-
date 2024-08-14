package main

import (
	"fmt"       // Importing the fmt package for formatted I/O (input/output)
	"math/rand" // Importing the math/rand package to generate random numbers
	"time"      // Importing the time package to seed the random number generator
)

func main() {
	// Seed the random number generator to get different random numbers each time
	rand.Seed(time.Now().UnixNano())

	// Generate a random number between 1 and 100
	target := rand.Intn(10) + 1

	// Introduce the game
	fmt.Println("Welcome to  Guess the Number game!")
	fmt.Println("I have selected a random number between 1 and 100.")
	fmt.Println("Can you guess it?")

	var guess int
	attempts := 0

	// Start a loop to allow multiple attempts
	for {
		fmt.Print("Enter your guess: ")
		fmt.Scanln(&guess) // Read user input

		attempts++ // Increment the attempt counter

		// Check if the guess is correct
		if guess < target {
			fmt.Println("Too low! Try again.")
		} else if guess > target {
			fmt.Println("Too high! Try again.")
		} else {
			// Correct guess
			fmt.Printf("Congratulations! You guessed the number in %d attempts.\n", attempts)
			break // Exit the loop since the game is over
		}
	}
}
