package design_mode

import (
	design_create "algor/design_mode/type_create"
	design_struct "algor/design_mode/type_struct"
	"fmt"
	"testing"
)

func TestTypeStruct(t *testing.T) {
	//选项  （初始化时为对象的每个属性赋值）
	struct_op := design_struct.NewOption(design_struct.IdOpFunc(1), design_struct.NameOpFunc("name"))
	fmt.Println(struct_op)
}

func TestTypeCreate(t *testing.T) {
	//简单工厂	(根据传参生成不同对象)
	simpleFactoryApi1 := design_create.NewSimpleFactoryApi(1)
	simpleFactoryApi2 := design_create.NewSimpleFactoryApi(2)
	simpleFactoryApi3 := design_create.NewSimpleFactoryApi(3)
	fmt.Println(simpleFactoryApi1.Say("simple1"), simpleFactoryApi2.Say("simple2"), simpleFactoryApi3)

	//工厂方法	(类似于子类继承父类方法)
	var computeNormalFactoryPrint = func(factory design_create.OperatorFactory, a, b int) {
		op := factory.CreateOperator()
		op.SetA(a)
		op.SetB(b)
		fmt.Println(op.GetResult())
	}
	normalPlusFactory := design_create.PlusOperatorFactory{}
	normalMinusFactory := design_create.MinusOperatorFactory{}
	computeNormalFactoryPrint(normalPlusFactory, 1, 2)
	computeNormalFactoryPrint(normalMinusFactory, 3, 1)

	//抽象工厂方法	(对象要有关联，用于生成产品族的工厂，否则退化成工厂方法)
	var getMainAndDetail = func(factory design_create.OrderDaoAbstractFactory) {
		factory.CreateOrderMainDao().SaveOrderMain()
		factory.CreateOrderDetailDao().SaveOrderDetail()
	}
	rdbFactory := &design_create.RDBOrderDaoAbstract{}
	getMainAndDetail(rdbFactory)
	xmlFactory := &design_create.XMLOrderDaoAbstract{}
	getMainAndDetail(xmlFactory)

	//创建者  (复杂对象分离成多个简单对象构建组合)
	builder := design_create.Builder1{}
	director := design_create.NewDirector(&builder)
	director.ConstructDirector()
	fmt.Println(builder.GetResult())

	//原型  (克隆对象)
	prototypeManager := design_create.NewPrototypeManager()
	a := &design_create.TestStruct{Id: 10}
	prototypeManager.Set("A", a)
	a1 := prototypeManager.Get("A")
	b := a1.Clone()
	fmt.Println(a == b, b.(*design_create.TestStruct).Id)

	//单例	(只初始化一次)
	singleton := design_create.GetSingletonInstance()
	fmt.Println(singleton)
}
