// Experiment 1: Insertion Sort
// Run with: go run experiments/insertion_sort/main.go
package main

import "fmt"

func insertionSort(arr []int) {
	for i := 1; i < len(arr); i++ {
		key := arr[i] // the number we pick up this round
		j := i - 1
		fmt.Printf("Round %d: pick up %d\n", i, key)

		// shift bigger numbers one step to the right
		for j >= 0 && arr[j] > key {
			fmt.Printf("  compare %d > %d -> shift %d to the right\n", arr[j], key, arr[j])
			arr[j+1] = arr[j]
			j--
		}
		if j >= 0 {
			fmt.Printf("  compare %d > %d -> no, stop\n", arr[j], key)
		}

		arr[j+1] = key
		fmt.Printf("  insert %d at position %d -> %v\n\n", key, j+1, arr)
	}
}

func runExperiment(name string, input []int) {
	fmt.Println("==============================")
	fmt.Println(name)
	fmt.Println("Start:", input)
	fmt.Println("==============================")

	// copy so the original input stays unchanged
	data := make([]int, len(input))
	copy(data, input)

	insertionSort(data)
	fmt.Println("Final result:", data)
	fmt.Println()
}

func main() {
	// Step 1: PREDICT the final result before running!
	runExperiment("Original input", []int{5, 2, 8, 1, 4})

	// Step 4: CHANGE the input, then Step 5: EXPLAIN what changed
	runExperiment("Changed input", []int{9, 3, 7, 2, 1})

	// Extra test ideas (uncomment to try):
	// runExperiment("Already sorted", []int{1, 2, 3, 4, 5})
	// runExperiment("Duplicates", []int{4, 2, 4, 1, 2})
	// runExperiment("Single element", []int{7})
}
