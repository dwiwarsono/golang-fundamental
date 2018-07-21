package main

import (
	"fmt"
)

func main () {
	// nextValue := getNumber()

	// fmt.Println(nextValue()) // 0 + 1 = 1
	// fmt.Println(nextValue()) // 1 + 1 = 2
	// fmt.Println(nextValue()) //
	// fmt.Println(nextValue())


	lv := love("Dylan Abinaya")
	fmt.Println(lv("Golang"))
	fmt.Println(lv("Java"))
}

//FUNCTION yang mengembalikan FUNCTION lain (FUNGSI RETURN FUNGSI(Closure))

func getNumber() func () int {
	x := 0
	return func() int {
		x = x + 1
		return x
	}
}

func love (name string) func (string) string {
	return func(things string) string{
		return fmt.Sprintf("%s love %s", name, things)
	}
}