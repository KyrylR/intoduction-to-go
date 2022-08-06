//--Summary:
//  Create a program to calculate the area and perimeter
//  of a rectangle.
//
//--Requirements:
//* Create a rectangle structure containing its coordinates
//* Using functions, calculate the area and perimeter of a rectangle,
//  - Print the results to the terminal
//  - The functions must use the rectangle structure as the function parameter
//* After performing the above requirements, double the size
//  of the existing rectangle and repeat the calculations
//  - Print the new results to the terminal
//
//--Notes:
//* The area of a rectangle is length*width
//* The perimeter of a rectangle is the sum of the lengths of all sides

package main

import "fmt"

type Coordinates struct {
	x, y int
}

type Rectangle struct {
	a, b Coordinates
}

func width(figure Rectangle) int {
	return figure.b.x - figure.a.x
}

func length(figure Rectangle) int {
	return figure.b.y - figure.a.y
}

//* Using functions, calculate the area and perimeter of a rectangle,
//  - Print the results to the terminal
//  - The functions must use the rectangle structure as the function parameter

func calculatePerimeter(figure Rectangle) {
	fmt.Println("Perimeter of Rectangle is:", 2*(width(figure)+length(figure)))
}

func calculateArea(figure Rectangle) {
	fmt.Println("Area of Rectangle is:", width(figure)*length(figure))
}

func printInfo(figure Rectangle) {
	calculateArea(figure)
	calculatePerimeter(figure)
}

func main() {
	// Create a rectangle structure containing its coordinates
	fig1 := Rectangle{Coordinates{0, 0}, Coordinates{2, 2}}
	printInfo(fig1)

	//* After performing the above requirements, double the size
	//  of the existing rectangle and repeat the calculations
	fmt.Println("Double sides.")
	fig1.b.x *= 5
	fig1.b.y *= 2

	//  - Print the new results to the terminal
	printInfo(fig1)
}
