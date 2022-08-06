//--Summary:
//  Implement the classic "FizzBuzz" problem using a `for` loop.
//
//--Requirements:
//* Print integers 1 to 50, except:
//  - Print "Fizz" if the integer is divisible by 3
//  - Print "Buzz" if the integer is divisible by 5
//  - Print "FizzBuzz" if the integer is divisible by both 3 and 5
//
//--Notes:
//* The remainder operator (%) can be used to determine divisibility

package main

import "fmt"

func getValue(num int) any {
	switch {
	case num%3 == 0 && num%5 == 0:
		// Print "FizzBuzz" if the integer is divisible by both 3 and 5
		return "FizzBuzz"
	case num%3 == 0:
		// Print "Fizz" if the integer is divisible by 3
		return "Fizz"
	case num%5 == 0:
		// Print "Buzz" if the integer is divisible by 5
		return "Buzz"
	default:
		// Print just a number otherwise
		return num
	}
}

func main() {
	for i := 1; i <= 50; i++ {
		fmt.Println(getValue(i))
	}
}
