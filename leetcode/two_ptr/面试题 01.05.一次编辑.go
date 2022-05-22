package leetcode

func oneEditAway(first string, second string) bool {
	if len(first) == len(second) ||
		len(first)-len(second) == 1 ||
		len(second)-len(first) == 1 {
		count, fi, si := 0, 0, 0
		for fi < len(first) && si < len(second) {
			if first[fi] != second[si] {
				if len(first) > len(second) {
					si--
				} else if len(first) < len(second) {
					fi--
				}
				count++
			}
			fi++
			si++
		}
		return count < 2
	}
	return false
}

/*
	双指针， fi和 si 分别指向两个str
	然后依次向后移动，如果当前两个字符不相同，那么len小的
	指针向后移动，然后记录不相同的个数，小于2的为符合条件的
*/
