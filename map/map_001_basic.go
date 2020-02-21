package main

import (
	"fmt"
)

func main() {
	statePopulation := make(map[string]int)
	statePopulation = map[string]int{
		"tirol":    1212,
		"wien":     12131,
		"salzburg": 1122,
		"noe":      4433,
	}
	fmt.Println(statePopulation)
}
