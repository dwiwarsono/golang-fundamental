package main

import (
	"fmt"
)

func main() {
	// LOGICAL OPERATOR
	a := true
	b := false

	//AND &&
	fmt.Println(a && b) //HUKUM LOGICAL &&, KALAU SALAH SATU NYA BERNILAI FALSE MAKA HASILNYA AKAN FALSE

	//OR ||
	fmt.Println(a || b) // HUKUM LOGICAL ||, KALAU SALAH SATU BERNILAI TRUE MAKA HASILNYA AKAN TRUE

	//NOT !
	fmt.Println(!a) // MEMBALIKAN NILAI AWALNYA, JIKA NILAI AWALNYA TRUE MAKA HASILNYA AKAN FALSE, BEGITU SEBALIKNYA
}