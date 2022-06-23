package leetcode

/*
 * @lc app=leetcode.cn id=30 lang=golang
 *
 * [30] 串联所有单词的子串
 */

// @lc code=start
func findSubstring(s string, words []string) []int {
	firstCount := map[byte][]int{}
	for k := range words {
		firstCount[words[k][0]] = append(firstCount[words[k][0]], k)
	}
	memoUsedFlag := make([]bool, len(words))
	source := []byte(s)
	needLength := len(words[0]) * len(words)
	var isEquals func(beginIndex, count int) bool
	isEquals = func(beginIndex, count int) bool {
		if count == len(words) {
			return true
		}
		if wds, ok := firstCount[s[beginIndex]]; ok {
			for _, idx := range wds {
				if !memoUsedFlag[idx] {
					for kk := range words[idx] {
						if words[idx][kk] != source[beginIndex+kk] {
							goto lab
						}
					}
					memoUsedFlag[idx] = true
					if isEquals(beginIndex+len(words[idx]), count+1) {
						return true
					}
				}
			lab:
			}
		}
		return false
	}
	result := []int{}
	for i := 0; i <= len(s)-needLength; i++ {
		memoUsedFlag = make([]bool, len(words))
		if isEquals(i, 0) {
			result = append(result, i)
		}
	}
	return result
}

// @lc code=end

/*
	类似回溯，递归求解
	firstCount 记录相同0索引字节的单词索引
	memoUsedFlag 备忘录记录每个已经使用的单词
	source 将原s转化成字节切片，防止每次比较的时候复制
	needLength 剪枝，当走到索引i时，i及后面的单词加起来长度都不够words的所有字节长度，那么直接跳出
	isEquals(beginIndex, count)  ,beginIndex 当前比较的索引， count 当前已经比较的字符串个数

	思路： 判断s[i] 当前i索引的字节，从 fistCount 找到对应的所有 单词，依次比较，
		  如果当前单词完全吻合，那么i索引后移，再次进行递归，直到所有单词比较都相同，即可以退出

		  只要有一个单词的某个字节不相同，那么直接跳过当前单词，重新匹配。如果当前 s[i] 的i索引
		  对应的单词都不匹配，那么本轮匹配终止，向上层回溯
*/
