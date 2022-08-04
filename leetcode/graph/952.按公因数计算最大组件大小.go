package leetcode

import (
	"math"
)

/*
 * @lc app=leetcode.cn id=952 lang=golang
 *
 * [952] 按公因数计算最大组件大小
 */

// @lc code=start
func largestComponentSize(nums []int) int {
	result := 1
	pms, elds := []int{}, make([]bool, int(math.Sqrt(1e5))+1)
	for i, temp := 2, 0; i < len(elds); i++ {
		if !elds[i] {
			pms = append(pms, i)
			for temp = i; temp <= len(elds); temp += i {
				elds[temp] = true
			}
		}
	}
	us, sz := map[int]int{}, map[int]int{}
	for k := range pms {
		us[pms[k]] = pms[k]
	}
	var find func(i int) int
	find = func(i int) int {
		if us[i] != i {
			us[i] = find(us[i])
		}
		return us[i]
	}
	union := func(a, b int) {
		fa, fb := find(a), find(b)
		if fa == fb {
			return
		}
		sz[fa] += sz[fb]
		us[fb] = us[fa]
	}
	seq := []int{}
	for i := 0; i < len(nums); i++ {
		seq = seq[:0]
		for j := 0; j < len(pms) && nums[i] >= pms[j]; j++ {
			for nums[i]%pms[j] == 0 {
				nums[i] /= pms[j]
				if len(seq) == 0 || seq[len(seq)-1] != pms[j] {
					seq = append(seq, pms[j])
				}
			}
		}
		if nums[i] > 1 {
			if us[nums[i]] == 0 {
				us[nums[i]] = nums[i]
			}
			seq = append(seq, nums[i])
		}
		if len(seq) > 0 {
			sz[find(seq[0])]++
			for i := 1; i < len(seq); i++ {
				union(seq[0], seq[i])
			}
		}
	}
	for k := range sz {
		if sz[k] > result {
			result = sz[k]
		}
	}
	return result
}

// @lc code=end

/*

-------------- 解1

 因为要求大于1的公因数，索引只需要找公共的质因数
 使用厄拉多塞筛法找质数可以大大减少时间，因为厄拉多塞需要保存范围内的所有数
 而最大数 1e5 范围内，能组成的质数对的的数 最大至 math.Sqrt(1e5))，所以仅需找到
 的最大质数范围是至313，所以仅需要求出到313为止的质数，极大的缩小了厄拉多塞筛法所占的空间
 dfs重复路径会超时，所以可以使用并查集来做搜索，降具有相同质因数的 数作为同一个集合


pms 存储质数, elds 存储厄拉多塞筛法结果
us 保存并查集父节点，sz保存每个父节点的个数
seq 保存每个数的质数集合

us 映射  pms 质数的父节点为当前质数  例如 us[2]=2
每当遍历一个数 nums[i] 时候，给 seq[0] 做自增，因为可能有多个质数，所以只需要给一个数累加
if len(seq) > 0 {} 是为了排除  nums[i] ==1 的情况

-----------------------------------------
解2：

暴力，dfs暴搜路径 超时
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
