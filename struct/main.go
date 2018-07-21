// STRUCT kita bisa meiliki 1 structur dengan tipe yang berbeda

package main

import (
	"fmt"
)

// INI UNTUK MEMBUAT STRUCT
type Vector struct{
	X int
	Y int
}

type Player struct{
	ID int
	Name string
}

func main () {

	var v Vector
	v.X = 10
	v.Y = 12	

	fmt.Println(v)

	fmt.Println("X = ", v.X)
	fmt.Println("Y = ", v.Y)

	player1 := Player{ ID: 1, Name: "Widya"}

	fmt.Println(player1.ID)
	fmt.Println(player1.Name)

}