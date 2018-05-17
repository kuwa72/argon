package argon

import (
	"fmt"
	"reflect"
)

// Parent is base
type Parent struct {
	ID int
}

// Child is embed parent object
type Child struct {
	Parent
	Name string
}

func (p *Parent) Generic() {
	fmt.Println("call generic\n")
	typeorig := reflect.TypeOf(*p)
	fmt.Println(typeorig)
	fmt.Println(typeorig.NumField())
	fmt.Println(typeorig.Field(0))
}
