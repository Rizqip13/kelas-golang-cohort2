package main

import "fmt"

func main() {
	n := 25
	var isFizz, isBuzz bool
	for i := 1; i < n; i++ {
		isBuzz = (i % 5) == 0
		isFizz = (i % 3) == 0

		switch {
		case isBuzz && isFizz:
			fmt.Println("FizzBuzz")
		case isBuzz:
			fmt.Println("Buzz")
		case isFizz:
			fmt.Println("Fizz")
		default:
			fmt.Println(i)
		}
	}
}
