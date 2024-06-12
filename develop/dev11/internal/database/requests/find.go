package requests

import (
	"calendar/internal/database/connection"

	_ "github.com/lib/pq"
)

func findByEventId(event_id int) error {
	conn, err := connection.DB()
	if err != nil {
		return err
	}
	defer conn.Close()

	_, err = conn.Query("SELECT * FROM events WHERE id = $1", event_id)
	if err != nil {
		return err
	}
	return nil
}
