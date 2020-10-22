package main

import "fmt"

func main() {
	fmt.Printf("%v", partitionLabels("qiejxqfnqceocmy"))
}

func partitionLabels(S string) []int {
	// 遍历字符串记录每个字符的最大下标
	indexMap := make(map[uint8]int)
	for i := 0; i < len(S); i++ {
		indexMap[S[i]] = i
	}

	// 再次遍历获取每个字符区间内的最大下标
	result := make([]int, 0, 0)
	lastIndex := 0
	for i := 0; i < len(S); {
		iMaxIndex := indexMap[S[i]]
		intervalMaxIndex := iMaxIndex
		for j := i + 1; j <= iMaxIndex; j++ {
			if indexMap[S[j]] > intervalMaxIndex {
				intervalMaxIndex = indexMap[S[j]]
				// 新扩展的字符也可能触发边界
				iMaxIndex = intervalMaxIndex
			}
		}
		tmp := intervalMaxIndex - lastIndex + 1
		lastIndex = intervalMaxIndex + 1
		result = append(result, tmp)
		i = intervalMaxIndex + 1
	}
	return result
}
