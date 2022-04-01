package leetcode

/*
 * @lc app=leetcode.cn id=438 lang=golang
 *
 * [438] 找到字符串中所有字母异位词
 */

// @lc code=start
func findAnagrams(s string, p string) []int {
	result := []int{}
	pMap := map[byte]int{}
	for k := range p {
		pMap[p[k]]++
	}
	iMap := map[byte]int{}
	for i := 0; i < len(s); i++ {
		if i < len(p) {
			iMap[s[i]]++
			if i < len(p)-1 {
				continue
			}
		} else {
			iMap[s[i]]++
			iMap[s[i-len(p)]]--
		}
		for k := range pMap {
			if pMap[k] != iMap[k] {
				goto lab
			}
		}
		result = append(result, i-len(p)+1)
	lab:
	}
	return result
}

// @lc code=end

/*
	将p存入 pMap，记录下每个字符的数量
	然后遍历 s，将每个 s[i] 存入iMap中
	iMap扫描长度 小于 len(p) 时，先存够
	iMap扫描长度 大于等于 len(p) 时，使用
		iMap[s[i]]++
		iMap[s[i-len(p)]]--
		来保证当前扫描存储的长度是 len(p),
	然后每个元素通过iMap后，来扫描个数是否相同
	如果相同， i-len(p)+1 就是起始元素位置
*/
