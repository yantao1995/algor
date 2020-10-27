package easy

/*
 * @lc app=leetcode.cn id=38 lang=golang
 *
 * [38] 外观数列
 */

// @lc code=start
func countAndSay(n int) string {
	m := map[int]string{}
	m[1] = "1"
	for i := 1; i < n; i++ {
		str := []byte{}
		for j := 0; j < len(m[i]); j++ {
			tempByte := m[i][j] //记录当前值
			tempCorsur := j + 1 //前进游标
			temp := 1           //数量
			breakFlag := false
			for tempCorsur < len(m[i]) {
				if m[i][tempCorsur] != tempByte {
					breakFlag = true
					break
				} else {
					temp++
					tempCorsur++
				}
			}
			j = tempCorsur
			if breakFlag { //break出来的 后退一步
				j--
			}
			str = append(str, []byte{byte(temp + 48), tempByte}...)
		}
		m[i+1] = string(str)
	}
	return m[n]
}

// @lc code=end
