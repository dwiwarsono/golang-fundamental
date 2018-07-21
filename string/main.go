// MANIPULASI STRING

package main

import (
	"fmt"
	"strings"
	"strconv" // untuk mengconversi int ke string atau sebaliknya ataupun tipe data lain
)

func main (){

	var myString = "Hello Golang"
	var myStringTwo = "10"

	fmt.Println(myString)
	fmt.Println(len(myString)) // FUngsi len untuk mengetahui berapa panjang value (Di bahasa pemograman lain dikenal dengan length)

	// for i := 0; i < len(myString); i++{
	// 	fmt.Println(myString[i])
	// }

	myStringUpper := strings.ToUpper(myString) //ToUpper berfungsi merubah huruf besar semua
	fmt.Println(myStringUpper)

	myStringLower := strings.ToLower(myString)
	fmt.Println(myStringLower)

	resultContains := strings.Contains(myString, "Go") // Fungsi Contains untuk mengecek apakah kata yang dimasukan ada di dalam variabel, dan menghasilkan nilai boolean
	fmt.Println(resultContains)

	resultSplit := strings.Split(myString, " ") //return dari split adalah dynamyc array
	for _, v := range resultSplit{
		fmt.Println(v)
	}

	// Cara mengkonversi string menjadi integer
	resultConvInt, err := strconv.Atoi(myStringTwo) // merubah nilai string pada var myStringTwo menjadi int
	if err != nil{
		fmt.Println(err)
	}
	fmt.Println(resultConvInt * 5)

	// Cara mengkonversi int ke string
	resultConvStr := strconv.Itoa(10)
	fmt.Println(resultConvStr + " Dylan")
}