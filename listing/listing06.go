package main

import (
	"fmt"
)

func main() {
	var s = []string{"1", "2", "3"}
	modifySlice(s)
	fmt.Println(s)
}

func modifySlice(i []string) {
	i[0] = "3"
	i = append(i, "4")
	i[1] = "5"
	i = append(i, "6")
}

/*
Программа выведет [3 2 3].

При передаче слайса в функцию как аргумент, копия ссылки на оригинальный слайс
создается, но сам слайс не копируется. Поэтому изменения, внесенные в элементы слайса
внутри функции, будут отражаться на оригинальном слайсе за пределами функции.

В данном случае, в функции modifySlice, первый элемент слайса i изменяется на "3",
затем к слайсу добавляются элементы "4" и "6". Но изменение элемента в оригинальном слайсе s
также отражается, поэтому первый элемент становится "3". Однако,
второй элемент "2" не изменяется, потому что после изменения первого элемента,
длина слайса i увеличивается, и второй элемент "2" остается без изменений.

Таким образом, при выводе fmt.Println(s) результат будет [3 2 3],
где изменения внесены только в первый элемент слайса.
*/
