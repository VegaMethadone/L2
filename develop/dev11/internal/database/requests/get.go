package requests

import (
	"calendar/internal/structs/cal"
	"time"

	_ "github.com/lib/pq"
)

func (store *DBEventSotre) GetEventsForDay(userid int, date time.Time) ([]*cal.Event, error) {
	conn, err := store.DB()
	if err != nil {
		return nil, err
	}
	defer conn.Close()

	startOfDay := time.Date(date.Year(), date.Month(), date.Day(), 0, 0, 0, 0, date.Location())
	endOfDay := startOfDay.Add(24 * time.Hour)

	rows, err := conn.Query(
		"SELECT id, user_id, date, title, body FROM events WHERE user_id = $1 AND date >= $2 AND date < $3",
		userid, startOfDay, endOfDay,
	)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var events []*cal.Event
	for rows.Next() {
		var event cal.Event
		if err := rows.Scan(&event.ID, &event.UserID, &event.Date, &event.Title, &event.Body); err != nil {
			return nil, err
		}
		events = append(events, &event)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return events, nil
}

func (store *DBEventSotre) GetEventsForWeek(userid int, date time.Time) ([]*cal.Event, error) {
	conn, err := store.DB()
	if err != nil {
		return nil, err
	}
	defer conn.Close()

	year, week := date.ISOWeek()
	startOfWeek := time.Date(year, time.January, 1, 0, 0, 0, 0, time.UTC)
	for startOfWeek.Weekday() != time.Monday {
		startOfWeek = startOfWeek.AddDate(0, 0, 1)
	}
	startOfWeek = startOfWeek.AddDate(0, 0, (week-1)*7)

	endOfWeek := startOfWeek.AddDate(0, 0, 7)

	rows, err := conn.Query(
		"SELECT id, user_id, date, title, body FROM events WHERE user_id = $1 AND date >= $2 AND date < $3",
		userid, startOfWeek, endOfWeek,
	)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var events []*cal.Event
	for rows.Next() {
		var event cal.Event
		if err := rows.Scan(&event.ID, &event.UserID, &event.Date, &event.Title, &event.Body); err != nil {
			return nil, err
		}
		events = append(events, &event)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return events, nil
}

func (store *DBEventSotre) GetEventsForMonth(userid int, date time.Time) ([]*cal.Event, error) {
	conn, err := store.DB()
	if err != nil {
		return nil, err
	}
	defer conn.Close()

	year, month, _ := date.Date()
	startOfMonth := time.Date(year, month, 1, 0, 0, 0, 0, date.Location())
	endOfMonth := startOfMonth.AddDate(0, 1, 0)

	rows, err := conn.Query(
		"SELECT id, user_id, date, title, body FROM events WHERE user_id = $1 AND date >= $2 AND date < $3",
		userid, startOfMonth, endOfMonth,
	)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var events []*cal.Event
	for rows.Next() {
		var event cal.Event
		if err := rows.Scan(&event.ID, &event.UserID, &event.Date, &event.Title, &event.Body); err != nil {
			return nil, err
		}
		events = append(events, &event)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return events, nil
}
