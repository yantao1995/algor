package leetcode

/*
 * @lc app=leetcode.cn id=873 lang=golang
 *
 * [873] 最长的斐波那契子序列的长度
 */

// @lc code=start
func lenLongestFibSubseq(arr []int) int {
	m := map[int]int{} // 值 - 索引
	for k := range arr {
		m[arr[k]] = k
	}
	var maxLength, currentLength, last, rememberJ int
	for i := 0; i < len(arr); i++ {
		currentLength = 2
		last = arr[i]
		rememberJ = i + 1
		for j := i + 1; j < len(arr); j++ {
			if currentLength == 2 {
				rememberJ = j
			}
			if idx, ok := m[last+arr[j]]; ok {
				last = arr[j]
				currentLength++
				if currentLength >= 3 && currentLength > maxLength {
					maxLength = currentLength
				}
				j = idx - 1
			} else {
				currentLength = 2
				last = arr[i]
				j = rememberJ
			}
		}
	}
	return maxLength
}

// @lc code=end

/*
	暴力枚举每一次结果
	m 记录  map[当前值]值索引

	按次序枚举出两个值，然后找两个数的和为第三个值，
	第三个值是否存在，如果存在，那么就成为长度为3的序列，
	然后记录末尾两个值，再依次向后找

	如果当前链条断了，那么j应该从 currentLength==2 时的位置向后枚举
	所以用rememberJ来记录那时的位置

*/
