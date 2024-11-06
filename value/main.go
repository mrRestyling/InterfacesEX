package main

import (
	"fmt"
)

type I interface {
	M()
}

type T struct {
	S string
}

func (t *T) M() {
	fmt.Println(t.S)
}

// Интерфейс J с встроенным интерфейсом I
type J interface {
	I
}

func main() {

	var i J

	i = &T{"Hello"}
	fmt.Printf("Value %v, type %T)\n", i, i)
	i.M()
}
