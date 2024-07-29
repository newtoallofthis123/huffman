package tree

type HuffmanPQueue []*HuffmanNode

func (pq HuffmanPQueue) Len() int { return len(pq) }

func (pq HuffmanPQueue) Less(i, j int) bool {
	return pq[i].count < pq[j].count
}

func (pq HuffmanPQueue) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
}

func (pq *HuffmanPQueue) Push(x any) {
	item := x.(*HuffmanNode)
	*pq = append(*pq, item)
}

func (pq *HuffmanPQueue) Pop() any {
	old := *pq
	n := len(old)
	item := old[n-1]
	old[n-1] = nil // avoid memory leak
	*pq = old[0 : n-1]
	return item
}
