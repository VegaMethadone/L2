package requests

import "database/sql"

type DBEventSotre struct {
	DB func() (*sql.DB, error)
}
