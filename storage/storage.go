package storage

import (
	"database/sql"
	"fmt"
	"log"
	"sync"

	_ "github.com/lib/pq"
)

var (
	db   *sql.DB
	once sync.Once
)

const (
	conPos = "postgres://postgres:Renzo08@localhost:5432/gbdb?sslmode=disable"
)

func NewPostgreDB() {

	once.Do(func() {
		var err error
		db, err = sql.Open("postgres", conPos)
		if err != nil {
			log.Fatalf("can't open db: %v", err)
		}

		if err = db.Ping(); err != nil {
			log.Fatalf("can't do ping: %v", err)
		}
		fmt.Println("Connection successfull ")
	})
}

// Pool return a unique instance of db
func Pool() *sql.DB {
	return db
}
