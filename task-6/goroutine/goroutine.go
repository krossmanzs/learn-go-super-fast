package goroutine

import "fmt"

func printMessage(msg string) {
	for i := 0; i < 5; i++ {
		fmt.Println(msg, i)
	}
}

func RunGoroutine() {
	go printMessage("Goroutine 1")
	go printMessage("Goroutine 2")

	// supaya program ga langsung berhenti
	fmt.Scanln()
}

