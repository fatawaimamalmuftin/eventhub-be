package service

import (
	"context"
	"encoding/json"
	"log"
	"time"

	cuserror "github.com/fatawaimamalmuftin/eventhub-be/internal/CusError"
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

func (e *EventsFilterSrvS) GetEvents(c context.Context, eventQuery dto.EventQuery) ([]model.Event, error) {
	eventRedis, err := e.RDB.Get(c, "eventhub:eventfilter").Result()
	if err != nil {
		event, err := e.EFr.GetEvents(c, eventQuery)
		if err != nil {
			return []model.Event{}, cuserror.InternalError
		}

		eventFromDb, err := json.Marshal(event)
		if err != nil {
			return []model.Event{}, err
		}
		er := e.RDB.Set(c, "eventhub:eventfilter", eventFromDb, 5*time.Minute)
		if er != nil {
			return []model.Event{}, cuserror.InternalError
		}
	}

	var event []model.Event
	// _ = e.RDB.Del(c, "eventhub:eventfilter").Err()

	er := json.Unmarshal([]byte(eventRedis), &event)
	if er != nil {
		log.Println(er)
	}

	return event, nil
}
