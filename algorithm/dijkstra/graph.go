package main

import (
	"fmt"
)

type edge struct {
	node   string
	weight int
}

type graph struct {
	nodes map[string][]edge
}

func newGraph() *graph {
	return &graph{nodes: make(map[string][]edge)}
}

func (g *graph) addEdge(origin, destiny string, weight int) {
	g.nodes[origin] = append(g.nodes[origin], edge{node: destiny, weight: weight})
	g.nodes[destiny] = append(g.nodes[destiny], edge{node: origin, weight: weight})
}

func (g *graph) getEdges(node string) []edge {
	return g.nodes[node]
}

func (g *graph) getPath(origin, destiny string) (int, []string) {
	h := newHeap()
	h.push(path{value: 0, nodes: []string{origin}})
	visited := make(map[string]bool)

	for len(*h.values) > 0 {
		// Find the nearest yet to visit node
		p := h.pop()
		fmt.Println("\nheap pop: ", p)
		node := p.nodes[len(p.nodes)-1]
		fmt.Println("node:", node, "->", p.value , p.nodes)
		if visited[node] {
			fmt.Println("\nvisited1: ", visited)
			continue
		}

		if node == destiny {
			fmt.Println("Destination found: ", p.value, p.nodes)
			return p.value, p.nodes
		}

		for _, e := range g.getEdges(node) {
			if !visited[e.node] {
				// We calculate the total spent so far plus the cost and the path of getting here
				// You can concatenate two slices using the three dots notation
				fmt.Println("unvisited node:", e.node , ": Push path -> ", p.value + e.weight, append([]string{}, append(p.nodes, e.node)...))
				h.push(path{value: p.value + e.weight, nodes: append([]string{}, append(p.nodes, e.node)...)})
			}else{
				fmt.Println("  visited node:", e.node)
			}
		}

		visited[node] = true
		fmt.Println("\nvisited2: ", visited)
	}

	return 0, nil
}