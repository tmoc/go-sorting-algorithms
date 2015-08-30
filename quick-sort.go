package main

import "fmt"

func quickSort(items []int) []int {
	length := len(items)

	if length < 2 {
		return items
	}

	left, right := 0, length-1
	pivot := (left + right) / 2

	items[right], items[pivot] = items[pivot], items[right]

	for i := range items {
		if items[i] < items[right] {
			items[left], items[i] = items[i], items[left]
			left++
		}
	}

	items[left], items[right] = items[right], items[left]

	quickSort(items[:left])
	quickSort(items[left+1:])

	return items
}

func main() {
	stuff := []int{200, 42, 1, 1, 15, 1, 32, 4, 3999, 10}

	fmt.Println("before: ", stuff)

	quickSort(stuff)

	fmt.Println("after: ", stuff)
}
