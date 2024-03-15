package config

import (
	"database/sql"
	"os"
	"strconv"
	"time"
)

const (
	defaultMaxIdleConns = 5
	defaultMaxOpenConns = 10
	defaultMaxLifetime  = 5
)

func ConfigDB(db *sql.DB) {
	maxIdleConns, err := strconv.Atoi(os.Getenv("POSTGRES_MAX_IDLE_CONNS"))
	if err != nil {
		maxIdleConns = defaultMaxIdleConns
	}

	maxOpenConns, err := strconv.Atoi(os.Getenv("POSTGRES_MAX_OPEN_CONNS"))
	if err != nil {
		maxOpenConns = defaultMaxOpenConns
	}

	maxLifetime, err := strconv.Atoi(os.Getenv("POSTGRES_MAX_LIFETIME"))
	if err != nil {
		maxLifetime = defaultMaxLifetime
	}

	db.SetMaxIdleConns(maxIdleConns)
	db.SetMaxOpenConns(maxOpenConns)
	db.SetConnMaxLifetime(time.Second * time.Duration(maxLifetime))

}
