package main

import (
	"fmt"
)

func main() {
	a := [5]int{76, 77, 78, 79, 80}
	//с первого по четвёртый не включительно, поэтому будут выведены с 77 по 79
	var b []int = a[1:4]
	fmt.Println(b)
}
