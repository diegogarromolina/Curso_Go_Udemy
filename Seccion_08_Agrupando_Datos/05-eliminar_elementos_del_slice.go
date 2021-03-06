package main

import (
	"fmt"
)

func main() {
	x := []int{4, 5, 7, 8, 42}
	fmt.Println(x)

	x = append(x, 44, 55, 66)
	fmt.Println(x)

	y := []int{1, 2, 3}
	x = append(x, y...)
	fmt.Println(x)

	x = append(x[:2], x[4:]...)
	fmt.Println(x)
}
