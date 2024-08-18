package main

import (
	"github.com/urfave/cli/v2"
)

var (
	flagPostgresURL string
)

var (
	flagAuthSecret string
)

var flags = []cli.Flag{
	&cli.StringFlag{
		Name:        "postgresql-url",
		Aliases:     []string{"pq-url", "pq"},
		Usage:       "Postgres URL",
		EnvVars:     []string{"POSTGRESQL_URL"},
		Category:    "Database",
		Required:    true,
		Destination: &flagPostgresURL,
	},
	&cli.StringFlag{
		Name:        "auth-secret",
		Aliases:     []string{"auth-key", "key"},
		Usage:       "Authentication secret, used in user validation and user register",
		EnvVars:     []string{"AUTH_SECRET"},
		Category:    "Authentication",
		Required:    true,
		Destination: &flagAuthSecret,
	},
}