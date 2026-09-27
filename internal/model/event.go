package model

import "time"

type Event struct {
	ID             int
	Title          string
	Images         string
	StartTime      time.Time
	EndTime        time.Time
	Location       string
	Attendees      int
	Capacity       int
	Description    string
	EventFormat    string
	CommunityID    int
	CommunityTitle string
}
