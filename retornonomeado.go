package main

import "fmt"

func split(sum int) (x, y int) {
	x = sum 
	y = sum - x
	return
}

func main() {
	fmt.Println(split(2))
}
