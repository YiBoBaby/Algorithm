package main

import "fmt"

func main() {
	arr := []int{91, 60, 96, 13, 35, 65, 46, 65, 10, 30, 20, 31, 77, 81, 22}
	fmt.Println("排序前: ", arr)

	// 构建大顶堆
	buildMaxHeap(arr)

	// 遍历
	for i := len(arr) - 1; i >= 0; i-- {
		swap(arr, 0, i)
		heapUp(arr, 0, i)
	}

	fmt.Println("排序后: ", arr)
}

func buildMaxHeap(arr []int) {
	for i := len(arr) / 2; i >= 0; i-- {
		heapUp(arr, i, len(arr)-1)
	}
}

// 堆化数组
func heapUp(arr []int, nodeIndex, limitLen int) {
	nodeChildL := 2*nodeIndex + 1
	nodeChildR := 2*nodeIndex + 2
	largestIndex := nodeIndex

	if nodeChildL < limitLen && arr[largestIndex] < arr[nodeChildL] {
		largestIndex = nodeChildL
	}
	if nodeChildR < limitLen && arr[largestIndex] < arr[nodeChildR] {
		largestIndex = nodeChildR
	}
	if nodeIndex != largestIndex {
		swap(arr, nodeIndex, largestIndex)
		heapUp(arr, largestIndex, limitLen)
	}
}

func swap(arr []int, a, b int) {
	arr[a], arr[b] = arr[b], arr[a]
}
