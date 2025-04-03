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

	log.Print("Init config...")
	cfg, err := config.MustLoad()
	if err != nil {
		log.Fatalf("Error: %+v", err)
	}

	log.Print("Init database...")
	db, err := New("pgx", cfg.DBUrl)
	if err != nil {
		log.Printf("Error: %v", err)
		os.Exit(1)
	}

	defer db.Close()
	if err := db.Ping(); err != nil {
		log.Fatal("База данных не отвечает")
		os.Exit(1)
	}

	step := flag.Int("step", 1, "шаг миграции")
	command := flag.String("command", "up", "команда миграции (up, down, upto, downto, status)")
	flag.Parse()
	goose.SetDialect(cfg.Driver)
	log.Print("Set goose dialect: complete")

	var out error
	switch *command {
	case "up":
		out = goose.Up(db.DB, "./migrations")
	case "down":
		out = goose.Down(db.DB, "./migrations")
	case "upto":
		out = goose.UpTo(db.DB, "./migrations", int64(*step))
	case "downto":
		out = goose.DownTo(db.DB, "./migrations", int64(*step))
	case "status":
		out = goose.Status(db.DB, "./migrations")
	default:
		log.Fatal("Неверная команда")
	}
	log.Printf("%v", out)
	log.Print("Migration complete")
}
