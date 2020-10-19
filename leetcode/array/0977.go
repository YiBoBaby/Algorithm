package main

func sortedSquares(A []int) []int {
	if A == nil || len(A) == 0 {
		return nil
	}
	for A[0] < 0 {
		tmp := -A[0]
		for i := 1; i < len(A); i++ {
			// 全负数情况
			if i == len(A) {
				A[i] = tmp
			}
			if A[i] < 0 || A[i] < tmp {
				A[i-1] = A[i]
			} else {
				A[i] = tmp
			}
		}
	}
	for i, n := range A {
		A[i] = n * n
	}
	return A
}
