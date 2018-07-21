package main

import (
	"fmt"
)

func main() {
	x := 100
	y := 100


	// EXPRESSION SWITCH
	switch x {
	case 60: //APAKAH X NILAINYA 60
		fmt.Println("NILAI = C")
	case 80:
		fmt.Println("NILAI = B")
	case 100:
		fmt.Println("NILAI = A")
	default:
		fmt.Println("NILAI TIDAK DIKETAHUI !!")
	}


	switch {
	case y == 60:
		fmt.Println("NILAI = C")
	case y == 80:
		fmt.Println("NILAI = B")
	case y == 100:
		fmt.Println("NILAI = A")
	default:
		fmt.Println("NILAI TIDAK DIKETAHUI !!")
	}


	//TYPE SWITCH
	var z interface{} //AWALNYA VAR Y TIDAK DIKETAHUI TIPE DATANYA
	z = 5

	switch z.(type) { // CEK TYPE DATA
	case int: // APAKAH TIPE DATA Y ADALAH INT
		fmt.Println("TIPE DATA Z ADALAH INTEGER")
	case string:  // APAKAH TIPE DATA Y ADALAH STRING
		fmt.Println("TIPE DATA Z ADALAH STRING")
	case float64:  // APAKAH TIPE DATA Y ADALAH FLOAT64
		fmt.Println("TIPE DATA Z ADALAH FLOAT 64")
	default:
		fmt.Println("TIPE DATA TIDAK DIKETAHUI")
	}
}