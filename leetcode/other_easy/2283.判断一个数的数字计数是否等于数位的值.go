package leetcode

/*
 * @lc app=leetcode.cn id=2283 lang=golang
 *
 * [2283] 判断一个数的数字计数是否等于数位的值
 */

// @lc code=start
func digitCount(num string) bool {
	m := make([]int, len(num))
	count := 0
	for k, v := range num {
		m[k] = int(v - '0')
		if m[k] > 0 {
			count++
		}
	}
	for _, v := range num {
		if int(v-'0') >= len(m) || m[v-'0'] == 0 {
			return false
		}
		m[v-'0']--
		if m[v-'0'] == 0 {
			count--
		}
	}
	return count == 0
}

// @lc code=end

/*
解法1:
	数组记录每位的个数，然后遍历num，对该数字进行相减，
	如果所有大于0的位都减完了，那么就可以满足条件

解法2：
	map 记录数字的个数，然后循环直接比较是否相等

	func digitCount(num string) bool {
		m := map[int]int{}
		for _, v := range num {
			m[int(v-'0')]++
		}
		for k, v := range num {
			if int(v-'0') != m[k] {
				return false
			}
		}
		return true
	}

*/
