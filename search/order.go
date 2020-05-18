package search

import (
	"algor/vals"
)

//SeqBinarySearch 有序表二分查找 ,返回查找结果和index下标
func SeqBinarySearch(seqs []vals.AlgorType, value vals.AlgorType) (bool, int) {
	low, high := 0, len(seqs)-1
	for low <= high {
		mid := (high + low) / 2
		//println("low: ", low, " mid:", mid, " high:", high)
		if seqs[mid] == value {
			return true, mid
		} else if vals.GetInt(seqs[mid]) > vals.GetInt(value) {
			high = mid - 1
		} else {
			low = mid + 1
		}
	}
	return false, 0
}

func TestSearch() {
	// seqs := []vals.AlgorType{1, 3, 5, 6, 7, 9, 12, 15, 24, 31, 35}
	// isExists, index := SeqBinarySearch(seqs, 3)
	// fmt.Println("isExists:", isExists, " index:", index)
}
