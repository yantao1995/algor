package leetcode

/*
 * @lc app=leetcode.cn id=89 lang=golang
 *
 * [89] 格雷编码
 */

// @lc code=start
func grayCode(n int) []int {
	getBit := func(i, n int) []int {
		rlt := make([]int, n)
		for idx := len(rlt) - 1; i > 0 || idx >= 0; idx, i = idx-1, i>>1 {
			rlt[idx] = i & 1
		}
		return rlt
	}
	cmp := func(a, b []int) bool {
		count := 0
		for i := 0; i < len(a) && count < 3; i++ {
			if a[i] != b[i] {
				count++
			}
		}
		return count == 1
	}
	bitMap := map[int][]int{}
	for i := 0; i < 1<<n; i++ {
		bitMap[i] = getBit(i, n)
	}
	distinct := map[int]bool{}
	result := []int{}
	find := false
	var backtrace func(count int, temp []int)
	backtrace = func(count int, temp []int) {
		if count == 1<<n || find {
			if find {
				return
			}
			if cmp(bitMap[temp[len(temp)-1]], bitMap[temp[0]]) {
				result = append(result, temp...)
				find = true
			}
			return
		}
		for i := 0; i < 1<<n; i++ {
			if !distinct[i] {
				if len(temp) > 0 && !cmp(bitMap[temp[len(temp)-1]], bitMap[i]) {
					continue
				}
				distinct[i] = true
				backtrace(count+1, append(temp, i))
				distinct[i] = false
			}
		}
	}
	backtrace(0, []int{})
	return result
}

// @lc code=end
