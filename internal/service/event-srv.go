package service

import (
	"context"
	"encoding/json"
	"log"
	"time"

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

	if err == nil {
		var events []model.Event

		if err := json.Unmarshal([]byte(eventRedis), &events); err != nil {
			log.Println("error unmarshal redis:", err)
		}

		log.Println("get from redis")
		return events, nil
	}
	log.Println("get from db")

	eventsFromDB, err := e.EFr.GetEvents(c, eventQuery)
	if err != nil {
		return []model.Event{}, err
	}

	eventFromDB, err := json.Marshal(eventsFromDB)
	if err != nil {
		return []model.Event{}, err
	}

	err = e.RDB.Set(c, "eventhub:eventfilter", eventFromDB, 5*time.Minute).Err()

	if err != nil {
		log.Println("error set cache:", err)
	}

	return eventsFromDB, nil
}
