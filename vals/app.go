package vals

//AlgorType 全局数据类型
type AlgorType interface{}

//AlgorNilVaule 全局空nil
var AlgorNilVaule AlgorType = nil

//LessThanInt interface{}转int  小于
func LessThanInt(prev, next AlgorType) bool {
	return prev.(int) < next.(int)
}

//GreaterThanInt interface{}转int  大于
func GreaterThanInt(prev, next AlgorType) bool {
	return prev.(int) > next.(int)
}

//GreaterThanInt interface{}转int  等于
func EquelsInt(prev, next AlgorType) bool {
	return prev.(int) == next.(int)
}

//GetInt 获取自定义默认int类型
func GetInt(at AlgorType) int {
	return at.(int)
}
