// --Summary:
//
//	Create a program that directs vehicles at a mechanic shop
//	to the correct vehicle lift, based on vehicle size.
//
// --Requirements:
// * The shop has lifts for multiple vehicle sizes/types:
//   - Motorcycles: small lifts
//   - Cars: standard lifts
//   - Trucks: large lifts
//   - Write a single function to handle all of the vehicles
//     that the shop works on.
//   - Vehicles have a model name in addition to the vehicle type:
//   - Example: "Truck" is the vehicle type, "Road Devourer" is a model name
//   - Direct at least 1 of each vehicle type to the correct
//     lift, and print out the vehicle information.
//
// --Notes:
// * Use any names for vehicle models
package main

import "fmt"

type Lifter interface {
	printLiftType()
}

type Motorcycle string
type Car string
type Truck string

func (m Motorcycle) printLiftType() {
	fmt.Println("Small lift")
}

func (m Car) printLiftType() {
	fmt.Println("Standard lift")
}

func (m Truck) printLiftType() {
	fmt.Println("Large lift")
}

func printAllInfo(vehicles []Lifter) {
	fmt.Println("Show vehicles lifts:")
	for _, vehicle := range vehicles {
		fmt.Printf("--Vehicle: %v--\n", vehicle)
		vehicle.printLiftType()
	}
	fmt.Println()
}

func main() {
	vehicles := []Lifter{Motorcycle("A"), Car("B"), Truck("C")}
	printAllInfo(vehicles)
}
