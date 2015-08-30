package main

import "fmt"

func insertionSort(items []int) []int {
	var j int

	for i := range items {
		current := items[i]

		for j = i - 1; j >= 0 && items[j] > current; j-- {
			items[j+1] = items[j]
		}

		items[j+1] = current
	}

	return items
}

func main() {
	stuff := []int{4, 1, 38, 1, 14, 900, 424, 2}

	fmt.Println("before: ", stuff)

	insertionSort(stuff)

	fmt.Println("after: ", stuff)
}
