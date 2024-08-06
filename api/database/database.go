package database

import (
	"fmt"
	"os"
	"path/filepath"
	"sync"

	"database/sql"
	"log/slog"

	_ "github.com/glebarez/go-sqlite"
)

type Database struct {
	m    sync.Mutex
	conn *sql.DB
}

var internal *Database = nil

func Get() *Database {
	return internal
}

func Set(d *Database) {
	internal = d
}

func New() (*Database, error) {
	config, err := os.UserConfigDir()
	if err != nil {
		return nil, err
	}
	path := filepath.Join(config, "mahlzeit", "mahlzeit.db")
	if err := os.MkdirAll(filepath.Dir(path), os.ModePerm); err != nil {
		return nil, err
	}
	slog.Info(path)
	conn, err := sql.Open("sqlite", path)
	if err != nil {
		return nil, fmt.Errorf("Failed to connect to database: %w", err)
	}
	if err := conn.Ping(); err != nil {
		return nil, err
	}
	return &Database{
		conn: conn,
		m:    sync.Mutex{},
	}, nil
}

func (d *Database) Destroy() error {
	d.m.Lock()
	return d.conn.Close()
}
