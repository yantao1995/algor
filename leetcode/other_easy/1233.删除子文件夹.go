package leetcode

import (
	"sort"
	"strings"
)

/*
 * @lc app=leetcode.cn id=1233 lang=golang
 *
 * [1233] 删除子文件夹
 */

// @lc code=start
func removeSubfolders(folder []string) []string {
	sort.Strings(folder)
	bit := make([]bool, len(folder))
	result := []string{}
	for i := 0; i < len(folder); i++ {
		if !bit[i] {
			result = append(result, folder[i])
			for j := i + 1; j < len(folder); j++ {
				if !bit[j] {
					if !strings.HasPrefix(folder[j], folder[i]) {
						break
					}
					if folder[j][len(folder[i])] == '/' {
						bit[j] = true
					}
				}
			}
		}
	}
	return result
}

// @lc code=end

/*

优化，替换每次的删除操作，用bit位来表示每一个元素是否删除
无前缀时直接跳过当前循环


排序后直接可以做前缀判断,然后删除对应的元素
超时 case 31/32
func removeSubfolders(folder []string) []string {
	sort.Strings(folder)
	result := []string{}
	for i := 0; i < len(folder); i++ {
		result = append(result, folder[i])
		for j := i + 1; j < len(folder); j++ {
			if strings.HasPrefix(folder[j], folder[i]) && folder[j][len(folder[i])] == '/' {
				folder = append(folder[:j], folder[j+1:]...)
				j--
			}
		}
	}
	return result
}
*/
