package main

import (
	"fmt"
)

func main() {

	//CONTOH 1
	// x := 15
	// y := 10

	// if x < y {
	// 	fmt.Println("NILAI X LEBIH KECIL DARI Y")
	// } else {
	// 	fmt.Println("NILAI X LEBIH BESAR DARI Y")
	// }


	//CONTOH 2

	// if false {
	// 	fmt.Println("INI BERNILAI TRUE")
	// } else {
	// 	fmt.Println("INI BERNILAI FALSE")
	// }

	//CONTOOH 3
	// x := true
	// y := false

	// if x && y {
	// 	fmt.Println("NILAI X ADALAH TRUE")
	// } else {
	// 	fmt.Println("NILAI X ADALAH FALSE")
	// }

	//CONTOH 4
	x := 100

	if x < 10{
		fmt.Println("X LEBIH KECIL DARI 10")
	} else if x >= 100{
		fmt.Println("X LEBIH BESAR ATAU SAMA DENGAN 100")
	} else {
		fmt.Println("BERAPAKAH X?")
	}

}