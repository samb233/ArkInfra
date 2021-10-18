// workerarray.go
// worker聚合，用数组的形式

package worker

type WorkerArray struct {
	Workers []*Worker
}

func NewWorkerArray(workers ...*Worker) *WorkerArray {
	return &WorkerArray{
		Workers: workers,
	}
}

// 确认工人是否有人疲劳
func (wa *WorkerArray) IsOneExhausted() bool {
	for _, worker := range wa.Workers {
		if worker.IsExhausted() {
			return true
		}
	}

	return false
}

// 确认工人是否全部疲劳
// 工人全部疲劳时，房间将停止运作
func (wa *WorkerArray) IsAllExhausted() bool {
	for _, worker := range wa.Workers {
		if !worker.IsExhausted() {
			return false
		}
	}

	return true
}

// 获取奖励库存
// TODO: 合并两个方法
func (wa *WorkerArray) GetBonusStorage(itemID int) (bonusStorage int) {
	for _, worker := range wa.Workers {
		bonusStorage = bonusStorage + worker.GetBonusStorage(itemID)
	}
	return
}

// 获取奖励生产力，用于展示
func (wa *WorkerArray) GetBonusProductivity(itemID int) (bonusProductivity int) {
	for _, worker := range wa.Workers {
		bonusProductivity = bonusProductivity + worker.GetBonusProductivity(itemID)
	}

	return
}

// 生产
// item: 生产的物品ID
// minutes: 生产的时间，单位为分钟
// maxWorkingMinutes: 实际生产的最大时间
// production: 工人合计生产的生产力
func (wa *WorkerArray) Produce(item int, minutes int) (maxWoringMinutes int, production int) {
	maxWoringMinutes = 0

	for _, worker := range wa.Workers {
		workingTime, workerProduction := worker.Produce(item, minutes)
		production = production + workerProduction

		if workingTime > maxWoringMinutes {
			maxWoringMinutes = workingTime
		}
	}

	return
}
