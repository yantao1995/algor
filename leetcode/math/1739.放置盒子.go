package leetcode

/*
 * @lc app=leetcode.cn id=1739 lang=golang
 *
 * [1739] 放置盒子
 */

// @lc code=start
func minimumBoxes(n int) int {
	cur, i, j := 1, 1, 1
	for n > cur {
		n -= cur
		i++
		cur += i
	}
	cur = 1
	for n > cur {
		n -= cur
		j++
		cur++
	}
	return (i-1)*i/2 + j
}

// @lc code=end

/*
	参考官方题解
	找规律，如果每层都铺满的话，那么依次层数和个数为
	1 : 1
	2 : 1+2
	3 : 1+2+3
	4 : 1+2+3+4
	所以需要找到刚好铺满大于等于n的层数，然后找刚好铺满时最底下一层的个数
*/
