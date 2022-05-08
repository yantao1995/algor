package leetcode

/*
 * @lc app=leetcode.cn id=433 lang=golang
 *
 * [433] 最小基因变化
 */

// @lc code=start
func minMutation(start string, end string, bank []string) int {
	if start == end {
		return 0
	}
	relationM := map[string][]string{}
	canExchange := func(a, b string) bool {
		count := 0
		for i := 0; i < len(a) && count < 2; i++ {
			if a[i] != b[i] {
				count++
			}
		}
		return count == 1
	}
	bank = append(bank, start)
	for i := 0; i < len(bank); i++ {
		for j := 0; j < len(bank); j++ {
			if i != j {
				if canExchange(bank[i], bank[j]) {
					if _, ok := relationM[bank[i]]; !ok {
						relationM[bank[i]] = []string{}
					}
					relationM[bank[i]] = append(relationM[bank[i]], bank[j])
				}
			}
		}
	}
	result := -1
	distinct := map[string]bool{start: true}
	var dfs func(from string, count int)
	dfs = func(from string, count int) {
		if len(from) > 0 {
			if result != -1 && result <= count {
				return
			}
			if from == end {
				result = count
				return
			}
			for i := 0; i < len(relationM[from]); i++ {
				if !distinct[relationM[from][i]] {
					distinct[relationM[from][i]] = true
					dfs(relationM[from][i], count+1)
					distinct[relationM[from][i]] = false
				}
			}
		}
	}
	dfs(start, 0)
	return result
}

// @lc code=end

/*
	深度优先遍历 + 回溯的思想。
	先从bank与start得到 当前基因可以变成哪些基因，relationM 记录基因关系，避免重复计算
	distinct记录路径，防止走回头路
	然后从start开始进行dfs遍历，只要能连通到end，就表示可以生成
	剪枝：在找到result后，遇到count比当前result大，就可以直接退出
	感觉这道题用广度会比较简单一点，因为在当前层数找到了，就可以直接返回
*/
