package leetcode

/*
 * @lc app=leetcode.cn id=855 lang=golang
 *
 * [855] 考场就座
 */

// @lc code=start
/* type ExamRoom struct {
	seats []int
	n     int
	minf  func(a, b int) int
}

func Constructor(n int) ExamRoom {
	return ExamRoom{
		seats: []int{},
		n:     n,
		minf: func(a, b int) int {
			if a < b {
				return a
			}
			return b
		},
	}
}

func (this *ExamRoom) Seat() int {
	rMid, rIdx, rMaxLength := 0, 0, 0
	mid, length := 0, 0
	prev := 0
	if len(this.seats) > 0 {
		rMaxLength = this.seats[0] //处理头
		rMid = 0
		rIdx = 0
		length2 := this.n - 1 - this.seats[len(this.seats)-1] //处理尾
		if length2 > rMaxLength {
			rMid = this.n - 1
			rIdx = len(this.seats)
		}
	}
	for k, tail := range this.seats {
		mid = (tail + prev) / 2
		length = this.minf(mid-prev, tail-mid)
		if length > rMaxLength {
			rMaxLength = length
			rMid = mid
			rIdx = k
		}
		prev = tail
	}
	this.seats = append(this.seats[:rIdx], append([]int{rMid}, this.seats[rIdx:]...)...)
	return rMid
}

func (this *ExamRoom) Leave(p int) {
	for i := 0; i < len(this.seats); i++ {
		if this.seats[i] == p {
			this.seats = append(this.seats[:i], this.seats[i+1:]...)
			return
		}
	}
} */

/**
 * Your ExamRoom object will be instantiated and called as such:
 * obj := Constructor(n);
 * param_1 := obj.Seat();
 * obj.Leave(p);
 */
// @lc code=end

/*
内存超过限制的改进版
用seats存放已经占座的坐标
例如作为 0 9 4 已经占座，那么seats =[]int{0,4,9}
然后直接可以计算相邻两个坐标的距离，只要当前中点与两端的距离大于记录的
最大距离rMaxLength那么就保存当前的坐标
for循环之前预先处理头尾结点，这样可以简化for内的处理逻辑
leave则直接删除对应的坐标


内存超过限制
Runtime Error
116/126 cases passed (N/A)
cannot allocate 1002438656-byte block (8159232 in use)

type ExamRoom struct {
	seats    []bool
	count, n int
	minf     func(a, b int) int
}

func Constructor(n int) ExamRoom {
	return ExamRoom{
		seats: make([]bool, n),
		count: 0,
		n:     n,
		minf: func(a, b int) int {
			if a < b {
				return a
			}
			return b
		},
	}
}

func (this *ExamRoom) Seat() int {
	this.count++
	if this.count == 1 {
		this.seats[0] = true
		return 0
	}
	rMid, rMaxLength := 0, 0
	mid, length := 0, 0
	i, j := 0, 0
	for !this.seats[j] { //预处理头
		j++
	}
	rMaxLength = j
	for ; j < this.n; j++ {
		if this.seats[j] { //占座
			mid = (j + i) / 2
			length = this.minf(j-mid, mid-i)
			if length > rMaxLength {
				rMaxLength = length
				rMid = mid
			}
			i = j
		} else if j == this.n-1 {
			length = this.n - i - 1
			if length > rMaxLength {
				rMid = j
			}
		}
	}
	this.seats[rMid] = true
	return rMid
}

func (this *ExamRoom) Leave(p int) {
	this.count--
	this.seats[p] = false
}
*/
