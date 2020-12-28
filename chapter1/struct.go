package main

import (
	"fmt"
)

// Program : intro to struct and interface in go

type Person struct {
	name string
	Age int
}

type Dog struct {}

//define a contract to register

type Friend interface {
	SayHello()
}

func Greet(f Friend) {
	// defines how to use
	f.SayHello();
}

// this got register under Friend
func (p *Person) SayHello() {
	fmt.Println("Hello", p.name);
}

// this method also got register under friend
func (d *Dog) SayHello() {
	fmt.Println("Woof Woof");
}

func main() {
	var guy = new(Person);
	guy.name = "Saurabh";
	Greet(guy);
	var rusty = new(Dog)
	Greet(rusty);
}