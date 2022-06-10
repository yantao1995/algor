package design_mode

import "fmt"

/*
	访问者模式可以给一系列对象透明的添加功能，并且把相关代码封装到一个类中。
	对象只要预留访问者接口Accept则后期为对象添加功能的时候就不需要改动对象
*/

// 访问者接口
type Customer interface {
	Accept(Visitor)
}

// 访问类别抽象接口
type Visitor interface {
	Visit(Customer)
}

// 访问对象列表实现
type CustomerCol struct {
	customers []Customer
}

func (c *CustomerCol) Add(customer Customer) {
	c.customers = append(c.customers, customer)
}

func (c *CustomerCol) Accept(visitor Visitor) {
	for _, customer := range c.customers {
		customer.Accept(visitor)
	}
}

// 对象组1  （企业类型）
type EnterpriseCustomer struct {
	name string
}

func NewEnterpriseCustomer(name string) *EnterpriseCustomer {
	return &EnterpriseCustomer{name: name}
}

func (c *EnterpriseCustomer) Accept(visitor Visitor) {
	visitor.Visit(c)
}

// 对象组2  （个人类型）
type IndividualCustomer struct {
	name string
}

func NewIndividualCustomer(name string) *IndividualCustomer {
	return &IndividualCustomer{name: name}
}

func (c *IndividualCustomer) Accept(visitor Visitor) {
	visitor.Visit(c)
}

// 实现类1  （服务类型）
type ServiceRequestVisitor struct{}

func (*ServiceRequestVisitor) Visit(customer Customer) {
	switch c := customer.(type) {
	case *EnterpriseCustomer:
		fmt.Println("ServiceRequestVisitor Enterprise customer:", c.name)
	case *IndividualCustomer:
		fmt.Println("ServiceRequestVisitor Individual customer:", c.name)
	}
}

//实现类2  （分析类型）
type AnalysisVisitor struct{}

func (*AnalysisVisitor) Visit(customer Customer) {
	switch c := customer.(type) {
	case *EnterpriseCustomer:
		fmt.Println("AnalysisVisitor Enterprise customer:", c.name)
	}
}
