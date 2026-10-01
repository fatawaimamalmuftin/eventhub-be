package service

import (
	"context"

	"github.com/fatawaimamalmuftin/eventhub-be/internal/dto"
	"github.com/fatawaimamalmuftin/eventhub-be/internal/model"
	"github.com/fatawaimamalmuftin/eventhub-be/internal/repo"
	"github.com/redis/go-redis/v9"
)

type EventsFilterSrvS struct {
	EFr *repo.DbEventFilterRepo
	RDB *redis.Client
}

func EventsFilterService(efr *repo.DbEventFilterRepo, rdb *redis.Client) *EventsFilterSrvS {
	return &EventsFilterSrvS{
		EFr: efr,
		RDB: rdb,
	}
}

func (e *EventsFilterSrvS) GetEvents(
	c context.Context,
	eventQuery dto.EventQuery,
) ([]model.Event, error) {
	return e.EFr.GetEvents(c, eventQuery)
}
