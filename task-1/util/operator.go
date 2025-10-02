package util

import "fmt"

// huruf kapital di awal menandakan fungsinya di eksport/public
func PrintSum(a, b int) {
	sum(a,b)
}

// fungsi ini tidak di eksport
func sum(a int, b int) {
	sum := a + b 
	fmt.Printf("Hasil %d + %d = %d\n", a, b, sum)
}