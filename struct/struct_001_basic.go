package main

import (
	"fmt"
)

type Doctor struct {
	number int
	actorName string
	companions []string 
	episodes []string
}

func main() {
	aDoctor := Doctor{
		// die Dekleration kann auch ohne "key" erfolgen,
		// allerdings muss dann die Order eingehalten werden.
		number: 2,
		actorName: "john doe",
		episodes: []string {},
		companions: []string {
			"jane doe",
			"max muster",
		},
	}
	fmt.Println(aDoctor)
	fmt.Println(aDoctor.companions)
}
