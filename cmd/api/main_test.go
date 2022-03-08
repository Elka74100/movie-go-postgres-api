package main

import (
	"flag"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestOpenDB(t *testing.T) {
	var cfg config

	flag.StringVar(&cfg.db.dsn, "dsn", "postgres://postgres:magical_password@localhost/my_db_name?sslmode=disable", "Postgres connection string")

	_, err := openDB(cfg)

	assert.Nil(t, err, "Open DB should not raise an error")
}
