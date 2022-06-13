package leetcode

/*
 * @lc app=leetcode.cn id=27 lang=golang
 *
 * [27] 移除元素
 */

// @lc code=start
func removeElement(nums []int, val int) int {
	if len(nums) == 0 {
		return 0
	}
	length := 0
	next := 0
	for i := 0; i < len(nums); i++ {
		if nums[i] == val {
			for next = i; next < len(nums); next++ {
				if nums[next] != val {
					nums[i], nums[next] = nums[next], nums[i]
					length++
					break
				}
			}
			if next == len(nums) {
				break
			}
			continue
		}
		length++
	}
	return length
}

// @lc code=end
