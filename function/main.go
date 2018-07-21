package main

import (
	"fmt"
)

func main() {
	fmt.Println("Fungsi main")

	x := 5
	y := 5

	z := add(x, y)
	fmt.Println(z)

	// INI UNTUK CONTOH 2
	name := "Dylan Abinaya"
	hasil := hello(name)
	fmt.Println(hasil)

	//INI UNTUK CONTOH 3
	xSwap := "Hello"
	ySwap := "Golang"
	resultX, resultY := swap(xSwap, ySwap)
	fmt.Println(resultX, resultY)
}

// CONTOH FUNCTION 1
func add(x int, y int) int {
	return x + y //Return atau mengembalikan nilai
}

// CONTOH FUNCTION 2
func hello(name string) string { //name adalah argument
	return fmt.Sprintf("Hello %s", name) //FUNGSI Sprintf ini concat (MENGGABUNGKAN DUA STRING atau TIPE DATA LAIN)
}

// CONTOH FUNCTION 3, multi value
func swap(x string, y string) (string, string) {
	return y, x
}