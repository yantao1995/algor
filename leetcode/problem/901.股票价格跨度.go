package leetcode

/*
 * @lc app=leetcode.cn id=901 lang=golang
 *
 * [901] 股票价格跨度
 */

// @lc code=start
type StockSpanner struct {
	log, span []int
}

// func Constructor() StockSpanner {
// 	return StockSpanner{
// 		make([]int, 0),
// 		make([]int, 0),
// 	}
// }

func (this *StockSpanner) Next(price int) int {
	this.log = append(this.log, price)
	sp := 1
	n := len(this.log) - 1
	for i := n - 1; i >= 0; {
		if this.log[n] < this.log[i] {
			break
		}
		sp += this.span[i]
		i -= this.span[i]
	}
	this.span = append(this.span, sp)
	return sp
}

/**
 * Your StockSpanner object will be instantiated and called as such:
 * obj := Constructor();
 * param_1 := obj.Next(price);
 */
// @lc code=end

/*
	span 存储跨度序列
	log 存储当日报价

	i为前一天，如果当前报价大于i的报价，那么当前跨度直接加上i的跨度，
		然后i索引向前移动到i减去跨度的位置,然后继续
*/
