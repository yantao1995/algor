package leetcode

/*
 * @lc app=leetcode.cn id=997 lang=golang
 *
 * [997] 找到小镇的法官
 */

// @lc code=start
func findJudge(N int, trust [][]int) int {
	if len(trust) == 0 && N == 1 {
		return N
	}
	allHt := map[int]bool{} //所有数字
	tHt := map[int][]int{}  //单键多值
	for k := range trust {
		tHt[trust[k][0]] = append(tHt[trust[k][0]], trust[k][1])
		allHt[trust[k][0]] = true
		allHt[trust[k][1]] = true
	}
	for k := range allHt {
		count := 0
		for _, v := range tHt {
			find := false
			for m := range v {
				if v[m] == k {
					count++
					find = true
					break
				}
			}
			if !find {
				break
			}
		}
		if count == N-1 {
			if _, ok := tHt[k]; !ok {
				return k
			}

		}
	}
	return -1
}

// @lc code=end
