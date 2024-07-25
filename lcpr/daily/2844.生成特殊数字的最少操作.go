package lcpr

import (
	"strconv"
	"strings"
)

/*
 * @lc app=leetcode.cn id=2844 lang=golang
 * @lcpr version=30204
 *
 * [2844] 生成特殊数字的最少操作
 */

// @lcpr-template-start

// @lcpr-template-end
// @lc code=start
func minimumOperations(num string) int {
	sliStr := strings.Split(num, "")
	sli := make([]int, len(num))
	maxOp := len(num)
	for k, v := range sliStr {
		sli[k], _ = strconv.Atoi(v)
		if v == "0" {
			maxOp = len(num) - 1
		}
	}
	m := map[int]bool{
		00: true,
		25: true,
		50: true,
		75: true,
	}
	currentOp := len(sli)
	for i := len(sli) - 1; i > 0; i-- {
		for j := i - 1; j >= 0; j-- {
			if m[sli[j]*10+sli[i]] {
				currentOp = min(currentOp, i-j-1+len(sli)-1-i)
			}
		}
	}
	return min(currentOp, maxOp)
}

// @lc code=end

/*
// @lcpr case=start
// "2245047"\n
// @lcpr case=end

// @lcpr case=start
// "2908305"\n
// @lcpr case=end

// @lcpr case=start
// "10"\n
// @lcpr case=end

*/

/*
	贪心法，删末尾的
	分情况 1.全删变成0
		  2.删的只剩下一个0
		  3.从末尾删，然后尾数留 00,25,50,75  。
		  	i-j-1 是i到j之间要删除的个数
			len(sli)-1-i 是i到末尾要删除的个数
	三种情况取最小值即可
*/
