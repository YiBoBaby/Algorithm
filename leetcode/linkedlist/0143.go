package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func main() {
	//n5 := ListNode{Val: 5}
	//n4 := ListNode{Val: 4, Next: &n5}
	n4 := ListNode{Val: 4}
	n3 := ListNode{Val: 3, Next: &n4}
	n2 := ListNode{Val: 2, Next: &n3}
	n1 := ListNode{Val: 1, Next: &n2}

	reorderList(&n1)
}

func reorderList(head *ListNode) {
	if head == nil || head.Next == nil {
		return
	}

	// 定义快慢指针找到中心点
	slow, fast := head, head
	for fast != nil && fast.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
	}

	// 将中间点之后的部分存入切片
	tmp := slow.Next
	nodes := make([]*ListNode, 0, 0)
	for tmp != nil {
		nodes = append(nodes, tmp)
		tmp = tmp.Next
	}

	// 连接整个链表
	l := len(nodes)
	first, second := head, head.Next
	for first != slow && l > 0 {
		first.Next = nodes[l-1]
		nodes[l-1].Next = second

		first = second
		second = second.Next
		l--
	}
	// 防止链表循环
	slow.Next = nil
}
