package main

import "fmt"

func mergeSorted(a []int, b []int) []int {
	var result []int
	lenA := len(a)
	lenB := len(b)
	idxA := 0
	idxB := 0

	for idxA < lenA && idxB < lenB {
		if a[idxA] < b[idxB] {
			result = append(result, a[idxA])
			idxA++
		} else {
			result = append(result, b[idxB])
			idxB++
		}
	}

	return append(result, append(a[idxA:], b[idxB:]...)...)
}

func mergeSort(items []int) []int {
	length := len(items)

	if length < 2 {
		return items
	}

	middle := length / 2

	return mergeSorted(mergeSort(items[:middle]), mergeSort(items[middle:]))
}

func main() {
	stuff := []int{4, 2, 9, 200, 1, 1, 38, 1, 4, 7, 23, 8}

	fmt.Println("before: ", stuff)

	stuff = mergeSort(stuff)

	fmt.Println("after: ", stuff)
}
