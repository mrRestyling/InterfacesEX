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

func MakeSound(a Animal) {
	fmt.Println((a.Speak()))
}

func (d *Dog) Bark() string {
	return fmt.Sprintf("%s громко лает", d.Name)
}

func (c *Cat) Purr() string {
	return fmt.Sprintf("%s мурлычет", c.Name)
}

func processTypeAssertion(a Animal) {
	if dog, ok := a.(*Dog); ok {
		// fmt.Printf("Type: %T , Value: %v\n", dog, dog)
		fmt.Println(dog.Bark())
	}

	if cat, ok := a.(*Cat); ok {
		// fmt.Printf("Type: %T , Value: %v\n", cat, cat)
		fmt.Println(cat.Purr())
	}
}

func processTypeSwitch(a Animal) {

	switch t := a.(type) {
	case *Dog:
		fmt.Println(t)
	case *Cat:
		fmt.Println(t)
	default:
		fmt.Println(t)
	}

}

func main() {

	var animal Animal

	cat := &Cat{
		Name: "Murka",
	}

	dog := &Dog{
		Name: "Lucky",
	}

	fmt.Println("-- 1 --")
	processTypeAssertion(cat)
	processTypeAssertion(dog)

	fmt.Println("-- 2 --")
	processTypeSwitch(cat)
	processTypeSwitch(dog)

	fmt.Println("-- 3 --")

	animal = cat
	animal.Speak()

	animal = dog
	animal.Speak()

}
