package easy

/*
 * @lc app=leetcode.cn id=1299 lang=golang
 *
 * [1299] 将每个元素替换为右侧最大元素
 */

// @lc code=start
func replaceElements(arr []int) []int {
	for k := range arr {
		max := -1<<32 - 1
		for i := k + 1; i < len(arr); i++ {
			if max < arr[i] {
				max = arr[i]
			}
		}
		arr[k] = max
	}
	arr[len(arr)-1] = -1
	return arr
}

// @lc code=end
