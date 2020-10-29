package main

import (
	"strconv"
)

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func sumNumbers(root *TreeNode) int {
	if root == nil {
		return 0
	}
	var t int64
	DFSPre(root, "", &t)
	return int(t)
}

func DFSPre(n *TreeNode, pns string, t *int64) {
	pns = pns + strconv.FormatInt(int64(n.Val), 10)
	if n.Left == nil {
		if n.Right != nil {
			DFSPre(n.Right, pns, t)
		} else {
			if v, err := strconv.ParseInt(pns, 10, 64); err == nil {
				*t = v + *t
			}
		}
	} else {
		DFSPre(n.Left, pns, t)
		if n.Right != nil {
			DFSPre(n.Right, pns, t)
		}
	}
}
