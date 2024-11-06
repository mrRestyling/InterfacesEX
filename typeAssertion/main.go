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
func (d *Dog) Bark() string {
	return fmt.Sprintf("%s громко лает", d.Name)
}

func (c *Cat) Speak() string {
	log.Println("Speak (cat)")
	return "MEOW"
}

func (c *Cat) Purr() string {
	return fmt.Sprintf("%s мурлычет", c.Name)
}

func TA(a Animal) {
	if dog, ok := a.(*Dog); ok {
		// fmt.Printf("Type: %T , Value: %v\n", dog, dog)
		fmt.Println(dog.Bark())
	}

	if cat, ok := a.(*Cat); ok {
		// fmt.Printf("Type: %T , Value: %v\n", cat, cat)
		fmt.Println(cat.Purr())
	}
}

func main() {
	dog := &Dog{
		Name: "Lucky",
	}

	cat := &Cat{
		Name: "Murka",
	}

	TA(dog)
	TA(cat)
}
