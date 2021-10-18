// room.go
// 房间对象的定义

package room

import (
	"time"

	"github.com/samb233/arkinfra/internal/domain/item"
	"github.com/samb233/arkinfra/internal/domain/worker"
)

const (
	// TODO: 弄清需求中是否有根据房间等级调整这些值

	// 默认房间的生产力，固定为100
	DefaultRoomProductivity int = 100

	// 默认房间的库存，固定为50
	DefaultRoomStorage int = 50
)

type Room struct {
	// 房间ID
	ID int

	// 房间类型
	// TODO: 类型定义
	RoomType int

	// 房间等级
	Level int

	// 上次操作时间
	LastOpTime time.Time

	// 房间库存
	Storage int

	// 已使用的库存
	StorageUsed int

	// 当前生产进度（还没达到物品生产的部分）
	Production int

	// 当前生产的物品
	// 包含物品ID、单个物品占用库存、单个物品需要消耗生产力
	Item *item.Item

	// 工人聚合
	Workers *worker.WorkerArray
}

// 新建房间
func NewRoom(id, roomType, level int, lastOpTime time.Time, options ...RoomOptions) *Room {
	room := &Room{
		ID:         id,
		RoomType:   roomType,
		Level:      level,
		LastOpTime: lastOpTime,
		Storage:    DefaultRoomStorage,
	}

	for _, option := range options {
		option(room)
	}

	room.AddBonusStorage()

	return room
}

// 刷新房间信息
func (r *Room) Flush() {
	minutes := int(time.Since(r.LastOpTime).Minutes())
	r.Produce(minutes)
	r.UpdateOpTime()
}

// 更新上次操作时间
func (r *Room) UpdateOpTime() {
	r.LastOpTime = time.Now()
}

// 将奖励库存加到库存中
func (r *Room) AddBonusStorage() {
	if r.Workers == nil {
		return
	}

	// 获取奖励库存
	bonusStorage := r.Workers.GetBonusStorage(r.Item.ID)
	r.Storage = DefaultRoomProductivity + bonusStorage
}

// 设置工人
func (r *Room) SetWorkers(workers *worker.WorkerArray) {
	r.Flush()
	r.Workers = workers
	r.AddBonusStorage()
}

// 获取物品
func (r *Room) GetItem() (itemID int, amount int) {
	r.Flush()

	itemID = r.Item.ID
	amount = r.StorageUsed / r.Item.Storage
	r.StorageUsed = 0
	return
}

// 重设物品
// 重设时会清空库存，获取物品
func (r *Room) GetItemAndReSet(item *item.Item) (itemID int, amount int) {
	itemID, amount = r.GetItem()

	r.Item = item
	r.AddBonusStorage()

	return
}

// 生产
// 根据经过的时间，进行生产物品的操作
// minutes: 从上次操作到现在经过的分钟数
func (r *Room) Produce(minutes int) {
	if r.IsNotWorking() {
		return
	}

	// 工人贡献的生产力和最大生产时间
	workerProduction, maxWorkingMinutes := r.Workers.Produce(r.Item.ID, minutes)

	// 房间自身的生产力
	roomProduction := DefaultRoomProductivity * maxWorkingMinutes

	// 两者求和，增加库存
	roomProduction = roomProduction + workerProduction + r.Production

	r.StorageUsed = r.StorageUsed + roomProduction/r.Item.Production
	r.Production = roomProduction % r.Item.Production
	if r.StorageUsed > r.Storage {
		r.StorageUsed = r.Storage
		r.Production = 0
	}
}

// 是否停工
func (r *Room) IsNotWorking() bool {
	if r.Workers == nil || r.IsFull() {
		return true
	}

	if r.Workers.IsAllExhausted() {
		return true
	}

	return false
}

// 库存是否已经用尽
func (r *Room) IsFull() bool {
	return r.Storage == r.StorageUsed
}

// 是否有工人疲劳
func (r *Room) IsOneExhausted() bool {
	return r.Workers.IsOneExhausted()
}

// 房间初始化选项
type RoomOptions func(*Room)

func ProductOption(storageUsed, production int) RoomOptions {
	return func(r *Room) {
		r.StorageUsed = storageUsed
		r.Production = production
	}
}

func ItemOption(item *item.Item) RoomOptions {
	return func(r *Room) {
		r.Item = item
	}
}

func WorkerOption(workers *worker.WorkerArray) RoomOptions {
	return func(r *Room) {
		r.Workers = workers
	}
}
