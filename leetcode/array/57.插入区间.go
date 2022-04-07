package leetcode

/*
 * @lc app=leetcode.cn id=57 lang=golang
 *
 * [57] 插入区间
 */

// @lc code=start
func insert(intervals [][]int, newInterval []int) [][]int {
	if len(intervals) == 0 {
		return [][]int{newInterval}
	}
	flagMap := map[int]int{}
	max, min := intervals[len(intervals)-1][1], intervals[0][0]
	if newInterval[1] > max {
		max = newInterval[1]
	}
	if newInterval[0] < min {
		min = newInterval[0]
	}
	index1, index2 := newInterval[0], newInterval[1]
	for _, v := range intervals {
		flagMap[v[0]]++
		flagMap[v[1]]++
	}
	for i := index1 + 1; i < index2; i++ {
		delete(flagMap, i)
	}
	flagCheck1 := true
	if index1 != index2 || flagMap[index1] == 0 {

		flagMap[index1] = 1
		if index1 != index2 {
			flagMap[index2] = 1
		} else {
			flagMap[index2]++
		}
	} else {
		flagMap[index1] = 1
		flagCheck1 = false
	}
	// 前半段
	if flagCheck1 {
		prevCount, prevLast := 0, 0
		for i := min; i < index1; i++ {
			for m := flagMap[i]; m > 0; m-- {
				prevLast = i
				prevCount++
			}
		}
		if prevCount%2 == 1 {
			delete(flagMap, index1)
			index1 = prevLast
		}

		// 后半段
		tailCount, tailFirst := 0, min-1
		for i := index2 + 1; i <= max; i++ {
			for m := flagMap[i]; m > 0; m-- {
				if tailFirst == min-1 {
					tailFirst = i
				}
				tailCount++
			}
		}
		if tailCount%2 == 1 {
			delete(flagMap, index2)
			index2 = tailFirst
		}
	}
	count := 0
	result := [][]int{}
	last := min - 2
	for min <= max {
		for m := flagMap[min]; m > 0; m-- {
			count++
			if count%2 == 0 {
				temp := []int{last, min}
				result = append(result, temp)
			} else {
				last = min
			}
		}
		min++
	}
	return result
}

// @lc code=end
