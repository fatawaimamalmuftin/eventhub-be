package repo

import (
	"context"

	"github.com/fatawaimamalmuftin/eventhub-be/internal/dto"
	"github.com/jackc/pgx/v5/pgxpool"
)

type DbCreateEventRepo struct {
	Db *pgxpool.Pool
}

func ProviderCreateEventRepo(db *pgxpool.Pool) *DbCreateEventRepo {
	return &DbCreateEventRepo{
		Db: db,
	}
}

func (pr *DbCreateEventRepo) CreateEventRepo(data dto.CreateEvent) error {
	q := `
	INSERT INTO events ( 
	title, images, start_time, end_time, location, attendees, capacity, description, event_format, community_id)
	VALUES ($1,$2,$3,$4,$5,$6,$7,$8,$9,$10);
	`

	_, err := pr.Db.Exec(context.Background(), q,
		data.Title, data.Images, data.StartTime, data.EndTime, data.Location, 0, data.Capacity, data.Description, data.EventFormat, data.CommunityID,
	)

	return err
}
