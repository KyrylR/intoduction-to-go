package main

import "fmt"

type Room struct {
	name    string
	cleaned bool
}

func checkCleanliness(rooms [4]Room) {
	//rooms[0].cleaned = true
	for i := 0; i < len(rooms); i++ {
		room := rooms[i]
		if room.cleaned {
			fmt.Println(room.name, "is clean")
		} else {
			fmt.Println(room.name, "is dirty")
		}
	}
}

func main() {
	rooms := [...]Room{
		{name: "1"},
		{name: "2"},
		{name: "3"},
		{name: "4"},
	}
	checkCleanliness(rooms)
	//fmt.Println(rooms[0])
	fmt.Println("Performing cleaning...")
	rooms[2].cleaned = true
	rooms[3].cleaned = true
	checkCleanliness(rooms)
}
