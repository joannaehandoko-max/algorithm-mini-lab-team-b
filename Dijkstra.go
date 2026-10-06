// Experiment 3: Dijkstra's Algorithm
// Run with: go run experiments/dijkstra/main.go
package main

import (
	"fmt"
	"math"
	"strings"
)

type Graph map[string]map[string]int

// addEdge adds a two-way road with a cost
func (g Graph) addEdge(a, b string, cost int) {
	if g[a] == nil {
		g[a] = map[string]int{}
	}
	if g[b] == nil {
		g[b] = map[string]int{}
	}
	g[a][b] = cost
	g[b][a] = cost
}

func dijkstra(g Graph, nodes []string, start, end string) {
	const inf = math.MaxInt32
	dist := map[string]int{}
	prev := map[string]string{}
	visited := map[string]bool{}

	for _, n := range nodes {
		dist[n] = inf
	}
	dist[start] = 0

	for {
		// pick the unvisited node with the smallest known cost
		current := ""
		best := inf
		for _, n := range nodes {
			if !visited[n] && dist[n] < best {
				best = dist[n]
				current = n
			}
		}
		if current == "" {
			break
		}

		visited[current] = true
		fmt.Printf("Visit %s (cheapest cost so far: %d)\n", current, best)

		// check each neighbor: is going through current cheaper?
		for _, nb := range nodes {
			cost, connected := g[current][nb]
			if !connected || visited[nb] {
				continue
			}
			newDist := best + cost
			if newDist < dist[nb] {
				fmt.Printf("  update %s: new cost %d via %s\n", nb, newDist, current)
				dist[nb] = newDist
				prev[nb] = current
			} else {
				fmt.Printf("  keep %s: %d is not better than %d\n", nb, newDist, dist[nb])
			}
		}
	}

	// rebuild the path from end back to start
	path := []string{}
	for at := end; at != ""; at = prev[at] {
		path = append([]string{at}, path...)
	}
	fmt.Printf("\nShortest path: %s\n", strings.Join(path, " -> "))
	fmt.Printf("Total cost: %d\n\n", dist[end])
}

func runExperiment(name string, cdCost int) {
	fmt.Println("==============================")
	fmt.Printf("%s (C-D = %d)\n", name, cdCost)
	fmt.Println("==============================")

	g := Graph{}
	g.addEdge("A", "B", 4)
	g.addEdge("A", "C", 2)
	g.addEdge("B", "D", 2)
	g.addEdge("C", "D", cdCost)

	dijkstra(g, []string{"A", "B", "C", "D"}, "A", "D")
}

func main() {
	// Step 1: PREDICT: A -> ______ -> D
	runExperiment("Original network", 3)

	// Step 4: CHANGE one input (C-D becomes more expensive)
	runExperiment("Changed network", 5)
}
