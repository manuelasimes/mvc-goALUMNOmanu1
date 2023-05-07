package product_client

import (
	"math/rand"
	"time"
)

// definimos la estructura productClient sin campos que usamos para implementar la interfaz
// define un comportamiento (ejecucion de los metodos)
type productClient struct{}

// productclientinterface definimos la interfaz q expone los metodos getproductid y getproductname
type ProductClientInterface interface {
	GetProductId() int
	GetProductName() string
}

// productclient variable global del tipo productclientinterface
var (
	ProductClient ProductClientInterface
)

//init es una especie de constructor del paquete--> cuando corremos la app se ejecutan todos los init de los paquetes
//instancia un nuevo product client
func init() {
	ProductClient = &productClient{}
}

//Gretproductid devuelve un numero aleatorio q puese ser -1 o 0
func (s *productClient) GetProductId() int { //metodos publicos--> inicia en mayuscula
	rand.Seed(time.Now().UnixNano())
	return rand.Intn(2) - 1
}

//getproductname devuelve una cadena aleatoria q puede ser vacio o "Hello"
func (s *productClient) GetProductName() string {
	rand.Seed(time.Now().UnixNano())
	if rand.Intn(2) == 0 {
		return ""
	}
	return "Hello"
}

