package main

import "fmt"

type User struct {
	Username	string
	Email		string
}

func main() {
	// array
	arr := [3]int{1, 2, 3}
	fmt.Println("Array: ", arr)
	fmt.Println("Index ke-1: ", arr[1])

	// Slice
	nums := []int{1, 2, 3}
	nums = append(nums, 4, 5)
	fmt.Println("Slice: ", nums)

	// Map
	users := []User{
		{"Linux", "linux@mail.com"},
		{"Budi", "budi@mail.com"},
	}

	for _, u := range users {
		fmt.Println("User: ", u.Username, " Email: ", u.Email)
	}
}