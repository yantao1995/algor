package leetcode

import (
	"sort"
)

/*
 * @lc app=leetcode.cn id=621 lang=golang
 *
 * [621] 任务调度器
 */

// @lc code=start

func leastInterval(tasks []byte, n int) int {
	if len(tasks) == 0 || n == 0 {
		return len(tasks)
	}
	m := map[byte]int{}
	keys := []byte{}
	for k := range tasks {
		m[tasks[k]]++
	}
	for k := range m {
		keys = append(keys, k)
	}
	sort.Slice(keys, func(i, j int) bool {
		return m[keys[i]] > m[keys[j]]
	})
	//
	result := []byte{}
	min := len(tasks) * n * 2
	m[keys[0]]--
	var greedy func(total, onlyChooseLength int, result []byte, val byte)
	greedy = func(total, onlyChooseLength int, result []byte, val byte) {
		//补位
		for ; onlyChooseLength < n; onlyChooseLength++ {
			result = append(result, 'a')
		}
		//添加元素
		result = append(result, val)
		//fmt.Println(string(result))
		if total == len(tasks) {
			if len(result) < min {
				min = len(result)
			}
			return
		}
		onlyChooseLength = 0                 //当前的选择所占的位置个数
		maxMapCount := 0                     //当前map中元素还剩的个数
		currentAnswerIndexSlice := []int{}   //当前轮的所有最优解
		notBigLength, notBigMapCount := 0, 0 //大于的没有，就启用这个
		//找当前最优的全部解的特征值 大于 和 maxMapCount
		for i := len(keys) - 1; i >= 0; i-- {
			if m[keys[i]] > 0 {
				//当前不需要待命或者待命最小
				currentLength := 0
				last := len(result) - 1
				for ; last >= 0; last-- {
					if result[last] == keys[i] {
						break
					}
					currentLength++
					if currentLength >= n {
						break
					}
				}
				if last == -1 || currentLength >= n { //到第0个都没有
					if m[keys[i]] > maxMapCount {
						maxMapCount = m[keys[i]]
					}
					onlyChooseLength = n
				} else {
					if currentLength >= notBigLength || m[keys[i]] >= notBigMapCount {
						notBigLength = currentLength
						notBigMapCount = m[keys[i]]
					}
				}
			}
		}
		//fmt.Println(maxMapCount, onlyChooseLength, notBigMapCount, notBigLength)
		if maxMapCount == 0 || onlyChooseLength == 0 {
			maxMapCount, onlyChooseLength = notBigMapCount, notBigLength
		}
		//根据特征值找当前最优的全部解
		for i := len(keys) - 1; i >= 0; i-- {
			if m[keys[i]] > 0 {
				//当前不需要待命或者待命最小
				currentLength := 0
				last := len(result) - 1
				for ; last >= 0; last-- {
					if result[last] == keys[i] {
						break
					}
					currentLength++
					if currentLength >= n {
						break
					}
				}
				if (currentLength >= n || last == -1 || currentLength == onlyChooseLength) && m[keys[i]] == maxMapCount {
					currentAnswerIndexSlice = append(currentAnswerIndexSlice, i)
				}
			}
		}
		//fmt.Println("maxMapCount", maxMapCount, "onlyChooseLength",
		//	onlyChooseLength, "currentAnswerIndexSlice", currentAnswerIndexSlice)
		for index := range currentAnswerIndexSlice {
			m[keys[currentAnswerIndexSlice[index]]]--
			greedy(total+1, onlyChooseLength, result, keys[currentAnswerIndexSlice[index]])
			if maxMapCount == 1 {
				break
			}
			m[keys[currentAnswerIndexSlice[index]]]++
		}
	}
	greedy(1, n, result, keys[0])
	return min
}

// @lc code=end
