package search

import (
	"algor/achieve/tree"
	"algor/achieve/vals"
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

//BSTSearch 二叉搜索树
func BSTSearch(root *tree.BSTBinary, key vals.AlgorType) *tree.BSTBinary {
	if root != nil {
		if vals.EquelsInt(key, root.Data.Key) {
			return root
		} else if vals.LessThanInt(key, root.Data.Key) { //小于
			return BSTSearch(root.Left, key)
		} else {
			return BSTSearch(root.Right, key)
		}
	} else {
		return nil
	}
}

func TestSearch() {
	// seqs := []vals.AlgorType{1, 3, 5, 6, 7, 9, 12, 15, 24, 31, 35}
	// isExists, index := SeqBinarySearch(seqs, 3)
	// fmt.Println("isExists:", isExists, " index:", index)

	/*  测试结构
			  10a
		6b_____|_____15f
	2c___|___8d   13g___|___17h

	*/
	// bstroot := tree.BSTInitTestSearch()
	// tree.FirstErgodicBST(bstroot)
	// println()
	// node := BSTSearch(bstroot, 10)
	// if node != nil {
	// 	fmt.Println(node.Data.Value)
	// } else {
	// 	fmt.Println(node)
	// }

}
