package main

import "fmt"

type Product interface {
	Use() string
}
type Creator interface {
	CreateProduct() Product
}

type ConcreteProductA struct{}
type ConcreteProductB struct{}

func (p *ConcreteProductA) Use() string { return "Using Product A" }
func (p *ConcreteProductB) Use() string { return "Using Product B" }

type BaseCreator struct{}
type ConcreteCreatorA struct{ BaseCreator }
type ConcreteCreatorB struct{ BaseCreator }

func (c *BaseCreator) CreateProduct() Product      { return nil }
func (c *ConcreteCreatorA) CreateProduct() Product { return &ConcreteProductA{} }
func (c *ConcreteCreatorB) CreateProduct() Product { return &ConcreteProductB{} }

func main() {
	var creator Creator

	creator = &ConcreteCreatorA{}
	productA := creator.CreateProduct()
	fmt.Println(productA.Use())

	creator = &ConcreteCreatorB{}
	productB := creator.CreateProduct()
	fmt.Println(productB.Use())
}
