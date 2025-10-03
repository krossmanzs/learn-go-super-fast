package goroutine

import "fmt"

/*
	Channel = jalur komunikasi antar goroutine.
	Bayangin channel itu kayak pipa: satu goroutine bisa
	ngirim data ke pipa, goroutine lain bisa nerima data dari pipa.

	tanpa channel, kita harus pake Sleep biar goroutine sempat jalan. 
	Tapi dengan channel, program sinkron menunggu sampai ada data dikirim.
*/

func worker(ch chan string) {
	ch <- "pekerjaan selesai"
}

func RunChannel() {
	ch := make(chan string)
	go worker(ch)

	msg := <-ch // menerima pesan dari channel
	fmt.Println(msg)
}

func RunAnotherChannel1() {
	ch := make(chan string)

	go func() {
		ch <- "pesan dari goroutine"
	}()

	// penerima channel
	msg := <-ch
	fmt.Println(msg)
}

func RunAnotherChannel2() {
	ch := make(chan int)

	// pengirim
	go func() {
		for i := 1; i < 5; i++ {
			ch <- i
		}
		close(ch) // tutup channel setelah selesai
	}()

	// penerima
	fmt.Println("Pesan yang diterima: ")

	// terima angka dari channel
	for num := range ch {
		fmt.Println("Terima: ", num)
	}

	fmt.Println("Pesan yang diterima2: ")
	for num := range ch {
		fmt.Println("Terima: ", num)
	}
}