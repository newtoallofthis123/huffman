package huffman

import (
	"container/heap"
)

// HuffmanTree is a Huffman tree implementation.
type HuffmanTree struct {
	root *HuffmanNode
	// Lookup is a map that stores the Huffman codes for each data in the tree.
	Lookup Lookup
	data   []interface{}
}

// New creates a new instance of HuffmanTree.
//
// It initializes a new HuffmanTree with an empty root node and an empty Lookup map.
// The function returns a pointer to the newly created HuffmanTree.
func New() *HuffmanTree {
	return &HuffmanTree{
		root:   NewNode(nil),
		Lookup: make(Lookup),
	}
}

// Put adds the given data to the HuffmanTree.
//
// It appends the data to the existing data slice of the HuffmanTree and then
// calls the calibrate method to rebalance the tree.
func (t *HuffmanTree) Put(data []interface{}) {
	t.data = append(t.data, data)

	t.Calibrate()
}

// Update updates the HuffmanTree with new data.
//
// It replaces the existing data slice of the HuffmanTree with the given data
// and then calls the calibrate method to rebalance the tree.
func (t *HuffmanTree) Update(data []interface{}) {
	t.data = data

	t.Calibrate()
}

// DebugPrint: Prints the tree using BFS for debugging
func (t *HuffmanTree) DebugPrint() {
	t.root.debugPrint()
}

// FromList creates a new HuffmanTree from a list of data.
// This is the preferred way to create a HuffmanTree and is the most managed way.
func FromList(data []interface{}) *HuffmanTree {
	t := New()
	t.data = data

	t.Calibrate()

	return t
}

// FromBinary creates a new HuffmanTree from a list of binary data.
// Can be useful for low level operations, like file compression
func FromBinary(data []uint8) *HuffmanTree {
	return FromList(ConvertToInterface(data))
}

// Calibrate recalibrates the Huffman tree by rebalancing the tree.
// This can be used externally to rebalance the tree, however
// the Update method is preferred.
func (t *HuffmanTree) Calibrate() {
	pq := HuffmanPQueue{}
	heap.Init(&pq)

	freq := CalculateFreq(t.data)
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
