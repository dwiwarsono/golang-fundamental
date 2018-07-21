//Slice atau dynamic array
// Slice index nya dapat bertambah karena index array nya tidak di fenisikan panjang nya

package main

import (
	"fmt"
)

func main() {

	mySlice := [] int{10, 20, 30, 40, 50}// untuk slice kita tidak mendifinisikan panjang array nya []
	fmt.Println(mySlice[0])

	//mengakses dengan loop
	for i, v := range mySlice{ // range digunakan untuk array atau juga slice
		fmt.Println(i, v)
	}

	//CONTOH LAGI
	mySliceStr := [] string{"Alex", "Ben", "Den", "Dylan", "Dwi"}

	//CONTOH MERUBAH VALUE KE INDEX ARRAY
	mySliceStr[4] = "Widya"
	//CONTOH MENAMBAH VALUE KE INDEX ARRAY TREAKHIR
	mySliceStr = append(mySliceStr, "Warsono")

	for _, v := range mySliceStr{ //index nya kita tidak perlukan ya, kita ganti dengan underscore _
		fmt.Println(v)
	}

}

