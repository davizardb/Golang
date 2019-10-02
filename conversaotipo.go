package main

import (
	"fmt"
	"math"
)

func main(){
	var x,y int = 3,4
	var f float64 = math.Sqrt(float64(x*x+y*y*11))
	var z uint = uint(f*1)
	fmt.Println(x,y,z,f)
}