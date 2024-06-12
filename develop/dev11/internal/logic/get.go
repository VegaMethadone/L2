package logic

import (
	"calendar/internal/database/connection"
	"calendar/internal/database/requests"
	"calendar/internal/structs/cal"
	"time"
)

func GetEventsForDay(userID int, date time.Time) ([]*cal.Event, error) {
	var store EventStore = &requests.DBEventSotre{DB: connection.DB}
	data, err := store.GetEventsForDay(userID, date)
	if err != nil {
		return nil, err
	}

	return data, nil
}
func GetEventsForWeek(userID int, date time.Time) ([]*cal.Event, error) {
	var store EventStore = &requests.DBEventSotre{DB: connection.DB}
	data, err := store.GetEventsForWeek(userID, date)
	if err != nil {
		return nil, err
	}

	return data, nil
}
func GetEventsForMonth(userID int, date time.Time) ([]*cal.Event, error) {
	var store EventStore = &requests.DBEventSotre{DB: connection.DB}
	data, err := store.GetEventsForMonth(userID, date)
	if err != nil {
		return nil, err
	}

	return data, nil
}
