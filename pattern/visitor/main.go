package main

import "fmt"

/*
Паттерн Посетитель (Visitor) позволяет определить операцию для объектов других классов без изменения этих классов.
При использовании паттерна Посетитель определяются две иерархии классов: одна для элементов, для которых надо определить новую операцию, и вторая иерархия для посетителей, описывающих данную операцию.

Когда использовать данный паттерн?
Когда имеется много объектов разнородных классов с разными интерфейсами, и требуется выполнить ряд операций над каждым из этих объектов
Когда классам необходимо добавить одинаковый набор операций без изменения этих классов
Когда часто добавляются новые операции к классам, при этом общая структура классов стабильна и практически не изменяется

*/

type Visitor interface {
	VisitConcreteElementA(elementA *ConcreteElementA)
	VisitConcreteElementB(elementB *ConcreteElementB)
}

type ConcreteElementA struct {
	Name string
}
type ConcreteElementB struct {
	Value int
}

func (e *ConcreteElementA) Accept(visitor Visitor) {
	visitor.VisitConcreteElementA(e)
}
func (e *ConcreteElementB) Accept(visitor Visitor) {
	visitor.VisitConcreteElementB(e)
}

type ConcreteVisitor struct{}

func (v *ConcreteVisitor) VisitConcreteElementA(elementA *ConcreteElementA) {
	fmt.Printf("Visiting ConcreteElementA with Name: %s\n", elementA.Name)
}

func (v *ConcreteVisitor) VisitConcreteElementB(elementB *ConcreteElementB) {
	fmt.Printf("Visiting ConcreteElementB with Value: %d\n", elementB.Value)
}

func main() {
	elementA := &ConcreteElementA{Name: "Element A"}
	elementB := &ConcreteElementB{Value: 420}

	visitor := &ConcreteVisitor{}

	elementA.Accept(visitor)
	elementB.Accept(visitor)
}
