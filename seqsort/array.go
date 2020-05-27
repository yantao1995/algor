package seqsort

import (
	"algor/vals"
	"fmt"
	"time"
)

/*	n:数据规模   k:桶的个数
排序方法	时间复杂度（平均）	时间复杂度（最坏）	时间复杂度（最好）	空间复杂度			 稳定性		内存占用
选择排序		O(n2)				O(n2)				O(n2)				O(1)			不稳定		x
希尔排序		O(n1.3)				O(n2)				O(n)				O(1)			不稳定		x
快速排序		O(nlogn)			O(n2)				O(nlogn)			O(nlogn)		不稳定		x
堆排序			O(nlogn)			O(nlogn)			O(nlogn)			O(1)			不稳定		x
冒泡排序		O(n2)				O(n2)				O(n)				O(1)			稳定		x
插入排序		O(n2)				O(n2)				O(n)				O(1)			稳定		x
归并排序		O(nlogn)			O(nlogn)			O(nlogn)			O(n)			稳定        额外内存
|非|基数排序	O(n∗k)				O(n∗k)				 O(n∗k)	     	    O(n+k)			 稳定		额外内存
|比|桶排序		O(n+k)				O(n2)				O(n+k)				O(n+k)			稳定		额外内存
|较|计数排序	O(n+k)				O(n+k)				O(n+k)				O(k)			稳定		额外内存
*/

/*
桶排序：首先，数据应该分布比较均匀。一种较坏的情况，如果数据全部都被分到一个桶里，那么桶排序的时间复杂度就退化到O(nlogn)了.
		其次，要排序的数据应该很容易分成m个桶，每个桶也应该有大小顺序。
计数排序： 其中k是整数的范围，快于任何比较排序算法。这是一种牺牲空间换取时间的做法，
		而且当O(k)>O(n*log(n))的时候其效率反而不如基于比较的排序
*/

/*
堆是具有以下性质的完全二叉树：
1.每个结点的值都大于或等于其左右孩子结点的值，称为大顶堆
2.或者每个结点的值都小于或等于其左右孩子结点的值，称为小顶堆
再数组中时满足性质：
大顶堆：arr[i] >= arr[2i+1] && arr[i] >= arr[2i+2]
小顶堆：arr[i] <= arr[2i+1] && arr[i] <= arr[2i+2]
*/

// exchangeIJ 切片交换两个元素
func exchangeIJ(nums []vals.AlgorType, i, j int) {
	nums[i], nums[j] = nums[j], nums[i]
}

//QuickSort  挖坑+分治
//每个块选中第0个元素,左哨兵交换了就i++，直到左哨兵被交换出去就j--。分别从两头向中间移动
func QuickSort(nums []vals.AlgorType, low, high int) {
	if low >= high {
		return
	}
	i := low + 1
	j := high
	sufftoprev := true
	index := low
	for i <= j {
		if sufftoprev { //坑在前面，哨兵从后往前
			if vals.GreaterThanInt(nums[index], nums[j]) {
				exchangeIJ(nums, index, j)
				index = j
				sufftoprev = false
			}
			j--
		} else { //坑在后面，哨兵从前往后
			if vals.LessThanInt(nums[index], nums[i]) {
				exchangeIJ(nums, index, i)
				index = i
				sufftoprev = true
			}
			i++
		}
	}
	QuickSort(nums, low, index-1)
	QuickSort(nums, index+1, high)
}

//InsertSort 直接插入排序 O(n^2)  正序 O(n)
func InsertSort(nums []vals.AlgorType) {
	for i := 1; i < len(nums); i++ { //依次挑选未排序序列，默认第1个有序，第2个元素开始
		for j := 0; j < i; j++ { //遍历已排序好的序列
			if vals.LessThanInt(nums[i], nums[j]) { //小于nums[j]时就放到j的位置
				temp := nums[i]
				for k := i; k > j; k-- { //元素后移一位
					nums[k] = nums[k-1]
				}
				nums[j] = temp
			}
		}
	}
}

//InsertSortOptimization 直接插入优化
func InsertSortOptimization(nums []vals.AlgorType) {
	for i := 1; i < len(nums); i++ { //依次遍历未排序数据
		for j := i - 1; j >= 0 && vals.LessThanInt(nums[j+1], nums[j]); j-- { //当前数据从后往前冒泡
			exchangeIJ(nums, j+1, j)
		}
	}
}

