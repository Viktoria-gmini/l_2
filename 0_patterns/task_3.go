//патерн "посетитель"

package main

import "fmt"

type Visitor interface {
	Visit(int)
}

type Element struct {
	value int
}

func (e *Element) Accept(v Visitor) {
	v.Visit(e.value)
}

type ConcreteVisitor struct{}

func (cv *ConcreteVisitor) Visit(value int) {
	fmt.Printf("Visited element with value: %d\n", value)
}

func main() {
	elements := []Element{
		{value: 1},
		{value: 2},
		{value: 3},
	}

	visitor := &ConcreteVisitor{}

	for _, element := range elements {
		element.Accept(visitor)
	}
}
