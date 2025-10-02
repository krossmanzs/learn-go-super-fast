package main // nama package boleh bebas asal 1 folder

import (
	"fmt"
	"task-1/util" 
	"task-1/util/packageLain"
)// Import package lain

// cara klasik golang definisi variabel
var name string = "Cornelius Linux"

// tipe data tidak wajib ditulis
var age = 21

// multiple variabel
var x,y int = 10, 20

// konstanta
const Pi = 3.14
const Greeting = "Selamat datang!"

// main function
func main() {
	fmt.Println("Hello, world") // Perintah pertama

	/*
		bisa juga pagai shorthand
		tapi harus di dalam fungsi
	*/
	city := "Jakarta"

	fmt.Println(city)

	// Operator
	// a = 10
	// b = 3
	a, b := 10, 3

	penjumlahan := a + b
	pengurangan := a - b
	perkalian := a * b
	pembagian := a / b
	modulo := a % b


	// multiline printing menggunakan print format
	fmt.Printf(`
Penjumlahan: %d
Pengurangan: %d
Perkalian  : %d
Pembagian  : %d
Modulo	   : %d
	`, penjumlahan, pengurangan, perkalian, pembagian, modulo)

	util.PrintSum(3,5)
	lain.TestDariLain()

	var c int = 55
	var d float64 = 3.14
	var pi bool = true

	fmt.Println(c)
	fmt.Println(d)
	fmt.Println(pi)
	fmt.Println(len(name))
	fmt.Println(name[0])
	fmt.Println(name[:2])
	fmt.Println(name[2:])

	lain.PrintArray()
	lain.PrintMap()
	lain.PrintPointer()
	lain.PrintSlice()
	lain.PrintStruct()
}