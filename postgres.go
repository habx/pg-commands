package pgcommands

import (
	"fmt"
)

type Postgres struct {
	// DB Host (e.g. 127.0.0.1)
	Host string
	// DB Port (e.g. 5432)
	Port int
	// DB Name
	DB string
	// Connection Username
	Username string
	// Connection Password
	Password string
	// Connection Password ENV format PGPASSWORD=
	EnvPassword string
}

func (x *Postgres) Parse() []string {
	var options []string

	if x.DB != "" {
		options = append(options, fmt.Sprintf(`--dbname=%v`, x.DB))
	}

	if x.Host != "" {
		options = append(options, fmt.Sprintf(`--host=%v`, x.Host))
	}

	if x.Port != 0 {
		options = append(options, fmt.Sprintf(`--port=%v`, x.Port))
	}

	if x.Username != "" {
		options = append(options, fmt.Sprintf(`--username=%v`, x.Username))
	}

	if x.Password != "" {
		x.EnvPassword = fmt.Sprintf(`PGPASSWORD=%v`, x.Password)
	}

	return options
}
