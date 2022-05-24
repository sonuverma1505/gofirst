package main

import "fmt"

func main() {
	x := 5
	for {
		fmt.Println("hey sonu", x)
		x += 1
		if x > 25 {
			break
		}
	}
}
