package main

import (
	"fmt"
)

const s string = "constant"

func main() {
	fmt.Println("Hello world")

	wtf := func(i interface{}) {
		switch t := i.(type) {
		case bool:
			fmt.Println("It's bool")
		case int:
			fmt.Println("It's integer")
		case string:
			fmt.Println("It's string")
		default:
			fmt.Println("de fuk %T\n", t)
		}
	}

	wtf("abc")
}
