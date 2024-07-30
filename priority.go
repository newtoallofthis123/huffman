package huffman

// HuffmanPQueue is a priority queue implementation of a Huffman tree.
type HuffmanPQueue []*HuffmanNode

// Len returns the number of elements in the priority queue.
func (pq HuffmanPQueue) Len() int { return len(pq) }

// Less compares two elements in the priority queue.
func (pq HuffmanPQueue) Less(i, j int) bool {
	return pq[i].count < pq[j].count
}

// Swap swaps two elements in the priority queue.
func (pq HuffmanPQueue) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
}

// Push adds an element to the priority queue.
func (pq *HuffmanPQueue) Push(x any) {
	item := x.(*HuffmanNode)
	*pq = append(*pq, item)
}

// Pop removes and returns an element from the priority queue.
func (pq *HuffmanPQueue) Pop() any {
	old := *pq
	n := len(old)
	item := old[n-1]
	old[n-1] = nil // avoid memory leak
	*pq = old[0 : n-1]
	return item
}
