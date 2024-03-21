// патттерн состояние
package main

import "fmt"

// Интерфейс состояния
type State interface {
	Handle()
}

// ConcreteStateA - конкретное состояние A
type ConcreteStateA struct{}

func (csa ConcreteStateA) Handle() {
	fmt.Println("Handle state A")
}

// ConcreteStateB - конкретное состояние B
type ConcreteStateB struct{}

func (csb ConcreteStateB) Handle() {
	fmt.Println("Handle state B")
}

// Context - контекст, управляющий состояниями
type Context struct {
	state State
}

func (c *Context) SetState(state State) {
	c.state = state
}

func (c *Context) Request() {
	c.state.Handle()
}

func main() {
	context := Context{}

	stateA := ConcreteStateA{}
	context.SetState(stateA)
	context.Request()

	stateB := ConcreteStateB{}
	context.SetState(stateB)
	context.Request()
}
