package main

import (
	"errors"
	"fmt"
)

// harus di definisikan tipe data fungsinya
func add(a int, b int) int {
	return a + b
}

// variadic function
// menerimah jumlah argumen tak terbatas

func sum(nums ...int) int {
	total := 0

	// range = loop khusus untuk koleksi (slice, array, map, string, dsb).
	for _, n := range nums {
		total += n
	}
	return total
}

// multiple return
func minMax(a, b int) (int, int) {
	if a > b {
		return b,a
	}
	return a,b
}

func divide(a, b int) (int, error) {
	if b == 0 {
		// gaya penulisan error handling string harus huruf kecil pada golang
		return 0, errors.New("pembagi tidak boleh nol")
	}

	return a/b, nil
}

func main() {
	hasil := add(3,5)
	fmt.Println("Hasil penjumlahan: ", hasil)

	fmt.Println(sum(1,2,3))

	kecil, besar := minMax(8,5)
	fmt.Println("Kecil: ", kecil, " Besar: ", besar)


	hasil, err := divide(10, 0)
	if err != nil {
		fmt.Println("Error: ", err)
	} else {
		fmt.Println("Hasil: ", hasil)
	}
}