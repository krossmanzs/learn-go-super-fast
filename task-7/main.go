package main

import (
	"fmt"
	"math"
	"task-7/mymath"
	"task-7/cli"
	"time"
)

func main() {
	fmt.Println("Hello, Go!")
	fmt.Println("Akar kuadrat 16 =", math.Sqrt(16))
	fmt.Println("Waktu sekarang: ", time.Now())

	result := mymath.Add(3,5)
	fmt.Println("Hasil penjumlahan dari package mymath: ", result, "\n\n")

	cli.RunCli()
}