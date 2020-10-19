package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func removeNthFromEnd(head *ListNode, n int) *ListNode {
	if head == nil || n == 0 {
		return head
	}

	nodes := make([]*ListNode, 0, n)
	tmpNode := head
	num := 0
	for tmpNode != nil {
		nodes = append(nodes, tmpNode)
		num++
		tmpNode = tmpNode.Next
	}

	index := num - n
	if index == 0 {
		// 只有一个元素的情况
		if num > 1 {
			return nodes[1]
		} else {
			return nil
		}
	} else {
		// 要删除的是最后一位元素
		if index == num-1 {
			nodes[index-1].Next = nil
		} else {
			nodes[index-1].Next = nodes[index+1]
		}
		return head
	}
}
