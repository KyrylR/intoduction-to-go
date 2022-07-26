package main

import "fmt"

type Space struct {
	occupied bool
}

type ParkingLot struct {
	spaces []Space
}

func (lot *ParkingLot) occupySpace(spaceNum int) {
	lot.spaces[spaceNum-1].occupied = true
}

func (lot *ParkingLot) vacateSpace(spaceNum int) {
	lot.spaces[spaceNum-1].occupied = false
}

func main() {
	lot := ParkingLot{spaces: make([]Space, 5)}
	fmt.Println(lot)
	lot.occupySpace(3)
	fmt.Println(lot)
	lot.occupySpace(4)
	fmt.Println(lot)
	lot.occupySpace(1)
	fmt.Println(lot)
	lot.occupySpace(2)
	fmt.Println(lot)
	lot.vacateSpace(3)
	fmt.Println(lot)
}
