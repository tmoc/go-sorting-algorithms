package main

import "fmt"

func bubbleSort(items []int) []int {
	stop := len(items) - 1

	for range items {

		for i := 0; i < stop; i++ {
			if items[i] > items[i+1] {
				items[i], items[i+1] = items[i+1], items[i]
			}
		}
		// The last item is sorted, doesn't need to be checked again.
		stop--
	}

	return items
}

func main() {
	stuff := []int{7, 42, 3, 98, 2, 111, 24}

	fmt.Println("before: ", stuff)

	bubbleSort(stuff)

	fmt.Println("after: ", stuff)
}
