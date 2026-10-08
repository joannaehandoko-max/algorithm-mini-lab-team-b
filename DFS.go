
package main

import (
	"fmt"
	"strings"
)

// dfs goes as deep as possible before backtracking (uses recursion = the call stack)
func dfs(graph map[string][]string, node string, visited map[string]bool, order *[]string, depth int) {
	indent := strings.Repeat("  ", depth)
	visited[node] = true
	*order = append(*order, node)
	fmt.Printf("%sVisit %s\n", indent, node)

	for _, neighbor := range graph[node] {
		if visited[neighbor] {
			fmt.Printf("%s  skip %s (already visited)\n", indent, neighbor)
			continue
		}
		dfs(graph, neighbor, visited, order, depth+1)
	}
	fmt.Printf("%sBacktrack from %s\n", indent, node)
}

func runExperiment(name string, graph map[string][]string, start string) {
	fmt.Println("==============================")
	fmt.Println(name)
	fmt.Println("==============================")

	visited := map[string]bool{}
	order := []string{}
	dfs(graph, start, visited, &order, 0)

	fmt.Println("\nTraversal:", strings.Join(order, " -> "))
	fmt.Println()
}

func main() {
	// Original graph:   A -- B, A -- C, B -- D, B -- E
	graph1 := map[string][]string{
		"A": {"B", "C"},
		"B": {"A", "D", "E"},
		"C": {"A"},
		"D": {"B"},
		"E": {"B"},
	}

	// Step 1: PREDICT: A -> ___ -> ___ -> ___ -> ___
	runExperiment("Original graph", graph1, "A")

	// Step 4: CHANGE the graph: add F under C
	graph2 := map[string][]string{
		"A": {"B", "C"},
		"B": {"A", "D", "E"},
		"C": {"A", "F"},
		"D": {"B"},
		"E": {"B"},
		"F": {"C"},
	}
	runExperiment("Changed graph (added F under C)", graph2, "A")

	// Extra test idea: a graph with a cycle (the visited map prevents an infinite loop)
	// graph3 := map[string][]string{
	// 	"A": {"B", "C"},
	// 	"B": {"A", "C"},
	// 	"C": {"A", "B"},
	// }
	// runExperiment("Graph with a cycle", graph3, "A")
}
