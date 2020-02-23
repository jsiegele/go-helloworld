package main

import (
	"fmt"
)

func main() {
	aDoctor := struct{name string}{name: "john doe"}
	anotherDoctor := aDoctor
	anotherDoctor.name = "jane doe"
	fmt.Println(aDoctor)
	fmt.Println(anotherDoctor)
}
