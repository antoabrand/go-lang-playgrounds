package main

import (
	"fmt"
)

func main() {
	fizzBuzz()
}

func fizzBuzz() {
	for i := 1; i < 125; i++ {
		if i%15 == 0 {
			fmt.Println("Fizzbuzz")
		} else if i%5 == 0 {
			fmt.Println("Buzz")
		} else if i%3 == 0 {
			fmt.Println("Fizz")
		} else {
			fmt.Println(i)
		}
	}
}
