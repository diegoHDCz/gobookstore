package main

import "fmt"

func main() {
	x := sum(1, 3)
	fmt.Println("Hello", x)
}
func sum(x, y int) int {
	return x + y
}
