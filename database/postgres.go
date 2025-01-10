package database

import (
	"fmt"
	"os"
	"time"

	"github.com/jmoiron/sqlx"
	"github.com/joho/godotenv"
	log "github.com/sirupsen/logrus"
)

type Postgres struct {
	master *sqlx.DB
}

var (
	host   = os.Getenv("HOST")
	port   = os.Getenv("PORT")
	dbName = os.Getenv("DBNAME")
	pwd    = os.Getenv("DBPWD")
)

func New() *Postgres {
	_, err := os.Stat(".env")
	if !os.IsNotExist(err) {
		errEnv := godotenv.Load()
		if errEnv != nil {
			log.Fatal("Error loading .env file")
		}
	}

	config := fmt.Sprintf("host=%s port=%s dbname=%s password=%s sslmode=disable", host, port, dbName, pwd)
	db, err := sqlx.Open("postgres", config)
	db.SetConnMaxLifetime(time.Minute)
	if err != nil {
		log.Fatal("connection error")
	}

	err = db.Ping()
	if err != nil {
		log.Panic("Error connecting to database", err)
		panic(err)
	}

	return &Postgres{
		master: db,
	}
}

func (p *Postgres) GetActiveDB() *sqlx.DB {
	return p.master
}
