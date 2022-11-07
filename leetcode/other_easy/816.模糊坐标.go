package leetcode

import "strconv"

/*
 * @lc app=leetcode.cn id=816 lang=golang
 *
 * [816] 模糊坐标
 */

// @lc code=start
func ambiguousCoordinates(s string) []string {
	result := []string{}
	var err error
	temp, temp2, f := "", "", 0.00
	getCombines := func(str string) []string {
		res := []string{}
		for i := 0; i < len(str); i++ {
			if i > 0 {
				temp = str[:i] + "." + str[i:]
			} else {
				temp = str
			}
			f, err = strconv.ParseFloat(temp, 64)
			if err != nil {
				continue
			}
			temp2 = strconv.FormatFloat(f, 'f', -1, 64)
			if temp2 == temp {
				res = append(res, temp)
			}
		}
		return res
	}
	s = s[1 : len(s)-1]
	for i := 1; i < len(s); i++ {
		front, back := getCombines(s[:i]), getCombines(s[i:])
		if len(front) == 0 || len(back) == 0 {
			continue
		}
		for m := 0; m < len(front); m++ {
			for n := 0; n < len(back); n++ {
				result = append(result, "("+front[m]+", "+back[n]+")")
			}
		}
	}

	return result
}

// @lc code=end

/*
	直接将数据拆分乘x和y两部分。
	然后依次判断是否可以通过增加 . 组成数字(需要包含原数字)
	将加或未加 . 的数字转换成 浮点型，然后再转回去，如果相等，就说明可以转成数字
	然后将x的集合 和 y的集合 做交集映射成坐标即可
*/
