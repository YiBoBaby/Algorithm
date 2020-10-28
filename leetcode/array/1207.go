package main

func uniqueOccurrences(arr []int) bool {
	bucket := make([]int, 2001, 2001)
	for i := range arr {
		bucIndex := arr[i] + 1001
		bucket[bucIndex] = bucket[bucIndex] + 1
	}

	mp := make(map[int]int8)
	for i := 1; i < len(bucket); i++ {
		count := bucket[i]
		if count > 0 {
			if _, ok := mp[count]; ok {
				return false
			} else {
				mp[count] = 1
			}
		}
	}
	return true
}
