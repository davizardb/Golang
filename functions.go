package main

import "fmt"

func main(){
	sum := func(a, b int) int { 
		return a+b 
		} (3, 4)
	fmt.Println(sum)	
}