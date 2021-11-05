// worker.go
// 工人的定义

package worker

import (
	"strconv"
	"strings"
)

const (
	// 默认生产力
	DefaultProductivity int32 = 1

	// 默认每分钟消耗疲劳
	DefaultFtgPerMinute int32 = 100

	// 初始疲劳
	DefaultFtg int32 = 144000
)

// 工人
// 一个工人每分钟的产出 = 默认生产力 + 奖励生产力
// 默认生产力为1
// 奖励生产力在生产标记物品时生效，值比如为25
// 一个工人每分钟消耗的疲劳值 = 默认疲劳 - 奖励疲劳
// 默认疲劳为100
// 奖励疲劳在生产标记物品时生效，值比如为25
type Worker struct {
	// 角色ID
	ID int32

	// 当前疲劳值，初始为24x100x60=144000点
	Ftg int32

	// 标记物品，如果生产的是这些物品，则触发奖励(如果疲劳，则失效)
	// 通过编排物品编号减少计算
	// 实际上就是判断物品ID前缀
	ItemPrefix string

	// 奖励生产力
	BonusProductivity int32

	// 奖励疲劳
	BonusFtg int32

	// 标记房间类型
	// 如果处在这种房间类型中，该房间库存增加(不会因为疲劳而失效)
	// 实际上是判断房间类型前缀
	RoomTypePrefix string

	// 奖励库存
	BonusStorage int32
}

// 新建工人
func NewWorker(id, ftg int32, options ...WorkerOptions) *Worker {
	worker := &Worker{
		ID:  id,
		Ftg: ftg,
	}

	for _, option := range options {
		option(worker)
	}
	return worker
}

// 工人选项
type WorkerOptions func(*Worker)

// 奖励生产力、疲劳相关设置
// bonusFtg: 标记物品前缀
// bonusProductivity: 奖励生产力
// ftg: 奖励疲劳
func BonusOptions(itemPrefix string, bonusProductivity, bonusFtg int32) WorkerOptions {
	return func(w *Worker) {
		w.ItemPrefix = itemPrefix
		w.BonusProductivity = bonusProductivity
		w.BonusFtg = bonusFtg
	}
}

// 奖励库存相关设置
// roomTypePrefix: 房间类型前缀
// stroage: 奖励库存
func StorageOptions(roomTypePrefix string, bonusStorage int32) WorkerOptions {
	return func(w *Worker) {
		w.RoomTypePrefix = roomTypePrefix
		w.BonusStorage = bonusStorage
	}
}

// 获取奖励库存
// 应当在进驻房间的时候调用
// roomType: 房间类型
func (w *Worker) GetBonusStorage(roomType int32) int32 {
	roomTypeString := strconv.Itoa(int(roomType))
	if strings.HasPrefix(roomTypeString, w.RoomTypePrefix) {
		return w.BonusStorage
	}

	return 0
}

// 获取奖励生产力(显示用)
// 应当在进驻房间的时候调用
// 只做显示用，实际计算时使用CheckBonus方法
// itemID: 物品ID
func (w *Worker) GetBonusProductivity(itemID int32) int32 {
	itemIDString := strconv.Itoa(int(itemID))
	if strings.HasPrefix(itemIDString, w.ItemPrefix) {
		return w.BonusProductivity
	}

	return 0
}

// 生产
// 传入生产的物品ID、生产的分钟数
// 会先进行判断能否生产那么长时间
// 返回能生产的最大时间，贡献的生产力
// itemID: 生产的物品ID
// minutes: 想要生产的时间
// workingMinutes: 实际生产的时间
// production: 产出的生产力
// TODO: 优化计算
func (w *Worker) Produce(itemID int32, minutes int32) (workingMinutes int32, production int32) {
	// 判断是否疲劳
	if w.IsExhausted() {
		return 0, 0
	}

	// 得到生产该物品时的生产力、疲劳消耗
	productivity, ftgPerMinute := w.CheckBonus(itemID)

	// 判断能否生产minutes分钟
	maxWorkingMinutes := w.CheckWorkingTime(minutes, ftgPerMinute)

	return maxWorkingMinutes, maxWorkingMinutes * productivity
}

// 是否疲劳
func (w *Worker) IsExhausted() bool {
	return w.Ftg == 0
}

// 确认奖励生产力、疲劳
// 输入物品ID，返回生产这个物品时的生产力、疲劳值消耗
// itemID: 生产的物品ID
// productivity: 生产该物品时的生产力
// ftgPerMinute: 生产该物品时消耗的疲劳值
func (w *Worker) CheckBonus(itemID int32) (productivity int32, ftgPerMinute int32) {
	productivity = DefaultProductivity
	ftgPerMinute = DefaultFtgPerMinute

	itemIDString := strconv.Itoa(int(itemID))
	if strings.HasPrefix(itemIDString, w.ItemPrefix) {
		productivity = productivity + w.BonusProductivity
		ftgPerMinute = ftgPerMinute - w.BonusFtg
	}

	return productivity, ftgPerMinute
}

// 确认工作时间
// 输入时间，如果能工作这么长时间，则返回该时间，并减去相应疲劳
// 如果不能，则返回最大能工作的时间，将疲劳置为0
// minutes: 想要让工人工作的时间
// ftgPerMinute: 单位时间疲劳值消耗
func (w *Worker) CheckWorkingTime(minutes int32, ftgPerMinute int32) int32 {
	ftgConsume := minutes * ftgPerMinute

	if ftgConsume > w.Ftg {
		minutes = w.Ftg / ftgPerMinute
		w.Ftg = 0
	} else {
		w.Ftg = w.Ftg - ftgConsume
	}

	return minutes
}
