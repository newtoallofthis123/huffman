package tree

import (
	"container/heap"

	"github.com/newtoallofthis123/huffman/utils"
)

type HuffmanTree struct {
	root   *HuffmanNode
	Lookup Lookup
	data   []interface{}
}

func New() *HuffmanTree {
	return &HuffmanTree{
		root:   NewNode(nil),
		Lookup: make(Lookup),
	}
}

func (t *HuffmanTree) Put(data []interface{}) {
	t.data = append(t.data, data)

	t.calibrate()
}

func (t *HuffmanTree) Update(data []interface{}) {
	t.data = data

	t.calibrate()
}

// DebugPrint: Prints the tree using BFS for debugging
func (t *HuffmanTree) DebugPrint() {
	t.root.debugPrint()
}

func FromList(data []interface{}) *HuffmanTree {
	t := New()
	t.data = data

	t.calibrate()

	return t
}

func (t *HuffmanTree) calibrate() {
	pq := HuffmanPQueue{}
	heap.Init(&pq)

	freq := utils.CalculateFreq(t.data)
	for k, v := range freq {
		heap.Push(&pq, NewWithCount(k, v))
	}

	for pq.Len() > 1 {
		n1 := heap.Pop(&pq).(*HuffmanNode)
		n2 := heap.Pop(&pq).(*HuffmanNode)
		heap.Push(&pq, Join(n1, n2))
	}

	t.root = heap.Pop(&pq).(*HuffmanNode)

	t.MakeLookUp(t.root, "", &t.Lookup)
}
