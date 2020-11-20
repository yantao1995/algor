package utils

// container/heap 接口实现

type Ints []int

func (n Ints) Swap(i, j int)      { n[i], n[j] = n[j], n[i] }
func (n Ints) Len() int           { return len(n) }
func (n Ints) Less(i, j int) bool { return n[i] < n[j] }

func (n *Ints) Push(x interface{}) {
	*n = append(*n, x.(int))
}

func (n *Ints) Pop() interface{} {
	old := *n
	rtn := old[len(old)-1]
	*n = old[:len(old)-1]
	return rtn
}
