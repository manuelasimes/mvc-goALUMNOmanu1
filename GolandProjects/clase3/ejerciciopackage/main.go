package main

import (
	"test/package_a"
	"test/package_b"
)
import "fmt"

func main() {
	fmt.Println(package_a.Hello())
	fmt.Println(package_b.Hello())
}
