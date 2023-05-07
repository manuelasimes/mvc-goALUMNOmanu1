package services

//logica de nuestro programa-->q es lo q hace!

import (
	"errors"
	client "ucc/clients"
)

// definimos la estructura q guarda la info del servicio, y q es la q usamos para implementar la interfaz
type productService struct {
	productClient client.ProductClientInterface
}

// definimos la interfaz del servicio
type productServiceInterface interface {
	GetProductId() (int, error)
	GetProductName() (string, error)
}

// Variable global de tipo productServiceInterface
var (
	ProductService productServiceInterface
)

// usamos esta funcion para que cuando se inicializa el paquete se guarda nuestra interfaz en la variable global
func init() {
	ProductService = initProductService(client.ProductClient) //inicializa product service con los parametros q le pasamos
	//ese parametro es lo q cambiaremos para los test --> inyeccion de dependencia (variar los client dependendiendo implemetancion del product service)

}

func initProductService(productClient client.ProductClientInterface) productServiceInterface {
	service := new(productService)
	service.productClient = productClient
	return service
}

// GetproductId busca(llama) un id del cliente y lo analiza para ver s devuelve error o no
// si la id es -1 devuelve error inicicando q no fue encontrado el id
// sino devuelve el id
func (s *productService) GetProductId() (int, error) {

	id := s.productClient.GetProductId()
	if id == -1 {
		return id, errors.New("no se encontro la id")
	}
	return id, nil
}

// Getproductname getproductid busca un nombre del cliente y lo analiza para ver si devuelve un error o no
// si el nombre esta vacio devuelve un error
// sino devuelve el nombre
func (s *productService) GetProductName() (string, error) {
	name := s.productClient.GetProductName()
	if len(name) <= 0 {
		return name, errors.New("el nombre esta vacio")
	}
	return name, nil
}
