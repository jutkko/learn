package minheap

import (
	"fmt"
	"os"
)

type MinHeap struct {
	array []int
}

func (m *MinHeap) Push(e int) {
	m.array = append(m.array, e)

	up(m.array, m.Size()-1)
}

func (m *MinHeap) Pop() int {
	if len(m.array) == 0 {
		fmt.Print("No elements in the heap")
		os.Exit(1)
	}

	result := m.array[0]
	m.array[0] = m.array[m.Size()-1]
	m.array = m.array[1:]
	down(m.array, 0)

	return result
}

func (m *MinHeap) Size() int {
	return len(m.array)
}

// put the element at the index up, to the correct place
func up(array []int, index int) {
	for {
		if index == 0 {
			break
		}

		parentIndex := (index - 1) / 2
		// float the current element up
		if array[parentIndex] > array[index] {
			array[parentIndex], array[index] = array[index], array[parentIndex]
		}

		index = parentIndex
	}
}

// heapify the new array
func down(array []int, index int) {
	for {
		var smallerChild int
		lChild := index*2 + 1
		rChild := index*2 + 2

		if lChild > len(array)-1 {
			break
		}

		if rChild <= len(array)-1 && array[rChild] < array[lChild] {
			smallerChild = rChild
		} else {
			smallerChild = lChild
		}

		if array[index] > array[smallerChild] {
			array[index], array[smallerChild] = array[smallerChild], array[index]
		}

		index = smallerChild
	}
}
