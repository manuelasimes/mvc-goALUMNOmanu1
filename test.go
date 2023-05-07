package main //funcion donde corre el prpgrama 

import "fmt" 
func Suma(values ...int) int {
		suma := 0
		for _, v := range values {
			suma += v
}
		return suma
} 


func main () {

ret:=Suma(5,2)
fmt.Println("Total es: ", ret)


}
