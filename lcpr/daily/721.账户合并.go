package lcpr

import "sort"

/*
 * @lc app=leetcode.cn id=721 lang=golang
 * @lcpr version=30204
 *
 * [721] 账户合并
 */

// @lcpr-template-start

// @lcpr-template-end
// @lc code=start
func accountsMerge(accounts [][]string) [][]string {
	result := [][]string{}
	name := map[string]string{}
	us := map[string]string{}

	//初始化并查集
	for k := range accounts {
		for _, ac := range accounts[k][1:] {
			us[ac] = ac
			name[ac] = accounts[k][0]
		}
	}

	var find func(i string) string
	find = func(i string) string {
		if us[i] != i {
			us[i] = find(us[i])
		}
		return us[i]
	}

	union := func(a, b string) {
		us[find(b)] = find(a)
	}

	//组内连通
	for _, account := range accounts {
		for _, mail := range account[1:] {
			union(account[1], mail)
		}
	}

	//更新并查集
	for k, v := range us {
		us[k] = find(v)
	}

	//归集
	m := map[string][]string{}
	for child, parent := range us {
		if len(m[parent]) == 0 {
			m[parent] = []string{name[parent]}
		}
		m[parent] = append(m[parent], child)
	}
	for k := range m {
		sort.Strings(m[k][1:])
		result = append(result, m[k])
	}

	return result
}

// @lc code=end

/*
// @lcpr case=start
// [["John", "johnsmith@mail.com", "john00@mail.com"], ["John", "johnnybravo@mail.com"], ["John","johnsmith@mail.com", "john_newyork@mail.com"], ["Mary", "mary@mail.com"]]\n
// @lcpr case=end

// @lcpr case=start
// [["Gabe","Gabe0@m.co","Gabe3@m.co","Gabe1@m.co"],["Kevin","Kevin3@m.co","Kevin5@m.co","Kevin0@m.co"],["Ethan","Ethan5@m.co","Ethan4@m.co","Ethan0@m.co"],["Hanzo","Hanzo3@m.co","Hanzo1@m.co","Hanzo0@m.co"],["Fern","Fern5@m.co","Fern1@m.co","Fern0@m.co"]]\n
// @lcpr case=end

*/

/*
	使用并查集来连通多个邮箱数组，更新并查集(让多个独立的邮箱数组合并)，归集处理结果

*/
