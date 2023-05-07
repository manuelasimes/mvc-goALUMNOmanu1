package main

import "fmt"

func Calculate(a int) int {
	return a + 2
}

func main() {
	s := Calculate(2)
	fmt.Println(s)
}
