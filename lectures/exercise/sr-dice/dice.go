//--Summary:
//  Create a program that can perform dice rolls. The program should
//  report the results as detailed in the requirements.
//
//--Requirements:
//* Print the sum of the dice roll
//* Print additional information in these circumstances:
//  - "Snake eyes": when the total roll is 2, and total dice is 2
//  - "Lucky 7": when the total roll is 7
//  - "Even": when the total roll is even
//  - "Odd": when the total roll is odd
//* The program must handle any number of dice, rolls, and sides
//
//--Notes:
//* Use packages from the standard library to complete the project

package main

import (
	"fmt"
	"math/rand"
	"time"
)

func rollTheDice(sides int) int {
	return rand.Intn(sides) + 1
}

func printResult(total, dice int) {
	fmt.Println("Result:", total)
	switch {
	case total == 2 && dice == 2:
		fmt.Println("Snake eyes")
	case total == 7:
		fmt.Println("Lucky 7")
	case total%2 == 0:
		fmt.Println("Even")
	case total%2 != 0:
		fmt.Println("Odd")
	}
}

func main() {
	//* Print the sum of the dice roll
	//* Print additional information in these circumstances:
	//  - "Snake eyes": when the total roll is 2, and total dice is 2
	//  - "Lucky 7": when the total roll is 7
	//  - "Even": when the total roll is even
	//  - "Odd": when the total roll is odd
	//* The program must handle any number of dice, rolls, and sides
	rand.Seed(time.Now().UnixNano())
	dice, sides := 2, 6
	rolls := 3

	for r := 0; r < rolls; r++ {
		total := 0
		for d := 0; d < dice; d++ {
			total += rollTheDice(sides)
		}
		printResult(total, dice)
	}

}
