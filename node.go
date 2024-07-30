package huffman

import "fmt"

type HuffmanNode struct {
	data  interface{}
	count int
	left  *HuffmanNode
	right *HuffmanNode
}

// NewNode creates a new HuffmanNode with the given data and initial count of 0.
func NewNode(data interface{}) *HuffmanNode {
	return &HuffmanNode{data: data, count: 0}
}

// NewWithCount creates a new HuffmanNode with the given data and count.
func NewWithCount(data interface{}, count int) *HuffmanNode {
	return &HuffmanNode{data: data, count: count}
}

// IsLeaf checks if the HuffmanNode is a leaf node.
// This function returns a boolean value indicating whether the HuffmanNode is a leaf node.
func (n *HuffmanNode) IsLeaf() bool {
	return n.left == nil && n.right == nil
}

// Join combines two HuffmanNodes into a new HuffmanNode.
func Join(n1, n2 *HuffmanNode) *HuffmanNode {
	return &HuffmanNode{count: n1.count + n2.count, left: n1, right: n2}
}

// debugPrint prints the Huffman tree using BFS for debugging.
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
