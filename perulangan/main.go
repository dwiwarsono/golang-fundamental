package main

import (
	"fmt"
)

// DI GOLANG TIDAK ADA WHILE

func main(){

	//CONTOH 1
	// for i := 0; i <= 10; i++ {
	// 	fmt.Printf("Perulangan ke %d\n", i)
	// }

	//CONTOH 2
	// x := 10
	// var y int
	
	// for y < x {
	// 	y++
	// 	fmt.Println(y)
	// } 

	//CONTOH 3
	// x := 1

	// for {
	// 	fmt.Println("Hello go")
	// 	x++
	// if x == 10 { // INI BERFUNGSI  AGAR TIDAK INFINITE LOOP (TERUS MELAKUKAN PENGULANGAN)
	// 	break
	// }
	// }

	//CONTOH 4
	x := 1

	for {
		x++

		if x == 5 { // INI BERFUNGSI NILAI 5 TIDAK AKAN TERCETAK
			continue // KETIKA NILAI X BERNILAI 5 KITA LEWATI
		}

		fmt.Println("Hello go %d\n", x)
		
		if x == 10 {
			break
		}
	}
}