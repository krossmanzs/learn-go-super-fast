package main

import "fmt"

func main() {
	umur := 21

	// kondisi if-else
	if umur >= 18 {
		fmt.Println("Dewasa")
	} else {
		fmt.Println("Anak-anak")
	}

	var a, b float64 = 2, 5

	sum := b / a

	// kondisi if-else if-else
	if sum == 5 {
		fmt.Println("penjumlahan = 5 ", sum)
	} else if sum < 5 {
		fmt.Println("penjumlahan < 5 ", sum)
	} else {
		fmt.Println("penjumlahan > 5 ", sum)
	}

	// switch
	hari := 3

	switch hari {
		case 1:
			fmt.Println("Senin")
		case 2:
			fmt.Println("Selasa")
		case 3:
			fmt.Println("Rabu")
		case 4:
			fmt.Println("Kamis")
		case 5:
			fmt.Println("Jumat")
		case 6:
			fmt.Println("Sabtu")
		case 7:
			fmt.Println("Minggu")
		default:
			fmt.Println("Nomor hari tidak valid")
	}

	// for loop
	fmt.Println("\nFor Loop:")
	for i := 1; i <= 5; i++ {
		fmt.Println(i)
	}

	for i := 1; i < 10; i++ {
		if i == 5 {
			break
		}

		if i % 2 == 0 {
			continue
		}

		fmt.Println(i)
	}

	// while style
	/*
		while (i < 5) {
			...
		}
	*/
	i := 0
	for i < 5 {
		fmt.Println(i)
		i++
	}

	// do while style
	/*
		do {
			...
		} while (i < 5)
	*/
	j := 0
	for {
		fmt.Println(j)
		j++
		if j >= 5 {
			break;
		}
	}

	// infinite loop
	for {
		fmt.Println("jalan terus...")
		break // jika tidak di break akan jalan terus
	}
}