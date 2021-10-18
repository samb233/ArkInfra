// domain.go
// 设置repo与usecase接口

package domain

import (
	"context"

	"github.com/samb233/arkinfra/internal/domain/room"
	"github.com/samb233/arkinfra/internal/domain/worker"
)

type InfraRepo interface {
	GetInfra(ctx context.Context, id int) (*Infra, error)
	SaveInfra(ctx context.Context, infra *Infra) error
}

type RoomRepo interface {
	GetRoom(ctx context.Context, id int) (*room.Room, error)
	SaveRoom(ctx context.Context, room *room.Room) error
}

type WorkerRepo interface {
	GetWorker(ctx context.Context, userID, charID int) (*worker.Worker, error)
	SaveWorker(ctx context.Context, userID int, worker *worker.Worker) error
}
