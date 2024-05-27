package lcpr

import (
	"bufio"
	"encoding/json"
	"fmt"
	"os"
	"strings"
	"testing"
)

func TestFunc(t *testing.T) {
	fmt.Println()
}

// 将lcpr插件生成的英文名文件重命名为中文
func TestFuncRenameChinese(t *testing.T) {
	pairs := [][2]string{}
	var recursive func(path string)
	recursive = func(path string) {
		files, err := os.ReadDir(path)
		if err != nil {
			panic(err)
		}
		for _, fs := range files {
			currentPath := path + "/" + fs.Name()
			if fs.Type().IsDir() {
				recursive(currentPath)
			}
			name := fs.Name()
			if nameSplit := strings.Split(name, "."); len(nameSplit) > 2 {
				file, err := os.Open(currentPath)
				if err != nil {
					panic(err)
				}
				defer file.Close()
				bScanner := bufio.NewScanner(file)
				signStr := fmt.Sprintf("[%s] ", nameSplit[0])
				for bScanner.Scan() {
					line := bScanner.Text()
					if strings.Contains(line, signStr) {
						zhName := line[strings.Index(line, signStr)+len(signStr):]
						if nameSplit[1] != zhName {
							pairs = append(pairs, [2]string{currentPath,
								path + "/" + nameSplit[0] + "." + zhName + ".go"})
						}
						break
					}
				}
			}
		}
	}
	recursive("./")
	for _, pair := range pairs {
		if err := os.Rename(pair[0], pair[1]); err != nil {
			panic(err)
		}
	}
}

// str 用例  转 一维 string 数组
func str2SliceString(str string) []string {
	result := []string{}
	json.Unmarshal([]byte(str), &result)
	return result
}

// str 用例  转 一维 string 数组
func str2SliceString2(str string) [][]string {
	result := [][]string{}
	json.Unmarshal([]byte(str), &result)
	return result
}

// str 用例  转 一维int数组
func str2SliceInt(str string) []int {
	result := []int{}
	json.Unmarshal([]byte(str), &result)
	return result
}

// str 用例  转 二维int数组
func str2SliceInt2(str string) [][]int {
	result := [][]int{}
	json.Unmarshal([]byte(str), &result)
	return result
}

// str 用例  转 单链表
func str2ListNode(str string) *ListNode {
	nodes := str2SliceInt(str)
	var listNode, h *ListNode
	for i, node := range nodes {
		nd := &ListNode{
			Val:  node,
			Next: nil,
		}
		if i == 0 {
			listNode = nd
			h = listNode
		}
		h.Next = nd
		h = h.Next
	}
	return listNode
}

// 打印单链表
func printListNode(list *ListNode) {
	fmt.Printf("%s", "[")
	for h := list; h != nil; h = h.Next {
		fmt.Printf("%d", h.Val)
		if h.Next != nil {
			fmt.Printf(",")
		}
	}
	fmt.Println("]")
}

// str 用例  转 完全二叉树
func str2FullTreeNode(str string) *TreeNode {
	nodes := str2SliceInt(str)
	if len(nodes) == 0 {
		return nil
	}
	root := &TreeNode{
		Val:   nodes[0],
		Left:  nil,
		Right: nil,
	}
	queue := []*TreeNode{root}
	nodeIndex := 1
	for i := 0; i < len(queue) && nodeIndex < len(nodes); i++ {
		if nodeIndex < len(nodes) {
			queue[i].Left = &TreeNode{
				Val:   nodes[nodeIndex],
				Left:  nil,
				Right: nil,
			}
			queue = append(queue, queue[i].Left)
			nodeIndex++
		}
		if nodeIndex < len(nodes) {
			queue[i].Right = &TreeNode{
				Val:   nodes[nodeIndex],
				Left:  nil,
				Right: nil,
			}
			queue = append(queue, queue[i].Right)
			nodeIndex++
		}
	}
	return root
}

// 打印 完全二叉树
func printFullTreeNode(root *TreeNode) {
	fmt.Printf("%s", "[")
	queue := []*TreeNode{root}
	for i := 0; i < len(queue); i++ {
		fmt.Printf("%d", queue[i].Val)
		if queue[i].Left != nil {
			queue = append(queue, queue[i].Left)
		}
		if queue[i].Right != nil {
			queue = append(queue, queue[i].Right)
		}
		if i+1 < len(queue) {
			fmt.Printf(",")
		}
	}
	fmt.Println("]")
}

//

// func TestAlgor() {

// 	fmt.Println()
// 	fmt.Println("______________")
// 	//
// 	lfu := Constructor(2)
// 	lfu.Put(1, 1)
// 	lfu.Put(2, 2)
// 	fmt.Println(lfu.Get(1))
// 	lfu.iterator()
// 	fmt.Println("lfu.KeyValueMap", lfu.KeyValueMap)
// 	fmt.Println("lfu.CountLinkMap", lfu.CountLinkMap)
// 	fmt.Println("______________")
// 	lfu.Put(3, 3)
// 	fmt.Println("lfu.KeyCountMap", lfu.KeyCountMap)
// 	fmt.Println("lfu.KeyValueMap", lfu.KeyValueMap)
// 	fmt.Println("lfu.KeyCountMap", lfu.KeyCountMap)
// 	fmt.Println("lfu.CountLinkMap", lfu.CountLinkMap)
// 	fmt.Println("______________")
// 	fmt.Println(lfu.Get(2))
// 	fmt.Println(lfu.Get(3))
// 	lfu.Put(4, 4)
// 	fmt.Println(lfu.Get(1))
// 	fmt.Println(lfu.Get(3))
// 	fmt.Println(lfu.Get(4))
// }

// func (this *LFUCache) iterator() {
// 	fmt.Print("CountSortList   ")
// 	for h := this.CountSortList.Front(); h != nil; h = h.Next() {
// 		fmt.Print("\t", h.Value, "\t")
// 	}
// 	fmt.Print("\nCountLinkMap{1}   ")
// 	for h := this.CountLinkMap{1}.Front(); h != nil; h = h.Next() {
// 		fmt.Print("\t", h.Value, "\t")
// 	}

// }

/////////////////////////////////////////////////////
