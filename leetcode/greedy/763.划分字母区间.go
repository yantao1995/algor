package leetcode

/*
 * @lc app=leetcode.cn id=763 lang=golang
 *
 * [763] 划分字母区间
 */

// @lc code=start
func partitionLabels(s string) []int {
	result := []int{}
	m := map[byte]int{} // 记录该字符出现的最大的索引
	for i := 0; i < len(s); i++ {
		m[s[i]] = i
	}
	max := 0
	for i := 0; i < len(s); i = max + 1 {
		max = m[s[i]]
		for j := i + 1; j < max; j++ {
			if max < m[s[j]] {
				max = m[s[j]]
			}
		}
		result = append(result, max-i+1)
	}
	return result
}

// @lc code=end

/*
	每次贪该字符的最大的索引
	类似于跳跃游戏-ii，每次都跳最长的索引
	"ababcbacadefegdehijhklij"
	例如当前索引0为a，a的最大索引是8
	那么 0-8内的所有元素都要查看，如果 b c大于8,那么右边界会到最大位置
*/
