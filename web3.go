package main

import (
	"fmt"
)

func inc(x *int){
	x++
}

func main(){
	i := 5
	inc(&i)
	fmt.Printlin(i)
}