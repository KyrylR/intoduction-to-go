package main

import "fmt"

func main() {
	sum := 0
	fmt.Println("Sum is: ", sum)
	for i := 0; i <= 10; i++ {
		sum += i
		fmt.Print(sum, " ")
	}

	fmt.Println("\nDecrement. Sum is: ", sum)
	for sum >= 10 {
		sum -= 5
		fmt.Print(sum, " ")
	}
}
