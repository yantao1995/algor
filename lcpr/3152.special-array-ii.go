package lcpr

/*
 * @lc app=leetcode.cn id=3152 lang=golang
 * @lcpr version=30204
 *
 * [3152] 特殊数组 II
 */

// @lcpr-template-start

// @lcpr-template-end
// @lc code=start
func isArraySpecial(nums []int, queries [][]int) []bool {
	left := map[int]bool{}
	for i := 1; i < len(nums); i++ {
		if i > 0 && nums[i]&1 == nums[i-1]&1 {
			left[i-1] = true
		}
	}
	result := make([]bool, 0, len(queries))
lab:
	for _, q := range queries {
		for i := q[0]; i <= q[1]; i++ {
			if left[i] && i+1 <= q[1] {
				result = append(result, false)
				continue lab
			}
		}
		result = append(result, true)
	}
	return result
}

// @lc code=end

/*
// @lcpr case=start
// [3,4,1,2,6]\n[[0,4]]\n
// @lcpr case=end

// @lcpr case=start
// [4,3,1,6]\n[[0,2],[2,3]]\n
// @lcpr case=end

*/

/*
再优化思路，将算法时间复杂度降低为O(n)


再优化思路，记录左坐标map
超时
func isArraySpecial(nums []int, queries [][]int) []bool {
	left := map[int]bool{}
	for i := 1; i < len(nums); i++ {
		if i > 0 && nums[i]&1 == nums[i-1]&1 {
			left[i-1] = true
		}
	}
	result := make([]bool, 0, len(queries))
lab:
	for _, q := range queries {
		for i := q[0]; i <= q[1]; i++ {
			if left[i] && i+1 <= q[1] {
				result = append(result, false)
				continue lab
			}
		}
		result = append(result, true)
	}
	return result
}


优化思路，记录相邻相同的坐标,然后比较
超时
func isArraySpecial(nums []int, queries [][]int) []bool {
	fs := [][]int{}
	for i := 1; i < len(nums); i++ {
		if i > 0 && nums[i]&1 == nums[i-1]&1 {
			fs = append(fs, []int{i - 1, i})
		}
	}
	result := make([]bool, 0, len(queries))
lab:
	for _, q := range queries {
		for _, f := range fs {
			if f[0] >= q[0] && f[1] <= q[1] {
				result = append(result, false)
				continue lab
			}
		}
		result = append(result, true)
	}
	return result
}


超时:
func isArraySpecial(nums []int, queries [][]int) []bool {
	for i := 0; i < len(nums); i++ {
		nums[i] &= 1
	}
	result := make([]bool, 0, len(queries))
lab:
	for _, q := range queries {
		for i := q[0] + 1; i <= q[1]; i++ {
			if nums[i] == nums[i-1] {
				result = append(result, false)
				continue lab
			}
		}
		result = append(result, true)
	}
	return result
}
*/
