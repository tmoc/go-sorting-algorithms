package main

import "fmt"

func maxHeapifyAtIndex(slice []int, parentIdx int) {
	length := len(slice)
	maxIdx := parentIdx
	leftChildIdx := 2*parentIdx + 1
	rightChildIdx := 2*parentIdx + 2

	if leftChildIdx < length && slice[leftChildIdx] > slice[maxIdx] {
		maxIdx = leftChildIdx
	}

	if rightChildIdx < length && slice[rightChildIdx] > slice[maxIdx] {
		maxIdx = rightChildIdx
	}

	if maxIdx != parentIdx {
		slice[maxIdx], slice[parentIdx] = slice[parentIdx], slice[maxIdx]
		maxHeapifyAtIndex(slice, maxIdx)
	}
}

func maxHeapify(slice []int) []int {
	for i := len(slice) / 2; i > -1; i-- {
		maxHeapifyAtIndex(slice, i)
	}
	return slice
}

func heapSort(items []int) []int {
	maxHeapify(items)
	sliceIdx := len(items)

	for i := sliceIdx - 1; i >= 1; i-- {
		items[0], items[i] = items[i], items[0]
		sliceIdx--
		maxHeapifyAtIndex(items[:sliceIdx], 0)
	}

	return items
}

func main() {
	stuff := []int{3, 1, 1, 3000, 9, 2, 8, 0, 42, 5}

	fmt.Println("before: ", stuff)

	heapSort(stuff)

	fmt.Println("after: ", stuff)
}
