// infra.go
// 定义基建对象
// 目前只实现了工厂类房间，所以目前可以视作工厂的聚合

package domain

import (
	"github.com/samb233/arkinfra/internal/domain/room"
)

type Infra struct {
	// 玩家ID
	ID int64

	// 房间数组
	Rooms []*room.Room

	// 产品的数量，用于展示
	Producted int32

	// 停工的数量，用于展示
	StopProduce int32
}

// 新建基建
// 在新建之后调用Flush方法刷新状态
func NewInfra(id int64, rooms ...*room.Room) *Infra {
	infra := &Infra{
		ID:    id,
		Rooms: rooms,
	}

	return infra
}

// 刷新状态
// 会根据当前的时间与上次时间的差调用房间的生产方法
// 更新Infra中用于展示的Producted、StopProduced等属性
func (infra *Infra) Flush() {
	var produced, stopProduced int32 = 0, 0
	for _, room := range infra.Rooms {
		room.Flush()
		produced = produced + room.GetItemAmount()
		if room.IsNotWorking() {
			stopProduced++
		}
	}

	infra.Producted = produced
	infra.StopProduce = stopProduced
}
