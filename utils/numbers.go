package utils

import (
	"algor/vals"
	"math/rand"
	"time"
)

// NonceSliceCreate 生成随机指定大小的切片数组，相邻两个不重复， 纯数字,用于排序测试
// 长度   范围   数据过大会造成循环等待
func NonceSliceCreate(len int, rangee int) (at []vals.AlgorType) {
	if len <= 0 {
		return at
	}
	last, rd := vals.IntMin, vals.IntMin
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	at = make([]vals.AlgorType, len)
	for i := 0; i < len; i++ {
		for last == rd {
			rd = r.Intn(rangee)
		}
		at[i] = rd
		last = rd
	}
	return at
}

// NonceSliceCreateRepate 生成随机指定大小的切片数组，相邻两个可重复， 纯数字,用于排序测试
// 长度   范围
func NonceSliceCreateRepate(len int, rangee int) (at []vals.AlgorType) {
	if len <= 0 {
		return at
	}
	at = make([]vals.AlgorType, len)
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	for i := 0; i < len; i++ {
		rd := r.Intn(rangee)
		at[i] = rd
	}
	return at
}

//GetPrimeNumberByLimit 获取一个数和0之间的所有素数  素数增量序列,大——>小
func GetPrimeNumberByLimit(max int) []int {
	if max < 2 {
		return []int{}
	}
	primeNumber := []int{}
	for max >= 2 {
		primeFlag := false
		for i := 2; i < max; i++ {
			if max%i == 0 {
				primeFlag = true
				break
			}
		}
		if !primeFlag {
			primeNumber = append(primeNumber, max)
		}
		max--
	}
	return primeNumber
}