//InsertSortByBinarySearch 折半插入排序 O(n^2)  相比InsertSort减少了查找时的时间消耗
func InsertSortByBinarySearch(nums []vals.AlgorType) {
	for i := 1; i < len(nums); i++ { //依次挑选未排序序列，默认第1个有序，第2个元素开始
		low, high := 0, i-1
		mid := 0
		for low <= high {
			mid = (high + low) / 2
			if vals.LessEquelsThanInt(nums[mid], nums[i]) {
				low = mid + 1
			} else {
				high = mid - 1
			}
		}
		temp := nums[i]
		for k := i; k > low; k-- { //元素后移一位
			nums[k] = nums[k-1]
		}
		nums[low] = temp
	}
}

//ShellSort 希尔排序 （直接插入基础上）  素数增量时 O(n^1.5)
func ShellSort(nums []vals.AlgorType, increment int) {
	for ; increment > 0; increment /= 2 { //增量递减
		for i := 0; i+increment < len(nums) && i < increment; i++ {
			for j := i + increment; j < len(nums); j += increment {
				for m := i; m < j; m += increment {
					if vals.LessThanInt(nums[j], nums[m]) {
						temp := nums[j]
						for k := j; k-increment >= m; k -= increment {
							nums[k] = nums[k-increment]
						}
						nums[m] = temp
					}
				}
			}
		}
	}
}

//ShellSortOptimization 希尔排序优化（直接插入优化基础上）   素数增量时 O(n^1.5)
func ShellSortOptimization(nums []vals.AlgorType, increment int) {
	for ; increment > 0; increment /= 2 { //增量递减
		for i := increment; i < len(nums); i++ { //从第一个递增量开始
			for j := i; j-increment >= 0 && vals.LessThanInt(nums[j], nums[j-increment]); j -= increment { //从后往前冒泡
				exchangeIJ(nums, j-increment, j)
			}
		}
	}
}

//SelectSort 选择排序  和冒泡区分
func SelectSort(nums []vals.AlgorType) {
	for i := 0; i < len(nums)-1; i++ {
		for j := i + 1; j < len(nums); j++ {
			if vals.LessThanInt(nums[j], nums[i]) {
				exchangeIJ(nums, i, j)
			}
		}
	}
}

//BubbleSort 冒泡排序   和选择区分
func BubbleSort(nums []vals.AlgorType) {
	for i := 0; i < len(nums)-1; i++ {
		for j := 0; j < len(nums)-1-i; j++ {
			if vals.LessThanInt(nums[j+1], nums[j]) {
				exchangeIJ(nums, j+1, j)
			}
		}
	}
}

//MergeSort 归并排序  分治 ,先分，后合    有临时数组。 空间换时间  耗时少
func MergeSort(nums, temp []vals.AlgorType, low, high int) {
	if low < high {
		mid := (low + high) / 2
		MergeSort(nums, temp, low, mid)
		MergeSort(nums, temp, mid+1, high)
		merge(nums, temp, low, high)
	}
}

//MergeSort 归并 合  空间换时间   耗时少
func merge(nums, temp []vals.AlgorType, low, high int) {
	tempt, lowt, midt, mid := low, low, (low+high)/2+1, (low+high)/2
	for tempt <= high {
		if lowt <= mid && midt <= high { //判断条件可以合并成 1个 if{}else{}
			if vals.LessThanInt(nums[lowt], nums[midt]) {
				temp[tempt] = nums[lowt]
				lowt++
			} else {
				temp[tempt] = nums[midt]
				midt++
			}
		} else if lowt > mid {
			temp[tempt] = nums[midt]
			midt++
		} else {
			temp[tempt] = nums[lowt]
			lowt++
		}
		tempt++
	}
	for ; low <= high; low++ {
		nums[low] = temp[low]
	}
}

// 合并ifelse  ---
func merge2(nums, temp []vals.AlgorType, low, high int) {
	tempt, lowt, midt, mid := low, low, (low+high)/2+1, (low+high)/2
	for tempt <= high {
		if lowt <= mid && midt <= high { //判断条件可以合并成 1个 if{}else{}
			if vals.LessThanInt(nums[lowt], nums[midt]) {
				temp[tempt] = nums[lowt]
				lowt++
			} else {
				temp[tempt] = nums[midt]
				midt++
			}
		} else if lowt > mid {
			temp[tempt] = nums[midt]
			midt++
		} else if midt > high {
			temp[tempt] = nums[lowt]
			lowt++
		}
		tempt++
	}
	for ; low <= high; low++ {
		nums[low] = temp[low]
	}
}

