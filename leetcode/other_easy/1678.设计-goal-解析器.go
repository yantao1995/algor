package leetcode

/*
 * @lc app=leetcode.cn id=1678 lang=golang
 *
 * [1678] 设计 Goal 解析器
 */

// @lc code=start
func interpret(command string) string {
	result := make([]byte, 0, len(command))
	for i := 0; i < len(command); i++ {
		if command[i] == 'G' {
			result = append(result, 'G')
		} else if command[i+1] == ')' {
			result = append(result, 'o')
			i++
		} else {
			result = append(result, 'a', 'l')
			i += 3
		}
	}
	return string(result)
}

// @lc code=end

/*
	直接遍历累加即可
*/
