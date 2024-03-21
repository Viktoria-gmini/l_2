//паттерн цепочка вызовов

package main

import "fmt"

type Handler interface {
	SetSuccessor(Handler)
	HandleRequest(int)
}

type ConcreteHandler struct {
	successor Handler
}

func (ch *ConcreteHandler) SetSuccessor(handler Handler) {
	ch.successor = handler
}

func (ch *ConcreteHandler) HandleRequest(request int) {
	// Проверяем, можем ли обработать запрос
	if request <= 10 {
		fmt.Println("Request handled by ConcreteHandler")
	} else if ch.successor != nil {
		fmt.Println("Passed to next handler")
		ch.successor.HandleRequest(request)
	} else {
		fmt.Println("No handler available for this request")
	}
}

func main() {
	handler1 := &ConcreteHandler{}
	handler2 := &ConcreteHandler{}

	handler1.SetSuccessor(handler2)

	handler1.HandleRequest(5)
	handler1.HandleRequest(15)
}
