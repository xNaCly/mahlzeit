package database

import (
	"fmt"
	"os"
	"path/filepath"

	"database/sql"

	_ "github.com/glebarez/go-sqlite"
)

type Database struct {
	conn *sql.DB
}

var internal *Database = &Database{}

func Get() *Database {
	return internal
}

func (d *Database) Init() error {
	home, err := os.UserHomeDir()
	if err != nil {
		return err
	}
	conn, err := sql.Open("sqlite", filepath.Join(home, "mahlzeit", "sqlite.db"))
	if err != nil {
		return fmt.Errorf("Failed to connect to database: %w", err)
	}
	d.conn = conn
	return nil
}

func (d *Database) Destroy() error {
	return d.conn.Close()
}
