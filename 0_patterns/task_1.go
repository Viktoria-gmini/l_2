// паттерн фасад
package main

import "fmt"

type Cow struct {
}

type Horse struct {
}

type Sheep struct {
}

func (c Cow) sayMoo() string {
	return "Cow is saying \"Mooo\""
}
func (h Horse) sayIgogog() string {
	return "Horse is saying \"Igogog\""
}
func (s Sheep) sayBeee() string {
	return "Sheep is saying \"Beeee\""
}

type Facade struct{}

func (f Facade) ActivateAnimalSounds() {
	c := Cow{}
	h := Horse{}
	s := Sheep{}
	fmt.Println(c.sayMoo())
	fmt.Println(h.sayIgogog())
	fmt.Println(s.sayBeee())
}
func main() {
	f := Facade{}
	f.ActivateAnimalSounds()
}
