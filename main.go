package main

import (
	"fmt"
)

func Fizz() string {
	return "Fizz"
}

func Buzz() string {
	return "Buzz"
}

func main() {
	for i := 1; i <= 100; i++ {
		switch {
		case i%3 == 0:
			fmt.Println("Fizz")
		case i%5 == 0:
			fmt.Println("Buzz")
		case i%15 == 0:
			fmt.Println("FizzBuzz")
		default:
			fmt.Println(i)
		}
	}
}
