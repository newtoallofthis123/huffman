package tree

import (
	"container/heap"

	"github.com/newtoallofthis123/huffman/utils"
)

type HuffmanTree struct {
	root *HuffmanNode
}

func NewHuffmanTree() *HuffmanTree {
	return &HuffmanTree{
		root: NewNode(nil),
	}
}

// DebugPrint: Prints the tree using BFS for debugging
func (t *HuffmanTree) DebugPrint() {
	t.root.debugPrint()
}

func FromList(data any) *HuffmanTree {
	pq := HuffmanPQueue{}
	heap.Init(&pq)

	freq := utils.CalculateFreq(data.([]interface{}))
	for k, v := range freq {
		heap.Push(&pq, NewWithCount(k, v))
	}

	for pq.Len() > 1 {
		n1 := heap.Pop(&pq).(*HuffmanNode)
		n2 := heap.Pop(&pq).(*HuffmanNode)
		heap.Push(&pq, Join(n1, n2))
	}

	t := NewHuffmanTree()
	t.root = heap.Pop(&pq).(*HuffmanNode)

	return t
}
