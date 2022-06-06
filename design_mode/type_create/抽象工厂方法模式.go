package design_mode

import "fmt"

/*
	抽象工厂模式用于生成产品族的工厂，所生成的对象是有关联的。
	如果抽象工厂退化成生成的对象无关联则成为工厂函数模式。
	比如本例子中使用RDB和XML存储订单信息，抽象工厂分别能生成相关的主订单信息和订单详情信息。
	如果业务逻辑中需要替换使用的时候只需要改动工厂函数相关的类就能替换使用不同的存储方式了。
*/

/*
  抽象接口定义
*/

// 订单主记录
type OrderMainDao interface {
	SaveOrderMain()
}

// 订单详细记录
type OrderDetailDao interface {
	SaveOrderDetail()
}

// 抽象工厂接口
type OrderDaoAbstractFactory interface {
	CreateOrderMainDao() OrderMainDao
	CreateOrderDetailDao() OrderDetailDao
}

/*
  RDB存储的实现
*/

//RDB抽象工厂实现
type RDBOrderDaoAbstract struct{}

func (*RDBOrderDaoAbstract) CreateOrderMainDao() OrderMainDao {
	return &rdbMainDao{}
}

func (*RDBOrderDaoAbstract) CreateOrderDetailDao() OrderDetailDao {
	return &rdbDetailDao{}
}

//  数据库的  OrderMainDao 的实现
type rdbMainDao struct{}

func (*rdbMainDao) SaveOrderMain() {
	fmt.Println("RDBMainDao 实现 SaveOrderMain()")
}

//  数据库的  RDBDetailDao 的实现
type rdbDetailDao struct{}

func (*rdbDetailDao) SaveOrderDetail() {
	fmt.Println("RDBDetailDao 实现 SaveOrderDetail()")
}

/*
   XML存储的实现
*/

//XML抽象工厂实现
type XMLOrderDaoAbstract struct{}

func (*XMLOrderDaoAbstract) CreateOrderMainDao() OrderMainDao {
	return &xmlMainDao{}
}

func (*XMLOrderDaoAbstract) CreateOrderDetailDao() OrderDetailDao {
	return &xmlDetailDao{}
}

//  XML的  OrderMainDao 的实现
type xmlMainDao struct{}

func (*xmlMainDao) SaveOrderMain() {
	fmt.Println("XMLMainDao 实现 SaveOrderMain()")
}

//  XML的  RDBDetailDao 的实现
type xmlDetailDao struct{}

func (*xmlDetailDao) SaveOrderDetail() {
	fmt.Println("XMLDetailDao 实现 SaveOrderDetail()")
}
