package leetcode

/*
 * @lc app=leetcode.cn id=1694 lang=golang
 *
 * [1694] 重新格式化电话号码
 */

// @lc code=start
func reformatNumber(number string) string {
	bts := make([]byte, 0, len(number))
	count := 0
	for i := 0; i < len(number); i++ {
		if number[i] != ' ' && number[i] != '-' {
			count++
			bts = append(bts, number[i])
			if count%3 == 0 && i != len(number)-1 && bts[len(bts)-1] != '-' {
				bts = append(bts, '-')
			}
		}
	}
	if bts[len(bts)-1] == '-' {
		bts = bts[:len(bts)-1]
	} else if bts[len(bts)-2] == '-' {
		bts[len(bts)-2], bts[len(bts)-3] = bts[len(bts)-3], bts[len(bts)-2]
	}
	return string(bts)
}

// @lc code=end

/*
	直接读取字符串，然后按3个分组
	然后判断末尾第一个是不是 - ，如果是，直接去掉
	如果末尾前一个是 - ，那么直接与前一个字节交换即可
*/
