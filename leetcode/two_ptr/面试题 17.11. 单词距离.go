package leetcode

func findClosest(words []string, word1 string, word2 string) int {
	minLength := len(words)
	min := func(a, b int) int {
		if a < b {
			return a
		}
		return b
	}
	p1, p2 := -1, -1
	for k := range words {
		if words[k] == word1 {
			p1 = k
			if p1 > 0 && p2 > 0 {
				minLength = min(minLength, p1-p2)
			}
		}
		if words[k] == word2 {
			p2 = k
			if p1 > 0 && p2 > 0 {
				minLength = min(minLength, p2-p1)
			}
		}
	}
	return minLength
}

/*
	双指针记录第一个单词和第二个单词的位置
	每次只要有一个匹配上了，然后给就与当前的位置做比较
	然后将 minLength 与当前两个指针的位置做比较，得到最小距离
*/
