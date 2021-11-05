// item.go
// 房间生产的物品的定义

package item

type Item struct {
	// 物品ID
	ID int32

	// 每个物品占用多少库存
	Storage int32

	// 每个物品需要消耗多少生产力
	Production int32
}
