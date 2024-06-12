package config

import "calendar/internal/structs/conf"

func GetConfig() (*conf.Config, error) {

	conf := &conf.Config{
		Version: 1,
		Env:     "dev",
		Network: conf.Network{
			Address:      "127.0.0.1:",
			Port:         "8080",
			WriteTimeout: 15,
			ReadTimeout:  15,
		},
		Postgres: conf.Postgres{
			Host:         "localhost",
			Port:         "5432",
			Username:     "postgres",
			Password:     "0000",
			DatabaseName: "testDB",
			Sslmode:      "disable",
		},
	}

	return conf, nil
}
