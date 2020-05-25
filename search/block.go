package search

import (
	"algor/vals"
	"fmt"
)

//BlockIndex 分块查找，索引块
type BlockIndex struct {
	Start    int            //开始位置
	MaxValue vals.AlgorType //块内最大值
}

//GetBlockIndexByOrder  对顺序表进行分块处理 并返回索引块
//blockSize 分块的大小 非正常输入默认1
func GetBlockIndexByOrder(blockSize int, at []vals.AlgorType) []BlockIndex {
	if blockSize < 1 || blockSize > len(at) {
		blockSize = 1
	}
	everyBlockElementsNums := (len(at) + 1) / 2 //每块内元素数量
	if everyBlockElementsNums == 0 {
		fmt.Println("无法获得索引块")
		return nil
	}
	var bidx []BlockIndex = make([]BlockIndex, everyBlockElementsNums)
	for i := 0; i < everyBlockElementsNums; i++ { //初始化默认值
		bidx[i].Start = everyBlockElementsNums * i
		bidx[i].MaxValue = vals.IntMin
	}
	//bidxJ := 0 //索引块计数器
	for i := 0; i < len(at); i++ {
		if (i+1)%blockSize != 0 {
			// if bidx[bidxJ].MaxValue =
			// bidx[bidxJ].MaxValue =
		}
	}
	return bidx
}

//quickSortBlock 分块
func quickSortBlock() {

}

//TestBlock 分块查找测试
func TestBlock() {
	at := []vals.AlgorType{0, 0}
	bidx := GetBlockIndexByOrder(3, at)
	fmt.Println(bidx)
	fmt.Println(vals.IntMin)
}
