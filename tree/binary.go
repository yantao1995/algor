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
func firstErgodic(root *Binary) {
	if root != nil {
		fmt.Printf("%s\t", root.Data)
		firstErgodic(root.Left)
		firstErgodic(root.Right)
	}
}

//中
func middleErgodic(root *Binary) {
	if root != nil {
		firstErgodic(root.Left)
		fmt.Printf("%s\t", root.Data)
		firstErgodic(root.Right)
	}
}

//后
func lastErgodic(root *Binary) {
	if root != nil {
		firstErgodic(root.Left)
		firstErgodic(root.Right)
		fmt.Printf("%s\t", root.Data)
	}
}

//深度
func deepErgodic() {

}

//广度
func wideErgodic(root *Binary) {
	for root != nil {
		fmt.Printf("%s\t", root.Data)
		root = root.Left
	}
}

//拷贝
func copyBinary() {

}

/*
		  a
	b_____|_____c
d___|___e   f___|___g

*/

func TestBinary() {
	root := InitBinary("a")
	LoadBinary(root)
	firstErgodic(root)
	println()
	middleErgodic(root)
	println()
	lastErgodic(root)
	println()
}

// 加载treenodes
func LoadBinary(root *Binary) {
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
