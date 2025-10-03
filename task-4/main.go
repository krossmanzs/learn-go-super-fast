package main

import (
	"fmt"
	"task-4/user"
)

func main() {
	// struct
	u1 := user.User{ID: 1, Name: "Linux"}
	u2 := user.User{ID: 2, Name: "Budi"}

	fmt.Println(u1)
	fmt.Println(u2.Name)

	u1.SayHello()
	u2.SayHello()

	// pointer
	x := 10
	p := &x					// pointer ke x
	fmt.Println("Nilai x: ", x)
	fmt.Println("Nilai dari pointer p: ", *p)	// akses nilai x lewat pointer
	fmt.Println("Alamat p: ", p)

	*p = 20					// ubah nilai x lewat pointer
	fmt.Println(x)	// x sekarnag menjadi 20
}

