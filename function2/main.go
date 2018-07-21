package main

import (
	"fmt"
)

// FUNGSI SEBAGAI VALUE

func main () {

	add := func(x, y int) int{ // Membuat fungsi ditampung dalam variabel
		return x + y
	}

	hello := func(name string) string{
		return fmt.Sprintf("Hello %s", name)
	}

	fmt.Println(add(5, 5))
	fmt.Println(hello("Dylan Abinaya"))
	
}

// func add(x, y int) int {
// 	return x + y
// }