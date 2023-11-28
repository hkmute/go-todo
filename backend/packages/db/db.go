package db

import (
	"context"
	"log"
	"os"

	"github.com/jackc/pgx/v5/pgxpool"
)

type dB struct {
	conn *pgxpool.Pool
}

var DB dB

func (d *dB) Connect() *pgxpool.Pool {
	conn, err := pgxpool.New(context.Background(), DSN())
	if err != nil {
		log.Printf("Unable to create connection pool: %v\n", err)
		os.Exit(1)
	}

	err = conn.Ping(context.Background())
	if err != nil {
		log.Printf("Unable to connect to database: %v\n", err)
		os.Exit(1)
	}

	log.Print("Connected to database")
	d.conn = conn

	return d.conn
}

func (d *dB) Close() {
	log.Println("Closing database connection")
	if d.conn != nil {
		d.conn.Close()
	}
}

func (d *dB) Conn() *pgxpool.Pool {
	return d.conn
}
