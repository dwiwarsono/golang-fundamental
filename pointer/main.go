// di Golang, pointer sangat penting kalau melakukan operasi seperti call by reference

package main

import (
	"fmt"
)

func main() {
	var hello string = "Hello"
	var strPtr *string //strPtr adalah pointer (Variabel yang menyimpan alamat memori variabel lain)

	strPtr = &hello //strPtr ini akan menyimpan alamat memori var hello, &hello adalah alamat memori yang diakses

	fmt.Println(&hello) // hasilnya adalah 0xc0420441c0 alamat memori
	fmt.Println(strPtr) // hasilnya adalah 0xc0420441c0 alamat memori

	fmt.Println(hello)	// hasil var hello sebelum dirubah
	change (hello) // rubah value var hello dengan value v
	fmt.Println(hello) // hasil var hello yang telah dirubah value v tetapi hasilnya akan sama dengan atas yaitu hello, hal ini terjadi var hello diatas hanya di copy kedalam func change, nah saat seperti ini dibutuhkan pointer

	changePtr(&hello) // pass by refer
	fmt.Println(hello) // pass by refere
}

// pass by value, valuenya hanya di copy saja bukan by reference
func change (v string) {
	v = v + " Golang"
}


// pass by reference 
func changePtr(v *string) {
	*v = *v + " Golang"
}