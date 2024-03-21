//фабричный метод

package main

import "fmt"

// Интерфейс продукта
type Product interface {
	Use()
}

// ConcreteProductA - конкретная реализация продукта A
type ConcreteProductA struct{}

func (cpa ConcreteProductA) Use() {
	fmt.Println("Using product A")
}

// ConcreteProductB - конкретная реализация продукта B
type ConcreteProductB struct{}

func (cpb ConcreteProductB) Use() {
	fmt.Println("Using product B")
}

// Интерфейс фабрики
type Factory interface {
	CreateProduct() Product
}

// ConcreteFactoryA - фабрика для создания продукта A
type ConcreteFactoryA struct{}

func (cfa ConcreteFactoryA) CreateProduct() Product {
	return ConcreteProductA{}
}

// ConcreteFactoryB - фабрика для создания продукта B
type ConcreteFactoryB struct{}

func (cfb ConcreteFactoryB) CreateProduct() Product {
	return ConcreteProductB{}
}

func main() {
	factoryA := ConcreteFactoryA{}
	productA := factoryA.CreateProduct()
	productA.Use()

	factoryB := ConcreteFactoryB{}
	productB := factoryB.CreateProduct()
	productB.Use()
}
