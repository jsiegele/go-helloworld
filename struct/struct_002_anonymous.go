package main

import (
	"fmt"
)

func main() {
	aDoctor := struct{name string}{name: "john doe"}
	fmt.Println(aDoctor)
	fmt.Println(aDoctor.name)
}
