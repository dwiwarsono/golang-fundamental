// Membuat sebuah fungsi menjadi sebuah member atau properi sebuah struct (super method)

package main

import (
	"fmt"
)

func main() {

	// kita membuat pointer dari struct Person
	p := &Person{
		ID: 1,
		Name: "Dylan",
		Age: 2,
	}
	fmt.Println(p.GetID())
	fmt.Println(p.GetName())
	fmt.Println(p.GetAge())

	p.ChangeName("Abinaya") // ini untuk memanggil method ChangeName dan di isi dengan value baru yang akan merubah Name pada struct Person
	fmt.Println(p.GetName()) //
}

type Person struct{
	ID int
	Name string
	Age int
}


// buat sebuah method, fungsi ini adalah member dari struct Person ini

func (p *Person) ChangeName(newName string){ //method ini akan merubah argument dari p.Name
	p.Name = newName
}

func (p *Person) GetID() int{ //p *Person kalau di Golang Ini namanya recevier
	return p.ID
}

func (p *Person) GetName() string{
	return p.Name
}

func (p *Person) GetAge()  int{
	return p.Age
}