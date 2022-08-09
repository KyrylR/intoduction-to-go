package main

import "fmt"

type Preparer interface {
	PrepareDish()
}

type Chicken string
type Salad string

func (c Chicken) PrepareDish() {
	fmt.Println("Cook chicken")
}

func (c Salad) PrepareDish() {
	fmt.Println("Chop salad")
}

func prepareDishes(dishes []Preparer) {
	fmt.Println("Prepare dishes:")
	for _, dish := range dishes {
		fmt.Printf("--Dish: %v--\n", dish)
		dish.PrepareDish()
	}
	fmt.Println()
}

func main() {
	dishes := []Preparer{Chicken("Chicken Wings"), Salad("Caesar")}
	prepareDishes(dishes)
}
