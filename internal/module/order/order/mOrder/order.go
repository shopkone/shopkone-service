package mOrder

type OrderType uint8

const (
	OrderTypeManual   OrderType = 1 // 手动补录
	OrderTypeCustomer OrderType = 2 // 客户下单
)

type OrderFeeType uint8

const (
	OrderFeeTypeAuto   OrderFeeType = 1 // 自动计算
	OrderFeeTypeManual OrderFeeType = 2 // 手动计算
)

type OrderStatus uint8

const (
	OrderStatusPending           OrderStatus = 1 // 待支付
	OrderStatusPaid              OrderStatus = 2 // 已支付
	OrderStatusShipped           OrderStatus = 3 // 已发货
	OrderStatusCompleted         OrderStatus = 4 // 已完成
	OrderStatusCancelled         OrderStatus = 5 // 已取消
	OrderStatusRefunded          OrderStatus = 6 // 已退款
	OrderStatusPartiallyRefunded OrderStatus = 7 // 部分退款
	OrderStatusCart              OrderStatus = 8 // 加入购物车
)

type Order struct {
	Amount      float32      `gorm:"default:0"` // 订单总金额
	Note        string       `gorm:"size:500"`  // 备注
	Type        OrderType    `gorm:"default:1"` // 类型
	FeeType     OrderFeeType `gorm:"default:1"` // 费用类型
	CreateEmail string       `gorm:"size:500"`  // 创建人
	Status      OrderStatus  `gorm:"index"`     // 状态
}
