package main

import "fmt"

/*
Паттерн Стратегия (Strategy) представляет шаблон проектирования, который определяет набор алгоритмов, инкапсулирует каждый из них и обеспечивает их взаимозаменяемость. В зависимости от ситуации мы можем легко заменить один используемый алгоритм другим. При этом замена алгоритма происходит независимо от объекта, который использует данный алгоритм.

Когда использовать стратегию?
Когда есть несколько родственных классов, которые отличаются поведением. Можно задать один основной класс, а разные варианты поведения вынести в отдельные классы и при необходимости их применять

Когда необходимо обеспечить выбор из нескольких вариантов алгоритмов, которые можно легко менять в зависимости от условий

Когда необходимо менять поведение объектов на стадии выполнения программы

Когда класс, применяющий определенную функциональность, ничего не должен знать о ее реализации
*/

type Strategy interface {
	Execute(a, b int) int
}

type ConcreteStrategyAdd struct{}
type ConcreteStrategySubtract struct{}
type ConcreteStrategyMultiply struct{}

func (s *ConcreteStrategyAdd) Execute(a, b int) int      { return a + b }
func (s *ConcreteStrategySubtract) Execute(a, b int) int { return a - b }
func (s *ConcreteStrategyMultiply) Execute(a, b int) int { return a * b }

type Context struct {
	strategy Strategy
}

func (c *Context) SetStrategy(strategy Strategy) { c.strategy = strategy }
func (c *Context) ExecuteStrategy(a, b int) int  { return c.strategy.Execute(a, b) }

func main() {
	contxt := &Context{}

	contxt.SetStrategy(&ConcreteStrategyAdd{})
	result := contxt.ExecuteStrategy(5, 5)
	fmt.Println(result)

	contxt.SetStrategy(&ConcreteStrategyMultiply{})
	result = contxt.ExecuteStrategy(result, 5)
	fmt.Println(result)

	contxt.SetStrategy(&ConcreteStrategySubtract{})
	result = contxt.ExecuteStrategy(result, 5)
	fmt.Println(result)
}
