package main

import (
	"fmt"
)

type Persona struct {
	Nombre string
	Edad   int
}

func (p Persona) Mostrar() {
	fmt.Println(p)
}

func byval(algo string) {

	ref := "hola" + algo
	algo = ref
}

func byref(algo *string) {
	ref := "hola " + *algo
	*algo = ref

}

func main() {

	//fmt.Println("hola ucc")
	//var p Persona
	//p.Nombre = "Manuela"
	//p.Edad = 20
	//p.Mostrar()
	//for i := 0; i < 5; i++ {
	//	defer fmt.Printf("%d", i) //ejecuta lo que esta despues del defer al final y lo pone en una pila
	//}

	s := "Esteban" //valor original
	byval(s)
	fmt.Println(s)
	byref(&s)
	fmt.Println(s)

}
