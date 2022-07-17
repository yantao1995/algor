package leetcode

/*
 * @lc app=leetcode.cn id=565 lang=golang
 *
 * [565] 数组嵌套
 */

// @lc code=start

func arrayNesting(nums []int) int {
	countMap := map[int]int{}
	temp, temp2, tempCount := 0, 0, 0
	maxLength := 0
	for k := range nums {
		if nums[k] != -1 {
			tempCount = 1
			temp, temp2 = nums[k], nums[k]
			for {
				if _, ok := countMap[temp]; ok || nums[temp] == -1 {
					countMap[k] = countMap[temp] + tempCount - 1
					break
				} else {
					tempCount++
					temp = nums[temp]
					nums[temp2] = -1
					temp2 = temp
				}
			}
			if countMap[k] > maxLength {
				maxLength = countMap[k]
			}
		}
	}
	return maxLength
}

// @lc code=end

/*
	将遍历过的值置为-1，这样k走到当前-1时，可以直接跳过，因为当前遍历过，那么从当前值开始走，一定小于maxLength
	countMap 记录当前索引开始的计数
	maxLength 记录当前找到的最大的长度
	tempCount 记录从当前k开始，走过的所有值的个数
	temp 和 temp2，为了用来记录下一个位置索引和给当前位置索引的值置为-1
	countMap[k] = countMap[temp] + tempCount - 1   减一 是因为每轮都会
		进行一次 tempCount ++,在遇到-1时，也会+1，所以此处减掉
*/
