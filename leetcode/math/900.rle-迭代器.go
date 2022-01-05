package leetcode

/*
 * @lc app=leetcode.cn id=900 lang=golang
 *
 * [900] RLE 迭代器
 */

// @lc code=start
type RLEIterator struct {
	Sequence           []int
	Cursor, Index, Len int
}

func Constructor(encoding []int) RLEIterator {
	return RLEIterator{
		Sequence: encoding,
		Cursor:   0,
		Index:    0,
		Len:      len(encoding),
	}
}

func (this *RLEIterator) Next(n int) int {
	result := -1
	for n > 0 {
		if this.Index >= this.Len {
			return -1
		}
		if this.Sequence[this.Index] >= n+this.Cursor {
			this.Cursor += n
			n = 0
			result = this.Sequence[this.Index+1]
		} else {
			n -= (this.Sequence[this.Index] - this.Cursor)
			this.Index += 2
			this.Cursor = 0
		}
	}
	return result
}

/**
 * Your RLEIterator object will be instantiated and called as such:
 * obj := Constructor(encoding);
 * param_1 := obj.Next(n);
 */
// @lc code=end
