package main

import (
	"fmt"
)

func main() {
	x := 10
	y := 5

	// Aritmatika operator


	// Penambahan / Additional
	zAdd := x + y
	fmt.Println(zAdd)

	// Pengurangan / Subtracts
	zSub := x - y
	fmt.Println(zSub)

	// Perkalian / Multiplies
	zMul := x * y
	fmt.Println(zMul)

	// Pembagian / Divides
	zDiv := x / y
	fmt.Println(zDiv)

	// Sisa hasil bagi / Modulus
	zMod := x % y
	fmt.Println(zMod)

	//RELATIONAL OPERATOR
	// RELATIONAL OPERATOR ini bernilai boolean (true, false)
	// ==, !=, >, <, >=, <=

	i := 10
	j := 5

	fmt.Println(i == j) // == apakah i sama dengan j, perbandingan nilai
	fmt.Println(i != j) // != apakah i tidak sama dengan j
	fmt.Println(i < j) // < apakah i lebih kecil dari j
	fmt.Println(i > j) // > apakah i lebih besar dari j
	fmt.Println(i >= j) // apakah i lebih besar atau sama dengan j
	fmt.Println(i <= j) // apakah i lebih kecil atau sama dengan j
}