package requests

import (
	"calendar/internal/structs/cal"
	"errors"

	_ "github.com/lib/pq"
)

func (store *DBEventSotre) NewEvent(data *cal.Event) error {
	conn, err := store.DB()
	if err != nil {
		return err
	}
	defer conn.Close()

	_, err = conn.Exec(
		"INSERT INTO events (id ,user_id, date, title, body) VALUES ($1, $2, $3, $4, $5)",
		data.ID,
		data.UserID,
		data.Date,
		data.Title,
		data.Body,
	)
	if err != nil {
		return err
	}

	return nil
}

func (store *DBEventSotre) UpdateEvent(data *cal.Event) error {
	conn, err := store.DB()
	if err != nil {
		return err
	}
	defer conn.Close()

	err = findByEventId(data.ID)
	if err != nil {
		return errors.New("no such data")
	}

	_, err = conn.Exec(
		"UPDATE events SET date = $1, title = $2, body = $3 WHERE id = $4",
		data.Date,
		data.Title,
		data.Body,
		data.ID,
	)
	if err != nil {
		return err
	}

	return nil
}

func (store *DBEventSotre) DeleteEvent(event_id int) error {
	conn, err := store.DB()
	if err != nil {
		return err
	}
	defer conn.Close()

	_, err = conn.Exec("DELETE FROM events WHERE id = $1", event_id)
	if err != nil {
		return err
	}

	return nil
}
