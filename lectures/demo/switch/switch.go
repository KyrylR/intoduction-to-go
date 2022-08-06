package main

import "fmt"

func price() int {
	return 12
}

const (
	Economy    = 0
	Business   = 1
	FirstClass = 2
)

func main() {
	switch p := price(); {
	case p < 2:
		fmt.Println("cheap item")
		fmt.Println("second")
	case p < 10:
		fmt.Println("Moderately price item")
	default:
		fmt.Println("Expensive item")
	}

	ticket := Business
	switch ticket {
	case Economy:
		fmt.Println("Economy ticket")
	case Business:
		fmt.Println("Economy ticket")
	case FirstClass:
		fmt.Println("Economy ticket")
	default:
		fmt.Println("Other ticket")
	}
}
