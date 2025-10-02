package lain

import "fmt"

// array
var arr [3]int = [3]int{1,2,3}

func PrintArray() {
	fmt.Println(arr)
}

// slice
func PrintSlice() {
	nums := []int{1,2,3}
	nums = append(nums, 4, 4)
	fmt.Println(nums)
}

// map
func PrintMap() {
	m := map[string]int {
		"linux": 21,
		"budi": 22,
	}

	fmt.Println(m)
}

// struct
type User struct {
	ID		int
	Name 	string
}

func PrintStruct() {
	u1 := User{1, "Linux"}
	u2 := User{ID: 2, Name: "Budi"}
	fmt.Println(u1)
	fmt.Println(u2)
}

// pointer
func PrintPointer() {
	x := 10
	p := &x					// pointer ke x
	fmt.Println(*p)	// akses nilai x lewat pointer
}

func TestDariLain() {
	fmt.Println("test dari lain")
}