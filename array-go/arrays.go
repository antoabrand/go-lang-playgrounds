package main

import "fmt"

func main() {
	arrays(12, 4543, 111)
}

func arrays(x int, y int, z int) {
	a := [5]int{
		y,
		x,
		z,
	}
	a[4] = 100
	fmt.Println(a)
}
