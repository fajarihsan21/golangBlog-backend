package database

import (
	"time"

	"github.com/jmoiron/sqlx"
	log "github.com/sirupsen/logrus"
)

type Postgres struct {
	master *sqlx.DB
}

func New() *Postgres {
	addrMaster := "postgresql://localhost:5432/postgres"
	db, err := sqlx.Open("postgres", addrMaster)
	db.SetConnMaxLifetime(time.Minute)
	if err != nil {
		log.Fatal("connection error")
	}

	return &Postgres{
		master: db,
	}
}

func (p *Postgres) GetActiveDB() *sqlx.DB {
	return p.master
}
