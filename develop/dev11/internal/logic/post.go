package logic

import (
	"calendar/internal/database/connection"
	"calendar/internal/database/requests"
	"calendar/internal/structs/cal"
)

func AddEvent(data *cal.Event) error {
	var store EventStore = &requests.DBEventSotre{DB: connection.DB}
	err := store.NewEvent(data)
	if err != nil {
		return err
	}

	return nil
}

func UpdateEvent(data *cal.Event) error {
	var store EventStore = &requests.DBEventSotre{DB: connection.DB}
	err := store.UpdateEvent(data)
	if err != nil {
		return err
	}

	return nil
}

func DeleteEvent(data *cal.Event) error {
	var store EventStore = &requests.DBEventSotre{DB: connection.DB}
	err := store.DeleteEvent(data.ID)
	if err != nil {
		return err
	}

	return nil
}
