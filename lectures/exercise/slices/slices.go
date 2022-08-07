//--Summary:
//  Create a program to manage parts on an assembly line.
//
//--Requirements:
//* Using a slice, create an assembly line that contains type Part
//* Create a function to print out the contents of the assembly line
//* Perform the following:
//  - Create an assembly line having any three parts
//  - Add two new parts to the line
//  - Slice the assembly line so it contains only the two new parts
//  - Print out the contents of the assembly line at each step
//--Notes:
//* Your program output should list 3 parts, then 5 parts, then 2 parts

package main

import "fmt"

type Part string

// * Create a function to print out the contents of the assembly line
func printStatus(title string, assemblyLine []Part) {
	fmt.Println()
	fmt.Println("---", title, "---")
	for i := 0; i < len(assemblyLine); i++ {
		elem := assemblyLine[i]
		fmt.Print(elem, " ")
	}
}

func main() {
	//* Using a slice, create an assembly line that contains type Part
	//  - Create an assembly line having any three parts
	assemblyLine := []Part{"1", "2", "3"}
	printStatus("Step 1", assemblyLine)

	//* Perform the following:
	//  - Add two new parts to the line
	assemblyLine = append(assemblyLine, "4", "5")
	printStatus("Step 2", assemblyLine)
	//  - Slice the assembly line so it contains only the two new parts
	assemblyLine = assemblyLine[3:]
	printStatus("Step 3", assemblyLine)
}
