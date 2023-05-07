package main

import (
	"errors"
	"fmt"
)

func Division(a int, b int) (int, error) {
	if b == 0 {
		return 0, errors.New("no puedo dividir por cero")
	}
	return a / b, nil
}

func main() {
	s, err := Division(9, 0)

	if err != nil {
		fmt.Println("ERROR", err.Error())
		return
	}

	fmt.Println(s)

}
