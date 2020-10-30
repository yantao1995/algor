package leetcode

/*
 * @lc app=leetcode.cn id=295 lang=golang
 *
 * [295] 数据流的中位数
 */

// @lc code=start
type MedianFinder struct {
	Data []int
	Len  int
}

/** initialize your data structure here. */
func Constructor() MedianFinder {
	return MedianFinder{
		Data: []int{},
		Len:  0,
	}
}

func (this *MedianFinder) AddNum(num int) {
	if len(this.Data) == 0 {
		this.Data = []int{num}
	}
	this.Data = append(this.Data, num)
	for i := len(this.Data) - 1; i > 1 &&
		this.Data[i] < this.Data[i-1]; i-- {
		this.Data[i], this.Data[i-1] = this.Data[i-1], this.Data[i]
	}
	this.Len++
}

func (this *MedianFinder) FindMedian() float64 {
	if this.Len%2 == 1 {
		return float64(this.Data[(this.Len+1)/2])
	}
	return (float64(this.Data[(this.Len+1)/2]) + float64(this.Data[(this.Len+1)/2+1])) / 2
}

/**
 * Your MedianFinder object will be instantiated and called as such:
 * obj := Constructor();
 * obj.AddNum(num);
 * param_2 := obj.FindMedian();
 */
// @lc code=end
