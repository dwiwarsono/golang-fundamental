package main

import (
	"fmt"
)

func main() {
	// Static type declaration (Tipe variabel satic)
	var x int
	x = 10

	var y float64
	y = 5.5

	fmt.Println(x)
	fmt.Println(y)

	fmt.Printf("X memiliki tipe data %T\n", x) //Printf = Untuk mengecek tipe data seuah variabel 
	fmt.Printf("Y memiliki tipe data %T\n", y)


	// Dynamic type declaration (Tipe variabel dynamic)
	z := "Dylan" // Dynamic tidak perlu memakai keyword var lagi, otomatis go akan tahu kalau ini adalah variabel
	i := 57
	d := 5.7
	fmt.Println(z)
	fmt.Println(i)


	fmt.Printf("Z memiliki tipe data %T\n", z) //Printf = Untuk mengecek tipe data seuah variabel 
	fmt.Printf("I memiliki tipe data %T\n", i)
	fmt.Printf("D memiliki tipe data %T\n", d)

	t := 5
	p := 10
	fmt.Println(t*p)
}