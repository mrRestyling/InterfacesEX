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

func (d *Dog) Speak() string {
	log.Println("Speak (dog)")
	return "WOOF"
}
func (d Dog) Bark() string {
	return fmt.Sprintf("%s громко лает!", d.Name)
}

type Cat struct {
	Name string
}

func (c *Cat) Speak() string {
	log.Println("Speak (cat)")
	return "MEOW"
}

func (c Cat) Purr() string {
	return fmt.Sprintf("%s мурлычет.", c.Name)
}

func main() {

	dog := &Dog{"Lucky"}
	cat := &Cat{"Murka"}

	TS(dog)
	TS(cat)

}

func TS(a Animal) {

	switch t := a.(type) {
	case *Dog:
		fmt.Println(t)
	case *Cat:
		fmt.Println(t)
	default:
		fmt.Println("type not found")
	}

}
