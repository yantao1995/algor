package vals

const IntMin = ^int(^uint(0) >> 1)
const IntMax = int(^uint(0) >> 1)

//AlgorType 全局数据类型   ---方便转换成 int,float(运算) 或 string(非运算)
type AlgorType interface{}

//AlgorNilVaule 全局空nil
var AlgorNilVaule AlgorType = nil

//LessThanInt interface{}转int  小于
func LessThanInt(prev, next AlgorType) bool {
	return prev.(int) < next.(int)
}

//LessThanInt interface{}转int  小于等于
func LessEquelsThanInt(prev, next AlgorType) bool {
	return prev.(int) <= next.(int)
}

//GreaterThanInt interface{}转int  大于
func GreaterThanInt(prev, next AlgorType) bool {
	return prev.(int) > next.(int)
}

//GreaterEquelsThanInt interface{}转int  大于等于
func GreaterEquelsThanInt(prev, next AlgorType) bool {
	return prev.(int) >= next.(int)
}

//GreaterThanInt interface{}转int  等于
func EquelsInt(prev, next AlgorType) bool {
	return prev.(int) == next.(int)
}

//GetInt 获取自定义默认int类型
func GetInt(at AlgorType) int {
	return at.(int)
}
