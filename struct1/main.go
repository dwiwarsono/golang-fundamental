// STRUCT sebagai argument atau parameter di sebuah function

package main

import (
	"fmt"
)

// FUNC MAIN INI ADALAH FUNC UTAMA YANG HARUS ADA UNTUK MENJALANKAN PROGRAM
func main() {
	p:= Person{
		ID : 1,
		Name : "Dwi Warsono",
		Age : 31,
	}

	//JAlankan fungsi printPerson
	printPerson(p)
}

type Person struct{
	ID int
	Name string
	Age int
}

func printPerson(p Person) {
	fmt.Println("ID = ", p.ID)
	fmt.Println("Name = ", p.Name)
	fmt.Println("Age = ", p.Age)
}