//MergeSortNoTemp 归并排序 无临时数组 耗时更多
func MergeSortNoTemp(nums []vals.AlgorType, low, high int) {
	if low < high {
		mid := (low + high) / 2
		MergeSortNoTemp(nums, low, mid)
		MergeSortNoTemp(nums, mid+1, high)
		mergeNoTemp(nums, low, high)
	}
}

//MergeSortNoTemp 归并无临时数组 合    耗时更多
func mergeNoTemp(nums []vals.AlgorType, low, high int) {
	mid := (low + high) / 2
	//插排思路
	for mid < high { //后驱块数据依次插入前驱块数据
		for j := mid + 1; j > low && vals.LessThanInt(nums[j], nums[j-1]); j-- { //后驱块数据冒泡到前驱块中
			exchangeIJ(nums, j, j-1)
		}
		mid++
	}
}

//大顶堆：arr[i] >= arr[2i+1] && arr[i] >= arr[2i+2]
//HeapSort 堆排序  升序[大顶堆]    降序[小顶堆]
func HeapSort(nums []vals.AlgorType) {
	heapBigTopInit(nums)
	// for j := len(nums) - 1; j > 0; j-- {
	// 	for i := 0; (2*i + 2) <= j; i++ { //构造堆
	// 		if vals.LessThanInt(nums[i], nums[2*i+1]) {
	// 			exchangeIJ(nums, i, 2*i+1)
	// 		}
	// 		if vals.LessThanInt(nums[i], nums[2*i+2]) {
	// 			exchangeIJ(nums, i, 2*i+2)
	// 		}
	// 	}
	// 	exchangeIJ(nums, 0, j)
	// }
}

//HeapSort 大顶堆初始化
func heapBigTopInit(nums []vals.AlgorType) {
	for i := 1; i < len(nums)-1; i++ {
		for j := i; j > 0; j = (j - 1) / 2 {
			if vals.LessThanInt(nums[j], nums[(j-1)/2]) {
				exchangeIJ(nums, j, (j-1)/2)
			}
		}
	}
}

//HeapSort 大顶堆调整
func heapBigTopAdjust(nums []vals.AlgorType, index int) {
	for i := 0; i < len(nums); i++ {
		changed := false
		if (2*i+1) < (len(nums)-1) && vals.LessThanInt(nums[i], nums[2*i+1]) {
			exchangeIJ(nums, i, 2*i+1)
			changed = true
		}
		if (2*i+2) < (len(nums)-1) && vals.LessThanInt(nums[i], nums[2*i+2]) {
			exchangeIJ(nums, i, 2*i+2)
			changed = true
		}
		if changed {
			for j := i; (j-1)/2 > 0; j = (j - 1) / 2 {
				if vals.LessThanInt(nums[(j-1)/2], nums[j]) {
					exchangeIJ(nums, j, (j-1)/2)
				}
			}
		}
	}
}

