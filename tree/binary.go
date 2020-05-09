package tree

import "fmt"

type Binary struct {
	Data  string
	Left  *Binary
	Right *Binary
}

func InitBinary(dt string) *Binary {
	return &Binary{
		Data:  dt,
		Left:  nil,
		Right: nil,
	}
}

//先
func FirstErgodic(root *Binary) {
	if root != nil {
		fmt.Printf("%s\t", root.Data)
		FirstErgodic(root.Left)
		FirstErgodic(root.Right)
	}
}

//中
func MiddleErgodic(root *Binary) {
	if root != nil {
		MiddleErgodic(root.Left)
		fmt.Printf("%s\t", root.Data)
		MiddleErgodic(root.Right)
	}
}

//后
func LastErgodic(root *Binary) {
	if root != nil {
		LastErgodic(root.Left)
		LastErgodic(root.Right)
		fmt.Printf("%s\t", root.Data)
	}
}

//深度
func DeepErgodic() {

}

//广度
func WideErgodic(root *Binary) {
	for root != nil {
		fmt.Printf("%s\t", root.Data)
		root = root.Left
	}
}

//拷贝
func CopyBinary() {

}

/*  测试结构
		  a
	b_____|_____c
d___|___e   f___|___g

*/

func TestBinary() {
	root := InitBinary("a")
	LoadTestBinary(root)
	FirstErgodic(root)
	println()
	MiddleErgodic(root)
	println()
	LastErgodic(root)
	println()
}

// 加载treenodes
func LoadTestBinary(root *Binary) {
	root.Left = &Binary{
		Data: "b",
		Left: &Binary{
			Data:  "d",
			Left:  nil,
			Right: nil,
		},
		Right: &Binary{
			Data:  "e",
			Left:  nil,
			Right: nil,
		},
	}

	root.Right = &Binary{
		Data: "c",
		Left: &Binary{
			Data:  "f",
			Left:  nil,
			Right: nil,
		},
		Right: &Binary{
			Data:  "g",
			Left:  nil,
			Right: nil,
		},
	}
}
