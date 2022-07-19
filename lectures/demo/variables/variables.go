package main

import "fmt"

func main() {
	var myName = "Alex"
	fmt.Println("My name is", myName)

	var name string = "Kathy"
	fmt.Println("name =", name)

	userName := "admin"
	fmt.Println("username =", userName)

	var sum int
	fmt.Println("The sum is", sum)

	part1, other := 1, 5
	fmt.Println("Part 1 is", part1, "other is", other)

	part2, other := 3, 52
	fmt.Println("Part 2 is", part2, "other is", other)

	sum = part1 + part2
	fmt.Println("The sum equal: ", sum)

	var (
		lessonName = "Variables"
		lessonType = "Demo"
	)
	fmt.Println("Lesson type: ", lessonType)
	fmt.Println("Lesson name: ", lessonName)

	word1, word2, _ := "Hello", "World", "!"
	fmt.Println(word1, word2)
}
