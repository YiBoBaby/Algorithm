package main

import "fmt"

func main() {
	arr := []int{91, 60, 96, 13, 35, 65, 46, 65, 10, 30, 20, 31, 77, 81, 22}
	fmt.Println("排序前: ", arr)
	heapSort(arr)
	fmt.Println("排序后: ", arr)
}

func heapSort(arr []int) []int {
	// 堆化
	buildMaxHeap(arr)
	for i := len(arr) - 1; i >= 0; i-- {
		swap(arr, 0, i)
		HeapUp(arr, 0, i-1)
	}
	return arr
}

func buildMaxHeap(arr []int) {
	for i := len(arr) / 2; i >= 0; i-- {
		HeapUp(arr, i, len(arr)-1)
	}
}

func HeapUp(arr []int, nonLeafNode, limit int) {
	left := nonLeafNode*2 + 1
	right := nonLeafNode*2 + 2
	largest := nonLeafNode

	if left <= limit && arr[left] > arr[largest] {
		largest = left
	}

	if right <= limit && arr[right] > arr[largest] {
		largest = right
	}

	if nonLeafNode != largest {
		swap(arr, largest, nonLeafNode)
		HeapUp(arr, largest, limit)
	}
}

func swap(arr []int, i, j int) {
	arr[i], arr[j] = arr[j], arr[i]
}
