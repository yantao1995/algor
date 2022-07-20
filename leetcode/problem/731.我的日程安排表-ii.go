package leetcode

/*
 * @lc app=leetcode.cn id=731 lang=golang
 *
 * [731] 我的日程安排表 II
 */

// @lc code=start
type MyCalendarTwo struct {
	store1, store2 [][]int
	min, max       func(int, int) int
}

// func Constructor() MyCalendarTwo {
// 	return MyCalendarTwo{
// 		store1: [][]int{},
// 		store2: [][]int{},
// 		min: func(a, b int) int {
// 			if a < b {
// 				return a
// 			}
// 			return b
// 		},
// 		max: func(a, b int) int {
// 			if a > b {
// 				return a
// 			}
// 			return b
// 		},
// 	}
// }

func (this *MyCalendarTwo) Book(start int, end int) bool {
	//check2
	for k := range this.store2 {
		if (this.store2[k][0] <= start &&
			this.store2[k][1] > start) ||
			(this.store2[k][0] >= start &&
				this.store2[k][0] < end) {
			return false
		}
	}
	//check1
	flag := true
	for k := range this.store1 {
		if (this.store1[k][0] <= start &&
			this.store1[k][1] > start) ||
			(this.store1[k][0] >= start &&
				this.store1[k][0] < end) {
			flag = false
			this.store2 = append(this.store2, []int{
				this.max(this.store1[k][0], start), this.min(this.store1[k][1], end)})
			if this.store1[k][0] != start {
				this.store1 = append(this.store1,
					[]int{this.min(this.store1[k][0], start), this.max(this.store1[k][0], start)})
			}
			if this.store1[k][1] != end {
				this.store1 = append(this.store1,
					[]int{this.min(this.store1[k][1], end), this.max(this.store1[k][1], end)})
			}
		}
	}
	if flag {
		this.store1 = append(this.store1, []int{start, end})
	}
	return true
}

/**
 * Your MyCalendarTwo object will be instantiated and called as such:
 * obj := Constructor();
 * param_1 := obj.Book(start,end);
 */
// @lc code=end

/*
	store1 存储单一预定
	store2 存储双重预定

	优先判断 store2 ，如果双重预定包含重复区间，那么就是三重预定
	如果不包含三重预定，那么判断store1，如果包含重复区间，那么就是双重预定
		如果双重预定，那么直接丢进store2
		去掉相交区间后，剩下的两个区间，如果长度都大于0，那么丢进store1内
*/
