package services

import (
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"testing"
)

// mockproductclient implementamos la estrcutura con mock es falso como de prueba
type mockProductClient struct {
	mock.Mock
}

// definimos las funciones con la interfaz Mock (se declara a si misma como una funcion q no hace nada)
func (c *mockProductClient) GetProductId() int {
	//cuando t llamen guarda los valores q retornarias en ret
	ret := c.Called() //guardamos en una variable tds los argumetos q pasemos en el test
	return ret.Int(0) //devuelve el primer return
}

func (c *mockProductClient) GetProductName() string {
	ret := c.Called()
	return ret.String(0)
}

// implementacion de los test (antes hicimos la interfaz)
func TestGetProductId(t *testing.T) {
	//creamos una instancia dle cliente mock
	mockClient := new(mockProductClient)

	//le definimos q nos tiene q devolver cuando llamamos al cliente
	mockClient.On("GetProductId").Return(1) //error logico

	//creamos una instancia del servicio con el cliente mock como parametro
	service := initProductService(mockClient)

	id, err := service.GetProductId()

	assert.Equal(t, 1, id)
	assert.Nil(t, err)

}

func TestProductIdWithError(t *testing.T) {
	mockClient := new(mockProductClient)
	mockClient.On("GetProductId").Return(-1)
	service := initProductService(mockClient)

	id, err := service.GetProductId()

	assert.Equal(t, -1, id)
	assert.NotNil(t, err)
}
func TestGetProductName(t *testing.T) {
	//creamos una instancia dle cliente mock
	mockClient := new(mockProductClient)

	//le definimos q nos tiene q devolver cuando llamamos al cliente
	mockClient.On("GetProductName").Return("Manu") //error logico

	//creamos una instancia del servicio con el cliente mock como parametro
	service := initProductService(mockClient)

	name, err := service.GetProductName()

	assert.Equal(t, "Manu", name)
	assert.Nil(t, err)
//basicamente le mando manu y veo si devuelve lo mism (si usca bien la bd) 
}
func TestProductNameWithError(t *testing.T) {
	mockClient := new(mockProductClient)
	mockClient.On("GetProductName").Return("")
	service := initProductService(mockClient)

	name, err := service.GetProductName()

	assert.Equal(t, "", name)
	assert.NotNil(t, err)
}
