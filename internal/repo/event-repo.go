package repo

import (
	"context"

	"github.com/fatawaimamalmuftin/eventhub-be/internal/dto"
	"github.com/fatawaimamalmuftin/eventhub-be/internal/model"
	"github.com/jackc/pgx/v5/pgxpool"
)

type DbEventFilterRepo struct {
	Db *pgxpool.Pool
}

func EventsFilterRepo(db *pgxpool.Pool) *DbEventFilterRepo {
	return &DbEventFilterRepo{
		Db: db,
	}
}

func (e *DbEventFilterRepo) GetEvents(
	c context.Context,
	eventQuery dto.EventQuery,
) ([]model.Event, error) {

	q := `
		SELECT
			e.id_event,
			e.title,
			e.images,
			e.start_time,
			e.end_time,
			e.location,
			e.attendees,
			e.capacity,
			e.description,
			e.event_format,
			e.community_id,
			c.title AS community_title
		FROM events e
		JOIN community c
			ON c.id_community = e.community_id
		WHERE e.title ILIKE '%' || $1 || '%'
		ORDER BY e.start_time ASC;
	`

	rows, err := e.Db.Query(c, q, eventQuery.Search)

	if err != nil {
		return nil, err
	}

	defer rows.Close()

	var events []model.Event

	for rows.Next() {
		var event model.Event

		err := rows.Scan(
			&event.ID,
			&event.Title,
			&event.Images,
			&event.StartTime,
			&event.EndTime,
			&event.Location,
			&event.Attendees,
			&event.Capacity,
			&event.Description,
			&event.EventFormat,
			&event.CommunityID,
			&event.CommunityTitle,
		)

		if err != nil {
			return nil, err
		}

		events = append(events, event)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return events, nil
}
