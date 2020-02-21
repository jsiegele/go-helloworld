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
	//order in maps werden nicht garantiert!
	fmt.Println(statePopulation)
	statePopulation["ooe"] = 4433
	fmt.Println(statePopulation["ooe"])
	fmt.Println(statePopulation)
	delete(statePopulation, "ooe")
	// komma ok syntax - ist unser key wirklich in der map
	pop, ok := statePopulation["ooe"]
	fmt.Println(pop, ok)
	_, ok = statePopulation["ooe"]
	fmt.Println(ok)
	fmt.Println(len(statePopulation))
}
