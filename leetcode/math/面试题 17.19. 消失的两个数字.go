package leetcode

/*
给定一个数组，包含从 1 到 N 所有的整数，但其中缺了两个数字。你能在 O(N) 时间内只用 O(1) 的空间找到它们吗？
以任意顺序返回这两个数字均可。

示例 1:
输入: [1]
输出: [2,3]

示例 2:
输入: [2,3]
输出: [1,4]
提示：
nums.length <= 30000

*/
func missingTwo(nums []int) []int {
	n := len(nums) + 2
	sumN := 0
	for i := 0; i < len(nums); i++ {
		sumN += nums[i]
	}
	sumI := n * (1 + n) / 2
	sum := sumI - sumN
	half := (sum + 1) / 2 // 必然有一个数小于half，一个数大于half
	sumH := 0
	for i := 0; i < len(nums); i++ {
		if nums[i] < half {
			sumH += nums[i]
		}
	}
	sumH2 := 0
	for i := 1; i < half; i++ {
		sumH2 += i
	}
	return []int{sumH2 - sumH, sum - (sumH2 - sumH)}
}

/*
	先找出两个数的和
	因为两个数肯定不一样，那么有一个数必然小于 和的一半(half)
	可以根据 half 来确定小的那个数是多少
*/
