package leetcode

/*
 * @lc app=leetcode.cn id=1436 lang=golang
 *
 * [1436] 旅行终点站
 */

// @lc code=start
func destCity(paths [][]string) string {
	ht := map[string]string{}
	for _, v := range paths {
		ht[v[0]] = v[1]
	}
	city := paths[0][1]
	for ht[city] != "" {
		city = ht[city]
	}
	return city
}

// @lc code=end
