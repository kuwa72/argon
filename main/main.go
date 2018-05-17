package main

import (
	"fmt"

	"github.com/kuwa72/argon"
)

func main() {
	fmt.Println("Test get inherited struct name from parent method")
	fmt.Println("Call method from parent object.")
	parent := &argon.Parent{}
	parent.Generic()

	fmt.Println("Call method from inherited object.")
	child := &argon.Child{}
	child.Generic()
}
