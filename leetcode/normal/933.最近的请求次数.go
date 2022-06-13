package leetcode

/*
 * @lc app=leetcode.cn id=933 lang=golang
 *
 * [933] 最近的请求次数
 */

// @lc code=start
type RecentCounter struct {
	Time []int
}

func Constructor() RecentCounter {
	return RecentCounter{
		Time: []int{},
	}
}

func (this *RecentCounter) Ping(t int) int {
	this.Time = append(this.Time, t)
	count := 0
	for i := len(this.Time) - 1; i >= 0 && t-3000 <= this.Time[i]; i-- {
		count++
	}
	return count
}

/**
 * Your RecentCounter object will be instantiated and called as such:
 * obj := Constructor();
 * param_1 := obj.Ping(t);
 */
// @lc code=end
