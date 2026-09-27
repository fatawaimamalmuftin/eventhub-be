package service

import (
	"context"

	"github.com/fatawaimamalmuftin/eventhub-be/internal/model"
	"github.com/fatawaimamalmuftin/eventhub-be/internal/repo"
)

type MyEventSrv struct {
	MEr *repo.DbMyEventRepo
}

func MyEventService(mer *repo.DbMyEventRepo) *MyEventSrv {
	return &MyEventSrv{
		MEr: mer,
	}
}

func (m *MyEventSrv) GetMyEventsService(
	userID int,
	c context.Context,
) ([]model.MyEvent, error) {
	return m.MEr.GetMyEventsRepo(userID, c)
}
