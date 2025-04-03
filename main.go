package main

import (
	"context"
	"database/sql"
	"flag"
	"log"
	"os"
	"sync"
	"time"

	_ "github.com/jackc/pgx/v5/stdlib"
	"github.com/pressly/goose/v3"

	"github.com/anon-d/migrator/config"
)

type PostgresDB struct {
	*sql.DB
	mu  sync.Mutex
	dsn string
}

func New(driver, dsn string) (*PostgresDB, error) {
	db, err := connect(driver, dsn)
	if err != nil {
		return nil, err
	}
	return &PostgresDB{
		DB:  db,
		dsn: dsn,
	}, nil
}

func connect(driver, dsn string) (*sql.DB, error) {
	db, err := sql.Open(driver, dsn)
	if err != nil {
		return nil, err
	}
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := db.PingContext(ctx); err != nil {
		return nil, err
	}
	db.SetMaxOpenConns(10)
	db.SetMaxIdleConns(5)
	db.SetConnMaxLifetime(5 * time.Minute)

	return db, nil
}

func main() {

	cfg, err := config.MustLoad()
	if err != nil {
		os.Exit(1)
	}
	log.Print("Init config: complete")

	db, err := New("pgx", cfg.DBUrl)
	if err != nil {
		log.Printf("Error: %v", err)
		os.Exit(1)
	}
	log.Print("Init database: complete")

	defer db.Close()
	if err := db.Ping(); err != nil {
		os.Exit(1)
	}

	command := flag.String("command", "up", "команда миграции (up, down, status)")
	flag.Parse()
	goose.SetDialect(cfg.Driver)
	log.Print("Set goose dialect: complete")

	var out error
	switch *command {
	case "up":
		out = goose.Up(db.DB, "./migrations")
	case "down":
		out = goose.Down(db.DB, "./migrations")
	case "status":
		out = goose.Status(db.DB, "./migrations")
	default:
		os.Exit(1)
	}
	log.Printf("%v", out)
	log.Print("Migration complete")
}
