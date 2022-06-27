package leetcode

/*
 * @lc app=leetcode.cn id=522 lang=golang
 *
 * [522] 最长特殊序列 II
 */

// @lc code=start
func findLUSlength(strs []string) int {
	maxLength := len(strs[0])
	m := make([]map[int]map[string]bool, len(strs)) //str索引 - 序列长度 - 子序列
	for k := range strs {
		if maxLength < len(strs[k]) {
			maxLength = len(strs[k])
		}
		m[k] = make(map[int]map[string]bool)
		m[k][len(strs[k])] = make(map[string]bool)
		m[k][len(strs[k])][strs[k]] = true
	}

	getLUSStrs := func(idx, length int, str string) {
		if _, ok := m[idx][length]; !ok {
			m[idx][length] = make(map[string]bool)
		}
		for k := range str {
			m[idx][length][str[:k]+str[k+1:]] = true
		}
	}
	for ; maxLength > 0; maxLength-- {
		for k := range m {
			for kk := range m[k][maxLength+1] {
				getLUSStrs(k, maxLength, kk)
			}
		}
		for k1 := range m {
			if _, ok := m[k1][maxLength]; !ok {
				continue
			}
		lab:
			for str1 := range m[k1][maxLength] {
				for k2 := range m {
					if _, ok := m[k2][maxLength]; !ok || k1 == k2 {
						continue
					}
					if m[k2][maxLength][str1] {
						continue lab
					}
				}
				return maxLength
			}
		}
	}
	return -1
}

// @lc code=end

/*
	思路：从最大的长度开始逐个递减，只要该长度符合条件，就返回长度
		 为了减少求子序列的次数，只在需要该长度的子序列的时候再求子序列集合

	m := []map[int]map[string]bool{}
	  第一维度：[]map 切片 存放与 strs 切片对应索引的 map 数据
	  第二维度：map[int] map 存放对应 int 长度的 map 子序列集合
	  第三维度：map[string] map 存放包含的子序列集合

	getLUSStrs 用户生成对应的 m ， 将字符串按顺序删除对应索引的字节，得到子序列

	逐个将对应长度的子序列取出来，去和其他子序列比较，如果有相同的，那么则抛弃本次循环
*/
