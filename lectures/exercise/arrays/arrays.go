//--Summary:
//  Create a program that can store a shopping list and print
//  out information about the list.
//
//--Requirements:
//* Using an array, create a shopping list with enough room
//  for 4 products
//  - Products must include the price and the name
//* Insert 3 products into the array
//* Print to the terminal:
//  - The last item on the list
//  - The total number of items
//  - The total cost of the items
//* Add a fourth product to the list and print out the
//  information again

package main

import "fmt"

// Product include the price and the name
type Product struct {
	name  string
	price int
}

func getProductsStatus(products [4]Product) (int, int) {
	totalNumberOfItems, totalPrice := 0, 0
	for i := 0; i < len(products); i++ {
		item := products[i]
		totalPrice += item.price

		if item.name != "" {
			totalNumberOfItems += 1
		}
	}
	return totalNumberOfItems, totalPrice
}

func printInfo(products [4]Product) {
	totalNumberOfItems, totalPrice := getProductsStatus(products)
	fmt.Println("The last item on the list:", products[totalNumberOfItems-1])
	fmt.Println("The total number of items:", totalNumberOfItems)
	fmt.Println("The total cost of the items:", totalPrice)
}

func main() {
	products := [4]Product{
		{name: "1", price: 12},
		{name: "2", price: 2},
		{name: "3", price: 13},
	}
	printInfo(products)
	fmt.Println("Adding fourth element...")
	products[3] = Product{name: "4", price: 4}
	printInfo(products)
}
