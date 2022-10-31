package leetcode

/*
 * @lc app=leetcode.cn id=481 lang=golang
 *
 * [481] 神奇字符串
 */

// @lc code=start
func magicalString(n int) int {
	seq := []byte{'1', '2', '2'}
	for i, tail := 2, 2; len(seq) < n; i++ {
		if seq[i] == '2' {
			if seq[tail] == '2' {
				seq = append(seq, '1', '1')
			} else {
				seq = append(seq, '2', '2')
			}
			tail += 2
		} else if seq[i] == '1' {
			if seq[tail] == '2' {
				seq = append(seq, '1')
			} else {
				seq = append(seq, '2')
			}
			tail++
		}
	}
	count1 := 0
	for i := 0; i < n; i++ {
		if seq[i] == '1' {
			count1++
		}
	}
	return count1
}

// @lc code=end

/*
	直接模拟出所有字符串长度，然后统计出里面的数字个数即可
	索引i记录当前判断的数,索引tail记录当前数字串的末尾数字，
	用于处理尾部添加是1还是2
*/
