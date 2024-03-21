package main

import "fmt"

func main() {
	arr := []string{
		"робка",
		"борка",
		"акроб",
		"суп",
		"ПУС",
		"Разборка",
		"борка",
		"почка",
		"суп",
		"кочка",
		"кочка",
	}
	v := FindAnnagrams(&arr)
	fmt.Println(v)
}
