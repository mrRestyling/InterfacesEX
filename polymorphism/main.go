package main

import (
	"fmt"
	"log"
)

type Animal interface {
	Speak() string
}

type Dog struct {
	Name string
}

type Cat struct {
	Name string
}

func (d *Dog) Speak() string {
	log.Println("Speak (dog)")
	return "WOOF"

}

func (c *Cat) Speak() string {
	log.Println("Speak (cat)")
	return "MEOW"
}

func MakeSpeak(a Animal) {
	fmt.Println(a.Speak())
}

func main() {

	dog := &Dog{
		Name: "Lucky",
	}

	cat := &Cat{
		Name: "Murka",
	}

	MakeSpeak(dog)
	MakeSpeak(cat)

}
