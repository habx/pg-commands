package fixtures

import (
	"os"
	"strconv"

	pg "github.com/habx/pg-commands"
)

const PostgresPort = 5432

func Setup() *pg.Postgres {
	pgHost := "localhost"
	if os.Getenv("HABX_ENV") == "test" && os.Getenv("CI") == "true" {
		pgHost = "postgres"
	}
	config := &pg.Postgres{
		Host:     pgHost,
		Port:     PostgresPort,
		DB:       "dev_example",
		Username: "test",
		Password: "azerty",
	}
	if os.Getenv("TEST_DB_HOST") != "" {
		config.Host = os.Getenv("TEST_DB_HOST")
	}
	if os.Getenv("TEST_DB_PORT") != "" {
		port, err := strconv.Atoi(os.Getenv("TEST_DB_PORT"))
		if err != nil {
			panic(err)
		}
		config.Port = port
	}
	if os.Getenv("TEST_DB_NAME") != "" {
		config.DB = os.Getenv("TEST_DB_NAME")
	}
	if os.Getenv("TEST_DB_USER") != "" {
		config.Username = os.Getenv("TEST_DB_USER")
	}
	if os.Getenv("TEST_DB_PASS") != "" {
		config.Password = os.Getenv("TEST_DB_PASS")
	}

	return config
}
