// паттерн стратегия
package main

import (
	"fmt"
)

// Интерфейс стратегии
type Strategy interface {
	Execute()
}

// ConcreteStrategyAdd - конкретная стратегия сложения
type ConcreteStrategyAdd struct{}

func (csa ConcreteStrategyAdd) Execute() {
	fmt.Println("Executing add strategy")
}

// ConcreteStrategySubtract - конкретная стратегия вычитания
type ConcreteStrategySubtract struct{}

func (css ConcreteStrategySubtract) Execute() {
	fmt.Println("Executing subtract strategy")
}

// Context - контекст, использующий стратегию
type Context struct {
	strategy Strategy
}

func (c *Context) SetStrategy(strategy Strategy) {
	c.strategy = strategy
}

func (c *Context) ExecuteStrategy() {
	c.strategy.Execute()
}

func main() {
	context := Context{}

	strategyAdd := ConcreteStrategyAdd{}
	context.SetStrategy(strategyAdd)
	context.ExecuteStrategy()

	strategySubtract := ConcreteStrategySubtract{}
	context.SetStrategy(strategySubtract)
	context.ExecuteStrategy()
}
