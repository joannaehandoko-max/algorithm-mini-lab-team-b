package main

import "fmt"

func dijkstra() {
	graph := map[string]map[string]int{
		"A": {"B": 4, "C": 2},
		"B": {"A": 4, "D": 2},
		"C": {"A": 2, "D": 3},
		"D": {"B": 2, "C": 3},
	}

	dist := map[string]int{
		"A": 0,
		"B": 999,
		"C": 999,
		"D": 999,
	}

	visited := map[string]bool{}

	for i := 0; i < 4; i++ {

		// Cari node dengan jarak paling kecil
		current := ""
		smallest := 999

		for node, distance := range dist {
			if !visited[node] && distance < smallest {
				smallest = distance
				current = node
			}
		}

		// Tandai sudah dikunjungi
		visited[current] = true

		fmt.Println("Visit:", current)

		// Cek tetangga
		for next, cost := range graph[current] {

			newDistance := dist[current] + cost

			if newDistance < dist[next] {
				dist[next] = newDistance
			}
		}
	}

	fmt.Println("Jarak:", dist)
}

func main() {
	dijkstra()
}
