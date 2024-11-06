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

func main() {

	var animal Animal

	// 1
	fmt.Printf("val: %v , type: %T \n", animal, animal)

	// 2
	dog := &Dog{
		Name: "Lucky",
	}
	animal = dog

	fmt.Println(animal.Speak())

	// dog.Name = "Plaki"

}

func (d *Dog) Speak() string {
	log.Println("Собака гавкает")
	return "WOOF"

}
