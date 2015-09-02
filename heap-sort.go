package main

import "fmt"

func maxHeapify(slice []int, parentIdx int) {
	maxIdx := parentIdx
	leftChildIdx := 2*parentIdx + 1
	rightChildIdx := 2*parentIdx + 2

	// Check if a child node is larger than parent, skipping out-of-bounds
	// indexes when maxHeapify is called on a leaf.

	if leftChildIdx < len(slice) && slice[leftChildIdx] > slice[maxIdx] {
		maxIdx = leftChildIdx
	}

	if rightChildIdx < len(slice) && slice[rightChildIdx] > slice[maxIdx] {
		maxIdx = rightChildIdx
	}

	if maxIdx != parentIdx {
		slice[maxIdx], slice[parentIdx] = slice[parentIdx], slice[maxIdx]
		maxHeapify(slice, maxIdx)
	}
}

func heapSort(items []int) []int {
	length := len(items)

	// Create initial max heap.
	for i := length / 2; i > -1; i-- {
		maxHeapify(items, i)
	}

	// Sort.
	for i := length - 1; i >= 1; i-- {
		items[0], items[i] = items[i], items[0]
		length--
		// Call maxHeapify on progressively smaller slices, leaving more
		// and the original slice sorted each time.
		maxHeapify(items[:length], 0)
	}

	return items
}

func main() {
	stuff := []int{3, 1, 1, 3000, 9, 2, 8, 0, 42, 5}

	fmt.Println("before: ", stuff)

	heapSort(stuff)

	fmt.Println("after: ", stuff)
}
