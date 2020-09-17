package ch2

// go 在标准库中有container/heap 可供参考
// heap 的定义：

type IntHeap []int

func (h IntHeap) Len() int {
	return len(h)
}

func (h IntHeap) Less(i, j int) bool {
	return h[i] < h[j]
}

func (h IntHeap) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
}

func (h *IntHeap) Pop() interface{} {
	old := *h
	n := len(old)
	value = old[n-1]
	*h = old[0 : n-1]
	return value
}

func (h *IntHeap) push(x interface{}) {
	*h = append(*h, x.(int))
}
