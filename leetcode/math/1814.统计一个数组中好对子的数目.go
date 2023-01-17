package leetcode

/*
 * @lc app=leetcode.cn id=1814 lang=golang
 *
 * [1814] 统计一个数组中好对子的数目
 */

// @lc code=start
func countNicePairs(nums []int) int {
	rev := func(a int) int {
		result := 0
		for a > 0 {
			result *= 10
			result += a % 10
			a /= 10
		}
		return result
	}
	m := map[int]int{}
	for i := 0; i < len(nums); i++ {
		m[nums[i]-rev(nums[i])]++
	}
	count := 0
	mod := int(1e9 + 7)
	for _, v := range m {
		count += v * (v - 1) / 2
		count %= mod
	}
	return count
}

// @lc code=end

/*

原来是判题系统是数for的个数，删去 rev 内的 for 即可通过

	for a%10 == 0 {
		a /= 10
	}

原式子中，这个 for 虽然多余了，仅仅只多了几次计算 a%10==0

	for a > 0 {
		result *= 10
		result += a % 10
		a /= 10
	}



优化公式 : nums[i] - rev(nums[i]) == nums[j] - rev(nums[j])
优化计算逻辑 超时
Time Limit Exceeded
83/84 cases passed (N/A)
func countNicePairs(nums []int) int {
	rev := func(a int) int {
		result := 0
		for a%10 == 0 {
			a /= 10
		}
		for a > 0 {
			result *= 10
			result += a % 10
			a /= 10
		}
		return result
	}
	sort.Ints(nums)
	m := map[int]int{}
	for i := 0; i < len(nums); i++ {
		if i > 0 && nums[i] == nums[i-1] {
			nums[i] = nums[i-1]
		} else {
			nums[i] -= rev(nums[i])
		}
		m[nums[i]]++
	}
	count := 0
	mod := int(1e9 + 7)
	for _, v := range m {
		count += v * (v - 1) / 2
		count %= mod
	}
	return count
}



枚举所有组合可能 超时
Time Limit Exceeded
76/84 cases passed (N/A)

func countNicePairs(nums []int) int {
	rev := func(a int) int {
		result := 0
		for a%10 == 0 {
			a /= 10
		}
		for a > 0 {
			result *= 10
			result += a % 10
			a /= 10
		}
		return result
	}
	revs := make([]int, len(nums))
	for i := 0; i < len(nums); i++ {
		revs[i] = rev(nums[i])
	}
	count := 0
	mod := int(1e9 + 7)
	for i := 0; i < len(nums); i++ {
		for j := i + 1; j < len(revs); j++ {
			if nums[i]+revs[j] == nums[j]+revs[i] {
				count++
				count %= mod
			}
		}
	}
	return count
}
*/
