package main

import (
	"fmt"
	"os"
)

/*
переменная err имеет тип *os.PathError, а nil представляет собой значение по
умолчанию для указателей, а не сам указатель.
*/
func Foo() error {
	var err *os.PathError = nil
	return err
}

func main() {
	err := Foo()
	fmt.Println(err)
	fmt.Println(err == nil)
}

/*
Интерфейсы в Go представляют собой набор методов. Пустой интерфейс
не имеет ни одного метода и по сути является "контейнером" для любого значения,
так как любой тип в Go реализует хотя бы 0 методов. Значения пустого интерфейса
могут содержать любой тип данных. В то время как указанный тип интерфейса
может содержать только типы, которые реализуют определенные методы.
*/