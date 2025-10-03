package main

import (
	"fmt"
	"task-6/goroutine"
)

// definisi interface
type Greeter interface {
	Greet()
}

// struct A
type User struct {
	Name string
}

// struct B
type Admin struct {
	Name string
}

// implementasi method Greet() untuk user
func (u User) Greet() {
	fmt.Println("Hello, saya user: ", u.Name)
}

// implementasi method Greet() untuk admin
func (a Admin) Greet() {
	fmt.Println("halo, saya admin: ", a.Name)
}

func main() {
	var g Greeter
	g = User{"Linux"}
	g.Greet()
	g = User{"Budi"}
	fmt.Println(g)

	goroutine.RunGoroutine()
	goroutine.RunChannel()
	goroutine.RunAnotherChannel1()
	goroutine.RunAnotherChannel2()
}