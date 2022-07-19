package main

import "fmt"

func double(x int) int {
	return x + x
}

func add(lhs, rhs int) int {
	return lhs + rhs
}

func greet() {
	fmt.Println("Hello from greet function!")
}

func main() {
	greet()

	dozen := double(6)
	fmt.Println("A dozen is", dozen)

	bakerDozen := add(dozen, 1)
	fmt.Println("A baker's dozen is", bakerDozen)

	cookDozen := add(double(dozen), 1)
	fmt.Println("A cook's dozen is", cookDozen)
}
