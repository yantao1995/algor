package leetcode

/*
 * @lc app=leetcode.cn id=729 lang=golang
 *
 * [729] 我的日程安排表 I
 */

// @lc code=start
type MyCalendar struct {
	store [][]int
}

func Constructor() MyCalendar {
	return MyCalendar{
		store: [][]int{},
	}
}

func (this *MyCalendar) Book(start int, end int) bool {
	for k := range this.store {
		if this.store[k][0] <= start &&
			this.store[k][1] > start {
			return false
		}
		if this.store[k][0] > start &&
			this.store[k][0] < end {
			return false
		}
	}
	this.store = append(this.store, []int{start, end})
	return true
}

/**
 * Your MyCalendar object will be instantiated and called as such:
 * obj := Constructor();
 * param_1 := obj.Book(start,end);
 */
// @lc code=end
