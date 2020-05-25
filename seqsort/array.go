package seqsort

import (
	"algor/utils"
	"algor/vals"
	"fmt"
)

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

//ShellSort 希尔排序   素数增量时 O(n^1.5)
func ShellSort(nums []vals.AlgorType) {
	incrementArray := utils.GetPrimeNumberByLimit(len(nums)) //获取素数增量序列,大——>小
	increment := incrementArray[0]
	for ; increment > 0; increment /= 2 { //增量递减
		fmt.Println("increment:", increment)
		for i := 0; i+increment < len(nums) && i < increment; i++ {
			fmt.Print(" i:", i)
			for j := i; j < len(nums); j += increment {
				fmt.Print(" j:", j)
				for k := i + increment; k < len(nums); k += increment {
					if vals.LessThanInt(nums[k], nums[j]) {
						temp := nums[k]
						for l := k; l > increment; l -= increment {
							nums[l] = nums[l-increment]
						}
						nums[j] = temp
					}
				}
			}
			print("\t")
		}
		println()
	}
}

func TestSorts() {
	// nums := []vals.AlgorType{60, 10, 42, 80, 5, 58, 74, 49, 86, 44, 38, 12, 13, 8, 9, 20, 24, 48, 22, 60, 53}
	// QuickSort(nums, 0, len(nums)-1)
	// fmt.Println("快排：", nums)

	nums2 := utils.NonceSliceCreate(10, 100)
	fmt.Println("    原始序列：", nums2)

	// QuickSort(nums2, 0, len(nums2)-1)
	// fmt.Println("    快速排序：", nums2)

	// InsertSort(nums2)
	// fmt.Println("直接插入排序：", nums2)

	// InsertSortByBinarySearch(nums2)
	// fmt.Println("折半插入排序：", nums2)

	ShellSort(nums2)
	fmt.Println("希尔增量排序：", nums2)
}
