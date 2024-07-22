package lcpr

/*
 * @lc app=leetcode.cn id=2101 lang=golang
 * @lcpr version=30204
 *
 * [2101] 引爆最多的炸弹
 */

// @lcpr-template-start

// @lcpr-template-end
// @lc code=start
func maximumDetonation(bombs [][]int) int {
	isBomb := func(a, b []int) (bool, bool) {
		dis2 := (a[0]-b[0])*(a[0]-b[0]) + (a[1]-b[1])*(a[1]-b[1])
		return dis2 <= a[2]*a[2], dis2 <= b[2]*b[2]
	}
	canBomb := make([][]bool, len(bombs))
	nums := make([]int, len(bombs))
	for i := 0; i < len(bombs); i++ {
		canBomb[i] = make([]bool, len(bombs))
		nums[i] = 1
	}

	for i := range bombs {
		canBomb[i][i] = true
		for j := i + 1; j < len(bombs); j++ {
			canBomb[i][j], canBomb[j][i] = isBomb(bombs[i], bombs[j])
		}
	}
	result := 0
	var dfs func(i int, road map[int]bool, count *int)
	dfs = func(i int, road map[int]bool, count *int) {
		for j := 0; j < len(bombs); j++ {
			if canBomb[i][j] && !road[j] {
				road[j] = true
				*count++
				dfs(j, road, count)
			}
		}
		if *count > result {
			result = *count
		}
	}
	for i := 0; i < len(bombs); i++ {
		dfs(i, map[int]bool{}, new(int))
	}
	return result
}

// @lc code=end

/*
// @lcpr case=start
// [[2,1,3],[6,1,4]]\n
// @lcpr case=end

// @lcpr case=start
// [[1,1,5],[10,10,5]]\n
// @lcpr case=end

// @lcpr case=start
// [[1,2,3],[2,3,1],[3,4,2],[4,5,3],[5,6,4]]\n
// @lcpr case=end

*/
