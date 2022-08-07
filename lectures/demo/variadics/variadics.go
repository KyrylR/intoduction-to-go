package main

import "fmt"

func sum(nums ...int) int {
	sum := 0
	for _, item := range nums {
		sum += item
	}
	return sum
}

func main() {
	fmt.Println(sum(1, 2, 3, 4, 5))
	fmt.Println(sum(1, 2))

	a := []int{1, 2, 3}
	b := []int{4, 5, 6}
	combination := append(a, b...)
	fmt.Println(sum(a...))
	fmt.Println(sum(b...))
	fmt.Println(sum(combination...))
}
