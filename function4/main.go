package main

import (
	"fmt"
)



// FUNGSI SEBAGAI ARGUMENT FUNGSI LAIN ATAU FUNGSI SEBAGAI PARAMETER
func main (){

	f := func (v string) bool{
		return v == "Golang"
	}
	result := match("GO", f)

	fmt.Println(result)

}

func match(v string, f func(string) bool) bool{ // Anonimus function (FUNGSI TANPA NAMA)
	if f(v) {
		return true
	}

	return false
}