func TestSorts() {
	nums2 := []vals.AlgorType{60, 10, 10, 42, 38, 80, 5, 58, 74, 49, 1, 86, 44, 38, 12, 13, 8, 9, 20, 24, 48, 22, 60, 53}
	//nums = nums2
	// QuickSort(nums, 0, len(nums)-1)
	// fmt.Println("快排：", nums)

	//start0 := time.Now().UnixNano() / 1e6
	////nums2 := utils.NonceSliceCreate(50000, 500000) //相邻不重复  耗时高
	//nums2 := utils.NonceSliceCreateRepate(30, 300) //相邻可重复
	// end0 := time.Now().UnixNano() / 1e6
	// fmt.Println("    创建切片耗时：", end0-start0, "毫秒")

	//var nums3 []vals.AlgorType = make([]vals.AlgorType, len(nums2))
	// var nums4 []vals.AlgorType = make([]vals.AlgorType, len(nums2))
	// var nums5 []vals.AlgorType = make([]vals.AlgorType, len(nums2))
	// var nums6 []vals.AlgorType = make([]vals.AlgorType, len(nums2))
	//var nums7 []vals.AlgorType = make([]vals.AlgorType, len(nums2))
	// var nums8 []vals.AlgorType = make([]vals.AlgorType, len(nums2))
	//var nums9 []vals.AlgorType = make([]vals.AlgorType, len(nums2))
	var nums10 []vals.AlgorType = make([]vals.AlgorType, len(nums2))
	//var nums11 []vals.AlgorType = make([]vals.AlgorType, len(nums2))
	var nums12 []vals.AlgorType = make([]vals.AlgorType, len(nums2))
	var nums13 []vals.AlgorType = make([]vals.AlgorType, len(nums2))
	var nums14 []vals.AlgorType = make([]vals.AlgorType, len(nums2))
	//copy(nums3, nums2)
	// copy(nums4, nums2)
	// copy(nums5, nums2)
	// copy(nums6, nums2)
	//copy(nums7, nums2)
	// copy(nums8, nums2)
	//copy(nums9, nums2)
	copy(nums10, nums2)
	//copy(nums11, nums2)
	copy(nums12, nums2)
	copy(nums13, nums2)
	copy(nums14, nums2)

	// fmt.Println("---start---")
	// //fmt.Println("    原始序列：", nums2)

	// start2 := time.Now().UnixNano() / 1e6
	// QuickSort(nums2, 0, len(nums2)-1)
	// end2 := time.Now().UnixNano() / 1e6
	// fmt.Println("    快速排序耗时：", end2-start2, "毫秒")
	// //fmt.Println("    快速排序：", nums2)

	// start3 := time.Now().UnixNano() / 1e6
	// InsertSort(nums3)
	// end3 := time.Now().UnixNano() / 1e6
	// fmt.Println("直接插入排序耗时：", end3-start3, "毫秒")
	// //fmt.Println("直接插入排序：", nums3)

	// start7 := time.Now().UnixNano() / 1e6
	// InsertSortOptimization(nums7)
	// end7 := time.Now().UnixNano() / 1e6
	// fmt.Println("直接插入优化耗时：", end7-start7, "毫秒")
	// //fmt.Println("直接插入优化：", nums3)

	// start4 := time.Now().UnixNano() / 1e6
	// InsertSortByBinarySearch(nums4)
	// end4 := time.Now().UnixNano() / 1e6
	// fmt.Println("折半插入排序耗时：", end4-start4, "毫秒")
	// // fmt.Println("折半插入排序：", nums4)

	// incrementArray := utils.GetPrimeNumberByLimit(len(nums5)) //获取素数增量序列,大——>小
	// start5 := time.Now().UnixNano() / 1e6
	// ShellSort(nums5, incrementArray[0])
	// end5 := time.Now().UnixNano() / 1e6
	// fmt.Println("希尔增量排序耗时：", end5-start5, "毫秒")
	// //fmt.Println("希尔增量排序：", nums5)

	// incrementArray2 := utils.GetPrimeNumberByLimit(len(nums6)) //获取素数增量序列,大——>小
	// start6 := time.Now().UnixNano() / 1e6
	// ShellSortOptimization(nums6, incrementArray2[0])
	// end6 := time.Now().UnixNano() / 1e6
	// fmt.Println("希尔增量优化耗时：", end6-start6, "毫秒")
	// //fmt.Println("希尔增量优化：", nums6)

	// start8 := time.Now().UnixNano() / 1e6
	// SelectSort(nums8)
	// end8 := time.Now().UnixNano() / 1e6
	// fmt.Println("    选择排序耗时：", end8-start8, "毫秒")
	// //fmt.Println("    选择排序：", nums8)

	// start9 := time.Now().UnixNano() / 1e6
	// BubbleSort(nums9)
	// end9 := time.Now().UnixNano() / 1e6
	// fmt.Println("    冒泡排序耗时：", end9-start9, "毫秒")
	// fmt.Println("    冒泡排序：", nums9)

	start10 := time.Now().UnixNano() / 1e6
	temp := make([]vals.AlgorType, len(nums10))
	MergeSort(nums10, temp, 0, len(nums10)-1)
	end10 := time.Now().UnixNano() / 1e6
	fmt.Println("归并排序耗时：", end10-start10, "毫秒")
	fmt.Println("    归并排序：", nums10)

	// start11 := time.Now().UnixNano() / 1e6
	// MergeSortNoTemp(nums11, 0, len(nums11)-1)
	// end11 := time.Now().UnixNano() / 1e6
	// fmt.Println("归并无temp耗时：", end11-start11, "毫秒")
	// //fmt.Println("归并无temp排序：", nums11)

	start12 := time.Now().UnixNano() / 1e6
	HeapSort(nums12)
	end12 := time.Now().UnixNano() / 1e6
	fmt.Println("大顶堆排序耗时：", end12-start12, "毫秒")
	fmt.Println("    大顶堆排序：", nums12)

}
