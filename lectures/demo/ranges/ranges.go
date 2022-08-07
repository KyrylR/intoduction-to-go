package main

import "fmt"

func main() {
	slice := []string{"Hello", "world", "!"}

	for i, elem := range slice {
		fmt.Println(i, elem)
		for _, ch := range elem {
			fmt.Printf("%q - %c - %v - %T ", ch, ch, ch, ch)
		}
		fmt.Println()
	}
}
