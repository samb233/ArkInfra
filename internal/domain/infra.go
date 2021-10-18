// infra.go
// 定义基建对象
// 目前只实现了工厂类房间，所以目前可以视作工厂的聚合

package domain

import (
	"time"

	"github.com/samb233/arkinfra/internal/domain/room"
)

const (
	// 最大计算时间：36小时
	// 超过36小时的时间间隔会被视为36小时
	maxMinutes int = 2160
)

type Infra struct {
	// 玩家ID
	ID int

	// 房间数组
	Rooms []*room.Room

	// 上次操作的时间
	LastOpTime time.Time

	// 产品的数量，用于展示
	Producted int

	// 停工的数量，用于展示
	StopProduce int
}

// 新建基建
// 在新建之后调用Flush方法刷新状态
func NewInfra(id int, lastOpTime time.Time, rooms ...*room.Room) *Infra {
	infra := &Infra{
		ID:         id,
		LastOpTime: lastOpTime,
		Rooms:      rooms,
	}

	return infra
}

// 刷新状态
// 会根据当前的时间与上次时间的差调用房间的生产方法
// 更新Infra中用于展示的Producted、StopProduced等属性
func (infra *Infra) Flush() {
	minutes := int(time.Since(infra.LastOpTime).Minutes())
	if minutes > maxMinutes {
		minutes = maxMinutes
	}

}
