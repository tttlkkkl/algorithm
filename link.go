package main

import (
	"fmt"

	"github.com/isdamir/gotype"
)

// reverse 带头节点的逆序
func reverse(node *gotype.LNode) {
	if node == nil || node.Next == nil {
		return
	}
	x := func(p *gotype.LNode) interface{} {
		if p != nil {
			return p.Data
		}
		return nil
	}
	var pre *gotype.LNode // 前驱节点
	var cur *gotype.LNode // 定义当前节点
	var next = node.Next
	for next != nil {
		cur = next.Next
		next.Next = pre
		pre = next
		next = cur
		fmt.Println(x(pre), x(cur), x(next))
	}
	node.Next = pre
}

func link() {
	head := &gotype.LNode{}
	gotype.CreateNode(head, 8)
	gotype.PrintNode("逆序前：", head)
	reverse(head)
	gotype.PrintNode("逆序后：", head)
}
