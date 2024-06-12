package logic

import (
	"calendar/internal/structs/cal"
	"time"
)

type EventStore interface {
	NewEvent(data *cal.Event) error
	UpdateEvent(data *cal.Event) error
	DeleteEvent(eventID int) error
	GetEventsForDay(userID int, date time.Time) ([]*cal.Event, error)
	GetEventsForWeek(userID int, date time.Time) ([]*cal.Event, error)
	GetEventsForMonth(userID int, date time.Time) ([]*cal.Event, error)
}
