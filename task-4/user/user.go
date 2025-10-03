package user

import "fmt"

type User struct {
	ID   int
	Name string
}

// method dengan receiver
func (u User) SayHello() {
	fmt.Println("Hello,", u.Name)
}