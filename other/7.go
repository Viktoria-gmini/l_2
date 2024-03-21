package main

import (
	"fmt"
	"math/rand"
	"time"
)

func asChan(vs ...int) <-chan int {
	c := make(chan int)

	go func() {
		for _, v := range vs {
			c <- v
			time.Sleep(time.Duration(rand.Intn(1000)) * time.Millisecond)
		}

		close(c)
	}()
	return c
}

func merge(a, b <-chan int) <-chan int {
	c := make(chan int)
	go func() {
		for {
			select {
			case v := <-a:
				c <- v
			case v := <-b:
				c <- v
			}
		}
	}()
	return c
}

func main() {

	a := asChan(1, 3, 5, 7)
	b := asChan(2, 4, 6, 8)
	c := merge(a, b)
	for v := range c {
		fmt.Println(v)
	}
}

/*
вывод программы будет представлять собой комбинацию значений из каналов a и b,
где порядок значений может быть случайным из-за случайной задержки при отправке.
*/
/*
в какой-то момент начинают выводиться только нули, это может
происходить из-за того, что один из каналов a или b закрывается раньше,
чем другой. Когда один из каналов закрывается, операция чтения из него
возвращает значение по умолчанию типа канала, то есть 0 для типа int.

Вероятно причина в том, что в функции merge() нет проверки на
закрытие каналов a и b, и программа продолжит выполнение цикла for даже
после закрытия одного из каналов. В результате программа будет пытаться
читать из закрытого канала и возвращать нули.
*/
