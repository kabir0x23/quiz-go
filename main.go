package main

import "fmt"

func main() {
	fmt.Println("Welcome to my quiz game")

	var name string
	fmt.Printf("Enter your name: ")
	fmt.Scan(&name)

	fmt.Printf("Namaste, %v, Welcome to the game!\n", name)

	var age uint
	fmt.Printf("Enter your age: ")
	fmt.Scan(&age)

	if age >= 10 {
		fmt.Println("Yah you can play!")
	} else {
		fmt.Println("You can't play!")
		return
	}

	score := 0
	numQues := 2

	op1 := 1
	op2 := 2

	fmt.Printf("What is bigger, %v or %v: ", op1, op2)

	var answer int
	fmt.Scan(&answer)

	if answer == op2 {
		fmt.Println("Correct!")
		score++
	} else {
		fmt.Println("Incorrect!")
	}

	fmt.Println("How many cores does the Ryzen 9 3900x have? ")
	var cores uint
	fmt.Scan(&cores)

	if cores == 12 {
		fmt.Println("Correct!")
		score++
	} else {
		fmt.Println("Incorrect!")
	}

	fmt.Printf("You scored %v out of %v.\n", score, numQues)
	scored := (float64(score) / float64(numQues)) * 100
	fmt.Printf("You scored %v%%", scored)

}
