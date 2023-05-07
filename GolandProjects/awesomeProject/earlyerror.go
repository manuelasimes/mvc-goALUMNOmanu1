package main

import (
	"errors"
	"fmt"
)

func division(a int, b int) (int, error) { // mandamos dos parametros y devuelve dos (int, error)
	if b == 0 { //early error
		return 0, errors.New("No podemos realizar una divison por cero ")
	}
	return a / b, nil
}

func main() {
	div, err := division(7, 4)
	if err != nil {
		fmt.Println("ERROR", err.Error())
		return
	}

	fmt.Println("Resultado: ", div)
}
