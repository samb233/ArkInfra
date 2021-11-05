// domain.go
// 设置repo与usecase接口

package domain

import (
	"context"

	"github.com/samb233/arkinfra/internal/domain/room"
	"github.com/samb233/arkinfra/internal/domain/worker"
)

type InfraRepo interface {
	GetInfra(ctx context.Context, id int64) (*Infra, error)
	SaveInfra(ctx context.Context, infra *Infra) error
}

type RoomRepo interface {
	GetRoom(ctx context.Context, id int32) (*room.Room, error)
	SaveRoom(ctx context.Context, room *room.Room) error
}

type WorkerRepo interface {
	GetWorker(ctx context.Context, userID int64, charID int32) (*worker.Worker, error)
	SaveWorker(ctx context.Context, userID int64, worker *worker.Worker) error
}
