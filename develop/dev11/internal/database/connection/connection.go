package connection

import (
	"calendar/config"
	"database/sql"
	"fmt"
	"log"

	_ "github.com/lib/pq"
)

func getConnectionStr() string {
	conf, err := config.GetConfig()
	if err != nil {
		log.Printf("Config is damaged: %v", err)
	}
	str := fmt.Sprintf("user=%s password=%s dbname=%s sslmode=%s host=%s", conf.Postgres.Username, conf.Postgres.Password, conf.Postgres.DatabaseName, conf.Postgres.Sslmode, conf.Postgres.Host)
	return str
}

func DB() (*sql.DB, error) {
	db, err := sql.Open("postgres", getConnectionStr())
	if err != nil {
		log.Fatal(err)
		return nil, err
	}

	return db, nil
}
