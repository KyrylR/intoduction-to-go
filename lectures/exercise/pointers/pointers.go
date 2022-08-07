//--Summary:
//  Create a program that can activate and deactivate security tags
//  on products.
//
//--Requirements:
//* Create a structure to store items and their security tag state
//  - Security tags have two states: active (true) and inactive (false)
//* Create functions to activate and deactivate security tags using pointers
//* Create a checkout() function which can deactivate all tags in a slice
//* Perform the following:
//  - Create at least 4 items, all with active security tags
//  - Store them in a slice or array
//  - Deactivate any one security tag in the array/slice
//  - Call the checkout() function to deactivate all tags
//  - Print out the array/slice after each change

package main

import "fmt"

//* Create a structure to store items and their security tag state
//  - Security tags have two states: active (true) and inactive (false)

const (
	active   = true
	inactive = false
)

type Tag bool

type Item struct {
	name string
	tag  Tag
}

// * Create functions to activate and deactivate security tags using pointers
func changeTagStatusOfItem(item *Item) {
	switch item.tag {
	case active:
		item.tag = inactive
	case inactive:
		item.tag = active
	default:
		panic("Unhandled security tag status")
	}
}

// * Create a checkout() function which can deactivate all tags in a slice
func checkout(items *[]Item) {
	for i := 0; i < len(*items); i++ {
		(*items)[i].tag = inactive
	}
	//for _, item := range *items {
	//	item.tag = inactive
	//}
}

func main() {
	//* Perform the following:
	//  - Create at least 4 items, all with active security tags
	item1 := Item{"1", active}
	item2 := Item{"2", active}
	item3 := Item{"3", active}
	item4 := Item{"4", active}
	//  - Store them in a slice or array
	items := []Item{item1, item2, item3, item4}
	fmt.Println(items)
	//  - Deactivate any one security tag in the array/slice
	changeTagStatusOfItem(&items[2])
	fmt.Println(items)
	//  - Call the checkout() function to deactivate all tags
	checkout(&items)
	fmt.Println(items)
	//  - Print out the array/slice after each change
}
