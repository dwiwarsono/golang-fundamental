package main

import (
	"fmt"
)

func main () {
	//CONTOH ARRAY TIPE DATA STRING
	var arr [5] int // [5] disini panjang array atau len (length)
	arr[0] = 1
	arr[1] = 2
	arr[2] = 3
	arr[3] = 4
	arr[4] = 5
	
	for i := 0; i < len(arr); i++ {
		fmt.Println(arr[i])
	}

	fmt.Println(arr[3]) // ini akan memanggil isi array index ke 3

	//CONTOH LAIN DENGAN ARRAY TIPE DATA STRING
	arrStr := [5] string{"Alex", "Ben", "Den", "Dylan", "Dwi"}
	fmt.Println(arrStr[4])

	// ARRAY MULTI DIMENSI (LEBIH DARI 1 ARRAY)
	arrMulti := [2][3] int{{1, 2, 3}, {4, 5, 6}}
	fmt.Println(arrMulti[0][2]) // ini akan mencetak pada array 0 value ke 2 yaitu 3
}