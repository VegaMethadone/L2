package main

import (
	"fmt"
)

/*
Состояние (State) - шаблон проектирования, который позволяет объекту изменять свое поведение в зависимости от внутреннего состояния.

Когда применяется данный паттерн?
Когда поведение объекта должно зависеть от его состояния и может изменяться динамически во время выполнения

Когда в коде методов объекта используются многочисленные условные конструкции, выбор которых зависит от текущего состояния объекта
*/

type State interface {
	Handle()
}

type ConcreteStateA struct{}
type ConcreteStateB struct{}
type Context struct{ state State }

func (c *Context) SetState(state State) { c.state = state }
func (c *Context) Request()             { c.state.Handle() }

func (s *ConcreteStateA) Handle() {
	fmt.Println("Handling in ConcreteStateA")
}

func (s *ConcreteStateB) Handle() {
	fmt.Println("Handling in ConcreteStateB")
}

func main() {
	context := &Context{}

	context.SetState(&ConcreteStateA{})

	context.Request()

	context.SetState(&ConcreteStateB{})

	context.Request()
}
