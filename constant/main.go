package main

import (
	"fmt"
)

// Cara membuat constant dengan keyword const
const A string = "INI ADALAH A CONSTANT"

func main(){

	fmt.Println(A)

	const x int = 10 //Kita juga dapat membuat const didalam function
	fmt.Println(x)

	z := 100 / x //Variabel z dibagi dengan const x
	fmt.Println(z)
}