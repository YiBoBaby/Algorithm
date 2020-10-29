package main

import "container/list"

func preorderTraversal(root *TreeNode) []int {
	stack := list.New()
	node := root
	res := make([]int, 0, 0)

	for stack.Len() > 0 || node != nil {
		if node != nil {
			stack.PushBack(node)
			res = append(res, node.Val)
			node = node.Left
		} else {
			ele := stack.Back()
			if v, ok := ele.Value.(*TreeNode); ok {
				var tail *TreeNode
				tail = v
				node = tail.Right
			}
			stack.Remove(ele)
		}
	}

	return res
}
