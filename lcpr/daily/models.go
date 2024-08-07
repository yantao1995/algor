package lcpr

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}
type ListNode struct {
	Val  int
	Next *ListNode
}

//四叉树node
// type Node struct {
// 	Val         bool
// 	IsLeaf      bool
// 	TopLeft     *Node
// 	TopRight    *Node
// 	BottomLeft  *Node
// 	BottomRight *Node
// }

type Node struct {
	Val   int
	Left  *Node
	Right *Node
	Next  *Node
}
