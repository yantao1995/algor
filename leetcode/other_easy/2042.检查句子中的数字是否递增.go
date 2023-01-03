package leetcode

import (
	"strconv"
	"strings"
)

/*
 * @lc app=leetcode.cn id=2042 lang=golang
 *
 * [2042] 检查句子中的数字是否递增
 */

// @lc code=start
func areNumbersAscending(s string) bool {
	seq := strings.Split(s, " ")
	last := -1
	for i := 0; i < len(seq); i++ {
		temp, err := strconv.Atoi(seq[i])
		if err == nil {
			if temp <= last {
				return false
			}
			last = temp
		}
	}
	return true
}

// @lc code=end

/*
	用空格切分字符串，然后挨个转换成int类型，
	如果转换成功，那么说明是整数，直接与上一个
	整数比较大小即可，因为整数始终大于零，所以
	可以初始化last为-1
*/
