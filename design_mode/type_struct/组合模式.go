package design_mode

import "fmt"

/*
	组合模式统一对象和对象集，使得使用相同接口使用对象和对象集。
	组合模式常用于树状结构，用于统一叶子节点和树节点的访问，并且可以用于应用某一操作到所有子节点。
*/
const (
	LeafNode = iota
	CompositeNode
)

//组件接口
type Component interface {
	Parent() Component
	SetParent(Component)
	Name() string
	SetName(string)
	AddChild(Component)
	Print(string)
}

func NewComponent(kind int, name string) Component {
	var c Component
	if kind == LeafNode {
		c = NewLeaf()
	} else if kind == CompositeNode {
		c = NewComposite()
	}
	c.SetName(name)
	return c
}

//组件实现
type componentImpl struct {
	parent Component
	name   string
}

func (c *componentImpl) Parent() Component {
	return c.parent
}
func (c *componentImpl) SetParent(parent Component) {
	c.parent = parent
}
func (c *componentImpl) Name() string {
	return c.name
}
func (c *componentImpl) SetName(name string) {
	c.name = name
}
func (c *componentImpl) AddChild(Component) {}
func (c *componentImpl) Print(string)       {}

//叶子结点
type Leaf struct {
	componentImpl
}

func NewLeaf() *Leaf {
	return &Leaf{}
}

func (c *Leaf) Print(pre string) {
	fmt.Println("leaf :", pre, c.Name())
}

//组合结点
type Composite struct {
	componentImpl
	children []Component
}

func NewComposite() *Composite {
	return &Composite{
		children: make([]Component, 0),
	}
}

func (c *Composite) AddChild(child Component) {
	child.SetParent(c)
	c.children = append(c.children, child)
}

func (c *Composite) Print(pre string) {
	fmt.Println("composite :", pre, c.Name())
	for _, child := range c.children {
		child.Print(pre + "  ")
	}
}
