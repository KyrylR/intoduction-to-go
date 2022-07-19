//--Summary:
//  Use functions to perform basic operations and print some
//  information to the terminal.
package main

import "fmt"

//--Requirements:
//* Write a function that accepts a person's name as a function
//  parameter and displays a greeting to that person.
func welcomeTheMan(name string) {
	fmt.Println("Welcome,", name)
}

//* Write a function that returns any message, and call it from within
//  fmt.Println()
func returnsAnyMessage() string {
	return "something"
}

//* Write a function to add 3 numbers together, supplied as arguments, and
//  return the answer
func tripleAdd(firstNum, secondNum, thirdNum int) int {
	return firstNum + secondNum + thirdNum
}

//* Write a function that returns any number
func returnsAnyNumber() int {
	return 5
}

//* Write a function that returns any two numbers
func returnsAnyTwoNumbers() (int, int) {
	return 3, 9
}

//* Add three numbers together using any combination of the existing functions.
//  * Print the result
//* Call every function at least once

func main() {
	welcomeTheMan("Alex")

	//* Write a function that returns any message, and call it from within
	//  fmt.Println()
	fmt.Println("Message from function is", returnsAnyMessage())

	firstAns := returnsAnyNumber()
	secondAns, thirdAns := returnsAnyTwoNumbers()
	fmt.Println("Tummy nested funcs result :) ->", tripleAdd(firstAns, secondAns, thirdAns))
}
