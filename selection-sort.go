package main

import "fmt"

func selectionSort(items []int) []int {
	length := len(items)

	for i := 0; i < length; i++ {

		for j := i + 1; j < length; j++ {
			if items[j] < items[i] {
				items[i], items[j] = items[j], items[i]
			}
		}

	}

	return items
}

func main() {
	stuff := []int{3, 42, 1, 400, 200, 84, 1, 1, 14}

	fmt.Println("before: ", stuff)

	selectionSort(stuff)

	fmt.Println("after: ", stuff)
}
