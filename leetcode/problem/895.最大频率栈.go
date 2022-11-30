package leetcode

/*
 * @lc app=leetcode.cn id=895 lang=golang
 *
 * [895] 最大频率栈
 */

// @lc code=start
type FreqStack struct {
	stacks  [][]int
	maxFreq int
	freq    map[int]int
}

// func Constructor() FreqStack {
// 	return FreqStack{
// 		stacks:  make([][]int, 1),
// 		maxFreq: 0,
// 		freq:    map[int]int{},
// 	}
// }

func (this *FreqStack) Push(val int) {
	this.freq[val]++
	if this.freq[val] > this.maxFreq {
		this.maxFreq = this.freq[val]
	}
	if len(this.stacks) <= this.freq[val] {
		this.stacks = append(this.stacks, []int{val})
	} else {
		this.stacks[this.freq[val]] = append(this.stacks[this.freq[val]], val)
	}
}

func (this *FreqStack) Pop() int {
	n := len(this.stacks[this.maxFreq]) - 1
	val := this.stacks[this.maxFreq][n]
	this.stacks[this.maxFreq] = this.stacks[this.maxFreq][:n]
	if n == 0 {
		this.maxFreq--
	}
	this.freq[val]--
	return val
}

/**
 * Your FreqStack object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(val);
 * param_2 := obj.Pop();
 */
// @lc code=end

/*
 空间换时间
 每一个频率都存一个栈
 然后记录下当前的最大频率，然后依次出栈即可，当前栈为空时，直接将最大频率减一


爆搜，超时 (30/38)

type FreqStack struct {
	m     map[int]int
	stack []int
	temp  map[int]bool
}

func Constructor() FreqStack {
	return FreqStack{
		m:     map[int]int{},
		stack: []int{},
		temp:  map[int]bool{},
	}
}

func (this *FreqStack) Push(val int) {
	this.m[val]++
	this.stack = append(this.stack, val)
}

func (this *FreqStack) Pop() int {
	maxCount := 0
	for _, v := range this.m {
		if maxCount < v {
			maxCount = v
		}
	}
	this.temp = make(map[int]bool)
	for k, v := range this.m {
		if maxCount == v {
			this.temp[k] = true
		}
	}
	for i := len(this.stack) - 1; i >= 0; i-- {
		if this.temp[this.stack[i]] {
			val := this.stack[i]
			this.stack = append(this.stack[:i], this.stack[i+1:]...)
			this.m[val]--
			return val
		}
	}
	return 0
}

*/
