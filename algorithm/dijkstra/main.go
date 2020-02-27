package main

import (
	"fmt"
)

func main() {
	fmt.Println("Dijkstra")
	// Example
	graph := newGraph()
	graph.addEdge("a", "c", 3)
	graph.addEdge("a", "b", 7)
	graph.addEdge("c", "b", 1)
	graph.addEdge("c", "d", 2)
	graph.addEdge("b", "d", 2)
	graph.addEdge("b", "e", 6)
	graph.addEdge("d", "e", 4)
	fmt.Println(graph.getPath("a", "e"))
}