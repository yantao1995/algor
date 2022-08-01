package leetcode

import "fmt"

/*
 * @lc app=leetcode.cn id=952 lang=golang
 *
 * [952] 按公因数计算最大组件大小
 */

// @lc code=start
func largestComponentSize(nums []int) int {
	maxNum := 0
	um := []int{}       //并查集
	sz := map[int]int{} //大小
	for k := range nums {
		if nums[k] > maxNum {
			maxNum = nums[k]
		}
	}
	result := 0
	pm := make([]bool, maxNum+1)
	pms := []int{}
	for i, temp := 2, 0; i < len(pm); i++ {
		if !pm[i] {
			pms = append(pms, i)
			for temp = i; temp <= maxNum; temp += i {
				pm[temp] = true
			}
		}
	}
	max := func(a, b int) int {
		if a > b {
			return a
		}
		return b
	}

	var find func(n int) int
	find = func(n int) int {
		if um[n] != n {
			um[n] = find(um[n])
		}
		return um[n]
	}
	union := func(a, b int) {
		if find(a) != find(b) {
			return
		}
		sz[find(a)] += sz[find(b)]
		um[find(b)] = um[find(a)]
		result = max(result, sz[find(a)])
	}
	fmt.Println(pms)
	temp := 0
	for i := 0; i < len(nums); i++ {
		temp = nums[i]
		for j := 0; j < len(pms) && pms[j] <= temp; j++ {
			for temp%pms[j] == 0 {
				um[pms[j]] = i
				temp /= pms[j]
			}
		}
		if temp > 1 {
			um[temp] = i
		}
	}
	fmt.Println(um)
	for i := 0; i < len(nums); i++ {
		for j := 0; j < len(nums); j++ {
			if i != j {
				union(nums[i], nums[j])
			}
		}
	}
	return result
}

// @lc code=end

/*



暴力，dfs 超时
厄拉多塞筛法，找到2到最大值的所有素数，然后构建每个nums之间的图关系
然后dfs搜图，找到连接的最大长度

func largestComponentSize(nums []int) int {
	maxNum := 0
	for k := range nums {
		if nums[k] > maxNum {
			maxNum = nums[k]
		}
	}
	result := 0
	pm := make([]bool, maxNum+1)
	pms := []int{}
	for i, temp := 2, 0; i < len(pm); i++ {
		if !pm[i] {
			pms = append(pms, i)
			for temp = i; temp <= maxNum; temp += i {
				pm[temp] = true
			}
		}
	}
	check := func(a, b int) bool {
		for i := 0; i < len(pms) && pms[i] <= a && pms[i] <= b; i++ {
			if a%pms[i] == 0 && b%pms[i] == 0 {
				return true
			}
		}
		return false
	}
	graph := map[int][]int{}
	for i := 0; i < len(nums); i++ {
		for j := 0; j < len(nums); j++ {
			if i != j && check(nums[i], nums[j]) {
				graph[nums[i]] = append(graph[nums[i]], nums[j])
			}
		}
	}
	//fmt.Println(graph)
	count := 1
	distinct := map[int]bool{}
	var dfs func(i int, count *int)
	dfs = func(i int, count *int) {
		if *count > result {
			result = *count
		}
		for k := range graph[i] {
			if !distinct[graph[i][k]] {
				*count++
				distinct[graph[i][k]] = true
				dfs(graph[i][k], count)
			}
		}
	}
	for k := range nums {
		distinct[nums[k]] = true
		count = 1
		dfs(nums[k], &count)
	}
	return result
}
*/
