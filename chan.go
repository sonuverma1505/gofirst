package main

import (
	"fmt"
)

func main() {
	c := make(chan string, 2)
	c <- "hello"
	c <- "baby"

	msg := <- c
	fmt.Println(msg)

	msg = <- c
	fmt.Println(msg)

}