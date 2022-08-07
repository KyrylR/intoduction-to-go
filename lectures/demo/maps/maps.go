package main

import "fmt"

func main() {
	shoppingList := make(map[string]int)
	shoppingList["eggs"] = 111
	shoppingList["milk"] = 1
	shoppingList["bread"] += 2
	shoppingList["eggs"] += 1
	fmt.Println(shoppingList)

	delete(shoppingList, "milk")
	fmt.Println("Milk deleted, new map:", shoppingList)

	fmt.Println("Need", shoppingList["eggs"], "eggs")

	cereal, found := shoppingList["cereal"]
	if !found {
		fmt.Println("Default val:", cereal)
	} else {
		fmt.Println("Number of cereals:", cereal)
	}

	totalItems := 0
	for _, value := range shoppingList {
		totalItems += value
	}
	fmt.Println("Total number of items:", totalItems)
}
