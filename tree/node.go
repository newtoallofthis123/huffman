package tree

import "fmt"

type HuffmanNode struct {
	data  interface{}
	count int
	left  *HuffmanNode
	right *HuffmanNode
}

func NewNode(data interface{}) *HuffmanNode {
	return &HuffmanNode{data: data, count: 0}
}

func NewWithCount(data interface{}, count int) *HuffmanNode {
	return &HuffmanNode{data: data, count: count}
}

func (n *HuffmanNode) IsLeaf() bool {
	return n.left == nil && n.right == nil
}

func Join(n1, n2 *HuffmanNode) *HuffmanNode {
	return &HuffmanNode{count: n1.count + n2.count, left: n1, right: n2}
}

func (n *HuffmanNode) debugPrint() {
	queue := []*HuffmanNode{n}
	for len(queue) > 0 {
		n := queue[0]
		queue = queue[1:]
		fmt.Printf("%c, %d", n.data, n.count)
		if n.left != nil {
			queue = append(queue, n.left)
		}
		if n.right != nil {
			queue = append(queue, n.right)
		}
		fmt.Println()
	}
}
